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
	"strconv"
)

type CourseHandler struct {
	service service.CourseService
}

func NewCourseHandler(service service.CourseService) *CourseHandler {
	return &CourseHandler{service: service}
}

// CreateCourse godoc
// @Summary Create a new course
// @Description Create a new course in LMS
// @Tags courses
// @Accept json
// @Produce json
// @Param course body dto.CreateCourseRequest true "Course payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /courses [post]
func (h *CourseHandler) Create(c *gin.Context) {
	var req dto.CreateCourseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, apperror.ErrInvalidRequestBody)
		return
	}

	course := entity.Course{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.service.Create(&course); err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	logrus.WithFields(logrus.Fields{
		"course_id": course.ID,
		"name":      course.Name,
	}).Info("course created")

	response.Success(c, http.StatusCreated, "course created successfully", course)
}

// GetAllCourses godoc
// @Summary Get all courses
// @Description Get list of all courses
// @Tags courses
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /courses [get]
func (h *CourseHandler) GetAll(c *gin.Context) {
	courses, err := h.service.GetAll()
	if err != nil {
		response.ErrorMessage(c, http.StatusInternalServerError, "failed to get courses")
		return
	}

	response.Success(c, http.StatusOK, "courses fetched successfully", courses)
}

// GetCourseByID godoc
// @Summary Get course by ID
// @Description Get course details by ID
// @Tags courses
// @Produce json
// @Param id path int true "Course ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /courses/{id} [get]
func (h *CourseHandler) GetByID(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		response.ErrorMessage(c, http.StatusBadRequest, "invalid course id")
		return
	}

	course, err := h.service.GetByID(id)
	if err != nil {
		status := apperror.StatusCode(err)

		response.Error(c, status, err)
		return
	}

	response.Success(c, http.StatusOK, "course fetched successfully", course)
}

// UpdateCourse godoc
// @Summary Update course
// @Description Update course by ID
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "Course ID"
// @Param course body dto.UpdateCourseRequest true "Course payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /courses/{id} [put]
func (h *CourseHandler) Update(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		response.ErrorMessage(c, http.StatusBadRequest, "invalid course id")
		return
	}

	var req dto.UpdateCourseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, apperror.ErrInvalidRequestBody)
		return
	}

	course := entity.Course{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.service.Update(id, &course); err != nil {
		status := apperror.StatusCode(err)
		response.Error(c, status, err)
		return
	}

	logrus.WithFields(logrus.Fields{
		"course_id": id,
		"name":      course.Name,
	}).Info("course updated")

	response.Success(c, http.StatusOK, "course updated successfully", course)
}

// DeleteCourse godoc
// @Summary Delete course
// @Description Delete course by ID
// @Tags courses
// @Produce json
// @Param id path int true "Course ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /courses/{id} [delete]
func (h *CourseHandler) Delete(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		response.ErrorMessage(c, http.StatusBadRequest, "invalid course id")
		return
	}

	if err := h.service.Delete(id); err != nil {
		status := apperror.StatusCode(err)

		response.Error(c, status, err)
		return
	}

	logrus.WithField("course_id", id).Info("course deleted")

	response.SuccessMessage(c, http.StatusOK, "course deleted successfully")
}

func parseID(idParam string) (uint, error) {
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
