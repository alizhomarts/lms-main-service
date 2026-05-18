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
}

func NewLessonHandler(service service.LessonService) *LessonHandler {
	return &LessonHandler{service: service}
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

	if err := h.service.Create(&lesson); err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	logrus.WithFields(logrus.Fields{
		"lesson_id":  lesson.ID,
		"name":       lesson.Name,
		"chapter_id": lesson.ChapterID,
	}).Info("lesson created")

	response.Success(c, http.StatusCreated, "lesson created successfully", lesson)
}

// GetAllLessons godoc
// @Summary Get all lessons
// @Description Get list of all lessons
// @Tags lessons
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /lessons [get]
func (h *LessonHandler) GetAll(c *gin.Context) {
	lessons, err := h.service.GetAll()
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
	id, err := parseID(c.Param("id"))
	if err != nil {
		response.ErrorMessage(c, http.StatusBadRequest, "invalid lesson id")
		return
	}

	lesson, err := h.service.GetByID(id)
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

	if err := h.service.Update(id, &lesson); err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	logrus.WithFields(logrus.Fields{
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
	id, err := parseID(c.Param("id"))
	if err != nil {
		response.ErrorMessage(c, http.StatusBadRequest, "invalid lesson id")
		return
	}

	if err := h.service.Delete(id); err != nil {
		status := apperror.StatusCode(err)

		response.Error(c, status, err)
		return
	}

	logrus.WithField("lesson_id", id).Info("lesson deleted")

	response.SuccessMessage(c, http.StatusOK, "lesson deleted successfully")
}
