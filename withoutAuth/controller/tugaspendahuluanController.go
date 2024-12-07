package controller

import (
	"net/http"
	"time"

	"withoutAuth/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TugasPendahuluanInput struct {
	Judul     string `gorm:"judul"`
	SubJudul  string `gorm:"sub_judul"`
	Kategori  string `gorm:"kategori"`
	Deadline  string `gorm:"deadline"`
	Deskripsi string `gorm:"deskripsi"`
}

// GetAllTugasPendahuluan godoc
// @Summary Get all TugasPendahuluan.
// @Description Get a list of TugasPendahuluan.
// @Tags TugasPendahuluan
// @Produce json
// @Success 200 {object} []models.TugasPendahuluan
// @Router /tugaspendahuluans [get]
func GetAllTugasPendahuluan(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var tugas []models.TugasPendahuluan
	db.Find(&tugas)

	c.JSON(http.StatusOK, gin.H{"data": tugas})
}

// GetTugasPendahuluanById godoc
// @Summary Get TugasPendahuluan.
// @Description Get an TugasPendahuluan by id.
// @Tags TugasPendahuluan
// @Produce json
// @Param id path string true "TugasPendahuluan id"
// @Success 200 {object} models.TugasPendahuluan
// @Router /tugaspendahuluans/id/{id} [get]
func GetTugasPendahuluanById(c *gin.Context) { // Get model if exist
	var tugas models.TugasPendahuluan

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&tugas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tugas})
}

// CreateTugasPendahuluan godoc
// @Summary Create New TugasPendahuluan.
// @Description Creating a new TugasPendahuluan.
// @Tags TugasPendahuluan
// @Param Body body TugasPendahuluanInput true "the body to create a new TugasPendahuluan"
// @Produce json
// @Success 200 {object} models.TugasPendahuluan
// @Router /tugaspendahuluans [post]
func CreateTugasPendahuluan(c *gin.Context) {
	// Validate input
	var input TugasPendahuluanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create TP
	tugas := models.TugasPendahuluan{
		Judul:     input.Judul,
		SubJudul:  input.SubJudul,
		Kategori:  input.Kategori,
		Deadline:  input.Deadline,
		Deskripsi: input.Deskripsi,
		CreatedAt: time.Now(),
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&tugas)

	c.JSON(http.StatusOK, gin.H{"data": tugas})
}

// UpdateTugasPendahuluan godoc
// @Summary Update TugasPendahuluan.
// @Description Update TugasPendahuluan by id.
// @Tags TugasPendahuluan
// @Produce json
// @Param id path string true "TugasPendahuluan id"
// @Param Body body TugasPendahuluanInput true "the body to update TugasPendahuluan"
// @Success 200 {object} models.TugasPendahuluan
// @Router /tugaspendahuluans/{id} [patch]
func UpdateTugasPendahuluan(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var tugas models.TugasPendahuluan
	if err := db.Where("id = ?", c.Param("id")).First(&tugas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input TugasPendahuluanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.TugasPendahuluan
	updatedInput.Judul = input.Judul
	updatedInput.SubJudul = input.SubJudul
	updatedInput.Kategori = input.Kategori
	updatedInput.Deadline = input.Deadline
	updatedInput.Deskripsi = input.Deskripsi

	db.Model(&tugas).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": tugas})
}

// DeleteTugasPendahuluan godoc
// @Summary Delete one TugasPendahuluan.
// @Description Delete a TugasPendahuluan by id.
// @Tags TugasPendahuluan
// @Produce json
// @Param id path int true "TugasPendahuluan id"
// @Success 200 {object} map[string]boolean
// @Router /tugaspendahuluans/{id} [delete]
func DeleteTugasPendahuluan(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var tugas models.TugasPendahuluan
	if err := db.Where("id = ?", c.Param("id")).First(&tugas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&tugas)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
