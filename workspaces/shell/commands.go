package shell

import (
	"fmt"
	"github.com/fawind/workspaces/workspaces/finder"
	"github.com/pkg/errors"
	"os"
	"os/exec"
)

func MayCloneRepo(repo finder.LocalRepo) error {
	hasDir, err := dirExists(repo.LocalDir)
	if err != nil {
		return err
	}

	if hasDir {
		return nil
	}

	return cloneRepo(repo)
}

func cloneRepo(repo finder.LocalRepo) error {
	gitAddress := repo.Organization.GetGitAddress(repo.RepoName)
	fmt.Printf("Cloning %s into %s\n", gitAddress, repo.LocalDir)
	cmd := exec.Command("git", "clone", gitAddress, repo.LocalDir)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		return errors.Wrap(err, "error cloning repo")
	}

	if err := cmd.Wait(); err != nil {
		return errors.Wrap(err, "error cloning repo")
	}

	return nil
}

func dirExists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, errors.Wrapf(err, "error checking if dir exist: %s", dir)
}
