package repositories

import (
	"fmt"
	"school/models"

	"gorm.io/gorm"
)

type SchoolRepository struct {
	DB *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) *SchoolRepository {
	return &SchoolRepository{DB: db}
}

func (r *SchoolRepository) CreateSchool(school *models.School) error {
	return r.DB.Create(school).Error
}

func (r *SchoolRepository) GetSchoolByID(id int) (*models.School, error) {
	var school models.School
	if err := r.DB.First(&school, id).Error; err != nil {
		return nil, err
	}
	return &school, nil
}

func (r *SchoolRepository) GetAllSchools() ([]models.School, error) {
	var schools []models.School
	if err := r.DB.Find(&schools).Error; err != nil {
		return nil, err
	}
	return schools, nil
}

func (r *SchoolRepository) UpdateSchool(id int, updatedSchool *models.School) error {
	var school models.School
	if err := r.DB.First(&school, id).Error; err != nil {
		return err
	}
	updatedSchool.ID = uint(id) // Ensure the correct ID is set
	return r.DB.Save(updatedSchool).Error
}

func (r *SchoolRepository) DeleteSchoolByID(id int) error {
	result := r.DB.Delete(&models.School{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no school found with ID %d", id)
	}
	return nil
}
