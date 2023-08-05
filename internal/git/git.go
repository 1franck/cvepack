package git

import (
	"errors"
	"os/exec"
	"runtime"
)

func CommandExists() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

func Pull(path string) (string, error) {
	osName := runtime.GOOS

	if osName == "windows" {
		return windowsPull(path)
	}

	if osName == "linux" {
		return linuxPull(path)
	}

	return "", errors.New("Unsupported OS")
}

func windowsPull(path string) (string, error) {
	cmd := exec.Command("cmd", "/C", "git", "pull")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}
	return string(out), nil
}

func linuxPull(path string) (string, error) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}
	return string(out), nil
}
