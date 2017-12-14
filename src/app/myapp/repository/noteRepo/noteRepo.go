package noteRepo

import (
	"database/sql"
	"fmt"

	mysql "github.com/go-sql-driver/mysql"
)

// Note data model struct
type NoteModel struct {
	NoteColor     int     `json:"NoteColor"`
	NoteContent   string  `json:"Notecontent"`
	NotePositionX float32 `json:"NotePositionX"`
	NotePositionY float32 `json:"NotePositionY"`
	FK_NoteKey    string  `json:"FK_NoteKey"`
}

// NoteMain data model struct
type NoteMainModel struct {
	NoteName string `json:"NoteName"`
	NoteKey  string `json:"NoteKey"`
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
func (noteRep *NoteRepository) GetNotes() []NoteModel {
	db, err := sql.Open("mysql", config.FormatDSN())
	checkErr(err)
	rows, err := db.Query(`SELECT 
		NoteColor,NoteContent
		,NotePositionX,NotePositionY 
		FROM NotesList`)
	checkErr(err)
	defer rows.Close()
	var notes []NoteModel
	var (
		NoteColor     int
		NoteContent   string
		NotePositionX float32
		NotePositionY float32
	)
	for rows.Next() {
		err := rows.Scan(&NoteColor, &NoteContent, &NotePositionX, &NotePositionX)
		checkErr(err)
		notes = append(notes, NoteModel{
			NoteColor:     NoteColor,
			NoteContent:   NoteContent,
			NotePositionX: NotePositionX,
			NotePositionY: NotePositionY})
	}
	return notes
}

// InsertNoteMain is insert note main data
func (noteRep *NoteRepository) InsertNoteMain(s NoteMainModel) {
	db, err := sql.Open("mysql", config.FormatDSN())
	checkErr(err)
	stmt, err := db.Prepare(`INSERT NotesMain
		SET NoteName=? , NoteKey=? `)
	checkErr(err)
	res, err := stmt.Exec(s.NoteName, s.NoteKey)
	checkErr(err)
	fmt.Println("Insert Success")
	fmt.Println(res.RowsAffected())
}

// InsertNote is insert note data
func (noteRep *NoteRepository) InsertNote(s NoteModel) {
	db, err := sql.Open("mysql", config.FormatDSN())
	checkErr(err)
	stmt, err := db.Prepare(`INSERT NotesList 
		SET NoteColor=? , NoteContent=? , 
		NotePositionX=? , NotePositionY=?`)
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
