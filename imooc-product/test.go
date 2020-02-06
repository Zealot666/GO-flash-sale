package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	string,err:=GetIntranceIp()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string)
}

func GetIntranceIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		//检查Ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("获取地址异常")
}
