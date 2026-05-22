package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lms-main-service/internal/apperror"
	"lms-main-service/internal/dto"
	"lms-main-service/internal/entity"
	"lms-main-service/internal/response"
	"lms-main-service/internal/service"
	"net/http"
)

type ChapterHandler struct {
	service service.ChapterService
	logger  *logrus.Logger
}

func NewChapterHandler(service service.ChapterService, logger *logrus.Logger) *ChapterHandler {
	return &ChapterHandler{
		service: service,
		logger:  logger,
	}
}

// CreateChapter godoc
// @Summary Create a new chapter
// @Description Create a new chapter for a course
// @Tags chapters
// @Accept json
// @Produce json
// @Param chapter body dto.CreateChapterRequest true "Chapter payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /chapters [post]
func (h *ChapterHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateChapterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, apperror.ErrInvalidRequestBody)
		return
	}

	chapter := entity.Chapter{
		Name:        req.Name,
		Description: req.Description,
		Order:       req.Order,
		CourseID:    req.CourseID,
	}

	if err := h.service.Create(ctx, &chapter); err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	h.logger.WithFields(logrus.Fields{
		"chapter_id": chapter.ID,
		"name":       chapter.Name,
		"course_id":  chapter.CourseID,
	}).Info("chapter created")

	response.Success(c, http.StatusCreated, "chapter created successfully", chapter)
}

// GetAllChapters godoc
// @Summary Get all chapters
// @Description Get paginated list of all chapters
// @Tags chapters
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /chapters [get]
func (h *ChapterHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset := parsePagination(c)

	chapters, err := h.service.GetAll(ctx, limit, offset)
	if err != nil {
		response.ErrorMessage(c, http.StatusInternalServerError, "failed to get chapters")
		return
	}

	response.Success(c, http.StatusOK, "chapters fetched successfully", chapters)
}

// GetChapterByID godoc
// @Summary Get chapter by ID
// @Description Get chapter details by ID
// @Tags chapters
// @Produce json
// @Param id path int true "Chapter ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /chapters/{id} [get]
func (h *ChapterHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := parseID(c.Param("id"))
	if err != nil {
		response.ErrorMessage(c, http.StatusBadRequest, "invalid chapter id")
		return
	}

	chapter, err := h.service.GetByID(ctx, id)
	if err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	response.Success(c, http.StatusOK, "chapter fetched successfully", chapter)
}

// UpdateChapter godoc
// @Summary Update chapter
// @Description Update chapter by ID
// @Tags chapters
// @Accept json
// @Produce json
// @Param id path int true "Chapter ID"
// @Param chapter body dto.UpdateChapterRequest true "Chapter payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /chapters/{id} [put]
func (h *ChapterHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := parseID(c.Param("id"))
	if err != nil {
		response.ErrorMessage(c, http.StatusBadRequest, "invalid chapter id")
		return
	}

	var req dto.UpdateChapterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, apperror.ErrInvalidRequestBody)
		return
	}

	chapter := entity.Chapter{
		Name:        req.Name,
		Description: req.Description,
		Order:       req.Order,
		CourseID:    req.CourseID,
	}

	if err := h.service.Update(ctx, id, &chapter); err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	h.logger.WithFields(logrus.Fields{
		"chapter_id": id,
		"name":       chapter.Name,
		"course_id":  chapter.CourseID,
	}).Info("chapter updated")

	response.Success(c, http.StatusOK, "chapter updated successfully", chapter)
}

// DeleteChapter godoc
// @Summary Delete chapter
// @Description Delete chapter by ID
// @Tags chapters
// @Produce json
// @Param id path int true "Chapter ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /chapters/{id} [delete]
func (h *ChapterHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := parseID(c.Param("id"))
	if err != nil {
		response.ErrorMessage(c, http.StatusBadRequest, "invalid chapter id")
		return
	}

	if err := h.service.Delete(ctx, id); err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	h.logger.WithField("chapter_id", id).Info("chapter deleted")

	response.SuccessMessage(c, http.StatusOK, "chapter deleted successfully")
}
