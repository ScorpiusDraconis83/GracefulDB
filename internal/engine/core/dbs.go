package core

import (
	"fmt"
	"os"
	"slices"
	"time"
)

// Marks the database as deleted, but does not delete files.
func RemoveDB(nameDB string) bool {
	// This function is complete
	dbInfo, ok := StorageInfo.DBs[nameDB]
	if ok {
		dbInfo.LastUpdate = time.Now()
		dbInfo.Deleted = true

		StorageInfo.Removed = append(StorageInfo.Removed, dbInfo)
		delete(StorageInfo.DBs, nameDB)
	}

	return dbInfo.Save()
}

// Deletes the folder and database files, if DB was mark as 'removed'
func StrongRemoveDB(nameDB string) bool {
	// This function is complete
	for indRange, dbInfo := range StorageInfo.Removed {
		if dbInfo.Name == nameDB {
			dbPath := fmt.Sprintf("%s%s", LocalCoreSettings.Storage, dbInfo.Folder)
			err := os.RemoveAll(dbPath)
			if err != nil {
				return false
			}

			StorageInfo.Removed = slices.Delete(StorageInfo.Removed, indRange, indRange+1)

			return true
		}
	}

	return false
}

// Creating a new database.
func CreateDB(nameDB string, secure bool) bool {
	// This function is complete
	if secure && !RegExpCollection["EntityName"].MatchString(nameDB) {
		return false
	}

	var folderDB string

	_, ok := StorageInfo.DBs[nameDB]
	if ok {
		return false
	}

	for {
		folderDB = GenerateName()
		if !CheckFolder(LocalCoreSettings.Storage, folderDB) {
			break
		}
	}

	fullNameFolderDB := fmt.Sprintf("%s%s", LocalCoreSettings.Storage, folderDB)
	err := os.Mkdir(fullNameFolderDB, 0666)
	if err != nil {
		return false
	}

	dbInfo := tDBInfo{
		Name:       nameDB,
		Folder:     folderDB,
		Tables:     make(map[string]tTableInfo),
		Removed:    make([]tTableInfo, 0),
		LastUpdate: time.Now(),
		Deleted:    false,
	}

	StorageInfo.DBs[nameDB] = dbInfo

	return dbInfo.Save()
}