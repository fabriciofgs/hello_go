package main 

import (
	"log" //Used to "print" errors
	"net/http" //Used to simplifly constants for http.StatusOK, http.StatusError, etc...
	"fmt" //Only for the ASCIIART

	"github.com/gin-gonic/gin"
)

func main(){
	// Create a Gin router with default middleware (logger and recovery)
	var router *gin.Engine = gin.Default()
  // Example of a condensed endpoint. Only for demonstration, check the other example that start on line 22
	router.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

  // router is the variable created earlier
	// .GET is the type of http call we are creating
	// "/hello" is the endpoint that we will call
	// helloFunction calls the function on the lines below
	/* Basically we are setting up a GET http request on the /hello endpoint
	that when called, will return run the helloFunction. */ 
	router.GET("/hello", helloFunction)

	// var err error is to catch any errors that the api can return.
	// router.Run() is the command that will start the rest api on the background
	printCoolASCIIArt()
	var err error = router.Run()
	// If the Server return any variable of the error type,
	// it will be catched by this if statement and log an error on the CMD
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

/* The gin module will be running on the background and it exposes a pointer
gin.Context, this pointer give us access to the running api to create responses */
func helloFunction (context *gin.Context) {
	// I am building a response, this can be one variable or a entire struct with slices and multiple variables.
	var myReturn string = "Hello World!"
	// I am building a response to the function, it uses the gin.H. Basically a map[string]any type.
	var jsonResponse gin.H = gin.H{"message": myReturn,}
  /* Using the context variable i call the JSON method to trigger a response on the api.
	This response have 2 parameters. First a http status like OK, Error, etc.
	Second is the response in Json built before */
	context.JSON(http.StatusOK, jsonResponse)
}

func printCoolASCIIArt () {
	fmt.Println(`
  /$$$$$$  /$$$$$$$  /$$$$$$        /$$$$$$   /$$                           /$$                     /$$
 /$$__  $$| $$__  $$|_  $$_/       /$$__  $$ | $$                          | $$                    | $$
| $$  \ $$| $$  \ $$  | $$        | $$  \__//$$$$$$    /$$$$$$   /$$$$$$  /$$$$$$    /$$$$$$   /$$$$$$$
| $$$$$$$$| $$$$$$$/  | $$        |  $$$$$$|_  $$_/   |____  $$ /$$__  $$|_  $$_/   /$$__  $$ /$$__  $$
| $$__  $$| $$____/   | $$         \____  $$ | $$      /$$$$$$$| $$  \__/  | $$    | $$$$$$$$| $$  | $$
| $$  | $$| $$        | $$         /$$  \ $$ | $$ /$$ /$$__  $$| $$        | $$ /$$| $$_____/| $$  | $$
| $$  | $$| $$       /$$$$$$      |  $$$$$$/ |  $$$$/|  $$$$$$$| $$        |  $$$$/|  $$$$$$$|  $$$$$$$
|__/  |__/|__/      |______/       \______/   \___/   \_______/|__/         \___/   \_______/ \_______/
	`)
}
