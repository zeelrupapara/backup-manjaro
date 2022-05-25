package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tradeData struct {
	Period         string
	Data_value     string
	Series_title_2 string
}

func main() {
	fmt.Printf("Enter Year: ")
	var year string
	fmt.Scanf("%s", &year)
	csvFile, err := os.Open("trade.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	totalTrade := make(map[string]float64)
	for _, line := range csvLines {
		data := tradeData{
			Period:         line[1],
			Data_value:     line[2],
			Series_title_2: line[9],
		}
		cutMonth := strings.Split(data.Period, ".")
		if year == cutMonth[0] {
			data_value, _ := strconv.ParseFloat(data.Data_value, 64)
			totalTrade[data.Series_title_2] = totalTrade[data.Series_title_2] + data_value
		}
	}
	for key, value := range totalTrade {
		fmt.Println(key, ":", value, "cr")
	}
}
