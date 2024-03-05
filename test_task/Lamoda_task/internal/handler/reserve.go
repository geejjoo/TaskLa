package handler

import (
	"github.com/geejjoo/lamoda_task/cmd/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) makeReservation(c *gin.Context) {
	var input db.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest,
			"Bad request")
		return
	}

	Reservation := db.Product{
		Articular: input.Articular,
		Quantity:  input.Quantity,
	}

	err := h.services.MakeReservation(&Reservation)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) cancelReservation(c *gin.Context) {
	var input db.CancelReservation
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest,
			"Bad request")
		return
	}

	CancelReservation := db.CancelReservation{
		Articular: input.Articular,
		Quantity:  input.Quantity,
		Storage:   input.Storage,
	}

	err := h.services.CancelReservation(&CancelReservation)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
