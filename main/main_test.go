package main

import (
	"fmt"
	"testing"
)

func TestGetline(t *testing.T) {
	getProcNetDev()
	getLine()
	//tmp := strings.Fields(netLin[0])
	//fmt.Println(tmp)
	fmt.Println(len(netLin))
}

func TestGetdevices(t *testing.T) {
	getProcNetDev()
	getLine()
	getNetDevices()
	fmt.Println(netDevices)
}

func TestShow(t *testing.T) {
	Show()
}
