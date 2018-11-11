package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
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
	bytes      int
	packets    int
	errs       int
	drop       int
	fifo       int
	frame      int
	compressed int
	multicast  int
}

func main() {
	getProcNetDev()
	fmt.Println(netInfo)
}

func Show() {
	for {
		netrefresh()
		for _, n := range netDevices {
			fmt.Printf("%s: %dKb/s\n", n.face, n.Receive.bytes/1024)
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func netrefresh() {
	netDevices = netDevices[:0]
	getProcNetDev()
	getLine()
	getNetDevices()
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
	netLin = netLin[2 : len(netLin)-1]
}

func getNetDevices() {
	for i := 0; i < len(netLin); i++ {
		ts := strings.Fields(netLin[i])
		var rn NetInfo
		rn.bytes, _ = strconv.Atoi(ts[1])
		rn.packets, _ = strconv.Atoi(ts[2])
		rn.errs, _ = strconv.Atoi(ts[3])
		rn.drop, _ = strconv.Atoi(ts[4])
		rn.fifo, _ = strconv.Atoi(ts[5])
		rn.frame, _ = strconv.Atoi(ts[6])
		rn.compressed, _ = strconv.Atoi(ts[7])
		rn.multicast, _ = strconv.Atoi(ts[8])
		var tn NetInfo
		tn.bytes, _ = strconv.Atoi(ts[9])
		tn.packets, _ = strconv.Atoi(ts[10])
		tn.errs, _ = strconv.Atoi(ts[11])
		tn.drop, _ = strconv.Atoi(ts[12])
		tn.fifo, _ = strconv.Atoi(ts[13])
		tn.frame, _ = strconv.Atoi(ts[14])
		tn.compressed, _ = strconv.Atoi(ts[15])
		tn.multicast, _ = strconv.Atoi(ts[16])

		dv := NetDevice{strings.Trim(ts[0], ":"), rn, tn}

		netDevices = append(netDevices, dv)
	}
}
