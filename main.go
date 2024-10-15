package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.com/railwayticket/controllers"
	"go.com/railwayticket/database"
)

func main() {

	//Load environment variables from .env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	//Set up the database connection
	database.SetUpDatabase()

	//create a Gin router
	router := gin.Default()

	//Define the routes
	router.POST("/bookticket", controllers.BookTicket)
	router.GET("/alltickets", controllers.GetAllTickets)
	router.GET("/ticketdetails", controllers.TicketDetails)

	//start the server
	router.Run(":8080")
}
