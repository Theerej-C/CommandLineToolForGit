package main

import (
	"flag"
)
func main()  {
	var folder string
	var email string
	flag.StringVar(&folder,"add","","add a new folder to scan")
	flag.StringVar(&email,"email","your@eamil.com","the email to scan")
	flag.Parse()
	if folder != ""{
		scan(folder)
	}
	stats(email)
}