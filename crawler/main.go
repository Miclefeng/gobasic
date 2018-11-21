package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/18 上午11:35
 */

 var Url = "http://www.zhenai.com/zhenghun"

func main() {
	var (
		err error
		resp *http.Response
		result []byte
		encodingReader *transform.Reader
	)

	if resp, err = http.Get(Url); err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error status code: ", resp.StatusCode)
		return
	}
	// 转换页面内容的编码格式
	encodingReader = transform.NewReader(resp.Body, determineEncoding(resp.Body).NewEncoder())
	// 读取页面内容
	if result, err = ioutil.ReadAll(encodingReader); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", result)
}

 // 判断页面内容的编码格式
func determineEncoding(r io.Reader) (e encoding.Encoding) {
	var (
		bytes []byte
		err error
	)

	if bytes, err = bufio.NewReader(r).Peek(1024); err != nil {
		log.Fatalf("bufio peek err: %v\n", err);
		e = unicode.UTF8
		return e
	}

	e, _, _ = charset.DetermineEncoding(bytes, "")
	return e
}