package main

import "fmt"
import "regexp"

func main() {
	timestamp_validation("2022/12/12")
}

func timestamp_validation(timestamp string) {
	re := regexp.MustCompile(`(\d{4})/(\d{2})/(\d{2})`)
	result := re.MatchString(timestamp)
	return result
}

func mail_address_validation(mail_address string){
	re := 	 regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@+^[a-zA-Z0-9_.+-]`)
	result := re.MatchString(mail_address)
	return result
}

func IsEmpty(strings []string) bool {
	return len(strings) == 0
}
func IsBtweenNum(min, max, i){
	if(min < i && i < max) reuturn true
	else return false
}

type Address struct {
	postCode string
	tel string
}
