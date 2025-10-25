package folder

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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

	err := filepath.WalkDir(baseFolder, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return errors.Wrapf(err, "The path %s is guarded by dark forces", path)
		}

		if err := verifyIgnoredFolder(info); err != nil {
			return err
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

func ListFolders(baseFolder string) error {
	var folders []string

	err := filepath.WalkDir(baseFolder, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return errors.Wrapf(err, "The path %s is guarded by dark forces", path)
		}

		if err := verifyIgnoredFolder(info); err != nil {
			return err
		}

		isInProjectsDepth, err := checkIfIsInProjectsDepth(path, baseFolder)
		if err != nil {
			return filepath.SkipAll
		}

		if info.IsDir() && path != baseFolder && isInProjectsDepth {
			folders = append(folders, info.Name())
		}

		return nil
	})

	if err != nil {
		return errors.Wrapf(err, "The search party has failed! Could not list folders in the realm of %s", baseFolder)
	}

	if len(folders) == 0 {
		return errors.New("Even the Palantír sees no projects in this land.")
	}

	log.Println("Legolas elf eyes saw the following projects being taken to Isengard:")
	for _, folder := range folders {
		log.Printf("- %s\n", folder)
	}

	return nil
}

func verifyIgnoredFolder(folderInfo os.DirEntry) error {
	defaultIgnoredFolders := []string{".git", "node_modules", "vendor", ".idea", ".vscode"}

	if folderInfo.IsDir() {
		for _, ignored := range defaultIgnoredFolders {
			if folderInfo.Name() == ignored {
				return filepath.SkipDir
			}
		}
	}

	return nil
}

func checkIfIsInProjectsDepth(path string, baseFolder string) (bool, error) {
	desiredDepth, err := strconv.Atoi(os.Getenv("PALANTIR_PROJECTS_DEPTH"))
	if err != nil {
		log.Printf("One does not simply convert PALANTIR_PROJECTS_DEPTH into an integer: %v", err)
		return false, err
	}

	pathSepsNum := strings.Count(path, string(os.PathSeparator))
	baseFolderSepsNum := strings.Count(baseFolder, string(os.PathSeparator))

	return pathSepsNum == baseFolderSepsNum+desiredDepth, nil
}
