package interface

import (
	"encoding/base64"
	"os"
)

// EncodeBase64 image Encode to base64
func EncodeBase64(savePath string, fileNama string) string {
	file, err := os.Open(savePath + fileNama)
	defer file.Close()

	if err != nil {
		return err.Error()
	}

	fi, _ := file.Stat() // interface
	size := fi.Size()    // file size

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}
