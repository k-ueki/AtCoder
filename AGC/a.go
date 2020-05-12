package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewScanner(os.Stdin)

func main(){
	var h,w int
	fmt.Scan(&h,&w)

for i:=0;i<h;i++{
	in.Scan()
	temp := strings.Split(in.Text(), " ")
	for j,v:=range temp{

	}
	li1, _ := strconv.Atoi(temp[0])
	li2, _ := strconv.Atoi(temp[1])
	a = append(a, li1)
	b = append(b, li2)
}

}


