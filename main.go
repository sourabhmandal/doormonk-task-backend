package main

import (
	"dmbackend/config"
	"dmbackend/database"
	"dmbackend/internal/student"
	"dmbackend/internal/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	config.LoadEnvConfig()
	fmt.Println(config.EnvConfig.GIN_MODE, config.EnvConfig.GIN_MODE == gin.DebugMode)
	// Validate the mode value
	switch config.EnvConfig.GIN_MODE {
	case gin.DebugMode, gin.ReleaseMode, gin.TestMode:
		// Set the Gin mode
		gin.SetMode(config.EnvConfig.GIN_MODE)
	default:
		// Default to debug mode if mode is invalid
		fmt.Println("Invalid GIN_MODE value. Defaulting to release mode.")
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Set up custom validator
	v, _ := binding.Validator.Engine().(*validator.Validate)
	v.RegisterValidation("phone", utils.ValidatePhone)

	db, err := database.NewSqliteDb(config.EnvConfig.DB_NAME)
	if err != nil {
		panic(fmt.Errorf(`unable to connect with db: %w`, err))
	}

	studentRepo := student.NewStudentRepository(db)
	studentController := student.NewStudentController(studentRepo)

	// Register dynamic routes
	r.GET("/student/:id", studentController.GetStudentById)
	r.GET("/student", studentController.GetStudents)
	r.POST("/student", studentController.CreateStudent)
	r.PUT("/student/:id", studentController.UpdateStudent)
	r.DELETE("/student/:id", studentController.DeleteStudent)

	// Start the server
	server_port := fmt.Sprintf(":%d", config.EnvConfig.HTTP_PORT)
	log.Println("Server started on ", server_port)
	log.Fatal(http.ListenAndServe(server_port, r))
}
