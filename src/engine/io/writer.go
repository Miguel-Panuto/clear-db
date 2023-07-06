package engine_io

import (
	"fmt"
	"os"
	"path"

	"github.com/miguel-panuto/clear-db/src/database"
)

func getPath(name string) string {
	return path.Join("data", name+".cdb")
}

func SaveData(db *database.Database) error {
	err := verifyDataFolder()

	if err != nil {
		fmt.Println(err)
		return err
	}

	lines := db.Name + "\n"

	for i, value := range db.Tables {
		if i > 0 {
			lines += "\n;;\n"
		}
		lines += value.Name + "\n"
		lines += value.GetFields()
	}

	file, err := os.Create(getPath(db.Name))

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()
	if _, err := file.Write([]byte(lines)); err != nil {
		fmt.Println("Database was not saved")
		return err
	}

	return nil
}

func UpdateFile(db *database.Database) error {
	filePath := getPath(db.Name)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = fmt.Errorf("database file %s does not exist: %w", filePath, err)
		fmt.Println(err)
		return err
	}

	if err := SaveData(db); err != nil {
		err = fmt.Errorf("failed to save updated data: %w", err)
		fmt.Println(err)
		return err
	}

	return nil
}
