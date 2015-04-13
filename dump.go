package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func dumpToCSV(user user) {
	csvfile, err := os.OpenFile("output.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	record := []string{user.name, user.email, user.url, user.username}
	err = writer.Write(record)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	writer.Flush()
}
