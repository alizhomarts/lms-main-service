package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"lms-main-service/internal/config"
	"lms-main-service/internal/database"
	"lms-main-service/internal/http/handler"
	"lms-main-service/internal/http/router"
	"lms-main-service/internal/repository"
	"lms-main-service/internal/service"
)

// @title LMS Main Service API
// @version 1.0
// @description Main Service for managing courses, chapters, and lessons in LMS System.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := database.NewPostgresDB(cfg.DB)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Successfully connected to Database")

	courseRepo := repository.NewCourseRepository(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	chapterRepo := repository.NewChapterRepository(db)
	chapterService := service.NewChapterService(chapterRepo)
	chapterHandler := handler.NewChapterHandler(chapterService)

	lessonRepo := repository.NewLessonRepository(db)
	lessonService := service.NewLessonService(lessonRepo)
	lessonHandler := handler.NewLessonHandler(lessonService)

	appRouter := router.NewRouter(courseHandler, chapterHandler, lessonHandler)
	appRouter.SetupRoutes()

	addr := fmt.Sprintf(":%s", cfg.AppPort)

	logrus.Infof("LMS Main Service started on port %s", cfg.AppPort)

	if err := appRouter.Run(addr); err != nil {
		logrus.Fatal(err)
	}
}
