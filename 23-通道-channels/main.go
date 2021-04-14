package main

import "fmt"

func main() {
	// 使用 `make(chan val-type)` 创建一个新的通道。
	// 通道类型就是他们需要传递值得类型。
	messages := make(chan string)

	// 使用 `channel <-` 语法 _发送(send)_ 一个新的值到通道中。
	// 这里我们在一个新的 go 协程中发送 `"ping"` 到上面创建的 `messages` 通道中。
	go func() { messages <- "ping" }()

	// 使用 `<-channel` 语法从通道中 _接收(receives)_ 一个值。
	// 这里将接收我们在上面发送的 `"ping"` 消息并打印出来。
	msg := <-messages
	fmt.Println(msg)

	// 注意：主协程不能直接往通道里发数据，会造成阻塞，即使后面有别的协程在等待，
	// 也会造成协程无法启动，因为主协程已经被阻塞。
}
