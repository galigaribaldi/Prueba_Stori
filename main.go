package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hola")
	file, err := os.Open("file.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	fmt.Println(reader)
	/*
		for {
			record, e := reader.Read()
			if e != nil {
				fmt.Println(e)
				break
			}
			fmt.Println(record)
		}
	*/
}
