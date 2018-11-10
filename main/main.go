package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const netDev = "/proc/net/dev"

var netInfo string

const baseineN = 3

var lineN int
var netLin []string
var netDevices []NetDevice

type NetDevice struct {
	face     string
	Receive  NetInfo
	Transmit NetInfo
}

type NetInfo struct {
	bytes      string
	packets    string
	errs       string
	drop       string
	fifo       string
	frame      string
	compressed string
	multicast  string
}

func main() {
	getProcNetDev()
	fmt.Println(netInfo)
}

func getProcNetDev() {
	b, err := ioutil.ReadFile(netDev)
	if err != nil {
		log.Fatal(err)
	}
	netInfo = string(b)
}

func getLine() {
	netLin = strings.Split(netInfo, "\n")
	netLin = netLin[2:]
}
