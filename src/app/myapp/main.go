package main

import (
	"app/myapp/hello"
	"app/myapp/notes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", notes.GetNoteInfo)
	router.POST("/main", notes.InsertNoteMain)
	router.POST("/notes", notes.InsertNoteInfo)
	router.GET("/notes", notes.GetNoteInfo)
	router.PUT("/notes", notes.PutNotes)
	router.POST("/notes/archived", notes.ArchivedNote)
	router.Run(":5566")
	fmt.Println(hello.BuildHello())
}
