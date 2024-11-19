package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("id")
	if !ok {
		return 0, errors.New("id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("id is of invalid type")
	}

	return idInt, nil
}
