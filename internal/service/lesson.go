package service

import (
	"lms-main-service/internal/apperror"
	"lms-main-service/internal/entity"
	"lms-main-service/internal/repository"
)

type (
	LessonService interface {
		Create(lesson *entity.Lesson) error
		GetAll() ([]entity.Lesson, error)
		GetByID(id uint) (*entity.Lesson, error)
		Update(id uint, lesson *entity.Lesson) error
		Delete(id uint) error
	}

	lessonService struct {
		repo repository.LessonRepository
	}
)

func NewLessonService(repo repository.LessonRepository) LessonService {
	return &lessonService{repo: repo}
}

func (s *lessonService) Create(lesson *entity.Lesson) error {
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

	exists, err := s.repo.ChapterExists(lesson.ChapterID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.ErrChapterNotFound
	}

	return s.repo.Create(lesson)
}

func (s *lessonService) GetAll() ([]entity.Lesson, error) {
	return s.repo.GetAll()
}

func (s *lessonService) GetByID(id uint) (*entity.Lesson, error) {
	return s.repo.GetByID(id)
}

func (s *lessonService) Update(id uint, lesson *entity.Lesson) error {
	existingLesson, err := s.repo.GetByID(id)
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

	exists, err := s.repo.ChapterExists(lesson.ChapterID)
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

	return s.repo.Update(existingLesson)
}

func (s *lessonService) Delete(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
