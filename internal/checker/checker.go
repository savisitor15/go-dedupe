package checker

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func checkFileExists(p string) error {
	pth := filepath.Clean(p)
	if fileInfo, err := os.Stat(pth); errors.Is(err, os.ErrNotExist) {
		// Directory not found!
		return fmt.Errorf("Path %s not found!\n", pth)
	} else if err != nil {
		return err
	} else {
		if fileInfo.IsDir() {
			return fmt.Errorf("Path %s must be a file nod directory!", pth)
		}
	}
	return nil
}

func chunkReadFileToSHA256(f *os.File) ([32]byte, error){
	shaHandler := sha256.New() // Initiate a checksum handler
	r := bufio.NewReader(f)
	for {
		buf := make([]byte, 48*1024) // Read chunk size
		n, err := r.Read(buf) // Read chunk
		buf = buf[:n] // ensure that trailing reads are all that's in the buffer
		if n == 0 {
			if err == io.EOF{
				// Just end of file! no error
				break
			}
			if err != nil{
				return [32]byte{}, err
			}
		}
		shaHandler.Write(buf)
	}
	return [32]byte(shaHandler.Sum(nil)), nil
}

func FileToSHA256(p string) ([32]byte, error){
	var err error
	if err := checkFileExists(p); err != nil {
		// Path is not good
		return [32]byte{}, err
	}
	var f *os.File
	if f, err = os.Open(p); err != nil {
		return [32]byte{}, err
	}
	defer f.Close()
	return chunkReadFileToSHA256(f)
}
