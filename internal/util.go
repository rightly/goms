package internal

import (
	"log"
	"encoding/json"
	"strings"
	"strconv"
)

func CheckErrWithLog(err error, msg string)  {
	logger := log.Logger{}
	if !CheckErr(err, msg) {
		logger.Println(err)
	}
}

func IndentedJson(v interface{}) []byte {
	indentedJson, err := json.MarshalIndent(v, "", "\t")
	CheckErr(err, "couldn't marshal json")

	return indentedJson
}

func IndentedString(v interface{}) string {
	indentedString := string(IndentedJson(v))

	return indentedString
}

func SplitAddress(s string) (string, uint16) {
	addr := strings.Split(s, ":")
	ip, port := addr[0], "80"
	if len(addr) > 1 {
		ip, port = addr[0], addr[1]
	}

	base, bitSize := 10, 16
	retPort, _ := strconv.ParseUint(port, base, bitSize)

	return ip, uint16(retPort)
}