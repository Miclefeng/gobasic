package FileOperation

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/1 下午4:46
 */

func ReadFile(file string) []string {
	inputFile, inputError := os.Open(file)
	defer inputFile.Close()

	if inputError != nil {
		fmt.Printf("An error occurred on opening the file :\n" + "Does the fiel exist?\n" + "Have you got acces to it?\n")
		return nil
	}
	words := make([]string, 0)
	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanWords)
	for inputScanner.Scan() {
		words = append(words, strings.ToLower(inputScanner.Text()))
	}
	return words
}