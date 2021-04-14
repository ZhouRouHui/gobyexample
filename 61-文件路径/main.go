package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

/**
filepath 包为文件路径提供了方便的跨操作系统的解析和构建函数；
比如：Linux 下的 dir/file 和 windows 下的 dir.ile
 */

func main() {
	// 应使用 `Join` 来构建可移植（跨操作系统）的路径。
	// 它接收任意数量的参数，并参照传入顺序构造一个对应层次结构的路径。
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// 你应该总是使用 `Join` 代替手动拼接 `/` 和 `\`。
	// 除了可移植性，`Join` 还会删除多余的分隔符和目录，使用路径更加规范。
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` 和 `Base` 可以被用于分割路径中的目录和文件。
	// 此外，`Split` 可以一次调用返回上面两个函数的结果
	fmt.Println("Dir:", filepath.Dir(p))
	fmt.Println("Base:", filepath.Base(p))
	fmt.Println(filepath.Split(p))

	// 判断路径是否为绝对路径。
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// 某些文件名包含了扩展名（文件类型）。
	// 我们可以用 `Ext` 将扩展名分割出来。
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// 想获取文件名清楚扩展名后的值，请使用 `strings.TrimSuffix`。
	fmt.Println(strings.TrimSuffix(p, ext))

	// `Rel` 寻找 `basepath` 与 `targpath` 之间的相对路径。
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b/t/file", "a/b")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
