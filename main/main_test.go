package main

import (
	"fmt"
	"testing"
	"time"
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

func TestNetrefresh(t *testing.T) {
	netrefresh()
	fmt.Println(netDevices)
	time.Sleep(time.Duration(1) * time.Second)
	netrefresh()
	fmt.Println(netDevices)
}

func TestShow(t *testing.T) {
	Show()
}

func TestSlice(t *testing.T) {
	var x []int

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
			//fmt.Println("我执行了？")
			x = append(x, 1, 2)
			fmt.Println(x[1])
		}()
		fmt.Println(x[1])
	}()
	fmt.Println("正常执行")
}
