package main
import(
	"net"
	"fmt"
	"runtime"
)
const MY_IP= "129.241.187.159"
const MY_PORT = "20011"
const TARGET_PORT = "33546"
const TARGET_IP = "129.241.187.161"

func connectTo(ipAdr string, port string)
{

	serverAddr, err := net.ResolveTCPAddr("tcp",ipAdr+":"+port)
	con, err := net.DialTCP("tcp", nil, serverAddr);
	
	if err != nil{
		fmt.Println("fail")
	}
	cd := make([]byte,1024)
	con.Read(cd)
	fmt.Printf("%s",cd)	
	stop :=0
	
	msg:= "Connect to: "+MY_IP+":"+MY_PORT+"\x00"
	con.Write([]byte(msg))	
	
	for stop !=-1{
		input := ""
		fmt.Scanf("%s",&input)
		if input=="stop"
		{
			stop=-1
		}
		con.Write([]byte(input+"\x00"))
		con.Read(cd)
		fmt.Printf("%s",cd)
	}
}

func ListenerCon(port string){
	psock, err := net.Listen("tcp", ":"+port)
	if err != nil { return }
	buf := make([]byte,1024)
 
  	for{
    		conn, err := psock.Accept()
    		if err != nil { return }
    		conn.Read(buf)
		fmt.Printf("%s",buf)
	 }		
}	

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	go connectTo(TARGET_IP,TARGET_PORT)
	go ListenerCon(MY_PORT)	
	
	
	for true
	{
	}
}
