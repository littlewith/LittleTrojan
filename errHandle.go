package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
)

var RESV_ACCOMPLISH string = "SQ223!!@@##$$77&&ttyy"

func chkError(err error) bool {
	if err != nil {
		log.Fatalln(err.Error())
		return false
	} else {
		return true
	}
}

func chkErrorNoExit(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		return true
	}
}

func covert2Utf8(raw []byte) string {
	output, err := simplifiedchinese.GB18030.NewDecoder().Bytes(raw)
	issuc := chkErrorNoExit(err)
	if issuc {
		result := string(output)
		return result
	} else {
		return "When convert some Error Occurred."
	}
}
