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

type LessonHandler struct {
	service service.LessonService
	logger  *logrus.Logger
}

func NewLessonHandler(service service.LessonService, logger *logrus.Logger) *LessonHandler {
	return &LessonHandler{
		service: service,
		logger:  logger,
	}
}

// CreateLesson godoc
// @Summary Create a new lesson
// @Description Create a new lesson for a chapter
// @Tags lessons
// @Accept json
// @Produce json
// @Param lesson body dto.CreateLessonRequest true "Lesson payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /lessons [post]
func (h *LessonHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateLessonRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, apperror.ErrInvalidRequestBody)
		return
	}

	lesson := entity.Lesson{
		Name:        req.Name,
		Description: req.Description,
		Content:     req.Content,
		Order:       req.Order,
		ChapterID:   req.ChapterID,
	}

	if err := h.service.Create(ctx, &lesson); err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	h.logger.WithFields(logrus.Fields{
		"lesson_id":  lesson.ID,
		"name":       lesson.Name,
		"chapter_id": lesson.ChapterID,
	}).Info("lesson created")

	response.Success(c, http.StatusCreated, "lesson created successfully", lesson)
}

// GetAllLessons godoc
// @Summary Get all lessons
// @Description Get paginated list of all lessons
// @Tags lessons
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /lessons [get]
func (h *LessonHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset := parsePagination(c)

	lessons, err := h.service.GetAll(ctx, limit, offset)
	if err != nil {
		response.ErrorMessage(c, http.StatusInternalServerError, "failed to get lessons")
		return
	}

	response.Success(c, http.StatusOK, "lessons fetched successfully", lessons)
}

// GetLessonByID godoc
// @Summary Get lesson by ID
// @Description Get lesson details by ID
// @Tags lessons
// @Produce json
// @Param id path int true "Lesson ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /lessons/{id} [get]
func (h *LessonHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := parseID(c.Param("id"))
	if err != nil {
		response.ErrorMessage(c, http.StatusBadRequest, "invalid lesson id")
		return
	}

	lesson, err := h.service.GetByID(ctx, id)
	if err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	response.Success(c, http.StatusOK, "lesson fetched successfully", lesson)
}

// UpdateLesson godoc
// @Summary Update lesson
// @Description Update lesson by ID
// @Tags lessons
// @Accept json
// @Produce json
// @Param id path int true "Lesson ID"
// @Param lesson body dto.UpdateLessonRequest true "Lesson payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /lessons/{id} [put]
func (h *LessonHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := parseID(c.Param("id"))
	if err != nil {
		response.ErrorMessage(c, http.StatusBadRequest, "invalid lesson id")
		return
	}

	var req dto.UpdateLessonRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, apperror.ErrInvalidRequestBody)
		return
	}

	lesson := entity.Lesson{
		Name:        req.Name,
		Description: req.Description,
		Content:     req.Content,
		Order:       req.Order,
		ChapterID:   req.ChapterID,
	}

	if err := h.service.Update(ctx, id, &lesson); err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	h.logger.WithFields(logrus.Fields{
		"lesson_id":  id,
		"name":       lesson.Name,
		"chapter_id": lesson.ChapterID,
	}).Info("lesson updated")

	response.Success(c, http.StatusOK, "lesson updated successfully", lesson)
}

// DeleteLesson godoc
// @Summary Delete lesson
// @Description Delete lesson by ID
// @Tags lessons
// @Produce json
// @Param id path int true "Lesson ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /lessons/{id} [delete]
func (h *LessonHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := parseID(c.Param("id"))
	if err != nil {
		response.ErrorMessage(c, http.StatusBadRequest, "invalid lesson id")
		return
	}

	if err := h.service.Delete(ctx, id); err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	h.logger.WithField("lesson_id", id).Info("lesson deleted")

	response.SuccessMessage(c, http.StatusOK, "lesson deleted successfully")
}
