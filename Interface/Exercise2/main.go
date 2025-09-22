// Section 09 - Lecture 03 : Method Sets - Final Reference
package main

import "fmt"

// تعريف struct
type Person struct {
	fname string
	lname string
	age   uint8
}

// ---------------- VALUE RECEIVERS ----------------

// Value receiver → بياخد نسخة (copy) من الـ struct
// ده معناه إن أي تعديل هنا مش هيأثر على النسخة الأصلية
func (p Person) Name() string {
	return fmt.Sprintf("%v, %v", p.lname, p.fname)
}

func (p Person) Age() uint8 {
	return p.age
}

// ---------------- POINTER RECEIVERS ----------------

// Pointer receiver → بياخد reference (عنوان في الذاكرة)
// أي تعديل هنا هيأثر على النسخة الأصلية
func (p *Person) SetAge(a uint8) {
	if a <= 150 && a > p.age {
		fmt.Printf("Changing age of %v from %v to %v\n", p.Name(), p.age, a)
		p.age = a
	}
}

func main() {
	// ---------------- USING VALUE ----------------
	// هنا mark متخزن كـ value (مش pointer)
	mark := Person{"Mark", "Smith", 35}

	// استدعاء method بـ value receiver ✅ (عادي جداً)
	fmt.Printf("Mark's name: %v\n", mark.Name())
	fmt.Printf("Mark's age: %v\n", mark.Age())

	// استدعاء method بـ pointer receiver ❗
	// Go هنا بيعمل أوتوماتيك تحويل: (&mark).SetAge(...)
	mark.SetAge(mark.Age() + 1)

	// هنلاحظ إن العمر اتغير في الـ object الأصلي
	fmt.Printf("Mark's age: %v\n\n", mark.Age())

	// ---------------- USING POINTER ----------------
	// هنا pPtr عبارة عن pointer
	pPtr := &Person{"Jane", "Doe", 28}

	// value receiver شغال على pointer برضه ✅
	fmt.Printf("Jane's name: %v\n", pPtr.Name())
	fmt.Printf("Jane's age: %v\n", pPtr.Age())

	// pointer receiver طبعاً شغال طبيعي ✅
	pPtr.SetAge(pPtr.Age() + 2)
	fmt.Printf("Jane's age: %v\n", pPtr.Age())
}
