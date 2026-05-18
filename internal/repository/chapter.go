package repository

import (
	"gorm.io/gorm"
	"lms-main-service/internal/entity"
)

type (
	ChapterRepository interface {
		Create(chapter *entity.Chapter) error
		GetAll() ([]entity.Chapter, error)
		GetByID(id uint) (*entity.Chapter, error)
		Update(chapter *entity.Chapter) error
		Delete(id uint) error
		CourseExists(courseID uint) (bool, error)
	}

	chapterRepository struct {
		db *gorm.DB
	}
)

func NewChapterRepository(db *gorm.DB) ChapterRepository {
	return &chapterRepository{db: db}
}

func (r *chapterRepository) Create(chapter *entity.Chapter) error {
	return r.db.Create(chapter).Error
}

func (r *chapterRepository) GetAll() ([]entity.Chapter, error) {
	var chapters []entity.Chapter

	err := r.db.
		Preload("Lessons").
		Find(&chapters).Error

	return chapters, err
}

func (r *chapterRepository) GetByID(id uint) (*entity.Chapter, error) {
	var chapter entity.Chapter

	err := r.db.
		Preload("Lessons").
		First(&chapter, id).Error

	if err != nil {
		return nil, err
	}

	return &chapter, nil
}

func (r *chapterRepository) Update(chapter *entity.Chapter) error {
	return r.db.Save(chapter).Error
}

func (r *chapterRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Chapter{}, id).Error
}

func (r *chapterRepository) CourseExists(courseID uint) (bool, error) {
	var count int64

	err := r.db.
		Model(&entity.Course{}).
		Where("id = ?", courseID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
