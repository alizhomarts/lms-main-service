package dto

type (
	CreateCourseRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	UpdateCourseRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}
)
