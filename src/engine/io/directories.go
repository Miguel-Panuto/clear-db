package engine_io

import (
	"fmt"
	"os"
)

func verifyDataFolder() error {
	errDir := os.MkdirAll("data", 0755)
	if errDir != nil {
		return fmt.Errorf("failed to create directory: %w", errDir)
	}
	return nil
}
