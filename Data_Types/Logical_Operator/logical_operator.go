package main

import "fmt"

/* ตัวดำเนินการเชิงตรรกะ */

func main() {
	number1 := 6  // กำหนดค่า number1 เป็น 6
	number2 := 12 // กำหนดค่า number2 เป็น 12
	number3 := 6  // กำหนดค่า number3 เป็น 6

	// AND operator
	result := (number1 > number2) && (number1 == number3) // ตรวจสอบว่า number1 เท่ากับ number2 และ number1 เท่ากับ number3
	fmt.Printf("Result of AND operator is %t \n", result)

	// OR operator
	result = (number1 > number2) || (number1 == number3) // ตรวจสอบว่า number1 เท่ากับ number2 หรือ number1 เท่ากับ number3
	fmt.Printf("Result of OR operator is %t \n", result)

	// NOT operator
	result = !(number1 == number3)                        // ตรวจสอบว่า number1 ไม่เท่ากับ number3
	fmt.Printf("Result of NOT operator is %t \n", result) // แสดงผลลัพธ์ของการตรวจสอบ NOT
}

/* ในตัวอย่างนี้ result = !(number1 == number3) จะตรวจสอบว่า number1 ไม่เท่ากับ number3
โดยใช้ตัวดำเนินการ NOT (!) เนื่องจาก number1 มีค่าเท่ากับ 6 และ number3 ก็มีค่าเท่ากับ 6
การเปรียบเทียบ number1 เท่ากับ number3 จะคืนค่าเป็น true ดังนั้น การใช้ NOT จะทำให้ result
มีค่าเป็น false */
/* สรุปคือ */
/* number1 == number3 คืนค่าเป็น true */
/* !(number1 == number3) คืนค่าเป็น false */
