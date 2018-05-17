package main

import (
	"fmt"
	"unsafe"

	"github.com/jinzhu/gorm"
)

type Struct struct {
	A int
	B float64
}

type GormModel struct {
	gorm.Model

	A int
	B float64
}

func main() {
	s := Struct{A: 1, B: 2.0}
	fmt.Println(unsafe.Sizeof(s)) // 16

	m := GormModel{A: 1, B: 2.0}
	fmt.Println(unsafe.Sizeof(m)) // 80
}
