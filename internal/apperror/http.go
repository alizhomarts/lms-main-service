package apperror

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
)

func StatusCode(err error) int {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return http.StatusNotFound

	case errors.Is(err, ErrCourseNotFound):
		return http.StatusNotFound

	case errors.Is(err, ErrChapterNotFound):
		return http.StatusNotFound

	case errors.Is(err, ErrLessonNotFound):
		return http.StatusNotFound

	case errors.Is(err, ErrCourseNameRequired),
		errors.Is(err, ErrChapterNameRequired),
		errors.Is(err, ErrChapterOrderRequired),
		errors.Is(err, ErrChapterIDRequired),
		errors.Is(err, ErrLessonNameRequired),
		errors.Is(err, ErrLessonContentRequired),
		errors.Is(err, ErrLessonOrderRequired),
		errors.Is(err, ErrCourseIDRequired):
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}
}
