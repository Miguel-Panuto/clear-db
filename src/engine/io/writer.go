package engine_io

import (
	"encoding/gob"
	"fmt"
	"os"
	"path"

	"github.com/miguel-panuto/clear-db/src/database"
)

func SaveData(db *database.Database) error {
	err := verifyDataFolder()

	if err != nil {
		return err
	}

	filePath := path.Join("data", db.Name+".cdb")
	file, err := os.Create(filePath)

	if err != nil {
		fmt.Print(err)
		return err
	}

	defer file.Close()
	encoder := gob.NewEncoder(file)

	if err := encoder.Encode(db); err != nil {
		fmt.Println("Database was not saved")
		return err
	}

	return nil
}
