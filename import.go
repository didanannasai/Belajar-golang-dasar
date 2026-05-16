package main

import (
	"golang-dasar/helper"
	"fmt"
)

func main() {
	fmt.Println(helper.SayHello("budi"))
	
	/* suatu function atau variabel dari package lain bisa diakses dari package lain dengan syarat
	   nama di variabel atau function itu huruf petamanya kapital */
	 fmt.Println(helper.Application)  // --> bisa diakses
	// fmt.Println(helper.version) --> tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("tono")) -- tidak bisa diakses
}