package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
	Bonus        Employee
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
	Bonus           Employee
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
	Bonus            Employee
}

// junior
func (j *Junior) GetBonus() float64 {
	prorata := float64(j.WorkingMonth) / 12
	bonus := float64(1*j.BaseSalary) * prorata
	return bonus
}

// senior
func (s *Senior) GetBonus() float64 {
	bonus := 2*float64(s.BaseSalary)*(float64(s.WorkingMonth)/12) + (s.PerformanceRate * float64(s.BaseSalary))
	return bonus
}

// manager
func (m *Manager) GetBonus() float64 {
	bonus := 2*float64(m.BaseSalary)*(float64(m.WorkingMonth)/12) + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
	return bonus
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()

}

func TotalEmployeeBonus(employees []Employee) float64 {
	res := 0.0
	for _, v := range employees {
		res += EmployeeBonus(v)
	}
	return res
}

func main() {

	var j = Junior{
		Name:         "fahmi",
		BaseSalary:   100000,
		WorkingMonth: 6,
		Bonus:        nil,
	}

	var s = Senior{
		Name:            "fahmi",
		BaseSalary:      100000,
		WorkingMonth:    12,
		Bonus:           nil,
		PerformanceRate: 0.5,
	}

	var m = Manager{
		Name:             "yulia",
		BaseSalary:       100000,
		WorkingMonth:     12,
		PerformanceRate:  0.5,
		BonusManagerRate: 0.5,
		Bonus:            nil,
	}

	// var employees = []Junior{
	// 	{
	// 		Name:         "A",
	// 		BaseSalary:   100000,
	// 		WorkingMonth: 12,
	// 		Bonus:        nil,
	// 	},
	// 	{
	// 		Name:         "B",
	// 		BaseSalary:   100000,
	// 		WorkingMonth: 12,
	// 		Bonus:        nil,
	// 	},
	// }

	var asd = []Employee{
		&j, &s, &m,
	}
	fmt.Println(TotalEmployeeBonus(asd))

	fmt.Println(j.GetBonus())
	fmt.Println(s.GetBonus())
	fmt.Println(m.GetBonus())
}