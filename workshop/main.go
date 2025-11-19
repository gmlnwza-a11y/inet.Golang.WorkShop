package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
	ex0()
	// ex1()
	// ex1_2()
	// ex2()
	// ex3()
	// ex3_1()
	// ex4()
	// ex4_1()
	ex5()
	// ex6()
	// exElse()
}

// work shop 0
func ex0() {
	i := 2
	// แบบ else if
	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("Other")
	}
	// แบบแยก if
	if i == 0 {
		fmt.Println("Zero")
	}

	if i == 1 {
		fmt.Println("One")
	}

	if i == 2 {
		fmt.Println("Two")
	}
	if i == 2 {
		fmt.Println("Three")
	}

}

// work shop 1
func ex1() {
	count := 0
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, " ")
			count++
		}
	}
	fmt.Printf("\nมีตัวเลขที่หาร 3 ลงตัวทั้งหมด %d ตัว\n", count)
}

// work shop 1_2
func ex1_2() {
	pow := 2
	num := 20
	result := 1

	for i := 0; i < pow; i++ {
		result = result * num
	}
	fmt.Printf("%d^%d = %d\n", num, pow, result)
}

// work shop 2
func ex2() {
	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	max := x[0]
	min := x[0]

	for _, value := range x {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	fmt.Printf("ค่ามากที่สุด: %d\n", max)
	fmt.Printf("ค่าน้อยที่สุด: %d\n", min)
}

// work shop 3
func ex3() {
	count := 0
	for i := 1; i <= 1000; i++ {
		s := strconv.Itoa(i)
		for _, ch := range s {
			if ch == '9' {
				count++
			}
		}
	}
	fmt.Println("จำนวนตัวเลข 9 ที่ปรากฏในเลข 1 ถึง 1000 คือ", count)
}

// work shop 3_1
func ex3_1() {
	result := countNines(1000)
	fmt.Println("จำนวนตัวเลข 9 ที่ปรากฏในเลข 1 ถึง 10000 คือ", result)
}

func countNines(limit int) int {
	count := 0
	for i := 1; i <= limit; i++ {
		s := strconv.Itoa(i)
		for _, ch := range s {
			if ch == '9' {
				count++
			}
		}
	}
	return count
}

// work shop 4
func ex4() {
	myWords := "AW SOME GO!"
	removedSpaces := ""
	for _, ch := range myWords {
		if ch != ' ' {
			removedSpaces += string(ch)
		}
	}
	fmt.Printf("ข้อความต้นฉบับ: '%s'\n", myWords)
	fmt.Printf("ข้อความหลังตัดช่องว่าง: '%s'\n", removedSpaces)
}

// work shop 4_1
func ex4_1() {
	fmt.Println(cutText("ine t")) // Output: "ine"
}
func cutText(s string) string {
	result := " "
	for _, ch := range s {
		if ch != ' ' {
			result += string(ch)
		}
	}
	return result
}

// work shop 5
func ex5() {
	peoples := map[string]map[string]string{
		"emp_01": {
			"fname":   "Guy",
			"lname":   "DuDu",
			"age":     "30",
			"address": "Jakarta",
		},
		"emp_02": {
			"fname":   "Max",
			"lname":   "Verstappen",
			"age":     "27",
			"address": "Los Angeles",
		},
		"emp_03": {
			"fname":   "Will",
			"lname":   "Smithy",
			"age":     "50",
			"address": "Bangkok",
		},
	}
	for _, person := range peoples {
		fmt.Println("Name ", person["fname"], person["lname"], "(Age ", person["age"], ")")
		fmt.Println("Address -:", person["address"])
	}
}

// work shop 5
func ex6() {
	exStructCompany()
}

type Company struct {
	Name        string
	CEO         string
	Employees   int
	Location    string
	Established int
}

func exStructCompany() {
	// สร้าง object จาก struct
	company := Company{
		Name:        "TechSoft Co., Ltd.",
		CEO:         "Alice Johnson",
		Employees:   120,
		Location:    "Bangkok, Thailand",
		Established: 2010,
	}

	// แสดงผล
	fmt.Println("=== Company Information ===")
	fmt.Println("Name        :", company.Name)
	fmt.Println("CEO         :", company.CEO)
	fmt.Println("Employees   :", company.Employees)
	fmt.Println("Location    :", company.Location)
	fmt.Println("Established :", company.Established)
}

func exElse() {
	// Row แถว
	for i := 1; i <= 6; i++ {
		// Column คอลัมน์
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
