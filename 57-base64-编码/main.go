package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	// 这是将要编解码的字符串。
	data := "abc123!?$*&()'-=@~"

	// go 同时进行标准的和 url 兼容的 base64 格式。
	// 编码需要使用 `[]byte` 类型参数，所以要将字符串转换成此类型。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码可能会返回错误，如果不确定输入信息格式是否正确，
	// 那么，那就需要进行错误检查了。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 使用 url 兼容的 base64 格式进行编解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
