package service

import (
	"lms-main-service/internal/apperror"
	"lms-main-service/internal/entity"
	"lms-main-service/internal/repository"
)

type (
	CourseService interface {
		Create(course *entity.Course) error
		GetAll() ([]entity.Course, error)
		GetByID(id uint) (*entity.Course, error)
		Update(id uint, course *entity.Course) error
		Delete(id uint) error
	}

	courseService struct {
		repo repository.CourseRepository
	}
)

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) Create(course *entity.Course) error {
	if course.Name == "" {
		return apperror.ErrCourseNameRequired
	}

	return s.repo.Create(course)
}

func (s *courseService) GetAll() ([]entity.Course, error) {
	return s.repo.GetAll()
}

func (s *courseService) GetByID(id uint) (*entity.Course, error) {
	return s.repo.GetByID(id)
}

func (s *courseService) Update(id uint, course *entity.Course) error {
	existingCourse, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if course.Name == "" {
		return apperror.ErrCourseNameRequired
	}

	existingCourse.Name = course.Name
	existingCourse.Description = course.Description

	return s.repo.Update(existingCourse)
}

func (s *courseService) Delete(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
