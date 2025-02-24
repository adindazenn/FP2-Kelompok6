package controller

import (
	"net/http"

	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/helper"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/input"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/response"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/service"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type photoController struct {
	photoService service.PhotoService
	userService  service.UserService
}

func NewPhotoController(photoService service.PhotoService, userService service.UserService) *photoController {
	return &photoController{photoService, userService}
}

func (h *photoController) AddNewPhoto(c *gin.Context) {
	var input input.PhotoCreateInput

	// Get user yang terotentikasi dari token JWT
	currentUser, tokenErr := GetUserFromToken(c)
	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	err := c.ShouldBindJSON(&input)

	photo, err := govalidator.ValidateStruct(input)

	if !photo {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// send to service
	newPhoto, err := h.photoService.CreatePhoto(input, currentUser.ID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPhotoResponse := response.CreatePhotoResponse{
		ID:        newPhoto.ID,
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		PhotoURL:  input.PhotoURL,
		UserID:    currentUser.ID,
		CreatedAt: newPhoto.CreatedAt,
	}

	c.JSON(http.StatusOK, newPhotoResponse)
}

func (h *photoController) DeletePhoto(c *gin.Context) {
	// Get user yang terotentikasi dari token JWT
	currentUser, tokenErr := GetUserFromToken(c)
	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	var idPhotoUri input.PhotoDeleteIDUser

	err := c.ShouldBindUri(&idPhotoUri)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	idPhoto := idPhotoUri.ID

	if idPhoto == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = h.photoService.DeletePhoto(idPhoto, currentUser.ID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deleteResponse := response.PhotoDeleteResponse{
		Message: "Your photo has been successfully deleted",
	}

	c.JSON(http.StatusOK, deleteResponse)
}

func (h *photoController) GetPhotos(c *gin.Context) {
	// Get user yang terotentikasi dari token JWT
	_, tokenErr := GetUserFromToken(c)
	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	photos, err := h.photoService.GetPhotosAll()
	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var photoResponse []response.GetPhotoUser

	for _, index := range photos {

		userTmp, _ := h.userService.GetUserByID(index.UserID)

		photoResponseTmp := response.GetPhotoUser{
			ID:        index.ID,
			Title:     index.Title,
			Caption:   index.Caption,
			PhotoURL:  index.PhotoURL,
			UserID:    index.UserID,
			CreatedAt: index.CreatedAt,
			UpdatedAt: index.UpdatedAt,
			User: response.UserInPhoto{
				Username: userTmp.Username,
				Email:    userTmp.Email,
			},
		}

		photoResponse = append(photoResponse, photoResponseTmp)
	}

	c.JSON(http.StatusOK, photoResponse)
}

func (h *photoController) GetPhoto(c *gin.Context) {

	var idPhotoUri input.PhotoDeleteIDUser

	err := c.ShouldBindUri(&idPhotoUri)
	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	idPhoto := idPhotoUri.ID

	if idPhoto == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	photo, err := h.photoService.GetPhotoByID(idPhoto)

	if err != nil {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, err := h.userService.GetUserByID(photo.UserID)
	if err != nil {
		response := helper.APIResponse("failed", "user not found!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// _, _ := h.commentService.GetCommentsByPhotoID(idPhoto)

	photoResponse := response.GetPhotoUser{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
		User: response.UserInPhoto{
			Username: user.Username,
			Email:    user.Email,
		},
	}

	c.JSON(http.StatusOK, photoResponse)
}

func (h *photoController) UpdatePhoto(c *gin.Context) {
	// Get user yang terotentikasi dari token JWT
	currentUser, tokenErr := GetUserFromToken(c)
	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	updatePhoto := input.PhotoUpdateInput{}

	err := c.ShouldBindJSON(&updatePhoto)

	photo, err := govalidator.ValidateStruct(updatePhoto)

	if !photo {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var idPhotoUri input.PhotoUpdateIDUser

	err = c.ShouldBindUri(&idPhotoUri)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	id_photo := idPhotoUri.ID

	_, err = h.photoService.UpdatePhoto(currentUser.ID, id_photo, updatePhoto)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	photoUpdated, _ := h.photoService.GetPhotoByID(id_photo)

	photoResponse := response.PhotoUpdateResponse{
		ID:        photoUpdated.ID,
		Title:     photoUpdated.Title,
		Caption:   photoUpdated.Caption,
		PhotoURL:  photoUpdated.PhotoURL,
		UserID:    photoUpdated.UserID,
		UpdatedAt: photoUpdated.UpdatedAt,
	}

	c.JSON(http.StatusOK, photoResponse)
}
