package engine_io

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/miguel-panuto/clear-db/src/database"
)

const dir string = "data"

func LoadDatabases() ([]string, error) {

	dirFile, err := os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to open directory: %w", err)
	}
	defer dirFile.Close()

	files, err := dirFile.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	dbs := []string{}

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".cdb") {
			dbs = append(dbs, strings.ReplaceAll(f.Name(), ".cdb", ""))
		}
	}

	return dbs, nil
}

func ReadDatabase(dbName string) (*database.Database, error) {
	filePath := filepath.Join(dir, dbName+".cdb")
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}

	db := parseToDatabase(string(file))

	if err != nil {
		return nil, err
	}

	return db, nil
}
