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

func TestCourseService_Create_Success(t *testing.T) {
	repo := new(mocks.CourseRepository)
	courseService := NewCourseService(repo)

	course := &entity.Course{
		Name:        "Golang Developer",
		Description: "Backend development course with Go",
	}

	repo.
		On("Create", course).
		Return(nil).
		Once()

	err := courseService.Create(course)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCourseService_Create_NameRequired(t *testing.T) {
	repo := new(mocks.CourseRepository)
	courseService := NewCourseService(repo)

	course := &entity.Course{
		Name:        "",
		Description: "Course without name",
	}

	err := courseService.Create(course)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperror.ErrCourseNameRequired))

	repo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestCourseService_GetAll_Success(t *testing.T) {
	repo := new(mocks.CourseRepository)
	courseService := NewCourseService(repo)

	expectedCourses := []entity.Course{
		{
			ID:          1,
			Name:        "Golang Developer",
			Description: "Go backend course",
		},
		{
			ID:          2,
			Name:        "Python Developer",
			Description: "Python backend course",
		},
	}

	repo.
		On("GetAll").
		Return(expectedCourses, nil).
		Once()

	courses, err := courseService.GetAll()

	assert.NoError(t, err)
	assert.Len(t, courses, 2)
	assert.Equal(t, "Golang Developer", courses[0].Name)
	assert.Equal(t, "Python Developer", courses[1].Name)

	repo.AssertExpectations(t)
}

func TestCourseService_GetByID_Success(t *testing.T) {
	repo := new(mocks.CourseRepository)
	courseService := NewCourseService(repo)

	expectedCourse := &entity.Course{
		ID:          1,
		Name:        "Golang Developer",
		Description: "Go backend course",
	}

	repo.
		On("GetByID", uint(1)).
		Return(expectedCourse, nil).
		Once()

	course, err := courseService.GetByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, course)
	assert.Equal(t, uint(1), course.ID)
	assert.Equal(t, "Golang Developer", course.Name)

	repo.AssertExpectations(t)
}

func TestCourseService_GetByID_NotFound(t *testing.T) {
	repo := new(mocks.CourseRepository)
	courseService := NewCourseService(repo)

	repo.
		On("GetByID", uint(999)).
		Return((*entity.Course)(nil), gorm.ErrRecordNotFound).
		Once()

	course, err := courseService.GetByID(999)

	assert.Error(t, err)
	assert.Nil(t, course)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)

	repo.AssertExpectations(t)
}

func TestCourseService_Update_Success(t *testing.T) {
	repo := new(mocks.CourseRepository)
	courseService := NewCourseService(repo)

	existingCourse := &entity.Course{
		ID:          1,
		Name:        "Golang Developer",
		Description: "Old description",
	}

	updateData := &entity.Course{
		Name:        "Advanced Golang Developer",
		Description: "Updated description",
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingCourse, nil).
		Once()

	repo.
		On("Update", mock.MatchedBy(func(course *entity.Course) bool {
			return course.ID == 1 &&
				course.Name == "Advanced Golang Developer" &&
				course.Description == "Updated description"
		})).
		Return(nil).
		Once()

	err := courseService.Update(1, updateData)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCourseService_Update_NameRequired(t *testing.T) {
	repo := new(mocks.CourseRepository)
	courseService := NewCourseService(repo)

	existingCourse := &entity.Course{
		ID:          1,
		Name:        "Golang Developer",
		Description: "Old description",
	}

	updateData := &entity.Course{
		Name:        "",
		Description: "Updated description",
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingCourse, nil).
		Once()

	err := courseService.Update(1, updateData)

	assert.Error(t, err)
	assert.ErrorIs(t, err, apperror.ErrCourseNameRequired)

	repo.AssertNotCalled(t, "Update", mock.Anything)
	repo.AssertExpectations(t)
}

func TestCourseService_Delete_Success(t *testing.T) {
	repo := new(mocks.CourseRepository)
	courseService := NewCourseService(repo)

	existingCourse := &entity.Course{
		ID:          1,
		Name:        "Golang Developer",
		Description: "Go backend course",
	}

	repo.
		On("GetByID", uint(1)).
		Return(existingCourse, nil).
		Once()

	repo.
		On("Delete", uint(1)).
		Return(nil).
		Once()

	err := courseService.Delete(1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCourseService_Delete_NotFound(t *testing.T) {
	repo := new(mocks.CourseRepository)
	courseService := NewCourseService(repo)

	repo.
		On("GetByID", uint(999)).
		Return((*entity.Course)(nil), gorm.ErrRecordNotFound).
		Once()

	err := courseService.Delete(999)

	assert.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)

	repo.AssertNotCalled(t, "Delete", mock.Anything)
	repo.AssertExpectations(t)
}
