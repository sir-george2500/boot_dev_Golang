package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	// Get the message content
	content := msg.getMessage()
	// Calculate the cost of the message
	cost := len(content) * 3
	return content, cost
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main() {
	birthday := birthdayMessage{
		birthdayTime:  time.Now(),
		recipientName: "John",
	}
	report := sendingReport{
		reportName:    "Monthly Summary",
		numberOfSends: 50,
	}

	birthdayContent, birthdayCost := sendMessage(birthday)
	fmt.Printf("Birthday Message: %s\nCost: %d\n", birthdayContent, birthdayCost)

	reportContent, reportCost := sendMessage(report)
	fmt.Printf("Report Message: %s\nCost: %d\n", reportContent, reportCost)
}

