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
		log.Fatal("The Seeing Stone requires alignment! Set PALANTIR_ROOT_FOLDER to gaze upon your realms")
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
			log.Fatalf("The vision has failed! %v", err)
		}
	case "list":
		if len(os.Args) < 2 {
			printUsageAndExit()
		}

		if err := folder.ListFolders(baseFolder); err != nil {
			log.Fatalf("The Seeing Stone grows dark! %v", err)
		}
	default:
		printUsageAndExit()
	}
}

func printUsageAndExit() {
	fmt.Println(`
    🔮 The Palantír - Seeing Stone of Code 🔮

    "Far-sighted it was, and its gaze pierced through shadow and distance"

    Commands:
      palantir open <project-name>    Peer into a realm and open it in VS Code
      palantir list                   Reveal all realms within your domain

    Required Environment Variables:
      PALANTIR_ROOT_FOLDER      The root path where your projects dwell
      PALANTIR_PROJECTS_DEPTH   The depth level where projects are found (e.g., 1, 2, 3)

    For detailed configuration and usage instructions, consult the sacred scrolls:
    📜 https://github.com/DevMatheusSilva/palantir/blob/main/README.md
    `)
	os.Exit(1)
}

func changeToHomeFolder() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("The path to your home is obscured by shadow and flame: %v\nTrace:\n%s",
			err, debug.Stack())
	}

	return homeDir
}
