package main

import (
	"fmt"
	"os"

	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/config"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/controller"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/entity"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/repository"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db := config.InitDB()
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Comment{})
	db.AutoMigrate(&entity.Photo{})
	db.AutoMigrate(&entity.SocialMedia{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoController := controller.NewPhotoController(photoService, userService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository, photoRepository)
	commentController := controller.NewCommentController(commentService, photoService)

	socialmediaRepository := repository.NewSocialMediaRepository(db)
	socialmediaService := service.NewSocialMediaService(socialmediaRepository)
	socialmediaController := controller.NewSocialMediaController(socialmediaService, userService)

	router := gin.Default()

	// Delete
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.Login)
	router.PUT("/users", userController.UpdateUser)
	router.DELETE("/users", userController.DeleteUser)

	// photos
	router.POST("/photos", photoController.AddNewPhoto)
	router.DELETE("/photos/:id", photoController.DeletePhoto)
	router.GET("/photos", photoController.GetPhotos)
	router.GET("/photos/:id", photoController.GetPhoto)
	router.PUT("/photos/:id", photoController.UpdatePhoto)

	// comments
	router.POST("/comments", commentController.AddNewComment)
	router.DELETE("/comments/:id", commentController.DeleteComment)
	router.GET("/comments", commentController.GetComment)
	router.PUT("/comments/:id", commentController.UpdateComment)

	// social media
	router.POST("/socialmedias", socialmediaController.AddNewSocialMedia)
	router.GET("/socialmedias", socialmediaController.GetSocialMedia)
	router.PUT("/socialmedias/:id", socialmediaController.UpdateSocialMedia)
	router.DELETE("/socialmedias/:id", socialmediaController.DeleteSocialmedia)

	router.Run(":" + os.Getenv("PORT"))

	// router.Run()
}
