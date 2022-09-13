package pinyin

import (
	"strings"
)

// Meta
const (
	Version   = "0.19.0"
	Author    = "mozillazg, 闲耘"
	License   = "MIT"
	Copyright = "Copyright (c) 2016 mozillazg, 闲耘"
)

// ReturnType 返回的拼音类型
type ReturnType int

var (
	// ReturnFirstLetter 仅返回首字母
	ReturnFirstLetter ReturnType = 0
	// ReturnNormal 仅返回全拼
	ReturnNormal ReturnType = 1
	// ReturnAll 返回首字母+全拼
	ReturnAll ReturnType = 2
	// ReturnInitials 返回声母。如： zh g
	ReturnInitials ReturnType = 3
)

// 拼音风格(推荐)
// const (
// 	Normal      = 0 // 普通风格，不带声调（默认风格）。如： zhong guo
// 	Tone        = 1 // 声调风格1，拼音声调在韵母第一个字母上。如： zhōng guó
// 	Tone2       = 2 // 声调风格2，即拼音声调在各个韵母之后，用数字 [1-4] 进行表示。如： zho1ng guo2
// 	Tone3       = 8 // 声调风格3，即拼音声调在各个拼音之后，用数字 [1-4] 进行表示。如： zhong1 guo2
// 	Initials    = 3 // 声母风格，只返回各个拼音的声母部分。如： zh g 。注意：不是所有的拼音都有声母
// 	FirstLetter = 4 // 首字母风格，只返回拼音的首字母部分。如： z g
// 	Finals      = 5 // 韵母风格，只返回各个拼音的韵母部分，不带声调。如： ong uo
// 	FinalsTone  = 6 // 韵母风格1，带声调，声调在韵母第一个字母上。如： ōng uó
// 	FinalsTone2 = 7 // 韵母风格2，带声调，声调在各个韵母之后，用数字 [1-4] 进行表示。如： o1ng uo2
// 	FinalsTone3 = 9 // 韵母风格3，带声调，声调在各个拼音之后，用数字 [1-4] 进行表示。如： ong1 uo2
// )

// 拼音风格(兼容之前的版本)
// const (
// 	NORMAL       = Normal
// 	TONE         = Tone
// 	TONE2        = Tone2
// 	INITIALS     = Initials
// 	FIRST_LETTER = FirstLetter
// 	FINALS       = Finals
// 	FINALS_TONE  = FinalsTone
// 	FINALS_TONE2 = FinalsTone2
// )

// 声母表
var initialArray = strings.Split(
	"b,p,m,f,d,t,n,l,g,k,h,j,q,x,r,zh,ch,sh,z,c,s",
	",",
)

// // 所有带声调的字符
// var rePhoneticSymbolSource = func(m map[string]string) string {
// 	s := ""
// 	for k := range m {
// 		s = s + k
// 	}
// 	return s
// }(phoneticSymbol)

// // 匹配带声调字符的正则表达式
// var rePhoneticSymbol = regexp.MustCompile("[" + rePhoneticSymbolSource + "]")

// // 匹配使用数字标识声调的字符的正则表达式
// var reTone2 = regexp.MustCompile("([aeoiuvnm])([1-4])$")

// // 匹配 Tone2 中标识韵母声调的正则表达式
// var reTone3 = regexp.MustCompile("^([a-z]+)([1-4])([a-z]*)$")

// Args 配置信息
// type Args struct {
// 	Style     int    // 拼音风格（默认： Normal)
// 	Heteronym bool   // 是否启用多音字模式（默认：禁用）
// 	Separator string // Slug 中使用的分隔符（默认：-)

// 	// 处理没有拼音的字符（默认忽略没有拼音的字符）
// 	// 函数返回的 slice 的长度为0 则表示忽略这个字符
// 	Fallback func(r rune, a Args) []string
// }

// // Style 默认配置：风格
// var Style = Normal

// // Heteronym 默认配置：是否启用多音字模式
// var Heteronym = false

// // Separator 默认配置： `Slug` 中 Join 所用的分隔符
// var Separator = "-"

// Fallback 默认配置: 如何处理没有拼音的字符(忽略这个字符)
// var Fallback = func(r rune, a Args) []string {
// 	return []string{string(r)}
// }

// var finalExceptionsMap = map[string]string{
// 	"ū": "ǖ",
// 	"ú": "ǘ",
// 	"ǔ": "ǚ",
// 	"ù": "ǜ",
// }
// var reFinalExceptions = regexp.MustCompile("^(j|q|x)(ū|ú|ǔ|ù)$")
// var reFinal2Exceptions = regexp.MustCompile("^(j|q|x)u(\\d?)$")

// NewArgs 返回包含默认配置的 `Args`
// func NewArgs() Args {
// 	return Args{Style, Heteronym, Separator, Fallback}
// }

// XPinyin 返回中文的拼音首字母和完整拼音字符串
//
// s: 需要转换的字符串
//
// t: 返回的格式
func XPinyin(s string, t ReturnType) string {
	var ss = make([]string, 0)
	switch t {
	case ReturnNormal:
		for _, r := range s {
			ss = append(ss, GetPinyin(r)[0])
		}
		return toneReplacer.Replace(strings.Join(ss, ""))
	case ReturnFirstLetter:
		for _, r := range s {
			ss = append(ss, string(GetPinyin(r)[0][0]))
		}
		return strings.Join(ss, "")
	case ReturnInitials:
		for _, r := range s {
			ss = append(ss, initial(GetPinyin(r)[0]))
		}
		return strings.Join(ss, "")
	default:
		for _, r := range s {
			ss = append(ss, GetPinyin(r)[0])
		}
		ss1 := make([]string, 0)
		for _, r := range s {
			ss1 = append(ss1, string(GetPinyin(r)[0][0]))
		}
		return toneReplacer.Replace(strings.Join(ss, "")) + " " + strings.Join(ss1, "")
	}
}

// XPinyinMatch 检查字符串的拼音首字母和全拼是否匹配输入值
// s: 中文字符串
//
// substr: 匹配项
//
// 返回： true:匹配, false: 不匹配
func XPinyinMatch(s, substr string) bool {
	if substr == "" {
		return true
	}
	return strings.Contains(XPinyin(s, ReturnFirstLetter), substr) || strings.Contains(XPinyin(s, ReturnNormal), substr)
}

// 获取单个拼音中的声母
func initial(p string) string {
	s := ""
	for _, v := range initialArray {
		if strings.HasPrefix(p, v) {
			s = v
			break
		}
	}
	return s
}

// // 处理 y, w
// func handleYW(p string) string {
// 	// 特例 y/w
// 	if strings.HasPrefix(p, "yi") {
// 		p = p[1:] // yi -> i
// 	} else if strings.HasPrefix(p, "y") {
// 		p = "i" + p[1:] // y -> i
// 	} else if strings.HasPrefix(p, "wu") {
// 		p = p[1:] // wu -> u
// 	} else if strings.HasPrefix(p, "w") {
// 		p = "u" + p[1:] // w -> u
// 	}
// 	return p
// }

// func toFixed(p string) string {
// 	p = handleYW(p)
// 	// 替换拼音中的带声调字符
// 	py := rePhoneticSymbol.ReplaceAllStringFunc(p, func(m string) string {
// 		symbol, _ := phoneticSymbol[m]
// 		// 去掉声调: a1 -> a
// 		m = reTone2.ReplaceAllString(symbol, "$1")
// 		return m
// 	})

// 	return py
// }

// // 获取单个拼音中的韵母
// func final(p string) string {
// 	n := initial(p)
// 	if n == "" {
// 		return handleYW(p)
// 	}

// 	// 特例 j/q/x
// 	matches := reFinalExceptions.FindStringSubmatch(p)
// 	// jū -> jǖ
// 	if len(matches) == 3 && matches[1] != "" && matches[2] != "" {
// 		v, _ := finalExceptionsMap[matches[2]]
// 		return v
// 	}
// 	// ju -> jv, ju1 -> jv1
// 	p = reFinal2Exceptions.ReplaceAllString(p, "${1}v$2")
// 	return strings.Join(strings.SplitN(p, n, 2), "")
// }

// func applyStyle(p []string, a Args) []string {
// 	newP := []string{}
// 	for _, v := range p {
// 		newP = append(newP, toFixed(v, a))
// 	}
// 	return newP
// }

// // SinglePinyin 把单个 `rune` 类型的汉字转换为拼音.
// func SinglePinyin(r rune, a Args) []string {
// 	if a.Fallback == nil {
// 		a.Fallback = Fallback
// 	}
// 	value, ok := PinyinDict[int(r)]
// 	pys := []string{}
// 	if ok {
// 		pys = strings.Split(value, ",")
// 	} else {
// 		pys = a.Fallback(r, a)
// 	}
// 	if len(pys) > 0 {
// 		if !a.Heteronym {
// 			pys = []string{pys[0]}
// 		}
// 		return applyStyle(pys, a)
// 	}
// 	return pys
// }

// // Pinyin 汉字转拼音，支持多音字模式.
// func Pinyin(s string, a Args) [][]string {
// 	pys := [][]string{}
// 	for _, r := range s {
// 		py := SinglePinyin(r, a)
// 		if len(py) > 0 {
// 			pys = append(pys, py)
// 		}
// 	}
// 	return pys
// }

// // LazyPinyin 汉字转拼音，与 `Pinyin` 的区别是：
// // 返回值类型不同，并且不支持多音字模式，每个汉字只取第一个音.
// func LazyPinyin(s string, a Args) []string {
// 	a.Heteronym = false
// 	pys := []string{}
// 	for _, v := range Pinyin(s, a) {
// 		pys = append(pys, v[0])
// 	}
// 	return pys
// }

// // Slug join `LazyPinyin` 的返回值.
// // 建议改用 https://github.com/mozillazg/go-slugify
// func Slug(s string, a Args) string {
// 	separator := a.Separator
// 	return strings.Join(LazyPinyin(s, a), separator)
// }

// // Convert 跟 Pinyin 的唯一区别就是 a 参数可以是 nil
// func Convert(s string, a *Args) [][]string {
// 	if a == nil {
// 		args := NewArgs()
// 		a = &args
// 	}
// 	return Pinyin(s, *a)
// }

// // LazyConvert 跟 LazyPinyin 的唯一区别就是 a 参数可以是 nil
// func LazyConvert(s string, a *Args) []string {
// 	if a == nil {
// 		args := NewArgs()
// 		a = &args
// 	}
// 	return LazyPinyin(s, *a)
// }
