package repository

import (
	"gorm.io/gorm"
	"lms-main-service/internal/entity"
)

type (
	LessonRepository interface {
		Create(lesson *entity.Lesson) error
		GetAll() ([]entity.Lesson, error)
		GetByID(id uint) (*entity.Lesson, error)
		Update(lesson *entity.Lesson) error
		Delete(id uint) error
		ChapterExists(chapterID uint) (bool, error)
	}
	lessonRepository struct {
		db *gorm.DB
	}
)

func NewLessonRepository(db *gorm.DB) LessonRepository {
	return &lessonRepository{db: db}
}

func (r *lessonRepository) Create(lesson *entity.Lesson) error {
	return r.db.Create(lesson).Error
}

func (r *lessonRepository) GetAll() ([]entity.Lesson, error) {
	var lessons []entity.Lesson

	err := r.db.Find(&lessons).Error

	return lessons, err
}

func (r *lessonRepository) GetByID(id uint) (*entity.Lesson, error) {
	var lesson entity.Lesson

	err := r.db.First(&lesson, id).Error
	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (r *lessonRepository) Update(lesson *entity.Lesson) error {
	return r.db.Save(lesson).Error
}

func (r *lessonRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Lesson{}, id).Error
}

func (r *lessonRepository) ChapterExists(chapterID uint) (bool, error) {
	var count int64

	err := r.db.
		Model(&entity.Chapter{}).
		Where("id = ?", chapterID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
