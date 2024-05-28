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


func getExistingDatabaseList(dbPath string, dbName string) ([]string) {
	entries, err := os.ReadDir()

}


func (db *CSVTaskDB) Exists bool {
	dbExists := false

}


func (db *CSVTaskDB) CreateDatabase (error) {
	error createDatabasesError := nil 
        
}
