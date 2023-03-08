package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// )

//  var pl=fmt.Println;

//  func main(){

//        f, err:=os.Create("data.txt")

// 	   if err!=nil{
// 		log.Fatal(err)
// 	   }
// 	   defer f.Close() //as long as we are in main file will remain open as we leave it gets closed

//       iPrimArr :=[]int{ 2,3,5,7,11}
// 	  var stringPrimarr []string

// for _,val:= range iPrimArr{

// 	stringPrimarr=append(stringPrimarr,strconv.Itoa(val)) //Itoa converts integer to string
// }

// for _, val:= range stringPrimarr{

// 	_,err:=f.WriteString(val + "\n")

// 	if err!=nil{

// log.Fatal(err)

// 	}
// }

// f ,err=os.Open("data.txt")

// if err!=nil {

// 	log.Fatal(err)
// }

// defer f.Close()

// scan1:=bufio.NewScanner(f)

// for scan1.Scan(){

// pl("Prime :" ,scan1.Text())

// }

// if err:=scan1.Err();err!=nil{
// 	log.Fatal(err)
// }

//  }

//Appending on File

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"os"
// )
 
// var pl=fmt.Println;

// func main(){

// /*
//         Exactly one of   O_RDONLY , O_WRONLY ,or O_RDWR must be specified 

// 		O_RDONLY : open the file read only
// 		O_WRONLY : open the file write only
// 		O_RDWR : open the file read and write 


// 		These can be or'ed //or'ed= oe edited 


// 		O_APPEND : append data to the file when writing 
// 		O_CREATE : create a new file if none exists
// 		O_EXCL : used with O_CREATE ,file must not exist
// 		O_SYNC : open for synchronous I/O
// 		O_TRUNC : truncate regular writable file when opened 
// */


// //First checking if file exists 
// _,err :=os.Stat("data.txt") // it exists 

// if errors.Is(err, os.ErrNotExist){
// pl("File Doesn't Exists")
// }else {

//     f,err:=os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)  // |=OR
  
// if err!=nil{

// 	log.Fatal(err)
// }

// defer f.Close()

// //adding data
// if _,err:=f.WriteString("13\n");err!=nil{

// 	log.Fatal(err)
// }


// }

// }