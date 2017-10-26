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
	router.POST("/notes", notes.InsertNoteInfo)
	router.Run()
	fmt.Println(hello.BuildHello())
}
