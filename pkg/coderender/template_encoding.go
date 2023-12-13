/*
----------------------------------------
@Create 2023/11/15
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe 多编码格式转换
----------------------------------------
@Version 1.0 2023/11/15
@Memo create this file
*/

package coderender

import "golang.org/x/text/encoding/simplifiedchinese"

const (
	GBK     string = "GBK"
	UTF8    string = "UTF8"
	UNKNOWN string = "UNKNOWN"
)

// EncodeAuto2Utf8 自动将各种格式转换为UTF-8
func EncodeAuto2Utf8(data []byte) []byte {
	coding := GetCoding(data)
	switch coding {
	case UNKNOWN:
		fallthrough
	case GBK:
		return EncodeGbk2Utf8(data)
	case UTF8:
		fallthrough
	default:
		return data
	}
}

// EncodeUtf82Gbk 将UTF-8转换为GBK
func EncodeUtf82Gbk(str string) []byte {
	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str)) //使用官方库将utf-8转换为gbk
	return gbkData
}

// EncodeGbk2Utf8 将Gbk转换为UTF-8
func EncodeGbk2Utf8(gbk []byte) []byte {
	utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(gbk) //将gbk再转换为utf-8
	return utf8Data
}

// GetCoding
// 需要说明的是，isGBK()是通过双字节是否落在gbk的编码范围内实现的，
// 而utf-8编码格式的每个字节都是落在gbk的编码范围内，
// 所以只有先调用isUtf8()先判断不是utf-8编码，再调用isGBK()才有意义
func GetCoding(data []byte) string {
	if isUtf8(data) == true {
		return UTF8
	} else if isGBK(data) == true {
		return GBK
	} else {
		return UNKNOWN
	}
}

// 判断是否是UTF8
func isUtf8(data []byte) bool {
	i := 0
	for i < len(data) {
		if (data[i] & 0x80) == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

// 判断是否是GBK
func isGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		if data[i] <= 0x7f {
			//编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		} else {
			//大于127的使用双字节编码，落在gbk编码范围内的字符
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

// 计数 8bit 中有多少个 前1
func preNUm(data byte) int {
	var mask byte = 0x80
	var num int = 0
	//8bit中首个0bit前有多少个1bits
	for i := 0; i < 8; i++ {
		if (data & mask) == mask {
			num++
			mask = mask >> 1
		} else {
			break
		}
	}
	return num
}
