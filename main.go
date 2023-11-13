package gosv

import (
	"os"
	"log"
)

type DataFrame struct {
	seperator string
	filepath string
	mat [][]float64
}

func Read_csv(path string) DataFrame{
	// read_csv reads the csv file and returns a dataframe
	file, err := os.Open(path)
  	
	if err != nil { log.Println("Failed Opening CSV file") }

	defer file.Close()

	log.Println("Successfully Opened CSV file")

	return DataFrame{filepath: path, seperator: ",", mat: [5][5]float64{}}
}