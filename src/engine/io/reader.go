package engine_io

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/miguel-panuto/clear-db/src/database"
)

func LoadDatabases() ([]*database.Database, error) {
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

	dbs := []*database.Database{}

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".cdb") {
			filePath := filepath.Join(dir, f.Name())

			file, err := os.ReadFile(filePath)
			if err != nil {
				return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
			}

			db := parseToDatabase(string(file))

			if err != nil {
				return nil, err
			}

			dbs = append(dbs, db)
		}
	}

	return dbs, nil
}
