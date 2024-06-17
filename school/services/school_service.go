package services

import (
	"school/models"
	"school/repositories"
)

type SchoolService struct {
	Repo *repositories.SchoolRepository
}

func NewSchoolService(repo *repositories.SchoolRepository) *SchoolService {
	return &SchoolService{Repo: repo}
}

func (s *SchoolService) CreateSchool(school *models.School) error {
	return s.Repo.CreateSchool(school)
}

func (s *SchoolService) GetSchoolByID(id int) (*models.School, error) {
	return s.Repo.GetSchoolByID(id)
}

func (s *SchoolService) GetAllSchools() ([]models.School, error) {
	return s.Repo.GetAllSchools()
}

func (s *SchoolService) UpdateSchool(id int, updatedSchool *models.School) error {
	return s.Repo.UpdateSchool(id, updatedSchool)
}

func (s *SchoolService) DeleteSchoolByID(id int) error {
	return s.Repo.DeleteSchoolByID(id)
}
