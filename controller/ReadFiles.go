package controller

import (
	"log"
	"os"
	"strconv"
	"strings"

	models "Prueba_Stori/models"

	"github.com/go-gota/gota/dataframe"
)

func ReadFile(route string) dataframe.DataFrame {
	file, err := os.Open("files/file.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	df := dataframe.ReadCSV(file)
	return df
}

func ConvertDFtoStruct(df dataframe.DataFrame) models.Transaction {
	ID, err := strconv.Atoi(df.Select("Id").Records()[1][0])
	if err != nil {
		log.Fatal(err)
	}
	Date := df.Select("Date").Records()[1][0]
	//log.Println("Date: ", Date)

	Month, err := strconv.Atoi(strings.Split(Date, "/")[1])
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("Month: ", Month)
	Day, err := strconv.Atoi(strings.Split(Date, "/")[0])
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("Day:", Day)
	Transact, err := strconv.ParseFloat(df.Select("Transaction").Records()[1][0], 32)
	if err != nil {
		log.Fatal(err)
	}
	TypeCredit := true
	if Transact > 0 {
		TypeCredit = true
	} else {
		TypeCredit = false
	}
	t := models.Transaction{
		ID:         ID,
		Month:      Month,
		Day:        Day,
		Date:       Date,
		Transact:   Transact,
		TypeCredit: TypeCredit,
	}
	return t
}
