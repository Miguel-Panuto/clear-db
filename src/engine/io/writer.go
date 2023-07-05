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

	lines := db.Name + ";;"

	for _, value := range db.Tables {
		lines += value.Name + ";;"
		lines += value.GetFields()
	}

	println(lines)

	// file, err := os.Create(getPath(db.Name))

	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// defer file.Close()
	// encoder := gob.NewEncoder(file)

	// if err := encoder.Encode(db); err != nil {
	// 	fmt.Println("Database was not saved")
	// 	return err
	// }

	return nil
}

func UpdateFile(db *database.Database) error {
	filePath := getPath(db.Name)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("database file %s does not exist: %w", filePath, err)
	}

	// The file exists, save the updated data directly
	if err := SaveData(db); err != nil {
		return fmt.Errorf("failed to save updated data: %w", err)
	}

	return nil
}
