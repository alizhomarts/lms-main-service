package dto

type (
	CreateLessonRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Content     string `json:"content" binding:"required"`
		Order       int    `json:"order" binding:"required"`
		ChapterID   uint   `json:"chapter_id" binding:"required"`
	}

	UpdateLessonRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Content     string `json:"content" binding:"required"`
		Order       int    `json:"order" binding:"required"`
		ChapterID   uint   `json:"chapter_id" binding:"required"`
	}
)
