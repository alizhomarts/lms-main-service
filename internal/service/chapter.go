package service

import (
	"lms-main-service/internal/apperror"
	"lms-main-service/internal/entity"
	"lms-main-service/internal/repository"
)

type (
	ChapterService interface {
		Create(chapter *entity.Chapter) error
		GetAll() ([]entity.Chapter, error)
		GetByID(id uint) (*entity.Chapter, error)
		Update(id uint, chapter *entity.Chapter) error
		Delete(id uint) error
	}
	chapterService struct {
		repo repository.ChapterRepository
	}
)

func NewChapterService(repo repository.ChapterRepository) ChapterService {
	return &chapterService{repo: repo}
}

func (s *chapterService) Create(chapter *entity.Chapter) error {
	if chapter.Name == "" {
		return apperror.ErrChapterNameRequired
	}

	if chapter.Order <= 0 {
		return apperror.ErrChapterOrderRequired
	}

	if chapter.CourseID == 0 {
		return apperror.ErrCourseIDRequired
	}

	exists, err := s.repo.CourseExists(chapter.CourseID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.ErrCourseNotFound
	}

	return s.repo.Create(chapter)
}

func (s *chapterService) GetAll() ([]entity.Chapter, error) {
	return s.repo.GetAll()
}

func (s *chapterService) GetByID(id uint) (*entity.Chapter, error) {
	return s.repo.GetByID(id)
}

func (s *chapterService) Update(id uint, chapter *entity.Chapter) error {
	existingChapter, err := s.repo.GetByID(id)
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

	exists, err := s.repo.CourseExists(chapter.CourseID)
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

	return s.repo.Update(existingChapter)
}

func (s *chapterService) Delete(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
