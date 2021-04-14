package main

import "fmt"

func main() {
	// 我们将遍历在 `queue` 通道中的两个值
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 这个 `range` 迭代从 `queue` 中得到的两个值。
	// 因为我们在前面 `close` 了这个通道，这个迭代会在接收完 2 个值之后结束。
	// 如果我们没有 `close` 它，我们将在这个循环中继续阻塞执行，等待接收第三个值
	// `range` 迭代一个通道的时候相当于一直在接收通道的消息，
	// 如果通道没有消息但是又没有 `close` 这是通道会阻塞等待
	for elem := range queue {
		fmt.Println(elem)
	}
}
