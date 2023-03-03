package output

import (
	"fmt"
	"os"
	"os/exec"
)

type Printer struct{}

func (p Printer) Output(input string) {
	p.ClearScreen()

	fmt.Println()
	fmt.Println(input)
	fmt.Println()

	fmt.Println("Please correct the input file & try again.")
}

func (p Printer) OutputBirthdays(birthdays []string) {
	p.ClearScreen()

	fmt.Println()
	fmt.Println("Total Number of Birthdays Today:", len(birthdays))
	fmt.Println()

	for _, b := range birthdays {
		fmt.Println(b)
	}
}

func (p Printer) ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
