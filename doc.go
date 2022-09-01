/*
Package pinyin : 汉语拼音转换工具.

Usage

	package main

	import (
		"github.com/xyzj/go-pinyin"
	)

	func main() {
		hans := "中国人"
		println(pinyin.XPinyin(hans, pinyin.ReturnNormal))
		// zhongguoren

		println(pinyin.XPinyin(hans, pinyin.ReturnFirstLetter))
		// zgr

		println(pinyin.XPinyin(hans, pinyin.ReturnAll))
		// zgr zhongguoren
	}
*/
package pinyin
