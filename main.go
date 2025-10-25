package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/DevMatheusSilva/palantir/pkg/folder"
)

func main() {
	if len(os.Args) < 2 {
		printUsageAndExit()
	}

	command := os.Args[1]

	rootFolder := os.Getenv("PALANTIR_ROOT_FOLDER")
	if rootFolder == "" {
		log.Fatal("By the power of the Istari! The Palantír requires its sacred environment path (PALANTIR_ROOT_FOLDER) before you may gaze into its depths")
	}

	homeDir := changeToHomeFolder()

	baseFolder := filepath.Join(homeDir, rootFolder)

	switch command {
	case "open":
		if len(os.Args) < 3 {
			printUsageAndExit()
		}

		folderName := os.Args[2]
		if err := folder.OpenFolderProject(folderName, baseFolder); err != nil {
			log.Fatalf("The Ring has betrayed us! %v", err)
		}
	case "list":
		if len(os.Args) < 2 {
			printUsageAndExit()
		}

		if err := folder.ListFolders(baseFolder); err != nil {
			log.Fatalf("The Ring has betrayed us! %v", err)
		}
	default:
		printUsageAndExit()
	}
}

func printUsageAndExit() {
	fmt.Println(`
    One Command to rule them all,
    One Path to find them,
    One Script to bring them all,
    and in the terminal bind them.

    Usage: palantir <command> <folder_name>
    `)
	os.Exit(1)
}

func changeToHomeFolder() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("The paths are blocked! Like the Mines of Moria: %v\nTrace of shadow and flame:\n%s",
			err, debug.Stack())
	}

	return homeDir
}
