package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func NewStudentController(repository *StudentRepository) *StudentController {
	return &StudentController{
		r:        repository,
		validate: validator.New(),
	}
}

type StudentController struct {
	r        *StudentRepository
	validate *validator.Validate
}

type IStudentController interface {
	GetStudentById(c *gin.Context)
	GetStudents(c *gin.Context)
	CreateStudent(c *gin.Context)
	UpdateStudent(c *gin.Context)
	DeleteStudent(c *gin.Context)
}

func (ctrl *StudentController) GetStudentById(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *StudentController) GetStudents(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *StudentController) CreateStudent(c *gin.Context) {
	// Parse request body into Student struct
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform validation checks
	if err := ctrl.validate.Struct(student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbStudent, err := ctrl.r.CreateStudent(student)
	if err != nil {
		// response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dbStudent)

}

func (ctrl *StudentController) UpdateStudent(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *StudentController) DeleteStudent(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
