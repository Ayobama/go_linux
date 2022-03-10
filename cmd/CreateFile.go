package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type File struct {
	ID            int
	Name          string
	Author        string
	Date_Created  string
	Last_Modified string
}

var emptyFile File

type Files []File

var dataFiles []File

//Creates a file using the os.Create Package. It returns an error.
//This function takes in two parameters
func (f File) CreateFile(name string, author string, content string) error {

	//Implement a check for if the directory exixts, before creating a new one
	err := os.MkdirAll("C:/Users/dell/github/golang/src/learning/go_linux/Creations/", 0777)
	if err != nil {
		ferr := errors.New("error message: string(err)")
		return ferr
	}

	var file *os.File
	file, err = os.OpenFile("./Creations/"+name, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer _ = file.Close()

	f.ID, _ = getNextID(dataFiles)
	f.Name = name
	f.Author = author
	f.Date_Created = time.Now().UTC().String()
	f.Last_Modified = f.Date_Created

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	dataFiles = append(dataFiles, f)
	fmt.Println("File", f.Name, "created")
	fmt.Printf("Other details:\n\t Author: %s\n\t Date Created: %s\n\t Last Modified: %s\n\t", f.Author, f.Date_Created, f.Last_Modified)

	return nil
}

func getNextID(dataFiles []File) (int, error) {
	if len(dataFiles) == 0 {
		//start the count
		return 1, nil
	} else {
		lastFile := dataFiles[len(dataFiles)-1]
		nextFile := lastFile.ID + 1

		return nextFile, nil
	}
	//I haven't imagined an error-case
	//return -1, nil

}

var FileNotFound error

func findFile(name string) (bool, File, error) {
	for i, file := range dataFiles {
		if name == file.Name {
			return true, file, nil
		} else {
			continue
		}
	}
	return false, emptyFile, FileNotFound
}

//touch the file: update the date modified, and create if not exists
func TouchFile(name string) error {
	bin, file, err := findFile(name)
	if err != nil {
		return err
	}
	if bin == true {
		file.Last_Modified = time.Now().UTC().String()
		return nil
	}

	return FileNotFound
}
