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

func TestSlice(t *testing.T) {
	var x []int

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
			fmt.Println("我执行了？")
		}()
		fmt.Println(x[1])
	}()
	fmt.Println("正常执行")
}

func t(x []int) {

}
