package engine_io

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/miguel-panuto/clear-db/src/database"
)

func LoadDatabases() (*[]database.Database, error) {
	dir := "data"

	dirFile, err := os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to open directory: %w", err)
	}
	defer dirFile.Close()

	files, err := dirFile.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	dbs := []database.Database{}

	// Iterate over the files in the directory
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".cdb") {
			// Construct the full file path
			filePath := filepath.Join(dir, f.Name())

			// Open the file
			file, err := os.Open(filePath)
			if err != nil {
				return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
			}
			defer file.Close()

			// Decode the file into a Database object
			var db database.Database
			decoder := gob.NewDecoder(file)
			if err := decoder.Decode(&db); err != nil {
				return nil, fmt.Errorf("failed to decode file %s: %w", filePath, err)
			}

			// Add the Database object to the slice
			dbs = append(dbs, db)
		}
	}

	return &dbs, nil
}
