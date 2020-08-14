package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(40 * time.Millisecond)

// 获取页面内容并返回
func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	// 获取 URL response 返回
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// 获取页面字符编码
	newReader := bufio.NewReader(resp.Body)
	encoder := determineEncoding(newReader)
	// 字符编码转换
	reader := transform.NewReader(newReader, encoder.NewEncoder())
	return ioutil.ReadAll(reader)
}

// 获取页面的字符编码
func determineEncoding(body *bufio.Reader) encoding.Encoding {
	// 获取前 1024 个字节
	bytes, err := body.Peek(1024)
	if err != nil {
		log.Printf("Fetch error: %v", err)
		return unicode.UTF8
	}

	// 获取字节编码
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
