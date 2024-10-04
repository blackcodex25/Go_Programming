package main

import "fmt"

/* ในภาษา Go ประเภทข้อมูล Boolean แทนค่าความจริงเชิงตรรกะ โดยมี
เพียงสองค่าคือ true หรือ false ใช้คีย์เวิร์ด bool ในการประกาศตัวแปร
ประเภทนี้ เช่น */

func main() {
	// ประกาศตัวแปร boolTrue และ boolFalse เป็นประเภท Boolean
	var boolTrue bool = true
	var boolFalse bool = false

	// แสดงค่า Boolean ทั้งสองตัว
	fmt.Println("The boolean values are", boolTrue, "and", boolFalse)
}

/* ในตัวอย่างนี้ boolTrue ถูกกำหนดค่าเป็น true */
/* และ boolFalse ถูกกำหนดเป็น false */
