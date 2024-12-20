package cmd

import (
	"os"
)

func Initialize() error {
	content := `
	root = "."
	tmp_dir = "tmp"
	[build]
	bin = "./tmp/main"
	cmd = "nvcc --std=c++17 -o ./tmp/main main.cu"
	log = "build-errors.log"
	`

	file, err := os.Create(".cudair.toml")
	if err != nil {

		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
