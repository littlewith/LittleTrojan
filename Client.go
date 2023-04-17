package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	code     int
	rHost    string
	rPort    int
	conn     net.Conn
	fullAddr string
	err      error
}

func (client *Client) connect(rHost string, rPort int) bool {
	client.rHost = rHost
	client.rPort = rPort
	client.fullAddr = client.rHost + ":" + strconv.Itoa(rPort)
	client.conn, client.err = net.Dial("tcp", client.fullAddr)
	chkError(client.err)
	if client.err != nil {
		return false
	} else {
		fmt.Println("Connect success to " + client.fullAddr)
		return true
	}
}

func (client *Client) close() bool {
	time.Sleep(750)
	client.sendData("Bye!" + RESV_ACCOMPLISH)
	err := client.conn.Close()
	isClose := chkErrorNoExit(err)
	if isClose {
		fmt.Println("Received Exit Signal. Exiting...")
		return true
	} else {
		fmt.Println("Warning! Exit failed!")
		return false
	}
}

func (client *Client) receiveData() (string, int) {
	var readBuffer bytes.Buffer
	var length int
	for {
		var tmp = make([]byte, 1024)
		MsgLen, err := bufio.NewReader(client.conn).Read(tmp)
		chkErrorNoExit(err)
		length += MsgLen
		readBuffer.Write(tmp)
		tmpStr := readBuffer.String()
		if strings.Contains(tmpStr, RESV_ACCOMPLISH) {
			finalString := strings.Split(readBuffer.String(), RESV_ACCOMPLISH)
			fmt.Println("Received: " + finalString[0])
			return finalString[0], length
		} else {
			continue
		}
	}
}

func (client *Client) sendData(data string) bool {
	writeBuffer := []byte(data)
	MsgLen, err := client.conn.Write(writeBuffer)
	isSend := chkErrorNoExit(err)
	if isSend {
		fmt.Printf("Sent %d length\n", MsgLen)
		return true
	} else {
		return false
	}
}

func (client *Client) sendPic(filename string) bool {
	fmt.Println("Current File: " + filename)
	//count the file length of the File
	length, err := ioutil.ReadFile(filename)
	chkErrorNoExit(err)
	fmt.Println("Tips:file length: " + strconv.Itoa(len(length)) + "{" + strconv.Itoa(len(length)/1024+1) + "}")

	isQuery := client.sendData(strconv.Itoa(len(length)/1024+1) + RESV_ACCOMPLISH)
	if !isQuery {
		fmt.Println("Can not send a query message!")
		return false
	}

	queryFileLength, _ := client.receiveData()
	if queryFileLength != "ok" {
		fmt.Println("Not query! exit.")
		return false
	} else {
		fmt.Println("Queried the file Length!")
	}

	time.Sleep(3000)

	fileBuffer, err := os.Open(filename)
	//Again to open the file
	isOpen := chkErrorNoExit(err)
	if !isOpen {
		fmt.Println("Can not open the file properly!")
	}

	for {
		var tmp []byte = make([]byte, 1024)
		//each for start a new tmp
		lengthFile, _ := fileBuffer.Read(tmp)
		//write in to the temp
		if lengthFile == 0 {
			//read file end!
			break
		}

		msgLen, err := client.conn.Write(tmp)
		eachQuery, _ := client.receiveData()
		if eachQuery != "ok" {
			fmt.Println("An error occurred when transmitting the file!")
			return false
		}

		isSend := chkErrorNoExit(err)
		if isSend {
			fmt.Print("Send Successfully--" + strconv.Itoa(msgLen) + " ")
		}
	}

	return true
}
