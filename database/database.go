package database

import (
	"fmt"
	"encoding/csv"
	"log"
	"os"
	"strings"
)


DB_PATH := "~/.tasker/db/"


type CSVTaskDB struct {
	Name string
}


type CSVTaskRecord struct {
	Title string
	CreateDate string
	CompleteDate string
	LastUpdateDate string
}


func getExistingDatabaseList() ([]string) {
	var dbList []string
	entries, err := os.ReadDir(DB_PATH)
	if err != nil {
		log.Fatalf("Failed to Read Databases located at: %s", DB_PATH)
	}
	for _, e := range entries {
		dbList = append(dbList, e.Name)
	}

	return dbList
}


func (db *CSVTaskDB) Exists bool {
	dbExists := false

}


func (db *CSVTaskDB) CreateDatabase (error) {
	error createDatabasesError := nil 
        
}
