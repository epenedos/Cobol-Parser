package main

import (
	"encoding/csv"
	"log"
	"os"
	"unicode/utf8"
)

func convertCSVToText(inputFile string, outputFile string) error {
	// Open input file
	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer input.Close()

	// Open output file
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	// Create a new CSV reader
	reader := csv.NewReader(input)

	// Read and process each line of the input file
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		// Build the output record
		outputRecord := SetStringToSize(record[0], 25) +
			SetFillerToSize(5) +
			SetStringToSize(record[1], 15) +
			SetFillerToSize(5) +
			SetStringToSize(record[2], 30) +
			SetFillerToSize(5) +
			SetStringToSize(record[3], 15) +
			SetFillerToSize(5) +
			SetStringToSize(record[4], 3) +
			SetFillerToSize(5) +
			SetStringToSize(record[5], 15)

		// Write the output record to the output file
		_, err = output.WriteString(outputRecord + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := convertCSVToText("./info.csv", "./output.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func SetStringToSize(input string, len int) string {
	bytes := make([]byte, len)
	len_input := utf8.RuneCountInString(input)
	for i := 0; i < len; i++ {
		if len_input > i {
			bytes[i] = input[i]
		} else {
			bytes[i] = byte(32)
		}
	}
	return string(bytes)
}

func SetFillerToSize(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(32)
	}
	return string(bytes)
}
