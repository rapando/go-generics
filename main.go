// Go Generics
// Rapando 2024.08.14

// If you want to create a function that takes any type of parameters, you can utilize interface{} (any) data type
// for example

package main

import (
	"fmt"
	"reflect"
)

func main() {
	PrintSomeValue(4)
	PrintSomeValue("James Doe")
	PrintSomeValue(false)

	var xInt int = 16
	var xInt32 int32 = 32
	var xInt64 int64 = 64
	var xFloat32 float32 = 100.56

	var _xInt = toFloat64(xInt)
	var _xInt32 = toFloat64(xInt32)
	var _xInt64 = toFloat64(xInt64)
	var _xFloat32 = toFloat64(xFloat32)

	// let's print out the data types of all the variables above
	fmt.Printf("xInt     : [%v],\t _xInt\t\t: [%v]\n", reflect.TypeOf(xInt), reflect.TypeOf(_xInt))
	fmt.Printf("xInt32   : [%v],\t _xInt64\t: [%v]\n", reflect.TypeOf(xInt32), reflect.TypeOf(_xInt32))
	fmt.Printf("xInt64   : [%v],\t _xInt32\t: [%v]\n", reflect.TypeOf(xInt64), reflect.TypeOf(_xInt64))
	fmt.Printf("xFloat32 : [%v],\t _xFloat32\t: [%v]\n", reflect.TypeOf(xFloat32), reflect.TypeOf(_xFloat32))

	// if you try to pass any other type other than the ones provided, your code will fail.
	// to try, uncomment the following 2 lines.

	// var xStr string = "Monday"
	// var _xStr = toFloat64(xStr)

}

// PrintSomeValue is a function that takes data of any type and prints it
func PrintSomeValue(x any) {
	fmt.Printf("We are printing: [%v]\n", x)
}

// say we want to create a function that converts a number into float64.
// this is a simple usecase, however, think of it as, we want to create a function that could
// allow us to pass different data types of data but limited to a few.
// in this case, we only want to take in int, int32, int64 and float32
// Enter, Go Generics

// NumericDataTypes contains the acceptable data types
// 1. Declare the acceptable data types.
type NumericDataTypes interface {
	~int | ~int32 | ~int64 | ~float32
}

// toFloat64 shows us how to declare a generic
// 2. We declare the function that will accept the input
// format is
// func functionName[N InterfaceWithTheType](x N)
// here, N is just a variable representing the NumericDataTypes interface
// therefore, when we pass x of type N,
// we're basically saying x can be of any data type as long as its declared in the NumericDataTypes interface
func toFloat64[N NumericDataTypes](x N) float64 {
	// this is a simple case, because we already know that the inputs will be numeric
	// and our custom type is also numeric, we can do a direct casting.
	return float64(x)
}
