package noteRepo

import (
	"database/sql"
	"fmt"

	mysql "github.com/go-sql-driver/mysql"
)

// Note is Note data struct
type Note struct {
	NoteColor     int     `json:"NoteColor"`
	NoteContent   string  `json:"Notecontent"`
	NotePositionX float32 `json:"NotePositionX"`
	NotePositionY float32 `json:"NotePositionY"`
	FK_NoteKey    string  `json:"FK_NoteKey"`
}

// NoteRepository is Public Struct for Use
type NoteRepository struct {
}

var config mysql.Config

func init() {
	config = mysql.Config{
		User:   "root",
		Passwd: "nfu123!@#",
		Addr:   "192.168.2.13",
		DBName: "Notes",
	}
}

// GetNotes will return mutiple Note data
func (noteRep *NoteRepository) GetNotes() []Note {
	db, err := sql.Open("mysql", config.FormatDSN())
	checkErr(err)
	rows, err := db.Query("SELECT NoteColor,NoteContent,NotePositionX,NotePositionY FROM NotesList")
	checkErr(err)
	defer rows.Close()
	var notes []Note
	var (
		NoteColor     int
		NoteContent   string
		NotePositionX float32
		NotePositionY float32
	)
	for rows.Next() {
		err := rows.Scan(&NoteColor, &NoteContent, &NotePositionX, &NotePositionX)
		checkErr(err)
		notes = append(notes, Note{
			NoteColor:     NoteColor,
			NoteContent:   NoteContent,
			NotePositionX: NotePositionX,
			NotePositionY: NotePositionY})
	}
	return notes
}

// InsertNote is insert note data
func (noteRep *NoteRepository) InsertNote(s Note) {
	db, err := sql.Open("mysql", config.FormatDSN())
	checkErr(err)
	stmt, err := db.Prepare("INSERT NotesList SET NoteColor=? , NoteContent=? , NotePositionX=? , NotePositionY=?")
	checkErr(err)
	res, err := stmt.Exec(s.NoteColor, s.NoteContent, s.NotePositionX, s.NotePositionY)
	checkErr(err)
	fmt.Println("Insert Success")
	fmt.Println(res.RowsAffected())
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
