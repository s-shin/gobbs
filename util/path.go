package util

import (
	"log"
	"os"
	"path"
	"runtime"
)

// In the production environtment, this function shouldn't be used.
func ProjectDir() string {
	return path.Dir(CallerDir())
}

// In the production environtment, this function shouldn't be used.
// http://andrewbrookins.com/tech/golang-get-directory-of-the-current-file/
func CallerDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func CWD() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
