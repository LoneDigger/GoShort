package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func init() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

//go:embed  public/*
var Fs embed.FS

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.NoRoute(noRoute)

	// router.StaticFS("/", http.FS(Fs))
	router.StaticFileFS("/public/css.css", "/public/css.css", http.FS(Fs))
	router.StaticFileFS("/public/js.js", "/public/js.js", http.FS(Fs))
	router.StaticFileFS("/public/index.html", "/public/index.html", http.FS(Fs))

	router.GET("/", index)
	router.POST("/", createUrl)
	router.GET("/:key", redirect)

	if err := router.Run(":8080"); err != nil {
		log.Panicf("error: %s", err)
	}
}
