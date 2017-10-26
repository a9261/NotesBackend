package notes

import (
	"app/myapp/repository/noteRepo"
	"log"

	"github.com/gin-gonic/gin"
)

//GetNoteInfo is Get Note Info from Notes
func GetNoteInfo(c *gin.Context) {
	c.String(200, "Success")
}

//InsertNoteInfo is Insert Note Info
func InsertNoteInfo(c *gin.Context) {
	log.Println("IN")
	var result noteRepo.Note
	if c.BindJSON(&result) == nil {
		log.Println("OK")
		log.Println(result)
		noteRepo.InsertNote(result)
	}
	//c.String(200, "Success")
}
