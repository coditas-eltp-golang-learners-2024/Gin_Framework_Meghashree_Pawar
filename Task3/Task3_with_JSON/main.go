package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Record struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	records, err := parseJSON(`C:\Users\Coditas-Admin\Desktop\hhh\task\data.json`)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	insertChan := make(chan Record)
	updateChan := make(chan Record)
	deleteChan := make(chan string)
	doneChan := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(3)

	go insertRecords(insertChan, &wg)
	go updateRecords(updateChan, &wg)
	go deleteRecords(deleteChan, &wg)

	for _, record := range records {
		insertChan <- record
		updateChan <- record
		deleteChan <- record.ID
	}

	close(insertChan)
	close(updateChan)
	close(deleteChan)

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	<-doneChan

	if err := writeCSV("data.csv", records); err != nil {
		fmt.Println("Error writing to CSV:", err)
	}
}

func parseJSON(filename string) ([]Record, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var records []Record
	err = json.NewDecoder(file).Decode(&records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func insertRecords(insertChan <-chan Record, wg *sync.WaitGroup) {
	defer wg.Done()

	for record := range insertChan {
		fmt.Println("Inserting record:", record)
	}
}

func updateRecords(updateChan <-chan Record, wg *sync.WaitGroup) {
	defer wg.Done()

	for record := range updateChan {
		fmt.Println("Updating record:", record)
	}
}

func deleteRecords(deleteChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for id := range deleteChan {
		fmt.Println("Deleting record with ID:", id)
	}
}

func writeCSV(filename string, records []Record) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	
	header := []string{"ID", "Name", "Age"}
	if err := writer.Write(header); err != nil {
		return err
	}

	
	for _, record := range records {
		row := []string{record.ID, record.Name, strconv.Itoa(record.Age)}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
