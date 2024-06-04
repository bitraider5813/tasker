package database

import (
	"fmt"
	"encoding/csv"
	"log"
	"os"
	"strings"
	"slices"
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
	dbList := getExistingDatabaseList()
	dbExists := slices.Contains(dbList, db.Name)
	return dbExists
}


// We should Initialize a Default Database if none already exist at DB_PATH
func InitializeDefaultDatabase() error {
	initError := nil
	return initError
}

func (db *CSVTaskDB) CreateDatabase (error) {
	error createDatabasesError := nil
}
