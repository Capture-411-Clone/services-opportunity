package collector

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func PopplerExtractTextFromPDF(pdfData []byte) (string, error) {
	hardLimit := "2"

	// write pdfData to a temp file WITH os.CreateTemp
	pdfFile, err := os.CreateTemp("", "pdf-*.pdf")
	if err != nil {
		return "", err
	}
	defer os.Remove(pdfFile.Name())
	defer pdfFile.Close()

	_, err = pdfFile.Write(pdfData)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("pdftotext", "-f", "1", "-l", hardLimit, pdfFile.Name(), "-")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	fmt.Println(out.String())
	return out.String(), nil
}
