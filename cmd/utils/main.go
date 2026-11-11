package utils

import (
	"math/rand"
	"strings"
)
type Req interface {
	
}
func GenToken() string{
	strPatten := "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678"
	var arr [32]string
	runes := []rune(strPatten)
	for i := range len(arr) {
		randIndex := rand.Intn(len(runes))
		arr[i] = string(runes[randIndex])
	}
	token := strings.Join(arr[:len(arr)],"")
	return token
}
func HttpReq(req Req){

}