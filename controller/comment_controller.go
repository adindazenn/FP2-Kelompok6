package controller

import (
	"net/http"

	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/response"

	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/helper"

	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/input"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/service"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentService service.CommentService
	photoService   service.PhotoService
	userService    service.UserService
}

func NewCommentController(commentService service.CommentService, photoService service.PhotoService, userService service.UserService) *commentController {
	return &commentController{commentService, photoService, userService}
}

// Post New Comment
func (h *commentController) AddNewComment(c *gin.Context) {
	var input input.CommentInput

	// Get user yang terotentikasi dari token JWT
	currentUser, tokenErr := GetUserFromToken(c)
	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	err := c.ShouldBindJSON(&input)

	comment, err := govalidator.ValidateStruct(input)

	if !comment {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// Send To Service
	newComment, err := h.commentService.CreateComment(input, currentUser.ID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCommentResponse := response.CreateCommentResponse{
		ID:        newComment.ID,
		Message:   newComment.Message,
		PhotoID:   newComment.PhotoID,
		UserID:    newComment.UserID,
		CreatedAt: newComment.CreatedAt,
	}

	c.JSON(http.StatusCreated, newCommentResponse)
}

// Delete Comment by id
func (h *commentController) DeleteComment(c *gin.Context) {
	// Get user yang terotentikasi dari token JWT
	currentUser, tokenErr := GetUserFromToken(c)
	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	var idCommentUri input.DeleteComment

	err := c.ShouldBindUri(&idCommentUri)

	comment, err := govalidator.ValidateStruct(idCommentUri)

	if !comment {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	idComment := idCommentUri.ID

	if idComment == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = h.commentService.DeleteComment(currentUser.ID, idComment)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	responseComment := response.CommentDeleteResponse{
		Message: "Your comment has been successfully deleted",
	}

	c.JSON(http.StatusOK, responseComment)
}

// Get Comment All Comment
func (h *commentController) GetComment(c *gin.Context) {
	// Get user yang terotentikasi dari token JWT
	_, tokenErr := GetUserFromToken(c)
	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	comments, err := h.commentService.GetCommentAll()
	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Query Photo
	var allCommentsPhoto []response.GetCommentResponse
	for _, item := range comments {
		userTmp, _ := h.userService.GetUserByID(item.UserID)
		photo, _ := h.photoService.GetPhotoByID(item.PhotoID)
		
		allCommentsPhotoTmp := response.GetAllComment(item, photo, userTmp)
		allCommentsPhoto = append(allCommentsPhoto, allCommentsPhotoTmp)
	}

	c.JSON(http.StatusOK, allCommentsPhoto)
}

// Edit/Update Comment (message) by id
func (h *commentController) UpdateComment(c *gin.Context) {
	// Get user yang terotentikasi dari token JWT
	currentUser, tokenErr := GetUserFromToken(c)
	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	UpdateComment := input.CommentUpdateInput{}

	err := c.ShouldBindJSON(&UpdateComment)

	comment, err := govalidator.ValidateStruct(UpdateComment)

	if !comment {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var idCommentUri input.UpdateComment

	err = c.ShouldBindUri(&idCommentUri)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	id_comment := idCommentUri.ID

	_, err = h.commentService.UpdateComment(currentUser.ID, id_comment, UpdateComment)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	Updated, _ := h.commentService.GetCommentByID(id_comment)

	if Updated.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, "comment not found")
		return
	}

	responseComment := response.CommentUpdateResponse{
		ID:        Updated.ID,
		Message:   Updated.Message,
		PhotoID:   Updated.PhotoID,
		UserID:    Updated.UserID,
		UpdatedAt: Updated.UpdatedAt,
	}

	c.JSON(http.StatusOK, responseComment)
}
