package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "lms-main-service/docs"
	"lms-main-service/internal/http/handler"
)

type Router struct {
	engine         *gin.Engine
	courseHandler  *handler.CourseHandler
	chapterHandler *handler.ChapterHandler
	lessonHandler  *handler.LessonHandler
}

func NewRouter(
	courseHandler *handler.CourseHandler,
	chapterHandler *handler.ChapterHandler,
	lessonHandler *handler.LessonHandler,
) *Router {
	r := gin.Default()

	if err := r.SetTrustedProxies(nil); err != nil {
		logrus.Fatal(err)
	}

	return &Router{
		engine:         r,
		courseHandler:  courseHandler,
		chapterHandler: chapterHandler,
		lessonHandler:  lessonHandler,
	}
}

func (r *Router) SetupRoutes() {
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/swagger/doc.json"),
	))

	r.engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	api := r.engine.Group("/api/v1")
	{
		courses := api.Group("/courses")
		{
			courses.POST("", r.courseHandler.Create)
			courses.GET("", r.courseHandler.GetAll)
			courses.GET("/:id", r.courseHandler.GetByID)
			courses.PUT("/:id", r.courseHandler.Update)
			courses.DELETE("/:id", r.courseHandler.Delete)
		}

		chapters := api.Group("/chapters")
		{
			chapters.POST("", r.chapterHandler.Create)
			chapters.GET("", r.chapterHandler.GetAll)
			chapters.GET("/:id", r.chapterHandler.GetByID)
			chapters.PUT("/:id", r.chapterHandler.Update)
			chapters.DELETE("/:id", r.chapterHandler.Delete)
		}

		lessons := api.Group("/lessons")
		{
			lessons.POST("", r.lessonHandler.Create)
			lessons.GET("", r.lessonHandler.GetAll)
			lessons.GET("/:id", r.lessonHandler.GetByID)
			lessons.PUT("/:id", r.lessonHandler.Update)
			lessons.DELETE("/:id", r.lessonHandler.Delete)
		}
	}
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}
