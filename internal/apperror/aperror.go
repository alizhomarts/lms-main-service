package apperror

import "errors"

var (
	ErrInvalidRequestBody = errors.New("invalid request body")

	ErrCourseNameRequired = errors.New("course name is required")
	ErrCourseNotFound     = errors.New("course not found")

	ErrChapterNameRequired  = errors.New("chapter name is required")
	ErrChapterOrderRequired = errors.New("chapter order must be greater than zero")
	ErrChapterIDRequired    = errors.New("chapter_id is required")
	ErrChapterNotFound      = errors.New("chapter not found")

	ErrLessonNameRequired    = errors.New("lesson name is required")
	ErrLessonContentRequired = errors.New("lesson content is required")
	ErrLessonOrderRequired   = errors.New("lesson order must be greater than zero")
	ErrLessonNotFound        = errors.New("lesson not found")

	ErrCourseIDRequired = errors.New("course_id is required")
)
