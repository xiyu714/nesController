package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

const netDev = "/proc/net/dev"

var netInfo string

func main() {
	for {
		getProcNetDev()
		fmt.Println(netInfo)
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func getProcNetDev() {
	b, err := ioutil.ReadFile(netDev)
	if err != nil {
		log.Fatal(err)
	}
	netInfo = string(b)
}
