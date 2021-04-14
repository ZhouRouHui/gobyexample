package main

import (
	"fmt"
	"net/url"
	"strings"
)

/*
url 提供一个统一资源定位方式。这里了解一下 go 中是如何解析 url 的。
 */

func main() {
	// 我们将解析这个 url 示例，它包含了一个 schema，
	// 认证信息，主机名，端口，路径，查询参数和片段。
	s := "postgres://user:pass@host.com:5432/path?k=v1&k=v2#f"

	// 解析这个 url 并确保解析没有出错
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 直接访问 schema。
	fmt.Println(u.Scheme)

	// `User` 包含了所有的认证信息，这里调用 `Username` 和 `Password` 来获取独立值。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())

	// `Host` 同时包括主机名和端口信息，如果端口存在的话，
	// 使用 `strings.Split()` 从 `Host` 中手动提取端口
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	// 这里我们提出路径和查询片段信息
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// 要得到字符串中的 `k=v` 这种格式的查询参数，可以使用 `RawQuery` 函数。
	// 你也可以将查询参数解析为一个 map。已解析的 map 以查询字符串为键，
	// 对应值字符串切片为值，所以如果只想得到一个键对应的第一个值，
	// 将索引位置设置为 `[0]` 就行了。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
	fmt.Println(m["k"][1])
}
