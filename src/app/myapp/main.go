package main

import (
	"app/myapp/hello"
	"app/myapp/notes"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	//https://github.com/gin-contrib/cors
	router.Use(cors.Default()) //Enable all origins
	router.GET("/ping", notes.GetNoteInfo)
	router.POST("/main", notes.InsertNoteMain)
	router.GET("/main", notes.GetNoteMain)

	router.POST("/notes", notes.InsertNoteInfo)
	router.GET("/notes", notes.GetNoteInfo)
	router.PUT("/notes", notes.PutNotes)

	router.POST("/notes/archived", notes.ArchivedNote)

	router.Run(":5566")
	fmt.Println(hello.BuildHello())
}
