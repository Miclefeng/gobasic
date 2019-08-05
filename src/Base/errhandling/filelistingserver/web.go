package main

import (
	"log"
	"net/http"
	"miclefeng/learngo/errhandling/filelistingserver/filelisting"
	"os"
)

// Type assertion
type appHandler func(writer http.ResponseWriter, request *http.Request) error

//noinspection GoUnresolvedReference 函数式编程
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {// recover 仅在defer中调用，获取 panic 的值			log.Printf("Panic : %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			logger := log.New(os.Stdout, "Warn : ", log.Lshortfile)
			logger.Printf("Error handling request: %s", err.Error());

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

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

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
