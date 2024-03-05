package handler

import (
	"github.com/geejjoo/lamoda_task/cmd/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) storageInformation(c *gin.Context) {
	var productList []db.Product
	storage := c.Param("storage")
	productList, err := h.services.StorageInformation(storage)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, productList)
}
