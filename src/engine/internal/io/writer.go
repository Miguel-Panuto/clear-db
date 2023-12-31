package engine_io

import (
	"fmt"
	"os"
	"strings"

	"github.com/miguel-panuto/clear-db/src/database"
	"github.com/miguel-panuto/clear-db/src/utils"
)

func SaveData(db *database.Database) error {
	err := verifyDataFolder()

	if err != nil {
		fmt.Println(err)
		return err
	}

	lines := db.Name

	for i, value := range db.Tables {
		lines += "\n"
		if i > 0 {
			lines += ";;\n"
		}
		lines += value.Name + "\n"
		lines += value.GetFieldsString()

		if len(*value.Rows) > 0 {
			for _, row := range *value.Rows {
				arrStr := utils.MakeStringArr(row)
				lines += "\n" + strings.Join(arrStr, ";")
			}
		}
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
