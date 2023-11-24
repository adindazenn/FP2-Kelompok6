package controller

import (
	"fmt"
	"net/http"
	"errors"

	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/helper"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/input"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/entity"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/response"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/service"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func GetUserFromToken(c *gin.Context) (*entity.User, error) {
    db := config.InitDB()    
    // Mendapatkan token dari header Authorization
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        return nil, errors.New("Token JWT tidak ditemukan")
    }

    // Menghapus "Bearer " dari token
    tokenString := authHeader[len("Bearer "):]

    // Memeriksa dan memverifikasi token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("password"), nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Mendapatkan user_id dari klaim token
        user_id, ok := claims["user_id"].(float64)
        if !ok {
            return nil, errors.New("User ID tidak valid")
        }

        // Temukan pengguna berdasarkan user_id
        var user entity.User
        if err := db.Where("id = ?", int(user_id)).First(&user).Error; err != nil {
            return nil, err
        }

        return &user, nil
    } else {
        return nil, errors.New("Token JWT tidak valid")
    }
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input input.UserRegisterInput

	err := c.ShouldBindJSON(&input)

	user, err := govalidator.ValidateStruct(input)

	if !user {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": err.Error(),
		})
		fmt.Println("error: " + err.Error())
		return
	}

	if input.Age <= 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": "Age must be above 8",
		})
		return
	}

	result, err := h.userService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": err.Error(),
		})
		return
	}

	registerResponse := response.UserRegisterResponse{
		ID:       result.ID,
		Age:      result.Age,
		Email:    result.Email,
		Username: result.Username,
	}

	response := helper.APIResponse("created", registerResponse)
	c.JSON(201, response)
}

func (h *userController) Login(c *gin.Context) {
	var input input.UserLoginInput

	err := c.ShouldBindJSON(&input)

	login, err := govalidator.ValidateStruct(input)

	if !login {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})

		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	// send to services
	// get user by email
	user, err := h.userService.GetUserByEmail(input.Email)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// return when user not found!
	if user.ID == 0 {
		errorMessages := "User not found!"
		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		response := helper.APIResponse("failed", "password not match!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	    // Buat token JWT
	    token := jwt.New(jwt.SigningMethodHS256)
	    claims := token.Claims.(jwt.MapClaims)
	    claims["user_id"] = user.ID
	    tokenString, err := token.SignedString([]byte("password")) 
	
	    if err != nil {
	        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
	        return
	    }

	loginResponse := response.UserLoginResponse{
		Token: tokenString,
	}

	// return token
	response := helper.APIResponse("ok", loginResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userController) UpdateUser(c *gin.Context) {
    // Get user yang terotentikasi dari token JWT
    user, err := GetUserFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
        return
    }

    // Parse data permintaan update akun
    var request input.UserUpdateInput
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update informasi akun user
    user.Username = request.Username
    user.Email = request.Email

    // Lakukan update data user dalam database
	_, err = h.userService.UpdateUser(user.ID, request)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

    // Response data user yang telah diperbarui
    updateResponse := response.UserUpdateResponse{
        ID:        user.ID,
        Username:  user.Username,
        Email:     user.Email,
	Age:       user.Age,
        UpdatedAt: user.UpdatedAt,
    }
    c.JSON(http.StatusCreated, updateResponse)
}

func (h *userController) DeleteUser(c *gin.Context) {
	// Get user yang terotentikasi dari token JWT
	user, err := GetUserFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	userDelete, err := h.userService.DeleteUser(user.ID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deleteResponse := response.UserDeleteResponse{
		Message: "Your account has been successfully deleted with id " + fmt.Sprint(userDelete.ID) + "!",
	}

	response := helper.APIResponse("ok", deleteResponse)
	c.JSON(http.StatusOK, response)
}
