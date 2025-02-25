package main

import (
	"fmt"
	"go-basic-logic/logic1"
	"go-basic-logic/logic3"
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

	//logic 1 no.2
	fmt.Println("Logic 1 no.2")
	slice1_2 := logic1.Logic1AscStep(10, 2, 2)
	go_print_slice.PrintSlice(slice1_2)
	fmt.Println()

	//logic 1 no.3
	fmt.Println("Logic 1 no.3")
	slice1_3 := logic1.Logic1AscStep(10, 3, 3)
	go_print_slice.PrintSlice(slice1_3)
	fmt.Println()

	//logic 1 no.4
	fmt.Println("Logic 1 no.4")
	slice1_4 := logic1.Logic1DescStep(10, 19, 2)
	go_print_slice.PrintSlice(slice1_4)
	fmt.Println()

	//logic 1 no.5
	fmt.Println("Logic 1 no.5")
	slice1_5 := logic1.Logic1DescStep(10, 20, 2)
	go_print_slice.PrintSlice(slice1_5)
	fmt.Println()

	//logic 1 no.6
	fmt.Println("Logic 1 no.6")
	slice1_6 := logic1.Logic1DescStep(10, 30, 3)
	go_print_slice.PrintSlice(slice1_6)
	fmt.Println()

	//logic 1 no.7
	fmt.Println("Logic 1 no.7")
	slice1_7_1 := logic1.MatriksPattern(10, 1, 2)
	slice1_7_2 := logic1.MatriksPattern(11, 1, 2)
	go_print_slice.PrintSlice(slice1_7_1)
	go_print_slice.PrintSlice(slice1_7_2)
	fmt.Println()

	//logic 1 no.8
	fmt.Println("Logic 1 no.8")
	slice1_8_1 := logic1.MatriksPattern(10, 2, 2)
	slice1_8_2 := logic1.MatriksPattern(11, 2, 2)
	go_print_slice.PrintSlice(slice1_8_1)
	go_print_slice.PrintSlice(slice1_8_2)
	fmt.Println()

	//logic 1 no.9
	fmt.Println("Logic 1 no.9")
	slice1_9_1 := logic1.MatriksPattern(10, 3, 3)
	slice1_9_2 := logic1.MatriksPattern(11, 3, 3)
	go_print_slice.PrintSlice(slice1_9_1)
	go_print_slice.PrintSlice(slice1_9_2)
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

	//logic 2 no.3
	fmt.Println("Logic 2 no.3")
	slice2_3 := logic2.Matriks3(9)
	go_print_slice.PrintSlice2(slice2_3)
	fmt.Println()

	//logic 2 no.4
	fmt.Println("Logic 2 no.4")
	slice2_4 := logic2.Matriks4(9)
	go_print_slice.PrintSlice2(slice2_4)
	fmt.Println()

	//logic 2 no.5
	fmt.Println("Logic 2 no.5")
	slice2_5 := logic2.Matriks5(9)
	go_print_slice.PrintSlice2(slice2_5)
	fmt.Println()

	//logic 2 no.6
	fmt.Println("Logic 2 no.6")
	slice2_6 := logic2.Matriks6(9)
	go_print_slice.PrintSlice2(slice2_6)
	fmt.Println()

	//logic 2 no.7
	fmt.Println("Logic 2 no.7")
	slice2_7 := logic2.Matriks7(9)
	go_print_slice.PrintSlice2(slice2_7)
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

	//logic 2 no.13
	fmt.Println("Logic 2 no.13")
	slice2_13 := logic2.Matriks13(9)
	go_print_slice.PrintSlice2(slice2_13)
	fmt.Println()

	//logic 3 no.1
	fmt.Println("Logic 3 no.1")
	slice3_1 := logic3.Matriks1(9)
	go_print_slice.PrintSlice2(slice3_1)
	fmt.Println()

	//logic 3 no.2
	fmt.Println("Logic 3 no.2")
	slice3_2 := logic3.Matriks2(9)
	go_print_slice.PrintSlice2(slice3_2)
	fmt.Println()

	//logic 3 no.3
	fmt.Println("Logic 3 no.3")
	slice3_3 := logic3.Matriks3(9)
	go_print_slice.PrintSlice2(slice3_3)
	fmt.Println()

	//logic 3 no.4
	fmt.Println("Logic 3 no.4")
	slice3_4 := logic3.Matriks4(9)
	go_print_slice.PrintSlice2(slice3_4)
	fmt.Println()

	//logic 3 no.7
	fmt.Println("Logic 3 no.7")
	slice3_7 := logic3.Matriks7(9)
	go_print_slice.PrintSlice2(slice3_7)
	fmt.Println()

	//logic 3 no.8
	fmt.Println("Logic 3 no.8")
	slice3_8 := logic3.Matriks8(9)
	go_print_slice.PrintSlice2(slice3_8)
	fmt.Println()

	//logic 3 no.14
	fmt.Println("Logic 3 no.14")
	slice3_14 := logic3.Matriks14(9)
	go_print_slice.PrintSlice2(slice3_14)
	fmt.Println()

}
