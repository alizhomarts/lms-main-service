package service

import (
	"context"
	"lms-main-service/internal/apperror"
	"lms-main-service/internal/entity"
	"lms-main-service/internal/repository"
)

type (
	ChapterService interface {
		Create(ctx context.Context, chapter *entity.Chapter) error
		GetAll(ctx context.Context, limit, offset int) ([]entity.Chapter, error)
		GetByID(ctx context.Context, id uint) (*entity.Chapter, error)
		Update(ctx context.Context, id uint, chapter *entity.Chapter) error
		Delete(ctx context.Context, id uint) error
	}
	chapterService struct {
		repo repository.ChapterRepository
	}
)

func NewChapterService(repo repository.ChapterRepository) ChapterService {
	return &chapterService{
		repo: repo,
	}
}

func (s *chapterService) Create(ctx context.Context, chapter *entity.Chapter) error {
	if chapter.Name == "" {
		return apperror.ErrChapterNameRequired
	}

	if chapter.Order <= 0 {
		return apperror.ErrChapterOrderRequired
	}

	if chapter.CourseID == 0 {
		return apperror.ErrCourseIDRequired
	}

	exists, err := s.repo.CourseExists(ctx, chapter.CourseID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.ErrCourseNotFound
	}

	return s.repo.Create(ctx, chapter)
}

func (s *chapterService) GetAll(ctx context.Context, limit, offset int) ([]entity.Chapter, error) {
	return s.repo.GetAll(ctx, limit, offset)
}

func (s *chapterService) GetByID(ctx context.Context, id uint) (*entity.Chapter, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *chapterService) Update(ctx context.Context, id uint, chapter *entity.Chapter) error {
	existingChapter, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if chapter.Name == "" {
		return apperror.ErrChapterNameRequired
	}

	if chapter.Order <= 0 {
		return apperror.ErrChapterOrderRequired
	}

	if chapter.CourseID == 0 {
		return apperror.ErrCourseIDRequired
	}

	exists, err := s.repo.CourseExists(ctx, chapter.CourseID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.ErrCourseNotFound
	}

	existingChapter.Name = chapter.Name
	existingChapter.Description = chapter.Description
	existingChapter.Order = chapter.Order
	existingChapter.CourseID = chapter.CourseID

	return s.repo.Update(ctx, existingChapter)
}

func (s *chapterService) Delete(ctx context.Context, id uint) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
