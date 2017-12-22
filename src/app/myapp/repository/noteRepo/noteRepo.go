package noteRepo

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/Masterminds/goutils"

	mysql "github.com/go-sql-driver/mysql"
)

// Note data model struct
type NoteModel struct {
	ID            int     `json:"Id"`
	NoteColor     int     `json:"NoteColor"`
	NoteContent   string  `json:"Notecontent"`
	NotePositionX float32 `json:"NotePositionX"`
	NotePositionY float32 `json:"NotePositionY"`
	NoteKey       string  `json:"NoteKey"`
	IsArchived    int     `json:"IsArchived"`
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
var conStr string

func init() {
	config = mysql.Config{
		User:   "root",
		Passwd: "abc123456",
		Addr:   "192.168.2.13:3306",
		DBName: "Notes",
	}
	conStr = "root:a5566%%^^@tcp(localhost:3333)/Notes"
}

// ArchivedNote will archived  Note data
func (noteRep *NoteRepository) ArchivedNote(item NoteModel) (isOk int) {
	db, err := sql.Open("mysql", conStr)
	defer db.Close()
	checkErr(err)
	tx, _ := db.Begin()
	tx.Exec(`
			UPDATE NotesList
			SET NoteColor=?,
			NoteContent=?,
			NotePositionY=?,
			NotePositionX=?
			WHERE FK_NoteKey=? AND idNoteList=?
			`, item.NoteColor, item.NoteContent,
		item.NotePositionX, item.NotePositionY,
		item.NoteKey, item.ID)
	tx.Commit()
	return 0
}

// PutNotes will update mutiple Note data
func (noteRep *NoteRepository) PutNotes(items []NoteModel) (isOk int) {
	db, err := sql.Open("mysql", conStr)
	defer db.Close()
	checkErr(err)
	tx, _ := db.Begin()
	for index, elem := range items {
		fmt.Println(index)
		fmt.Println(elem)
		tx.Exec(`
			UPDATE NotesList
			SET NoteColor=?,
			NoteContent=?,
			NotePositionY=?,
			NotePositionX=?
			WHERE FK_NoteKey=? AND idNoteList=?
			`, elem.NoteColor, elem.NoteContent,
			elem.NotePositionX, elem.NotePositionY,
			elem.NoteKey, elem.ID,
		)
	}
	tx.Commit()
	return 0
}

// GetNotes will return mutiple Note data. data is not archived
func (noteRep *NoteRepository) GetNotes(key string) []NoteModel {
	db, err := sql.Open("mysql", conStr)
	defer db.Close()
	checkErr(err)
	rows, err := db.Query(`SELECT 
		idNotesList,FK_NoteKey,
		NoteColor,NoteContent
		,NotePositionX,NotePositionY
		,IsArchived 
		FROM NotesList
		WHERE FK_NoteKey=? AND IsArchived=0
		`, key)
	checkErr(err)
	defer rows.Close()
	var notes []NoteModel
	var (
		ID            int
		NoteKey       string
		NoteColor     int
		NoteContent   string
		NotePositionX float32
		NotePositionY float32
		IsArchived    int
	)
	for rows.Next() {
		err := rows.Scan(
			&ID, &NoteKey,
			&NoteColor, &NoteContent,
			&NotePositionX, &NotePositionX)
		checkErr(err)
		notes = append(notes, NoteModel{
			ID:            ID,
			NoteKey:       NoteKey,
			NoteColor:     NoteColor,
			NoteContent:   NoteContent,
			NotePositionX: NotePositionX,
			NotePositionY: NotePositionY,
			IsArchived:    IsArchived,
		})
	}
	return notes
}

// InsertNoteMain is insert note main data
func (noteRep *NoteRepository) InsertNoteMain(s NoteMainModel) (result string) {
	db, err := sql.Open("mysql", conStr)
	defer db.Close()
	checkErr(err)
	stmt, err := db.Prepare(`INSERT NoteMain
		SET NoteName=? , NoteKey=? `)
	checkErr(err)
	// Generator RandomKey
	// https://github.com/Masterminds/goutils
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	key, _ := goutils.RandomSeed(10, 0, 0, true, true, nil, random)
	res, err := stmt.Exec(s.NoteName, key)
	checkErr(err)
	fmt.Println("Insert Success")
	fmt.Println(res.RowsAffected())
	return key
}

// InsertNote is insert note data
func (noteRep *NoteRepository) InsertNote(s NoteModel) {
	db, err := sql.Open("mysql", config.FormatDSN())
	defer db.Close()
	checkErr(err)
	stmt, err := db.Prepare(`INSERT NotesList 
		SET NoteColor=? , NoteContent=? , 
		NotePositionX=? , NotePositionY=?,
		FK_NoteKey=?
		 `)
	checkErr(err)
	res, err := stmt.Exec(s.NoteColor,
		s.NoteContent, s.NotePositionX,
		s.NotePositionY, s.NoteKey)
	checkErr(err)
	fmt.Println("Insert Success")
	fmt.Println(res.RowsAffected())
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
