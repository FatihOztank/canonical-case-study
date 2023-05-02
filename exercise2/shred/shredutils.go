package shred

import (
	"crypto/rand"
	"io"
	"os"
)

func Shred(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	// get fileinfo to query size of the file
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	defer file.Close()

	// pass the file 3 times with random data
	for pass := 0; pass < 3; pass++ {
		shredBuf := make([]byte, fileInfo.Size())
		_, err := rand.Read(shredBuf)
		if err != nil {
			return err
		}

		_, err = file.WriteAt(shredBuf, 0)
		if err != nil {
			return err
		}

	}

	err = file.Truncate(fileInfo.Size())
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	// for debugging, ensuring file has been filled with
	// random data. Uncomment below to print the file content

	// buffer := make([]byte, fileInfo.Size())
	// _, err = file.Read(buffer)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%s\n", string(buffer))

	// delete the file after filling with random data
	if err := os.Remove(fileName); err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

func CopyFile(source string, destination string) error {
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	err = dstFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
