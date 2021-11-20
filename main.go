package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	wordPtr := flag.String("test", "exemplo.txt", "Please specify the path of the test you want to take")

	in, err := ioutil.ReadFile(*wordPtr)

	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(strings.NewReader(string(in)))
	correct, total := 0, 0

	for {
		var response string
		record, err := r.Read()
		if err == io.EOF {
			fmt.Println("Correct answers: ", correct, "/", total)
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\nResponse: ", record[0])
		fmt.Scanln(&response)
		if response == record[1] {
			correct += 1
		}
		total += 1
	}

}
