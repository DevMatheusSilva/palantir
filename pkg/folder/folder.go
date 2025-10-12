package folder

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func OpenFolderProject(folderName string, baseFolder string) error {
	projectFolderPath, err := findProjectFolder(folderName, baseFolder)
	if err != nil {
		return errors.Wrap(err, "The quest to find the project folder has failed!")
	}

	path, err := exec.LookPath("code")
	if err != nil {
		return errors.Wrap(err, "Alas! The ancient scrolls speak of VS Code, but it cannot be found in the paths of your realm")
	}

	cmd := exec.Command(path, projectFolderPath)
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "The gates of VS Code remain sealed at %s. Even speaking 'mellon' would not work", projectFolderPath)
	}

	return nil
}

func findProjectFolder(folderName string, baseFolder string) (string, error) {
	var projectPath string

	err := filepath.Walk(baseFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrapf(err, "The path %s is guarded by dark forces", path)
		}

		if info.IsDir() {
			switch info.Name() {
			case ".git", "node_modules", "vendor":
				return filepath.SkipDir
			}
		}

		if info.IsDir() && info.Name() == folderName {
			projectPath = path
			return filepath.SkipAll
		}

		return nil
	})

	if err != nil {
		return "", errors.Wrapf(err, "The search party has failed! Could not locate %s in the realm of %s", folderName, baseFolder)
	}

	if projectPath == "" {
		return "", errors.Errorf("Not even the eyes of Legolas could find %s in the lands of %s", folderName, baseFolder)
	}

	log.Printf("By the light of Eärendil! Project found at: %s\nSummoning the editor...", projectPath)

	return projectPath, nil
}
