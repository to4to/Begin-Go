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

