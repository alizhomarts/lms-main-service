package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	defaultLimit  = 10
	maxLimit      = 100
	defaultOffset = 0
)

func parsePagination(c *gin.Context) (int, int) {
	limit := defaultLimit
	offset := defaultOffset

	if limitParam := c.Query("limit"); limitParam != "" {
		if parsedLimit, err := strconv.Atoi(limitParam); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	if offsetParam := c.Query("offset"); offsetParam != "" {
		if parsedOffset, err := strconv.Atoi(offsetParam); err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	return limit, offset
}
