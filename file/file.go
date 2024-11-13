package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type File struct {
	Url       string
	name      string
	createdAt time.Time
}

func NewFile(name string, content *string) *File {
	results := strings.Split(name, ".")
	url := results[0] + "_" + time.Now().Format("2006-01-01") + "." + results[1]

	file, err := os.Create(url)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(*content)
	if err != nil {
		file.Close()
		panic(err)
	}

	fmt.Println("File created: ", url)
	defer file.Close()

	return &File{
		Url:       url,
		name:      name,
		createdAt: time.Now(),
	}
}

func Read(url string) string {
	file, err := os.Open(url)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(content)
}
