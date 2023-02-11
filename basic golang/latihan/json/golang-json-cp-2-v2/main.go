package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type LoanData struct {
	StartBalance int
	Data         []Loan
	Employees    []Employee
}

type Loan struct {
	Date        string
	EmployeeIDs []string
}

type Employee struct {
	ID       string
	Name     string
	Position string
}

// json structure
type LoanRecord struct {
	MonthDate string `json:"month_date"`
	StartBalance int `json:"start_balance"`
	EndBalance int `json:"end_balance"`
	Borrowers []Borrower `json:"borrowers"`
}

type Borrower struct {
	Id string `json:"id"`
	TotalLoan int `json:"total_loan"`
}


func LoanReport(data LoanData) LoanRecord {
	var borrower = []Borrower{}
	var asset = data.StartBalance 
	var get bool
	var split []string
	for _, loan := range data.Data {
		split = strings.Split(loan.Date, "-")

		for _, employeeID := range loan.EmployeeIDs{
			if asset <= 0 {
				break
			}
			
			get = false

			for i, borrowerItem := range borrower{
				if employeeID == borrowerItem.Id {
					asset -= 50000
					borrower[i].TotalLoan += 50000
					get = true
				}

			}
			
			if !get {
				b := Borrower{
					Id: employeeID,
					TotalLoan: 50000,
				}
				borrower = append(borrower, b)
				asset -= 50000
			}
		}
	}

	loanRecord := LoanRecord{
		MonthDate: fmt.Sprintf("%s %s", split[1], split[2]),
		StartBalance: data.StartBalance,
		EndBalance: asset,
		Borrowers: borrower,
	}
	
	return loanRecord // TODO: replace this
}

func RecordJSON(record LoanRecord, path string) error {
	data, err := json.Marshal(record)
	if err !=nil {
		return err
	}

	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	
	return nil // TODO: replace this
}

// gunakan untuk debug
func main() {
	records := LoanReport(LoanData{
		StartBalance: 500000,
		Data: []Loan{
			{"01-January-2021", []string{"1", "2"}},
			{"02-January-2021", []string{"1", "2", "3"}},
			{"03-January-2021", []string{"2", "3"}},
			{"04-January-2021", []string{"1", "3"}},
		},
		Employees: []Employee{
			{"1", "Employee A", "Manager"},
			{"2", "Employee B", "Staff"},
			{"3", "Employee C", "Staff"},
		},
	})

	err := RecordJSON(records, "loan-records.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)
}
