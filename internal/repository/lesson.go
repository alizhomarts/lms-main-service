package repository

import (
	"context"
	"gorm.io/gorm"
	"lms-main-service/internal/entity"
)

type (
	LessonRepository interface {
		Create(ctx context.Context, lesson *entity.Lesson) error
		GetAll(ctx context.Context, limit, offset int) ([]entity.Lesson, error)
		GetByID(ctx context.Context, id uint) (*entity.Lesson, error)
		Update(ctx context.Context, lesson *entity.Lesson) error
		Delete(ctx context.Context, id uint) error
		ChapterExists(ctx context.Context, chapterID uint) (bool, error)
	}
	lessonRepository struct {
		db *gorm.DB
	}
)

func NewLessonRepository(db *gorm.DB) LessonRepository {
	return &lessonRepository{
		db: db,
	}
}

func (r *lessonRepository) Create(ctx context.Context, lesson *entity.Lesson) error {
	return r.db.WithContext(ctx).Create(lesson).Error
}

func (r *lessonRepository) GetAll(ctx context.Context, limit, offset int) ([]entity.Lesson, error) {
	var lessons []entity.Lesson

	err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&lessons).Error

	return lessons, err
}

func (r *lessonRepository) GetByID(ctx context.Context, id uint) (*entity.Lesson, error) {
	var lesson entity.Lesson

	err := r.db.WithContext(ctx).
		First(&lesson, id).Error

	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (r *lessonRepository) Update(ctx context.Context, lesson *entity.Lesson) error {
	return r.db.WithContext(ctx).Save(lesson).Error
}

func (r *lessonRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Lesson{}, id).Error
}

func (r *lessonRepository) ChapterExists(ctx context.Context, chapterID uint) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(&entity.Chapter{}).
		Where("id = ?", chapterID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
