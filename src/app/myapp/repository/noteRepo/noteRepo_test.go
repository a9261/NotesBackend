package noteRepo

import (
	"app/myapp/repository/noteRepo"
	"testing"
)

var note = new(noteRepo.NoteRepository)

func TestGetNotes(t *testing.T) {
	excepte := []Note{}
	result := note.GetNotes()
}
