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

func TestLessonService_Create_Success(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	lesson := &entity.Lesson{
		Name:        "If-else Statement in Golang",
		Description: "This lesson explains conditional branching in Go.",
		Content:     "In Go, the if statement is used to execute code only when a specific condition is true.",
		Order:       1,
		ChapterID:   1,
	}

	repo.
		On("ChapterExists", lesson.ChapterID).
		Return(true, nil).
		Once()

	repo.
		On("Create", lesson).
		Return(nil).
		Once()

	err := lessonService.Create(lesson)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestLessonService_Create_NameRequired(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	lesson := &entity.Lesson{
		Name:        "",
		Description: "This lesson explains conditional branching in Go.",
		Content:     "Lesson content",
		Order:       1,
		ChapterID:   1,
	}

	err := lessonService.Create(lesson)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrLessonNameRequired))

	repo.AssertNotCalled(t, "ChapterExists", mock.Anything)
	repo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestLessonService_Create_ContentRequired(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	lesson := &entity.Lesson{
		Name:        "If-else Statement in Golang",
		Description: "This lesson explains conditional branching in Go.",
		Content:     "",
		Order:       1,
		ChapterID:   1,
	}

	err := lessonService.Create(lesson)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrLessonContentRequired))

	repo.AssertNotCalled(t, "ChapterExists", mock.Anything)
	repo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestLessonService_Create_OrderRequired(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	lesson := &entity.Lesson{
		Name:        "If-else Statement in Golang",
		Description: "This lesson explains conditional branching in Go.",
		Content:     "Lesson content",
		Order:       0,
		ChapterID:   1,
	}

	err := lessonService.Create(lesson)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrLessonOrderRequired))

	repo.AssertNotCalled(t, "ChapterExists", mock.Anything)
	repo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestLessonService_Create_ChapterIDRequired(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	lesson := &entity.Lesson{
		Name:        "If-else Statement in Golang",
		Description: "This lesson explains conditional branching in Go.",
		Content:     "Lesson content",
		Order:       1,
		ChapterID:   0,
	}

	err := lessonService.Create(lesson)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrChapterIDRequired))

	repo.AssertNotCalled(t, "ChapterExists", mock.Anything)
	repo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestLessonService_Create_ChapterNotFound(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	lesson := &entity.Lesson{
		Name:        "If-else Statement in Golang",
		Description: "This lesson explains conditional branching in Go.",
		Content:     "Lesson content",
		Order:       1,
		ChapterID:   999,
	}

	repo.
		On("ChapterExists", uint(999)).
		Return(false, nil).
		Once()

	err := lessonService.Create(lesson)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrChapterNotFound))

	repo.AssertNotCalled(t, "Create", mock.Anything)
	repo.AssertExpectations(t)
}

func TestLessonService_GetAll_Success(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	expectedLessons := []entity.Lesson{
		{
			ID:          1,
			Name:        "If-else Statement in Golang",
			Description: "Conditional branching in Go",
			Content:     "Lesson content",
			Order:       1,
			ChapterID:   1,
		},
		{
			ID:          2,
			Name:        "Switch Statement in Golang",
			Description: "Switch statements in Go",
			Content:     "Lesson content",
			Order:       2,
			ChapterID:   1,
		},
	}

	repo.
		On("GetAll").
		Return(expectedLessons, nil).
		Once()

	lessons, err := lessonService.GetAll()

	assert.NoError(t, err)
	assert.Len(t, lessons, 2)
	assert.Equal(t, "If-else Statement in Golang", lessons[0].Name)
	assert.Equal(t, "Switch Statement in Golang", lessons[1].Name)

	repo.AssertExpectations(t)
}

func TestLessonService_GetByID_Success(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	expectedLesson := &entity.Lesson{
		ID:          1,
		Name:        "If-else Statement in Golang",
		Description: "Conditional branching in Go",
		Content:     "Lesson content",
		Order:       1,
		ChapterID:   1,
	}

	repo.
		On("GetByID", uint(1)).
		Return(expectedLesson, nil).
		Once()

	lesson, err := lessonService.GetByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, lesson)
	assert.Equal(t, uint(1), lesson.ID)
	assert.Equal(t, "If-else Statement in Golang", lesson.Name)

	repo.AssertExpectations(t)
}

func TestLessonService_GetByID_NotFound(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	repo.
		On("GetByID", uint(999)).
		Return((*entity.Lesson)(nil), gorm.ErrRecordNotFound).
		Once()

	lesson, err := lessonService.GetByID(999)

	assert.Error(t, err)
	assert.Nil(t, lesson)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)

	repo.AssertExpectations(t)
}

func TestLessonService_Update_Success(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	existingLesson := &entity.Lesson{
		ID:          1,
		Name:        "If-else Statement in Golang",
		Description: "Old description",
		Content:     "Old content",
		Order:       1,
		ChapterID:   1,
	}

	updateData := &entity.Lesson{
		Name:        "If-else and Conditional Statements in Go",
		Description: "Updated description",
		Content:     "Updated content",
		Order:       2,
		ChapterID:   1,
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingLesson, nil).
		Once()

	repo.
		On("ChapterExists", uint(1)).
		Return(true, nil).
		Once()

	repo.
		On("Update", mock.MatchedBy(func(lesson *entity.Lesson) bool {
			return lesson.ID == 1 &&
				lesson.Name == "If-else and Conditional Statements in Go" &&
				lesson.Description == "Updated description" &&
				lesson.Content == "Updated content" &&
				lesson.Order == 2 &&
				lesson.ChapterID == 1
		})).
		Return(nil).
		Once()

	err := lessonService.Update(1, updateData)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestLessonService_Update_NotFound(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	updateData := &entity.Lesson{
		Name:        "If-else and Conditional Statements in Go",
		Description: "Updated description",
		Content:     "Updated content",
		Order:       1,
		ChapterID:   1,
	}

	repo.
		On("GetByID", uint(999)).
		Return((*entity.Lesson)(nil), gorm.ErrRecordNotFound).
		Once()

	err := lessonService.Update(999, updateData)

	assert.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)

	repo.AssertNotCalled(t, "ChapterExists", mock.Anything)
	repo.AssertNotCalled(t, "Update", mock.Anything)
	repo.AssertExpectations(t)
}

func TestLessonService_Update_NameRequired(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	existingLesson := &entity.Lesson{
		ID:          1,
		Name:        "If-else Statement in Golang",
		Description: "Old description",
		Content:     "Old content",
		Order:       1,
		ChapterID:   1,
	}

	updateData := &entity.Lesson{
		Name:        "",
		Description: "Updated description",
		Content:     "Updated content",
		Order:       1,
		ChapterID:   1,
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingLesson, nil).
		Once()

	err := lessonService.Update(1, updateData)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrLessonNameRequired))

	repo.AssertNotCalled(t, "ChapterExists", mock.Anything)
	repo.AssertNotCalled(t, "Update", mock.Anything)
	repo.AssertExpectations(t)
}

func TestLessonService_Update_ContentRequired(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	existingLesson := &entity.Lesson{
		ID:          1,
		Name:        "If-else Statement in Golang",
		Description: "Old description",
		Content:     "Old content",
		Order:       1,
		ChapterID:   1,
	}

	updateData := &entity.Lesson{
		Name:        "If-else and Conditional Statements in Go",
		Description: "Updated description",
		Content:     "",
		Order:       1,
		ChapterID:   1,
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingLesson, nil).
		Once()

	err := lessonService.Update(1, updateData)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrLessonContentRequired))

	repo.AssertNotCalled(t, "ChapterExists", mock.Anything)
	repo.AssertNotCalled(t, "Update", mock.Anything)
	repo.AssertExpectations(t)
}

func TestLessonService_Update_OrderRequired(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	existingLesson := &entity.Lesson{
		ID:          1,
		Name:        "If-else Statement in Golang",
		Description: "Old description",
		Content:     "Old content",
		Order:       1,
		ChapterID:   1,
	}

	updateData := &entity.Lesson{
		Name:        "If-else and Conditional Statements in Go",
		Description: "Updated description",
		Content:     "Updated content",
		Order:       0,
		ChapterID:   1,
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingLesson, nil).
		Once()

	err := lessonService.Update(1, updateData)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrLessonOrderRequired))

	repo.AssertNotCalled(t, "ChapterExists", mock.Anything)
	repo.AssertNotCalled(t, "Update", mock.Anything)
	repo.AssertExpectations(t)
}

func TestLessonService_Update_ChapterIDRequired(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	existingLesson := &entity.Lesson{
		ID:          1,
		Name:        "If-else Statement in Golang",
		Description: "Old description",
		Content:     "Old content",
		Order:       1,
		ChapterID:   1,
	}

	updateData := &entity.Lesson{
		Name:        "If-else and Conditional Statements in Go",
		Description: "Updated description",
		Content:     "Updated content",
		Order:       1,
		ChapterID:   0,
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingLesson, nil).
		Once()

	err := lessonService.Update(1, updateData)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrChapterIDRequired))

	repo.AssertNotCalled(t, "ChapterExists", mock.Anything)
	repo.AssertNotCalled(t, "Update", mock.Anything)
	repo.AssertExpectations(t)
}

func TestLessonService_Update_ChapterNotFound(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	existingLesson := &entity.Lesson{
		ID:          1,
		Name:        "If-else Statement in Golang",
		Description: "Old description",
		Content:     "Old content",
		Order:       1,
		ChapterID:   1,
	}

	updateData := &entity.Lesson{
		Name:        "If-else and Conditional Statements in Go",
		Description: "Updated description",
		Content:     "Updated content",
		Order:       1,
		ChapterID:   999,
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingLesson, nil).
		Once()

	repo.
		On("ChapterExists", uint(999)).
		Return(false, nil).
		Once()

	err := lessonService.Update(1, updateData)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrChapterNotFound))

	repo.AssertNotCalled(t, "Update", mock.Anything)
	repo.AssertExpectations(t)
}

func TestLessonService_Delete_Success(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	existingLesson := &entity.Lesson{
		ID:          1,
		Name:        "If-else Statement in Golang",
		Description: "Conditional branching in Go",
		Content:     "Lesson content",
		Order:       1,
		ChapterID:   1,
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingLesson, nil).
		Once()

	repo.
		On("Delete", uint(1)).
		Return(nil).
		Once()

	err := lessonService.Delete(1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestLessonService_Delete_NotFound(t *testing.T) {
	repo := new(mocks.LessonRepository)
	lessonService := NewLessonService(repo)

	repo.
		On("GetByID", uint(999)).
		Return((*entity.Lesson)(nil), gorm.ErrRecordNotFound).
		Once()

	err := lessonService.Delete(999)

	assert.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)

	repo.AssertNotCalled(t, "Delete", mock.Anything)
	repo.AssertExpectations(t)
}
