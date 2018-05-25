package internal

import "fmt"

func CheckErr(err error, msg string) bool {
	if err == nil {
		return true
	}
	fmt.Printf("%s :%s\n", msg, err)
	return false
}