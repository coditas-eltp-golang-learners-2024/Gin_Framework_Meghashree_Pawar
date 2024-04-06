package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

type Record struct {
	ID   string
	Name string
	Age  int 
}

var (
	data    []Record
	dataMap map[string]int
	mutex   sync.Mutex
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 1000; i++ {
		data = append(data, Record{
			ID:   strconv.Itoa(i + 1),
			Name: randomName(),
			Age:  rand.Intn(50) + 20,
		})
	}

	dataMap = make(map[string]int)
	for i, record := range data {
		dataMap[record.ID] = i
	}

	insertRecord(Record{"1001", "Meghashree Pawar", 45})
	updateRecord("2", Record{"2", "Updated Jayesh", 30})
	deleteRecord("3")

	if err := writeToFile("data.csv"); err != nil {
		fmt.Println("Error writing to file:", err)
	}

	if err := readFromFile("data.csv"); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}

func randomName() string {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Henry", "Ivy", "Jack"}
	return names[rand.Intn(len(names))]
}

func insertRecord(record Record) {
	mutex.Lock()
	defer mutex.Unlock()

	data = append(data, record)
	dataMap[record.ID] = len(data) - 1

	fmt.Printf("Inserted record: %+v\n", record)
}

func updateRecord(id string, newRecord Record) {
	mutex.Lock()
	defer mutex.Unlock()

	index, ok := dataMap[id]
	if !ok {
		fmt.Println("Record not found")
		return
	}

	data[index] = newRecord

	fmt.Printf("Updated record with ID %s\n", id)
}

func deleteRecord(id string) {
	mutex.Lock()
	defer mutex.Unlock()

	index, ok := dataMap[id]
	if !ok {
		fmt.Println("Record not found")
		return
	}

	data = append(data[:index], data[index+1:]...)
	delete(dataMap, id)

	fmt.Printf("Deleted record with ID %s\n", id)
}

func writeToFile(filename string) error {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		err := writer.Write([]string{record.ID, record.Name, strconv.Itoa(record.Age)})
		if err != nil {
			return err
		}
	}

	fmt.Println("Data written to file successfully")
	return nil
}

func readFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	fmt.Println("Data read from file:")
	for _, record := range records {
		fmt.Println(record)
	}
	
	return nil
}
