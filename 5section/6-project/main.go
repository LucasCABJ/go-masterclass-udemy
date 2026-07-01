package main

import "fmt"

type Payable interface {
	CalculatePay() float64
	fmt.Stringer
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (s SalariedEmployee) CalculatePay() float64 {
	return s.AnnualSalary / 12
}

func (s SalariedEmployee) String() string {
	return fmt.Sprintf("SalariedEmployee: %s (Annual Pay: %.2f)", s.Name, s.AnnualSalary)
}

type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64
}

func (h HourlyEmployee) CalculatePay() float64 {
	return h.HourlyRate * h.HoursWorked
}

func (h HourlyEmployee) String() string {
	return fmt.Sprintf("HourlyEmployee: %s (Hourly Rate: %.2f, Hours Worked: %.2f)", h.Name, h.HourlyRate, h.HoursWorked)
}

type CommissionEmployee struct {
	Name           string
	BaseSalary     float64
	CommissionRate float64
	SalesAmount    float64
}

func (ce CommissionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommissionRate * ce.SalesAmount)
}

func (ce CommissionEmployee) String() string {
	return fmt.Sprintf("CommissionesEmployee: %s (Base salary: %.2f, CommissionRate: %.2f, SalesAmount: %.2f)",
		ce.Name, ce.BaseSalary, ce.CommissionRate, ce.SalesAmount)
}

func PrintEmployeeSummary[P fmt.Stringer](employee P) {
	fmt.Printf(" - Processing: %s\n", employee)
}

func ProcessPayroll(employess []Payable) {
	fmt.Println("\n---- Processing Payroll ----")
	totalPayroll := 0.0
	for _, emp := range employess {
		PrintEmployeeSummary(emp)
		pay := emp.CalculatePay()
		fmt.Printf("Monthly Pay: %.2f\n", pay)
		totalPayroll += pay
	}
	fmt.Printf("\nTotal Monthly Payroll: %.2f\n", totalPayroll)
	fmt.Println()
}

func main() {
	saEmp := SalariedEmployee{
		Name:         "John Doe",
		AnnualSalary: 120000,
	}

	hourlyEmp := HourlyEmployee{
		Name:        "Jane Smith",
		HourlyRate:  25,
		HoursWorked: 160,
	}

	commEmp := CommissionEmployee{
		Name:           "Bob Johnson",
		BaseSalary:     3000,
		CommissionRate: 0.1,
		SalesAmount:    50000,
	}

	employees := []Payable{saEmp, hourlyEmp, commEmp}
	ProcessPayroll(employees)
}
