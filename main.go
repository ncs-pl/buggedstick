package main

import (
	"flag"
	"log"
)

var (
	flagToEmail string
)

func init() {
	flag.StringVar(&flagToEmail, "to", "", "You victim email")
	flag.Parse()
}

func main() {
	if flagToEmail == "" {
		log.Fatalln("Missing To Email. Use flag '-to <victim email>'!")
	}

	mail := NewMail(flagToEmail)
	mail.Send()
}
