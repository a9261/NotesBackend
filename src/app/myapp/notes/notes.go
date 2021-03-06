package notes

import (
	"app/myapp/repository/noteRepo"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var note = new(noteRepo.NoteRepository)

type QueryModel struct {
	Key string `form:"key" json"key"`
}
type ArchivedModel struct {
	Key string `form:"key" json"key"`
	ID  int    `form:"id" json"id"`
}

func GetArchivedNote(c *gin.Context) {
	var model QueryModel
	if c.Bind(&model) == nil {
		result, err := note.GetArchivedNote(model.Key)
		if err == nil {
			c.JSON(200, result)
		} else {
			c.JSON(500, "GetArchivedNote is wrong")
		}
	} else {
		c.JSON(500, "Binding querystring is wrong")
	}
}
func ArchivedNote(c *gin.Context) {
	var noteItem ArchivedModel
	if c.BindJSON(&noteItem) == nil {
		log.Println("OK")
		log.Println(noteItem)
		note.ArchivedNote(noteItem.Key, noteItem.ID)
		c.JSON(200, "OK")
	} else {
		c.JSON(500, "archived fail")
	}

}

func PutNotes(c *gin.Context) {
	var result []noteRepo.NoteModel
	if c.BindJSON(&result) == nil {
		log.Println("OK")
		log.Println(result)
		note.PutNotes(result)
		c.JSON(200, "Update Ok")
	} else {
		c.JSON(500, "Update Fail")
	}
}

//Pong is response Ping
func Pong(c *gin.Context) {
	c.JSON(200, "Hi")
}

//GetNoteInfo is Get Note Info from Notes
func GetNoteInfo(c *gin.Context) {
	var model QueryModel
	var result []noteRepo.NoteModel
	if c.Bind(&model) == nil {
		fmt.Println("model.Key is ")
		fmt.Println(model.Key)
		result = note.GetNotes(model.Key)
	} else {
		result = note.GetNotes("axsfds")
	}
	c.JSON(200, result)
}

//GetNoteMain is Get NoteMain Info
func GetNoteMain(c *gin.Context) {
	log.Println("IN")
	var model QueryModel
	var result noteRepo.NoteMainModel
	if c.Bind(&model) == nil {
		log.Println("OK")
		log.Println(result)
		c.JSON(200, note.GetNoteMain(model.Key))
	} else {
		c.String(500, "GetNoteMain fail")
	}
}

//Insert NoteMain Info
func InsertNoteMain(c *gin.Context) {
	log.Println("IN")
	var result noteRepo.NoteMainModel
	if c.BindJSON(&result) == nil {
		log.Println("OK")
		log.Println(result)
		c.JSON(200, note.InsertNoteMain(result))
	} else {
		c.String(500, "Insert fail")
	}
}

//InsertNoteInfo is Insert Note Info
func InsertNoteInfo(c *gin.Context) {
	log.Println("IN InsertNoteInfo")
	var result noteRepo.NoteModel
	if c.BindJSON(&result) == nil {
		log.Println("OK")
		log.Println(result)
		note.InsertNote(result)
		c.String(200, "Success")
	} else {
		log.Println(result)
		c.String(500, "InsertNoteInfo fail")
	}
}
