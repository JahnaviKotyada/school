package controllers

import (
	"net/http"
	"school/models"
	"school/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SchoolController struct {
	Service *services.SchoolService
}

func NewSchoolController(service *services.SchoolService) *SchoolController {
	return &SchoolController{Service: service}
}

func (c *SchoolController) SetupRoutes(router *gin.Engine) {
	router.POST("/schools", c.CreateSchool)
	router.GET("/schools/:id", c.GetSchoolByID)
	router.GET("/schools", c.GetAllSchools)
	router.PUT("/schools/:id", c.UpdateSchool)
	router.DELETE("/schools/:id", c.DeleteSchoolByID)
}

func (c *SchoolController) CreateSchool(ctx *gin.Context) {
	var school models.School
	if err := ctx.ShouldBindJSON(&school); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.CreateSchool(&school); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create school"})
		return
	}

	ctx.JSON(http.StatusOK, school)
}

func (c *SchoolController) GetSchoolByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid school ID"})
		return
	}

	school, err := c.Service.GetSchoolByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}

	ctx.JSON(http.StatusOK, school)
}

func (c *SchoolController) GetAllSchools(ctx *gin.Context) {
	schools, err := c.Service.GetAllSchools()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schools"})
		return
	}

	ctx.JSON(http.StatusOK, schools)
}

func (c *SchoolController) UpdateSchool(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid school ID"})
		return
	}

	var updatedSchool models.School
	if err := ctx.ShouldBindJSON(&updatedSchool); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.UpdateSchool(id, &updatedSchool); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update school"})
		return
	}

	ctx.JSON(http.StatusOK, updatedSchool)
}

func (c *SchoolController) DeleteSchoolByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid school ID"})
		return
	}

	if err := c.Service.DeleteSchoolByID(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete school"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "School deleted"})
}
