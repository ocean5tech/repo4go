package main

//iota 特殊常量，可以被编译器修改的常量,从零开始，自动加一,只在当前const块生效，换块重新计算

const (
	ERR1 = iota
	ERR2
	ERR3 = "ha"
	ERR4
	ERR5 = iota
)
