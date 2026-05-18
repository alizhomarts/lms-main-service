package repository

import (
	"gorm.io/gorm"
	"lms-main-service/internal/entity"
)

type (
	CourseRepository interface {
		Create(course *entity.Course) error
		GetAll() ([]entity.Course, error)
		GetByID(id uint) (*entity.Course, error)
		Update(course *entity.Course) error
		Delete(id uint) error
	}

	courseRepository struct {
		db *gorm.DB
	}
)

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{
		db: db,
	}
}

func (r *courseRepository) Create(course *entity.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) GetAll() ([]entity.Course, error) {
	var courses []entity.Course

	err := r.db.
		Preload("Chapters").
		Find(&courses).Error

	return courses, err
}

func (r *courseRepository) GetByID(id uint) (*entity.Course, error) {
	var course entity.Course

	err := r.db.
		Preload("Chapters").
		First(&course, id).Error

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *courseRepository) Update(course *entity.Course) error {
	return r.db.Save(course).Error
}

func (r *courseRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Course{}, id).Error
}
