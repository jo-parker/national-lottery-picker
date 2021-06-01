package main

import (
	"fmt"
	"github.com/jpparker/euromillions-picker/internal/pkg/service"
)

func main() {
	service.PlayEuromillions()
	fmt.Println("Done generating first")
	ticket := service.GenerateTicket()
	fmt.Println("Done generating second")
	mainNumbers := make([]int, 0, len(ticket.MainNumbers))
	specialNumbers := make([]int, 0, len(ticket.SpecialNumbers))

	for key := range ticket.MainNumbers {
		mainNumbers = append(mainNumbers, key)
	}

	for key := range ticket.SpecialNumbers {
		specialNumbers = append(specialNumbers, key)
	}

	fmt.Println(mainNumbers)
	fmt.Println(specialNumbers)
}