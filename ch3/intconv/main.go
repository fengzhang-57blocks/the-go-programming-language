package main

import "fmt"

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	/*
	*  %后的谓词[1]告知Printf重复使用第一个操作数
	*  %o %x或%X之前的谓词#告知Printf输出相应的前缀0 0x或0X
	 */
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x\n", x)

	// 用%c输出文字符号，如果希望输出带有单引号则用%q

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)
}
