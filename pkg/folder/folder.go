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
		return errors.Wrap(err, "The Palantír grows dim... the vision fades")
	}

	path, err := exec.LookPath("code")
	if err != nil {
		return errors.Wrap(err, "The Seeing Stone reveals no path to VS Code in your realm")
	}

	cmd := exec.Command(path, projectFolderPath)
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "The gates of VS Code remain closed for %s... even speaking 'mellon' would not open them", projectFolderPath)
	}

	return nil
}

func findProjectFolder(folderName string, baseFolder string) (string, error) {
	var projectPath string

	err := filepath.WalkDir(baseFolder, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return errors.Wrapf(err, "The Seeing Stone cannot pierce the shadows at %s", path)
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
		return "", errors.Wrapf(err, "The Palantír's gaze failed to find %s within %s", folderName, baseFolder)
	}

	if projectPath == "" {
		return "", errors.Errorf("The Seeing Stone reveals no trace of %s in the lands of %s", folderName, baseFolder)
	}

	log.Printf("The Palantír has revealed your project! Found at: %s\nOpening the gates to your realm...", projectPath)

	return projectPath, nil
}

func ListFolders(baseFolder string) error {
	var folders []string

	err := filepath.WalkDir(baseFolder, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return errors.Wrapf(err, "The Seeing Stone cannot pierce the shadows at %s", path)
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
		return errors.Wrapf(err, "The Palantír's vision is clouded... could not survey the realm of %s", baseFolder)
	}

	if len(folders) == 0 {
		return errors.New("The Seeing Stone gazes far and wide, yet no projects appear in its vision")
	}

	log.Println("The Palantír reveals the following realms in your domain:")
	for _, folder := range folders {
		log.Printf("\t%s\n", folder)
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
