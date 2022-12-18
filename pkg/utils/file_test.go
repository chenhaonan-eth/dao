package utils

import (
	"fmt"
	"testing"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func TestIs(t *testing.T) {
	str := "月色真美，风也温柔，233333333，~！@#"                          //go字符串编码为utf-8
	fmt.Println("before convert:", str)                                 //打印转换前的字符串
	fmt.Println("isUtf8:", IsUtf8([]byte(str)))                         //判断是否是utf-8
	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str)) //使用官方库将utf-8转换为gbk
	fmt.Println("gbk直接打印会出现乱码:", string(gbkData))                       //乱码字符串
	fmt.Println("isGBK:", IsGBK(gbkData))                               //判断是否是gbk
	utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(gbkData)    //将gbk再转换为utf-8
	fmt.Println("isUtf8:", IsUtf8(utf8Data))                            //判断是否是utf-8
	fmt.Println("after convert:", string(utf8Data))                     //打印转换前的字符串

}
