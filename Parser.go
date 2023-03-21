package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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
		outputRecord := fmt.Sprintf("%-25s", record[0]) +
			fmt.Sprintf("%-5s", "") +
			fmt.Sprintf("%-15s", record[1]) +
			fmt.Sprintf("%-5s", "") +
			fmt.Sprintf("%-30s", record[2]) +
			fmt.Sprintf("%-5s", "") +
			fmt.Sprintf("%-15s", record[3]) +
			fmt.Sprintf("%-5s", "") +
			fmt.Sprintf("%-3s", record[4]) +
			fmt.Sprintf("%-5s", "") +
			fmt.Sprintf("%-10s", record[5]) +
			fmt.Sprintf("%-38s", "")

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
