package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	in, err := ioutil.ReadFile("exemplo.txt")

	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(strings.NewReader(string(in)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record[0], record[1])
	}

}
