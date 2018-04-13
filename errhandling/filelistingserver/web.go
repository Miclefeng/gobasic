package main

import (
	"log"
	"net/http"
	"miclefeng/learngo/errhandling/filelistingserver/filelisting"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

//noinspection GoUnresolvedReference
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			logger := log.New(os.Stdout, "Warn : ", log.Lshortfile)
			logger.Printf("Error handling request: %s", err.Error());
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				//http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}