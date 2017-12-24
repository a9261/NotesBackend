package noteRepo

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var note = new(NoteRepository)
var testKey = "ax95"

func init() {
	fmt.Println("--noteRepo Testing --")
}
func createInitData() (mockMain *NoteMainModel, mockNote *NoteModel) {
	main := &NoteMainModel{"Test", testKey}
	noteItem := &NoteModel{
		NoteColor:     "red",
		NoteContent:   "helllo",
		NotePositionX: 55.66,
		NotePositionY: 123.33,
		NoteKey:       testKey,
		IsArchived:    0,
	}
	note.InsertNoteMain(*main)
	note.InsertNote(*noteItem)
	return main, noteItem
}
func dropInitData() {
	db, err := sql.Open("mysql", conStr)
	defer db.Close()
	checkErr(err)
	_, err = db.Exec(`
			DELETE FROM NotesList
			WHERE FK_NoteKey=?
			`, "ax95")
	_, err = db.Exec(`
				DELETE FROM NoteMain
				WHERE NoteKey=?
				`, "ax95")
	checkErr(err)
}
func TestGetNotes(t *testing.T) {
	_, notes := createInitData()
	var excepted []NoteModel
	notes.ID = -1
	excepted = append(excepted, *notes)
	defer dropInitData()
	result := note.GetNotes(testKey)
	result[0].ID = -1
	assert.Equal(t, excepted, result, "they should be equal")
}
