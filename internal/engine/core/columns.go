package core

import (
	"fmt"
	"os"
	"time"
)

// Creating a new column.
func CreateColumn(nameDB, nameTable, nameColumn string) bool {
	// This function is complete
	var folderName string

	dbInfo, ok := StorageInfo.DBs[nameDB]
	if !ok {
		return false
	}

	tableInfo, ok := dbInfo.Tables[nameTable]
	if !ok {
		return false
	}

	pathTable := fmt.Sprintf("%s%s/%s/", LocalCoreSettings.Storage, tableInfo.Parent, tableInfo.Folder)

	for {
		folderName = GenerateName()
		if !CheckFolderOrFile(pathTable, folderName) {
			break
		}
	}

	fullColumnName := fmt.Sprintf("%s%s", pathTable, folderName)
	err := os.Mkdir(fullColumnName, 0666)
	if err != nil {
		return false
	}

	tNow := time.Now()

	columnInfo := tColumnInfo{
		Name:       nameColumn,
		Folder:     folderName,
		Parents:    fmt.Sprintf("%s/%s", tableInfo.Parent, tableInfo.Folder),
		BucketLog:  2,
		BucketSize: LocalCoreSettings.BucketSize,
		OldRev:     "",
		CurrentRev: GenerateRev(),
		LastUpdate: tNow,
		Deleted:    false,
	}

	tableInfo.Columns[nameColumn] = columnInfo
	tableInfo.Order = append(tableInfo.Order, nameColumn)
	tableInfo.LastUpdate = tNow

	dbInfo.Tables[nameTable] = tableInfo
	dbInfo.LastUpdate = tNow

	StorageInfo.DBs[nameDB] = dbInfo

	return StorageInfo.Save()
}
