package main

import (
	"fmt"
	"reflect"

	stuff "example/project/mypackage"
)


func main(){


	fmt.Println("Hello",stuff.Name)


	intArr:= []int{2,3,5,7,11}
	
	
	strArr :=stuff.IntArrToStrArr(intArr)

	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))

}