package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pelicanch1k/ProductGatewayAPI/structs"
	"net/http"
)

// @Summary Create product
// @Security ApiKeyAuth
// @Tags products
// @Description create product
// @ID create-product
// @Accept  json
// @Produce  json
// @Param input body structs.Product true "product info"
// @Success 200 {integer} id 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /products [post]
func (h *Handler) create(c *gin.Context) {
	var input structs.Product

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	productId, err := h.services.Products.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": productId,
	})
}

// @Summary Update product
// @Security ApiKeyAuth
// @Tags products
// @Description update product
// @ID update-list
// @Accept  json
// @Produce  json
// @Param id query string true "Product ID"
// @Param input body structs.Product true "product info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /products/ [put]
func (h *Handler) update(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input structs.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = h.services.Products.Update(input, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete product
// @Security ApiKeyAuth
// @Tags products
// @Description delete product
// @ID delete-list
// @Accept  json
// @Produce  json
// @Param  id query string true "Product ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /products/ [delete]
func (h *Handler) delete(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.Products.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Get products by ID
// @Security ApiKeyAuth
// @Tags products
// @Description get products by ID
// @ID getById
// @Accept  json
// @Produce  json
// @Param  id query string true "Product ID"
// @Success 200 {object} structs.Product
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /products/ [get]
func (h *Handler) getById(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	product, err := h.services.Products.Get(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary Get all products
// @Security ApiKeyAuth
// @Tags products
// @Description get all products
// @ID get-all
// @Accept  json
// @Produce  json
// @Success 200 {object} []structs.Product
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /products [get]
func (h *Handler) getAll(c *gin.Context) {
	products, err := h.services.Products.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}
