package engine_io

import (
	"encoding/gob"
	"fmt"
	"os"
	"path"

	"github.com/miguel-panuto/clear-db/src/database"
)

func getPath(name string) string {
	filePath := path.Join("data", name+".cdb")
	return filePath
}

func SaveData(db *database.Database) error {
	err := verifyDataFolder()

	if err != nil {
		fmt.Print(err)
		return err
	}

	file, err := os.Create(getPath(db.Name))

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

func UpdateFile(db *database.Database) error {
	filePath := getPath(db.Name)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := fmt.Errorf("database file %s does not exist: %w", filePath, err)
		fmt.Println(err)
		return err
	}

	if err := os.Remove(filePath); err != nil {
		err := fmt.Errorf("failed to remove file %s: %w", filePath, err)
		fmt.Println(err)
		return err
	}

	if err := SaveData(db); err != nil {
		err := fmt.Errorf("failed to save updated data: %w", err)
		fmt.Println(err)
		return err
	}

	return nil
}
