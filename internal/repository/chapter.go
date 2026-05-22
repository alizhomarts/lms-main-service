package repository

import (
	"context"
	"gorm.io/gorm"
	"lms-main-service/internal/entity"
)

type (
	ChapterRepository interface {
		Create(ctx context.Context, chapter *entity.Chapter) error
		GetAll(ctx context.Context, limit, offset int) ([]entity.Chapter, error)
		GetByID(ctx context.Context, id uint) (*entity.Chapter, error)
		Update(ctx context.Context, chapter *entity.Chapter) error
		Delete(ctx context.Context, id uint) error
		CourseExists(ctx context.Context, courseID uint) (bool, error)
	}

	chapterRepository struct {
		db *gorm.DB
	}
)

func NewChapterRepository(db *gorm.DB) ChapterRepository {
	return &chapterRepository{
		db: db,
	}
}

func (r *chapterRepository) Create(ctx context.Context, chapter *entity.Chapter) error {
	return r.db.WithContext(ctx).Create(chapter).Error
}

func (r *chapterRepository) GetAll(ctx context.Context, limit, offset int) ([]entity.Chapter, error) {
	var chapters []entity.Chapter

	err := r.db.WithContext(ctx).
		Preload("Lessons").
		Limit(limit).
		Offset(offset).
		Find(&chapters).Error

	return chapters, err
}

func (r *chapterRepository) GetByID(ctx context.Context, id uint) (*entity.Chapter, error) {
	var chapter entity.Chapter

	err := r.db.WithContext(ctx).
		Preload("Lessons").
		First(&chapter, id).Error

	if err != nil {
		return nil, err
	}

	return &chapter, nil
}

func (r *chapterRepository) Update(ctx context.Context, chapter *entity.Chapter) error {
	return r.db.WithContext(ctx).Save(chapter).Error
}

func (r *chapterRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Chapter{}, id).Error
}

func (r *chapterRepository) CourseExists(ctx context.Context, courseID uint) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(&entity.Course{}).
		Where("id = ?", courseID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
