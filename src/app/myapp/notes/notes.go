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

func ArchivedNote(c *gin.Context) {
	var noteItem noteRepo.NoteModel
	if c.BindJSON(&note) == nil {
		log.Println("OK")
		log.Panicln(noteItem)
		note.ArchivedNote(noteItem)
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
	}
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
	log.Println("IN")
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
