package main

import "fmt"

// การเปลี่ยนค่าของตัวแปร
// ใน Go เราสามารถเปลี่ยนค่าที่เก็บในตัวแปรได้ โดยใช้โค้ดดังนี้
func main() {
	// ค่าต้นฉบับ
	number := 10
	fmt.Println("Before change:", number) // แสดงผล 10

	// เปลี่ยนค่าตัวแปร
	number = 100                         // Assign 100 to number
	fmt.Println("After change:", number) // แสดงผล 100
}

// ในตัวอย่างนี้ ค่าเริ่มต้น 10 ถูกเก็บในตัวแปร number และหลังจากนั้น
// ค่าของมันถูกเปลี่ยนเป็น 100

/* หมายเหตุ */
/* ใน Go เราไม่สามารถเปลี่ยนประเภทของตัวแปรหลังจากที่มันถูกประกาศแล้ว */
/* ในตัวอย่างข้างต้น ตัวแปร number สามารถเก็บค่าเป็นจำนวนเต็มได้เท่านั้น */
/* และไม่สามารถใช้เพื่อเก็บข้อมูลประเภทอื่นได้ เช่น */
/*
number := 10

เกิดข้อผิดพลาด: กำหนดข้อมูลประเภทสติง
number = "Hello" // จะเกิดข้อผิดพลาด */
