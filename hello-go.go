package main

import (
	//"bufio"
	"fmt"
	//"math"
	//"math/rand"
	//"time"
	//"strings"
	//"unicode/utf8"
	//"strconv"
	//"reflect"
	//"log"
	//"os"
)


var pl=fmt.Println;

//func main(){

    //  //Part 1
	// pl("What is Your Name");
	
	// reader :=bufio.NewReader(os.Stdin);
	// name ,err :=reader.ReadString('\n');

	// if err==nil{
    //   pl("Hello ", name)


	// }else
	// {
	// 	log.Fatal(err);
	// }




	//Part 2 Variables


// var vName string="Himanshu"
// var v1,v2=1.2,1.3

// var v3="Hello"

// v4:=2.8




//Part 3 Data Type
// int, float64 ,bool,string,rune 
//Default type 0,0.0,false,""
// pl(reflect.TypeOf(25))
// pl(reflect.TypeOf(3.2))
// pl(reflect.TypeOf(false))
// pl(reflect.TypeOf("hello"))
// pl(reflect.TypeOf('m'))


//Part 3 Casting Strings

// cV1:= 1.5
// cV2:=int(cV1)
// pl(cV2)

// //String to int 
// cV3:="5000000"
// cV4,err:=strconv.Atoi(cV3)
// pl(cV4,err,reflect.TypeOf(cV4))


//int To String 


// cV5 :=500000
// cV6:=strconv.Itoa(cV5)
// pl(cV6)

//String to float

// cV7:="3.27"

// if cV8,err:=strconv.ParseFloat(cV7,64); err==nil{

// pl(cV8)

// }

//float to string

// cV9:= fmt.Sprintf("%f",3.14)

//}



//Part 2

// func main(){

// iAge:=8
// if(iAge>=1 && iAge<=18){
// pl("Your Birthday is important")
// } else if(iAge ==21 || iAge==50){
// pl("Important Birthday")
// } else if iAge>=65 {

// 	pl("Import ant Birthday")
// }else {
// 	pl("Not important")
// }



//Strings 

// func main(){

// sV1:=" A Word "

// replacer:=strings.NewReplacer("A" ,"Another")
// sV2:=replacer.Replace(sV1)

// pl(sV2)

// pl("Length " ,len(sV2))
// pl("Contains Another :" ,strings.Contains(sV2,"Another"))
// pl("o Index: ",strings.Index(sV2,"o"))
// pl("Replace :",strings.Replace(sV2,"o","0",-1))
// sV3:="\n Some Words\n"
// sV3=strings.TrimSpace(sV3)

// pl("Split :" ,strings.Split("a-b-c-d","-"))
// pl("Lower :" , strings.ToLower(sV2))
// pl("Upper :" , strings.ToUpper(sV2))

// pl("Prefix: ",strings.HasPrefix("tacocat","taco"))
// pl("Suffix: ",strings.HasSuffix("tacocat","cat"))

// }


//Runes are called characters and rune are character code

// func main(){

// rStr:="abcdefg"
// pl("Rune Count:" ,utf8.RuneCountInString(rStr))

// for i,runeVal:=range rStr{

// fmt.Printf("%d : %#U : %c \n",i,runeVal, runeVal)


// }


// }


//Time in go


// func main(){

// now:= time.Now();
// pl(now.Year(), now.Month() ,now.Day())
// pl(now.Hour(), now.Minute() , now.Second())

// }


//Math Functions
// func main(){


// // seedSec:=time.Now().Unix()
// // rand.Seed(seedSec)
// // randNum:=rand.Intn(50)+1
// // pl("Random Num: ", randNum)

// pl("Abs(-10) ",math.Abs(-10))
// pl("Pow(4,2) ",math.Pow(4,2))
// pl("Sqrt(16) ",math.Sqrt(16))
// pl("Cbrt(8) ",math.Cbrt(8))
// pl("Ceil(4.4) ",math.Ceil(4.4))
// pl("Floor(4.4) ",math.Floor(4.4))
// pl("Round(4.4) ",math.Round(4.4))
// pl("Log2(8) ",math.Log2(8))
// pl("Log10(100) ",math.Log10(100))
// //Get the log of e to the power of 2
// pl("Log(7.389) " ,math.Log(math.Exp(2)))
// pl("Max(5,4) ",math.Max(5,4))
// pl("Min(5,4) ",math.Min(5,4)) 

// //convert 90 degree into radians
// r90:=90*math.Pi/180

// //Radian into degree
// d90:=r90*(180/math.Pi)
// pl(d90)


// pl("Sin90 ",math.Sin(90))
// //Thre is also function for Cos,Tan,Acos,Asin
// //Atan,Asinh,Acosh,Atanh,Atan2,Cosh,Sinh,Sincos
// //Htpot

// }


// %d=Integer
// %c=Character
// %f=Float
// %t=boolaen
// %s=String
// %o=Base 8
// %x=Base 16
// %v=Gusses Based on Data Type
// %T= Type of supplied Value
// func main(){


// 	// fmt.Printf("%s %d %c %f %t %o %x \n","Stuff", 1,'A' ,3.14,true,1,1)
// 	// fmt.Printf("%9f\n",3.14)
// 	// fmt.Printf("%.2f\n",3.1423232)
//     // fmt.Printf("%9.f\n",3.14242424)

// 	// sp1:=fmt.Sprintf("%9.f\n",3.14434343)
// 	// pl(sp1)



// //for Loops 
     
//     //for initilisation;condition ;postStatement 
// 	              //{Body}
// 				  for x:=1;x<=5;x++{
// 					pl(x)
// 				  }	
// }




//Range And Array

// func main(){

// // aNums:=[]int{1,2,3}

// // for _, num:= range aNums{

// // 	pl(num)
// // }


// // var arr1 [5]int;
// // arr1[0]=1

// // arr2:=[5]int{1,2,3,4,5}
// // pl("Index 0: ",arr2[0])
// // pl("Array Length: ",len(arr2))



// //iterate without range 
// // for i:=0;i<len(arr2);i++{

// // 	pl("Index :", i, arr2[i])
// // }


// // //iterate with range 


// // for i,val := range arr2{

// // 	fmt.Printf("%d : %d \n",i,val)
// // }


// // //MultiDimensonal Array
// // arr3:= [2][2]int{
// // {1,2},
// // {0,9},
// // }

// // for i:=0;i<2;i++{
// // for j:=0;j<2;j++{

// // 	pl(arr3[i][j])
// // }
// //}






// //Slices 

// aStr1:="abcde"
// rArr:=[]rune(aStr1)

// for _,v:= range rArr{

// 	fmt.Printf("Rune Array : %d\n",v)
// }


// byteArr:=[]byte{'a','b','c','c'}
// bStr:=string(byteArr[:])
// pl("I Am a String :",bStr)



// }






//Slices


// func main(){
// //general syntax: var name []datatype

// sl1:=make([]string,6)
// sl1[0]="Society "
// sl1[1]="Of "
// sl1[2]="the "
// sl1[3]="Simulated "
// sl1[4]="Universe "

// pl("Slice size :",len(sl1))

// for i:=0;i<len(sl1);i++{

// pl(sl1[i])

// }


// for _,v :=range sl1{

// 	pl(v)
// }



// sArr:=[]int{1,2,3,4,5}
// sl3:=sArr[0:2]
// pl("First 3 ",sArr[:3])
// pl("Last Three ",sArr[2:])
// sArr[0]=10
// pl("sl3 ",sl3)


// sl3 = append(sl3, 12)
// pl("sl3 ",sl3)
// pl("sArr ",sArr)

// } 

//Functions 

// func sayHello(){

// 	pl("Say Hello")
// }

// func getSum(x int,y int) int //What type is to be Returned
// {

// return x+y
// }

// func getTwo(x int) (int ,int){

// 	return x+1,x+2
// }

// func getQuotient(x float64,y float64) (ans float64,err error){

// if y==0{
// 	return 0,fmt.Errorf("You can't divide by zero")
// }else {

// 	return x/y,nil
// }



// }

//Varadic Functions 



// func getSum(nums ...int) int{

// sum:=0
// for _,num:=range nums{
// 	sum+=num
// }
// return sum
// }


// func getArraySum(arr []int) int{


// sum:=0

// for _,num:=range arr{

// 	sum+=num
// }
// return sum
// }


// func changeValue(f3 int ) int {
// f3+=1
// return f3
// }

// func main(){
// //General Syntax: func funcName(parameters) returnType{ BODY }

// //sayHello()
// // pl(getSum(4,5))
// // pl(getTwo(5))
// //pl(getQuotient(2,4))

// // pl(getSum(1,2,3,4,5,6,7))
// // pl(getArraySum(1,2,3,4,5,6,7))

// f3:=5
// pl("f3 Before Function :",f3)
// changeValue(f3)
// pl("f3 After Function ", f3)  //Value will Remain SAME

// }




//Pointers 

// func changeValue2(myPtr *int) {
// *myPtr=22
// }


// func main(){



// f4:=10
// pl("f4 before change value: ",f4)
// changeValue2(&f4)
// pl("f4 after Change Value: ",f4)


// f5:=14
//  var  f5Ptr *int=&f5
// pl("Address of f5",f5Ptr)
// pl("Value of f5", *f5Ptr)

// //Assign value to ptr
// *f5Ptr=77
// pl("Value of f5", *f5Ptr) 
// }



//Passing Array using Pointers

// func dblArrVals(arr *[4]int ){
// for x:=0;x<4;x++{

// arr[x]*=2
// }
// }


// func getAverage(nums ...float64) float64{

// var sum float64=0.0
// var numSize float64=float64(len(nums))

// for _,val:= range nums{

// 	sum+=val
// }

// return (sum/numSize)
// }

// func main(){


// 	pArr:=[4]int{1,2,3,4}

// 	dblArrVals(&pArr)
// 	pl(pArr)


// 	// Slice to a Function
// 	iSlice:=[]float64{11,12,17}
// 	fmt.Printf("Average %.2f\n",getAverage(iSlice...) )
// }


