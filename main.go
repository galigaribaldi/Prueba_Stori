package main

import (
	"fmt"

	controller "Prueba_Stori/controller"
)

func main() {

	df := controller.ReadFile("files/file.csv")
	for i := 0; i < df.Nrow(); i++ {
		d := df.Subset(i)
		t := controller.ConvertDFtoStruct(d)
		fmt.Println(t)
	}
}
