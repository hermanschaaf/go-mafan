package mafan

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Compress() {
	// compress the dictionary file

	content, err := ioutil.ReadFile("data/dict.txt.big")
	if err != nil {
		//Do something
		log.Panic("Dictionary file not found")
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
}
