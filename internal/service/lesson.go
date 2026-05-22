package service

import (
	"context"
	"lms-main-service/internal/apperror"
	"lms-main-service/internal/entity"
	"lms-main-service/internal/repository"
)

type (
	LessonService interface {
		Create(ctx context.Context, lesson *entity.Lesson) error
		GetAll(ctx context.Context, limit, offset int) ([]entity.Lesson, error)
		GetByID(ctx context.Context, id uint) (*entity.Lesson, error)
		Update(ctx context.Context, id uint, lesson *entity.Lesson) error
		Delete(ctx context.Context, id uint) error
	}

	lessonService struct {
		repo repository.LessonRepository
	}
)

func NewLessonService(repo repository.LessonRepository) LessonService {
	return &lessonService{
		repo: repo,
	}
}

func (s *lessonService) Create(ctx context.Context, lesson *entity.Lesson) error {
	if lesson.Name == "" {
		return apperror.ErrLessonNameRequired
	}

	if lesson.Content == "" {
		return apperror.ErrLessonContentRequired
	}

	if lesson.Order <= 0 {
		return apperror.ErrLessonOrderRequired
	}

	if lesson.ChapterID == 0 {
		return apperror.ErrChapterIDRequired
	}

	exists, err := s.repo.ChapterExists(ctx, lesson.ChapterID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.ErrChapterNotFound
	}

	return s.repo.Create(ctx, lesson)
}

func (s *lessonService) GetAll(ctx context.Context, limit, offset int) ([]entity.Lesson, error) {
	return s.repo.GetAll(ctx, limit, offset)
}

func (s *lessonService) GetByID(ctx context.Context, id uint) (*entity.Lesson, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *lessonService) Update(ctx context.Context, id uint, lesson *entity.Lesson) error {
	existingLesson, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if lesson.Name == "" {
		return apperror.ErrLessonNameRequired
	}

	if lesson.Content == "" {
		return apperror.ErrLessonContentRequired
	}

	if lesson.Order <= 0 {
		return apperror.ErrLessonOrderRequired
	}

	if lesson.ChapterID == 0 {
		return apperror.ErrChapterIDRequired
	}

	exists, err := s.repo.ChapterExists(ctx, lesson.ChapterID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.ErrChapterNotFound
	}

	existingLesson.Name = lesson.Name
	existingLesson.Description = lesson.Description
	existingLesson.Content = lesson.Content
	existingLesson.Order = lesson.Order
	existingLesson.ChapterID = lesson.ChapterID

	return s.repo.Update(ctx, existingLesson)
}

func (s *lessonService) Delete(ctx context.Context, id uint) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
