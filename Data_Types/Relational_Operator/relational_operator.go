package main

import "fmt"

/* ตัวอย่าง การใช้ตัวดำเนินการเชิงสัมพันธ์ใน Go */
/* โปรแกรมแสดงการทำงานของตัวดำเนินการเชิงสัมพันธ์ */
func main() {
	number1 := 12   // กำหนดค่า 12 ให้กับตัวแปร number1
	number2 := 20   // กำหนดค่า 20 ให้กับตัวแปร number2
	var result bool // กำหนดตัวแปร result ให้เป็นตัวแปรประเภท Boolean

	// ตัวดำเนินการเท่ากับ
	result = (number1 == number2)
	fmt.Printf("%d == %d ให้ผลลัพธ์ %t \n", number1, number2, result)

	// ตัวดำเนินการไม่เท่ากับ
	result = (number1 != number2)
	fmt.Printf("%d != %d ให้ผลลัพธ์ %t \n", number1, number2, result)

	// ตัวดำเนินการมากกว่า
	result = (number1 > number2)
	fmt.Printf("%d > %d ให้ผลลัพธ์ %t \n", number1, number2, result)

	// ตัวดำเนินการน้อยกว่า
	result = (number1 < number2)
	fmt.Printf("%d < %d ให้ผลลัพธ์ %t \n", number1, number2, result)
}
