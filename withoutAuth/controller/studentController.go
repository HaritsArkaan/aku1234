package controller

import (
	"net/http"

	"withoutAuth/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentInput struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class uint   `json:"class"`
}

// GetAllStudent godoc
// @Summary Get all Student.
// @Description Get a list of Student.
// @Tags Student
// @Produce json
// @Success 200 {object} []models.Student
// @Router /students [get]
func GetAllStudent(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var student []models.Student
	db.Find(&student)

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// GetStudentById godoc
// @Summary Get Student.
// @Description Get an Student by id.
// @Tags Student
// @Produce json
// @Param id path string true "Student id"
// @Success 200 {object} models.Student
// @Router /students/id/{id} [get]
func GetStudentById(c *gin.Context) { // Get model if exist
	var student models.Student

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// GetIdByStudent godoc
// @Summary Get Id.
// @Description Get an id by nama Student.
// @Tags Student
// @Produce json
// @Param name path string true "Student Name"
// @Success 200 {object} models.Student
// @Router /students/name/{name} [get]
func GetIdByStudent(c *gin.Context) {
	var student models.Student

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("name = ?", c.Param("name")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// CreateStudent godoc
// @Summary Create New Student.
// @Description Creating a new Student.
// @Tags Student
// @Param Body body StudentInput true "the body to create a new Student"
// @Produce json
// @Success 200 {object} models.Student
// @Router /students [post]
func CreateStudent(c *gin.Context) {
	// Validate input
	var input StudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Student
	student := models.Student{
		Name:    input.Name,
		Age:     input.Age,
		ClassID: input.Class,
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&student)

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// UpdateStudent godoc
// @Summary Update Student.
// @Description Update Student by id.
// @Tags Student
// @Produce json
// @Param id path string true "Student id"
// @Param Body body StudentInput true "the body to update student"
// @Success 200 {object} models.Student
// @Router /student/{id} [patch]
func UpdateStudent(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var student models.Student
	if err := db.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input StudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Student
	updatedInput.Name = input.Name
	updatedInput.Age = input.Age
	updatedInput.ClassID = input.Class

	db.Model(&student).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// DeleteStudent godoc
// @Summary Delete one Student.
// @Description Delete a Student by id.
// @Tags Student
// @Produce json
// @Param id path int true "Student id"
// @Success 200 {object} map[string]boolean
// @Router /students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var student models.Student
	if err := db.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&student)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
