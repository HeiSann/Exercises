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

func getBroadCastIP(myIP string) string{
	temp :=strings.Split(myIP,".")
	fmt.Println(temp)
	broadCastIP :=temp[0]+"."+temp[1]+"."+temp[2]+"."+"255"
	return broadCastIP
	

}

func main(){
	var ip string
	ip = getMyIP()

	fmt.Println(ip)
	bcIP :=getBroadCastIP(ip)
	fmt.Println(bcIP)
	
}


