package dto

type (
	CreateChapterRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Order       int    `json:"order" binding:"required"`
		CourseID    uint   `json:"course_id" binding:"required"`
	}

	UpdateChapterRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Order       int    `json:"order" binding:"required"`
		CourseID    uint   `json:"course_id" binding:"required"`
	}
)
