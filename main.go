package main

import (
	"fmt"
	"os"
	"time"

	"github.com/iris-contrib/graceful"
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

func main() {
	web := iris.New()

	web.Use(logger.New())
	web.Use(recovery.Handler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "ACCEPT", "ORIGIN"},
		AllowCredentials: true,
		Debug:            true,
	})

	web.Use(c)
	web.UseTemplate(html.New()).Directory("./app/views", ".html")

	web.Get("/", indexHandler)
	web.Static("/assets", "./app/assets", 1)

	//Get port from environment variables, default port number is 7000
	port := os.Getenv("PORT")

	if port == "" {
		port = "7000"
	}

	fmt.Println("Service is listening at:" + port)
	graceful.Run(":"+port, time.Duration(10)*time.Second, web)
}

func indexHandler(ctx *iris.Context) {
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
	if err := ctx.Render("index.html", nil); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
