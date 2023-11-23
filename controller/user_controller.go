package controller

import (
	"fmt"
	"net/http"

	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/helper"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/middleware"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/input"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/response"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/service"

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

	// create token
	jwtService := middleware.NewService()
	token, err := jwtService.GenerateToken(user.ID)
	if err != nil {
		response := helper.APIResponse("failed", "failed to generate token!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginResponse := response.UserLoginResponse{
		Token: token,
	}

	// return token
	response := helper.APIResponse("ok", loginResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userController) UpdateUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputUserUpdate input.UserUpdateInput

	err := c.ShouldBindJSON(&inputUserUpdate)

	user, err := govalidator.ValidateStruct(inputUserUpdate)

	if !user {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	var idUserUri input.UserUpdateID

	err = c.ShouldBindUri(&idUserUri)

	if currentUser != idUserUri.ID {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	id_user := idUserUri.ID

	_, err = h.userService.UpdateUser(id_user, inputUserUpdate)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userUpdated, err := h.userService.GetUserByID(id_user)
	if err != nil {
		response := helper.APIResponse("failed", "Cannot fetch user!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updateResponse := response.UserUpdateResponse{
		ID:        userUpdated.ID,
		Email:     userUpdated.Email,
		Username:  userUpdated.Username,
		Age:       userUpdated.Age,
		UpdatedAt: userUpdated.UpdatedAt,
	}

	response := helper.APIResponse("ok", updateResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userController) DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var idUserUri input.UserDeleteID

	err := c.ShouldBindUri(&idUserUri)

	if currentUser != idUserUri.ID {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	userDelete, err := h.userService.DeleteUser(idUserUri.ID)

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

func (h *userController) TestUser(c *gin.Context) {
	id_user, err := c.Get("currentUser")

	if !err {
		c.JSON(http.StatusNotFound, helper.APIResponse("not created", err))
		return
	}

	c.JSON(http.StatusOK, helper.APIResponse("created", id_user))
}
