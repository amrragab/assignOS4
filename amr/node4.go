// TODO: In this file you need to build a graph of nodes
// and each node contains a set of files.

package main

import (
	"./core/student"
	"fmt"
	"time"
	"strings"
	//"strconv"
)

// Message struct.
type Message struct {
	From     int
	To       int
	UserName string
	Content  string
}

type Param struct {
	Files map[string]int
	//AdjList arr[5]string
}

var P Param
// Global Declarations.
var masterAddr string = "10.0.0.7:46321"
var connectedNodes = []int{2, 3, 5}
var fileList = []string{
	"3737941_134967132_n.jpg",
	"1935851_722579545_n.jpg",
	"1512714_286659690_n.jpg"}

/*var sendLoop int = 3
	
// var to know the sender 
var sender int = 0*/

// TODO: Change this to your current password.
var studentPassword string = "P6Hjqh"

// Implementing ReceiveHandler for student package.
type RcvHandler struct{}

//5leto global
var St= new(student.Student)

// to handle the msg after recive handler is caller
func handleMsg(from int, to int, username string,content string){
	fmt.Println(from," ",to)
	lines := strings.Split(content, " ")
	P.Files[lines[0]] = int(lines[1][0])
		sentstr := fmt.Sprintf("%s4",content)
		for c := 0; c < len(connectedNodes); c++ {
				if( ! strings.Contains(lines[1],string(connectedNodes[c])) ){
				error := St.SendMsg(connectedNodes[c],sentstr)
					x := 2
				time.Sleep(time.Second * time.Duration(x))
				if error != nil {
					fmt.Println("Failed to SendMsg to node",connectedNodes[c],": ", error)
					//return
					}	
				//fmt.Println("File ",c,": ", content)
				}
			}
}

func handleMsgFound(from int, to int, username string, content string){
	fmt.Println(from," ",to," ",content)
	lines := strings.Split(content, " ")
	c_sz := len(connectedNodes)
	sentstr := content + "4"
	for c := 0; c < c_sz; c++{
		if( ! strings.Contains(lines[1],string(connectedNodes[c])) ){
			error := St.SendMsg(connectedNodes[c],sentstr)
			x := 2
			time.Sleep(time.Second * time.Duration(x))
			if error != nil {
				fmt.Println("Failed to SendMsg to node",connectedNodes[c],": ", error)
				//return
			}	
		}
	}
}

func handleNode(from int, content string){
	tmp := strings.Count(content,"x")
	if tmp == 0 {
		return
	}
	content = strings.Replace(content, "x4", "25", -1)
	content = "node " + content
		fmt.Println("alone ",content)

	for _,c := range connectedNodes{
		if ( c != from){
			St.SendMsg(c,content)
		//	S := 5
	//time.Sleep(time.Second * time.Duration(S))
		}
	}
}

// Handle a message received.
func (rcvHand *RcvHandler) ReceiveHandler(from int, to int, username string,
	content string) {
	//tmp_str := fmt.Sprintf("%s is not connected at node %d",username,to)
	//fmt.Println(content)
	if(! strings.Contains(content,"not connected")  ){
			lines := strings.Split(content, " ")
		if(lines[0] == "node"){
						fmt.Println(lines[1])
			go handleNode(from,lines[1])
		}else{
			_, found := P.Files[lines[0]]
			if( ! found ){
					go handleMsg(from,to,username,content)
			} 
		}
	}
	// DONOT CHANGE PARAMENTERS OR FUNCTION HEADER.
	// TODO: Implement handling a message received.
	//go handleMsg(from,to,username,content)
}

func main() {
	// initializing current local struct
	P.Files = make(map[string]int)
	f_size := len(fileList)
	for i := 0 ; i < f_size ; i++ {
		P.Files[fileList[i]] = 4
	}
	//fmt.Println(P.Files)
	// Setup connection to master of current node.
	error := St.Connect(masterAddr, studentPassword)
	if error != nil {
		fmt.Println("Failed to connect to master node:", error)
		return
	}

	// Link implementation of ReceiveHandler to student.
	rcv := new(RcvHandler)
	go St.Receive(rcv)
	// End of Setup.


	S := 5
	time.Sleep(time.Second * time.Duration(S))

	// TODO: Broadcast your files to neighbours.
	
	c_sz := len(connectedNodes)
	f_sz := len(fileList)
	// send nodes to other connected nodes
	node := "node x1:x2:x3:23:x5"
	for c := 0; c < c_sz; c++{
		error := St.SendMsg(connectedNodes[c],node)
		x := 5
			time.Sleep(time.Second * time.Duration(x))
			if error != nil {
				fmt.Println("Failed to SendMsg to node 1: ", error)
				return
			}
	}

		fmt.Println("Intializing node 4\n");
	for j := 0; j < f_sz ; j++ {
		for c := 0; c < c_sz; c++ {
			sentstr := fmt.Sprintf("%s 4",fileList[j])
			error = St.SendMsg(connectedNodes[c],sentstr)
				x := 5
				time.Sleep(time.Second * time.Duration(x))
				if error != nil {
					fmt.Println("Failed to SendMsg to node",connectedNodes[c],": ", error)
					return
				} 
			}
	}

	// TODO: It's expected to converge after N second
	// To be able to print a stable graph and shortest
	// path for file.
	N := 20
	time.Sleep(time.Second * time.Duration(N))

	fmt.Println(P.Files,'\n',len(P.Files))
}
