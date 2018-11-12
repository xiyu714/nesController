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
	diff
}

type diff struct {
	bytes   int
	packets int
}

func main() {
	Show()
}

func Show() {

	for {
		netrefresh()
		for _, n := range netDevices {
			fmt.Printf("%s: %.2fKb/s\n", n.face, float32(n.Receive.diff.bytes)/1024)
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func netrefresh() {
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

		//处理diff
		var rdif diff
		var tdif diff
		func() {
			defer func() {
				err := recover()
				if err != nil {
					var x = diff{0, 0}
					rn.diff = x
					tn.diff = x
					dv := NetDevice{strings.Trim(ts[0], ":"), rn, tn}
					netDevices = append(netDevices, dv)
				}
			}()
			oldR := netDevices[i].Receive
			rdif.bytes = rn.bytes - oldR.bytes
			rdif.packets = rn.packets - oldR.packets
			rn.diff = rdif

			oldT := netDevices[i].Transmit
			tdif.bytes = tn.bytes - oldT.bytes
			tdif.packets = tn.packets - oldT.packets
			tn.diff = tdif

			dv := NetDevice{strings.Trim(ts[0], ":"), rn, tn}
			netDevices[i] = dv
		}()
	}
}
