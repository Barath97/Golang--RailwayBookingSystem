package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.com/railwayticket/models"
	"go.com/railwayticket/services"
)

func GetAllTickets(c *gin.Context) {

	passengers, err := services.GetAllTickets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't get all tickets"})
		return
	}
	c.JSON(http.StatusOK, passengers)
}

func BookTicket(c *gin.Context) {
	var passenger models.Passenger
	if err := c.BindJSON(&passenger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid JSON provided"})
		return
	}

	success, err := services.BookTicket(&passenger)
	if !success || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to book ticket"})
		return
	}

	c.JSON(http.StatusOK, passenger)
}

func TicketDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	passenger, err := services.GetTicketDetails(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting passenger details"})
		return
	}

	c.JSON(http.StatusOK, passenger)
}
