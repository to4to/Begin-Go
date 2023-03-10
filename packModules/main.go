package main

// import (
// 	"fmt"
// 	"reflect"

// 	stuff "example/project/mypackage"
// )

// func main(){

// 	fmt.Println("Hello",stuff.Name)

// 	intArr:= []int{2,3,5,7,11}

// 	strArr :=stuff.IntArrToStrArr(intArr)

// 	fmt.Println(strArr)
// 	fmt.Println(reflect.TypeOf(strArr))

// }

//Using getter setter from mypackage.go

//Encapsulation is also demonstrated in it

import (
	stuff "example/project/mypackage"
	"fmt"
	"log"
)

func main(){
date:=stuff.Date{}

err:=date.SetMonth(03)
if err!=nil{
log.Fatal(err)

}

err=date.SetDay(10)
if err!=nil{
log.Fatal(err)
}


err=date.SetYear(2023)
if err!=nil{
log.Fatal(err)

}


fmt.Printf("Today's Data: %d/%d/%d/",date.Day(),date.Month(),date.Year())
}

