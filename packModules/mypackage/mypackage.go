package stuff

import "strconv"

var Name string="Himanshu"

func IntArrToStrArr(intArr []int) []string {


	var strArr []string
	for _,i:= range intArr{


		strArr=append(strArr, strconv.Itoa(i))
	}

	return strArr
}