package main

import (
	"fmt"
	"time"
)

func addOneDay(dateStr string) (string, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", err
	}

	newDate := date.AddDate(0, 0, 1)
	newDateStr := newDate.Format("2006-01-02")
	if newDate.After(time.Now()) {
        return "", fmt.Errorf("incremented date is greater than current date")
    }
	return newDateStr, nil
}
func main() {
	newdate, err := addOneDay("2023-03-22")
	fmt.Println("new date", newdate)
	if err != nil {
		fmt.Println("err", err)

	}

}
