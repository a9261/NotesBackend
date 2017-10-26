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
}

// InsertNote is insert note data
func InsertNote(s Note) {
	config := mysql.Config{
		User:   "root",
		Passwd: "nfu123!@#",
		Addr:   "192.168.52.137",
		DBName: "Notes",
	}
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
