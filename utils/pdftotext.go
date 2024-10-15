package utils

import "os/exec"

// check if pdftotext is installed

func IsPdftotextInstalled() bool {
	_, err := exec.LookPath("pdftotext")
	return err == nil
}
