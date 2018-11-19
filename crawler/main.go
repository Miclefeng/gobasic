package main

import (
	"fmt"
	"io/ioutil"
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

	if result, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", result)
}