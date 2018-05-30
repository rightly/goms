package internal

import "fmt"

const (
	ConfigureError = iota + 9000
	PushError
)

func CheckErr(err error, msg string) bool {
	if err == nil {
		return false
	}
	fmt.Printf("%s :%s\n", msg, err)
	return true
}