package repository

import (
	"context"
	"gorm.io/gorm"
	"lms-main-service/internal/entity"
)

type (
	CourseRepository interface {
		Create(ctx context.Context, course *entity.Course) error
		GetAll(ctx context.Context, limit, offset int) ([]entity.Course, error)
		GetByID(ctx context.Context, id uint) (*entity.Course, error)
		Update(ctx context.Context, course *entity.Course) error
		Delete(ctx context.Context, id uint) error
	}

	courseRepository struct {
		db *gorm.DB
	}
)

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Create(ctx context.Context, course *entity.Course) error {
	return r.db.WithContext(ctx).Create(course).Error
}

func (r *courseRepository) GetAll(ctx context.Context, limit, offset int) ([]entity.Course, error) {
	var courses []entity.Course

	err := r.db.WithContext(ctx).
		Preload("Chapters").
		Limit(limit).
		Offset(offset).
		Find(&courses).Error

	return courses, err
}

func (r *courseRepository) GetByID(ctx context.Context, id uint) (*entity.Course, error) {
	var course entity.Course

	err := r.db.WithContext(ctx).
		Preload("Chapters").
		First(&course, id).Error

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *courseRepository) Update(ctx context.Context, course *entity.Course) error {
	return r.db.WithContext(ctx).Save(course).Error
}

func (r *courseRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Course{}, id).Error
}
