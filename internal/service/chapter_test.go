package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"lms-main-service/internal/apperror"
	"lms-main-service/internal/entity"
	"lms-main-service/mocks"
	"testing"
)

func TestChapterService_Create_Success(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	chapter := &entity.Chapter{
		Name:        "Control Structures",
		Description: "This chapter explains how to control program flow in Go using if statements, switch statements, loops, break, continue, and defer.",
		Order:       1,
		CourseID:    1,
	}

	repo.
		On("CourseExists", chapter.CourseID).
		Return(true, nil).
		Once()

	repo.
		On("Create", chapter).
		Return(nil).
		Once()

	err := chapterService.Create(chapter)

	assert.NoError(t, err)
	repo.AssertExpectations(t)

}

func TestChapterService_Create_NameRequired(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	chapter := &entity.Chapter{
		Name:        "",
		Description: "This chapter explains how to control program flow in Go using if statements, switch statements, loops, break, continue, and defer.",
		Order:       1,
		CourseID:    1,
	}

	err := chapterService.Create(chapter)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrChapterNameRequired))

	repo.AssertNotCalled(t, "CourseExists", mock.Anything)
	repo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestChapterService_Create_OrderRequired(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	chapter := &entity.Chapter{
		Name:        "Control Structures",
		Description: "Chapter description",
		Order:       0,
		CourseID:    1,
	}

	err := chapterService.Create(chapter)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrChapterOrderRequired))

	repo.AssertNotCalled(t, "CourseExists", mock.Anything)
	repo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestChapterService_Create_CourseIDRequired(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	chapter := &entity.Chapter{
		Name:        "Control Structures",
		Description: "Chapter description",
		Order:       1,
		CourseID:    0,
	}

	err := chapterService.Create(chapter)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrCourseIDRequired))

	repo.AssertNotCalled(t, "CourseExists", mock.Anything)
	repo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestChapterService_Create_CourseNotFound(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	chapter := &entity.Chapter{
		Name:        "Control Structures",
		Description: "Chapter description",
		Order:       1,
		CourseID:    999,
	}

	repo.
		On("CourseExists", uint(999)).
		Return(false, nil).
		Once()

	err := chapterService.Create(chapter)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrCourseNotFound))

	repo.AssertNotCalled(t, "Create", mock.Anything)
	repo.AssertExpectations(t)
}

func TestChapterService_GetAll_Success(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	expectedChapters := []entity.Chapter{
		{
			ID:          1,
			Name:        "Control Structures",
			Description: "Control flow in Go",
			Order:       1,
			CourseID:    1,
		},
		{
			ID:          2,
			Name:        "Functions",
			Description: "Functions in Go",
			Order:       2,
			CourseID:    1,
		},
	}

	repo.
		On("GetAll").
		Return(expectedChapters, nil).
		Once()

	chapters, err := chapterService.GetAll()

	assert.NoError(t, err)
	assert.Len(t, chapters, 2)
	assert.Equal(t, "Control Structures", chapters[0].Name)
	assert.Equal(t, "Functions", chapters[1].Name)

	repo.AssertExpectations(t)
}

func TestChapterService_GetByID_Success(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	expectedChapter := &entity.Chapter{
		ID:          1,
		Name:        "Control Structures",
		Description: "Control flow in Go",
		Order:       1,
		CourseID:    1,
	}

	repo.
		On("GetByID", uint(1)).
		Return(expectedChapter, nil).
		Once()

	chapter, err := chapterService.GetByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, chapter)
	assert.Equal(t, uint(1), chapter.ID)
	assert.Equal(t, "Control Structures", chapter.Name)

	repo.AssertExpectations(t)
}

func TestChapterService_GetByID_NotFound(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	repo.
		On("GetByID", uint(999)).
		Return((*entity.Chapter)(nil), gorm.ErrRecordNotFound).
		Once()

	chapter, err := chapterService.GetByID(999)

	assert.Error(t, err)
	assert.Nil(t, chapter)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)

	repo.AssertExpectations(t)
}

func TestChapterService_Update_Success(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	existingChapter := &entity.Chapter{
		ID:          1,
		Name:        "Control Structures",
		Description: "Old description",
		Order:       1,
		CourseID:    1,
	}

	updateData := &entity.Chapter{
		Name:        "Control Flow in Go",
		Description: "Updated description",
		Order:       2,
		CourseID:    1,
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingChapter, nil).
		Once()

	repo.
		On("CourseExists", uint(1)).
		Return(true, nil).
		Once()

	repo.
		On("Update", mock.MatchedBy(func(chapter *entity.Chapter) bool {
			return chapter.ID == 1 &&
				chapter.Name == "Control Flow in Go" &&
				chapter.Description == "Updated description" &&
				chapter.Order == 2 &&
				chapter.CourseID == 1
		})).
		Return(nil).
		Once()

	err := chapterService.Update(1, updateData)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestChapterService_Update_NotFound(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	updateData := &entity.Chapter{
		Name:        "Control Flow in Go",
		Description: "Updated description",
		Order:       1,
		CourseID:    1,
	}

	repo.
		On("GetByID", uint(999)).
		Return((*entity.Chapter)(nil), gorm.ErrRecordNotFound).
		Once()

	err := chapterService.Update(999, updateData)

	assert.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)

	repo.AssertNotCalled(t, "CourseExists", mock.Anything)
	repo.AssertNotCalled(t, "Update", mock.Anything)
	repo.AssertExpectations(t)
}

func TestChapterService_Delete_Success(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	existingChapter := &entity.Chapter{
		ID:          1,
		Name:        "Control Structures",
		Description: "Control flow in Go",
		Order:       1,
		CourseID:    1,
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingChapter, nil).
		Once()

	repo.
		On("Delete", uint(1)).
		Return(nil).
		Once()

	err := chapterService.Delete(1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestChapterService_Delete_NotFound(t *testing.T) {
	repo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(repo)

	repo.
		On("GetByID", uint(999)).
		Return((*entity.Chapter)(nil), gorm.ErrRecordNotFound).
		Once()

	err := chapterService.Delete(999)

	assert.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)

	repo.AssertNotCalled(t, "Delete", mock.Anything)
	repo.AssertExpectations(t)
}
