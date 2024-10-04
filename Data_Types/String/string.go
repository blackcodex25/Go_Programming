package main

import "fmt"

func main() {
	// ประกาศตัวแปร message เป็น string
	var message string

	// กำหนดค่า "Welcome to Programiz" ให้กับตัวแปร message
	message = "Welcome to Programiz"

	// แสดงค่าของตัวแปร message
	fmt.Println(message)

	// กำหนดค่า "Hello, World!" ให้กับตัวแปร message โดยใช้ backtick
	message = `Hello, World!`

	// แสดงค่าของตัวแปร message
	fmt.Println(message)
}

/* คำอธิบาย: ในตัวอย่างนี้ มีการประกาศตัวแปรชนิดข้อมูล string ชื่อ message */
/* และกำหนดค่าเป็น "Welcome to Programiz" ซึ่งค่าใน */
/* ตัวแปรจะถูกแสดงผลด้วยฟังก์ชัน fmt.Println() */
/* ในภาษา Go ตัวแปรชนิดข้อมูล string ใช้เก็บข้อมูลข้อความ โดยใช้คำว่า */
/* string เพื่อประกาศตัวแปรชนิดนี้ */
