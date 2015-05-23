// TODO: In this file you need to build a graph of nodes
// and each node contains a set of files.

package main

import (
	"./core/student"
	"fmt"
	"time"
)

// Message struct.
type Message struct {
	From     int
	To       int
	UserName string
	Content  string
}

// Global Declarations.
var masterAddr string = "10.0.0.5:46321"
var connectedNodes = []int{1, 3, 4}
var fileList = []string{
	"9978620_137577509_n.jpg",
	"1005860_104234756_n.jpg",
	"1099051_699791809_n.jpg"}

//var sendLoop int = 3

// var to know the sender 
//var sender int = 0

// TODO: Change this to your current password.
var studentPassword string = "P6Hjqh"

// Implementing ReceiveHandler for student package.
type RcvHandler struct{}

//5leto global
//student := new(student.Student)
var St= new(student.Student)

// to handle the msg after recive handler is caller
func handleMsg(from int,msg string){

	//for j := 0; j < 3 ; j++ {
			for c := 0; c < 3; c++ {
				if from == connectedNodes[c]{
					continue
				}
				error := St.SendMsg(connectedNodes[c],msg)
				if error != nil {
					fmt.Println("Failed to SendMsg to node",connectedNodes[c],": ", error)
					return
					}	
				fmt.Println("File ",c,": ", msg)
			}
//	}
}

// Handle a message received.
func (rcvHand *RcvHandler) ReceiveHandler(from int, to int, username string,
	content string) {
	// DONOT CHANGE PARAMENTERS OR FUNCTION HEADER.
	// TODO: Implement handling a message received.
	
	go handleMsg(from,content);
	//fmt.Printf("%+v\n", rcvHand)
	//fmt.Println("rcvHand struct: ",rcvHand)
}

func main() {
	
	S := 5
	time.Sleep(time.Second * time.Duration(S))
	// Setup connection to master of current node.
	error := St.Connect(masterAddr, studentPassword)
	if error != nil {
		fmt.Println("Failed to connect to master node:", error)
		return
	}

	rcv := new(RcvHandler)
	go St.Receive(rcv)
	// End of Setup.
	// TODO: Broadcast your files to neighbours.	

	for j := 0; j < 3 ; j++ {
		error = St.SendMsg(1,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 1: ", error)
			return
		}
		error = St.SendMsg(3,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 3: ", error)
			return
		}
		error = St.SendMsg(4,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 4: ", error)
			return
		} 		
	}
	

	// TODO: It's expected to converge after N second
	// To be able to print a stable graph and shortest
	// path for file.
	N := 10
	time.Sleep(time.Second * time.Duration(N))

	fmt.Println("node 2 done ")

}
