// read input from stdin

package main

import (
	"bufio"
	"fmt"
)
import "os"

import "io/ioutil"

func main() {
	//	bytes, err := readBytes()
	//	if err != nil {
	//		log.Println("error")
	//		os.Exit(1)
	//	}
	//	log.Println(err, string(bytes))

	indata := readByLine()
	for i := 0; i < len(indata); i++ {
		fmt.Println(indata[i])
	}

}

func readBytes() ([]byte, error) {
	data, err := ioutil.ReadAll(os.Stdin)
	return data, err
}

func readByLine() []string {
	var file []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		file = append(file, s.Text())
	}
	return file
}
