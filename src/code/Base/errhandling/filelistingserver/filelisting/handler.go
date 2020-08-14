package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {

	fmt.Println(strings.Index(request.URL.Path, prefix))
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("Path must start with " + prefix)
	}

	path := request.URL.Path[len(prefix):] // /list/fib.txt, 去除 /list/
	//fmt.Println(path, request.URL.Path)
	//writer.Write([]byte(path))
	//return nil
	//fmt.Println(os.Getwd())
	file, err := os.Open("src/code/resources/" + path)
	//fmt.Println(err)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	_, _ = writer.Write(all)
	return nil
}
