package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	latitude := 55.6167 // Пример широты Видное
	fmt.Printf("Широта (десятичный): %f\n", latitude)

	// Преобразование в IEEE 754 (float64)
	// bits := math.Float64bits(latitude)
	// fmt.Printf("Широта (IEEE 754, 64 бита): %064b\n", bits)

	// Преобразование в байты для хранения или передачи
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, latitude)
	byteRepresentation := buf.Bytes()
	fmt.Printf("Широта (IEEE 754, байты): %v\n", byteRepresentation)

	//Обратное преобразование (из байт в float64)
	var restoredLatitude float64
	buf2 := bytes.NewBuffer(byteRepresentation)
	binary.Read(buf2, binary.LittleEndian, &restoredLatitude)

	fmt.Printf("Восстановленная широта (float64): %f\n", restoredLatitude)

	lat32 := []byte{0x82, 0x8B, 0x5E, 0x42} //{0x95, 0x8B, 0x5E, 0x42}
	buf32 := bytes.NewBuffer(lat32)
	var restLat32 float32
	binary.Read(buf32, binary.LittleEndian, &restLat32)
	fmt.Printf("Восстановленная широта (float32): %f\n", restLat32)

	lon32 := []byte{0x8F, 0xD5, 0x14, 0x42}
	lbuf32 := bytes.NewBuffer(lon32)
	var restLon32 float32
	binary.Read(lbuf32, binary.LittleEndian, &restLon32)
	fmt.Printf("Восстановленная долгота (float32): %f\n", restLon32)

	fmt.Printf(" ")

	p := Person{}

	p.SetName("Ivan")
	p.SetAge(30)
	fmt.Println(p.String())
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)

}

type Person struct {
	Name string
	Age  int
}

// String returns a string representation of the Person struct
// This method implements the Stringer interface from the fmt package
func (p *Person) String() string {
	// Use fmt.Sprintf to format the person's name and age into a string
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p *Person) SetName(name string) {
	p.Name = name
}

// SetAge is a method that sets the age of a Person
// It takes an integer parameter representing the age
// and updates the Age field of the Person instance
func (p *Person) SetAge(age int) {
	p.Age = age // Assign the provided age to the Person's Age field
}
