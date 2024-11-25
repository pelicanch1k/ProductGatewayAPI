package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pelicanch1k/ProductGatewayAPI/structs"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user structs.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Auth.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})

}

func (h *Handler) signIn(c *gin.Context) {
	var user structs.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Auth.GenerateJWT(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
