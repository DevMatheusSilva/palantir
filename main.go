package main

import (
	"fmt"
	"github.com/DevMatheusSilva/palantir/pkg/folder"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	rootFolder := os.Getenv("PALANTIR_ROOT_FOLDER")
	if rootFolder == "" {
		log.Fatal("By the power of the Istari! The Palantír requires its sacred environment path (PALANTIR_ROOT_FOLDER) before you may gaze into its depths")
	}

	changeToDefaultFolder(rootFolder)

	command := os.Args[1]

	switch command {
	case "open":
		folderName := os.Args[2]
		if err := folder.OpenFolderProject(folderName); err != nil {
			log.Fatalf("The Ring has betrayed us! %v", err)
		}
	}
}

func printUsage() {
	fmt.Println(`
    One Command to rule them all,
    One Path to find them,
    One Script to bring them all,
    and in the terminal bind them.

    Usage: palantir <command> <folder_name>
    `)
}

func changeToDefaultFolder(rootFolder string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("The paths are blocked! Like the Mines of Moria: %v\nTrace of shadow and flame:\n%s",
			err, debug.Stack())
	}

	baseFolder := filepath.Join(homeDir, rootFolder)

	err = os.Chdir(baseFolder)
	if err != nil {
		log.Fatalf("Failed to traverse to %s! As Gandalf said of Moria: %v\nTrace of shadow and flame:\n%s",
			baseFolder, err, debug.Stack())
	}
}
