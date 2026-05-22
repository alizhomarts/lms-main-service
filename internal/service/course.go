package service

import (
	"context"
	"lms-main-service/internal/apperror"
	"lms-main-service/internal/entity"
	"lms-main-service/internal/repository"
)

type (
	CourseService interface {
		Create(ctx context.Context, course *entity.Course) error
		GetAll(ctx context.Context, limit, offset int) ([]entity.Course, error)
		GetByID(ctx context.Context, id uint) (*entity.Course, error)
		Update(ctx context.Context, id uint, course *entity.Course) error
		Delete(ctx context.Context, id uint) error
	}

	courseService struct {
		repo repository.CourseRepository
	}
)

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) Create(ctx context.Context, course *entity.Course) error {
	if course.Name == "" {
		return apperror.ErrCourseNameRequired
	}

	return s.repo.Create(ctx, course)
}

func (s *courseService) GetAll(ctx context.Context, limit, offset int) ([]entity.Course, error) {
	return s.repo.GetAll(ctx, limit, offset)
}

func (s *courseService) GetByID(ctx context.Context, id uint) (*entity.Course, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *courseService) Update(ctx context.Context, id uint, course *entity.Course) error {
	existingCourse, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if course.Name == "" {
		return apperror.ErrCourseNameRequired
	}

	existingCourse.Name = course.Name
	existingCourse.Description = course.Description

	return s.repo.Update(ctx, existingCourse)
}

func (s *courseService) Delete(ctx context.Context, id uint) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
