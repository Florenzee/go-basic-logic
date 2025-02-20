package main

import (
	"fmt"
	"go-basic-logic/logic1"
)

import (
	go_print_slice "github.com/Florenzee/go-print-slice"
	"go-basic-logic/logic2"
)

func main() {
	//logic 1 no.1
	fmt.Println("Logic 1 no.1")
	slice1_1 := logic1.Logic1AscStep(10, 1, 2)
	go_print_slice.PrintSlice(slice1_1)
	fmt.Println()

	//logic 2 no.1
	fmt.Println("Logic 2 no.1")
	slice2_1 := logic2.Matriks1(9)
	go_print_slice.PrintSlice2(slice2_1)
	fmt.Println()

	//logic 2 no.2
	fmt.Println("Logic 2 no.2")
	slice2_2 := logic2.Matriks2(9)
	go_print_slice.PrintSlice2(slice2_2)
	fmt.Println()

	//logic 2 no.6
	fmt.Println("Logic 2 no.6")
	slice2_6 := logic2.Matriks6(9)
	go_print_slice.PrintSlice2(slice2_6)
	fmt.Println()

	//logic 2 no.8
	fmt.Println("Logic 2 no.8")
	slice2_8 := logic2.Matriks8(9)
	go_print_slice.PrintSlice2(slice2_8)
	fmt.Println()

	//logic 2 no.9
	fmt.Println("Logic 2 no.9")
	slice2_9 := logic2.Matriks9(9)
	go_print_slice.PrintSlice2(slice2_9)
	fmt.Println()

	//logic 2 no.10
	fmt.Println("Logic 2 no.10")
	slice2_10 := logic2.Matriks10(9)
	go_print_slice.PrintSlice2(slice2_10)
	fmt.Println()

	//logic 2 no.11
	fmt.Println("Logic 2 no.11")
	slice2_11 := logic2.Matriks11(9)
	go_print_slice.PrintSlice2(slice2_11)
	fmt.Println()

	//logic 2 no.12
	fmt.Println("Logic 2 no.12")
	slice2_12 := logic2.Matriks12(9)
	go_print_slice.PrintSlice2(slice2_12)
	fmt.Println()

}
