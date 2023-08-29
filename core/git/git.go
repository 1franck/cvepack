package git

import (
	"os/exec"
	"runtime"
)

func CommandExists() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

func Pull(path string) (string, error) {
	args := buildCmdArgs("git", "pull")
	return execCommand(path, args...)
}

func Push(path string, origin string, branch string) (string, error) {
	args := buildCmdArgs("git", "push", origin, branch)
	return execCommand(path, args...)
}

func StageAllModified(path string) (string, error) {
	args := buildCmdArgs("git", "add", "-u")
	return execCommand(path, args...)
}

func Commit(path string, message string) (string, error) {
	args := buildCmdArgs("git", "commit", "-m", message)
	return execCommand(path, args...)
}

func buildCmdArgs(args ...string) []string {
	if runtime.GOOS == "windows" {
		return append([]string{"cmd", "/C"}, args...)
	}
	return args
}

func execCommand(path string, args ...string) (string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = path
	out, err := cmd.CombinedOutput()

	if err != nil {
		return string(out), err
	}
	return string(out), nil
}
