package main

import (
	"fmt"
	"testing"
)

func TestGetline(t *testing.T) {
	getProcNetDev()
	getLine()
	fmt.Println(netLin[0])
}
