package main

import (
	"fmt"
	cc "github.com/frenata/complexconv"
	"reflect"
	"strconv"
)

func main() {
	s := "1e+6i"
	c, e := cc.ParseComplex(s)
	fmt.Println(c, e)
	c, e = cc.ParseComplex("999")
	fmt.Println(c, e)
	c, e = cc.ParseComplex("2.33e+4+2i")
	fmt.Println(c, e)
	s = "2.4"
	fmt.Println(strconv.ParseFloat(s, 64))
	fmt.Println(strconv.ParseFloat(s, 32))
	s = fmt.Sprint(c)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)
}
