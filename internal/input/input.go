package input

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TheMorgz/gobirth/internal/output"
	"github.com/TheMorgz/gobirth/internal/validator"
	"github.com/spf13/cobra"
)

type Reader struct{}

var printer output.Printer

type User struct {
	Surname             string `json:"0"`
	Name                string `json:"1"`
	DateOfBirth         string `json:"2"`
	validator.Validator `json:"-"`
}

func readData(path string, isCSV bool) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]string
	if !isCSV {
		data, _ = processJson(file)
	} else {
		data, _ = processCsv(file)
	}

	return data, nil
}

func processJson(file *os.File) ([][]string, error) {
	var userData [][]string

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&userData); err != nil {
		return nil, err
	}

	return userData, nil
}

func processCsv(file *os.File) ([][]string, error) {
	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (r Reader) ReadFile(cmd *cobra.Command, args []string) {
	csvFlag, _ := cmd.Flags().GetBool("csv")

	var path string
	if !csvFlag {
		path = "./static/files/input.json"
	} else {
		path = "./static/files/input.csv"
	}

	userData, err := readData(path, csvFlag)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var users []User
	for _, ud := range userData {
		if csvFlag && ud[0] == "last_name" {
			continue
		}

		user := User{
			Surname:     strings.TrimSpace(ud[0]),
			Name:        strings.TrimSpace(ud[1]),
			DateOfBirth: strings.TrimSpace(ud[2]),
		}
		users = append(users, user)
	}

	var birthdays []string
	var leapYear string

	for record, u := range users {
		record++

		birthday, _ := time.Parse("2006/01/02", u.DateOfBirth)

		u.CheckInput(validator.NotBlank(u.Name), "Name", fmt.Sprintf("Cannot submit an empty Name. Please check record number: %d", record))
		u.CheckInput(validator.NotBlank(u.Surname), "Surname", fmt.Sprintf("Cannot submit an empty Surname. Please check record number: %d", record))
		u.CheckInput(validator.ValidDate(u.DateOfBirth), "DateOfBirth", fmt.Sprintf("Something went wrong with the DOB. Please check record number: %d", record))

		if len(u.InputErrors) > 0 {
			fmt.Println("There is an issue with the input file. Please see the error(s) below:")

			for _, input := range u.InputErrors {
				printer.Output(input)
				return
			}
		}

		if u.IsBirthday(birthday) {
			leapYear = u.IsFebTwentyNine(birthday, birthday.Year(), time.Now().Year())
			birthdays = append(birthdays, fmt.Sprintf("Happy Birthday %s %s! %s", u.Name, u.Surname, leapYear))
		}
	}

	printer.OutputBirthdays(birthdays)
}
