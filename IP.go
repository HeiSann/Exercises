package main

import (
	"net"
	"fmt"
	"strings"
)
func getMyIP() string{
	allIPs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error receiving IPs")
		return ""
	}
			
	uncutIP := allIPs[1].String()
	IP := strings.Split(uncutIP, "/")		
	myIP := IP[0]

	return myIP
}

func main(){
	var ip string
	ip = getMyIP()

	fmt.Println(ip)
	
}


