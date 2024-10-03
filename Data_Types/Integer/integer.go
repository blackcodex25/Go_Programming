package main

import "fmt"

/* ตัวอย่างที่ 1: ทำความเข้าใจประเภทข้อมูลจำนวนเต็มใน Go */
func main() {
	// ประกาศตัวแปรสองตัว integer1, integer2 ด้วยชนิดข้อมูล int
	var integer1 int
	var integer2 int

	integer1 = 5  // กำหนดค่า 5 ให้ตัวแปร integer1
	integer2 = 10 // กำหนดค่า 10 ให้ตัวแปร integer2

	fmt.Println(integer1) // แสดงค่า integer1 = 5
	fmt.Print(integer2)   // แสดงค่า integer2 = 10
}

/* ในตัวอย่างนี้ เราประกาศตัวแปรจำนวนเต็ม integer1 และ integer2 */
/* จากนั้นกำหนดค่าให้เป็น 5 และ 10 ตามลำดับ และใช้คำสั่ง fmt.Println */
/* และ fmt.Print เพื่อแสดงผลค่าของตัวแปรเหล่านั้นในคอนโซล */
