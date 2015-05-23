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

type Param struct {
	Files map[string]int
	//AdjList arr[5]string
}


var P Param

// Global Declarations.
var masterAddr string = "10.0.0.8:46321"
var connectedNodes = []int{4}
var fileList = []string{
	"2838472_973816382_n.jpg",
	"5834591_818124870_n.jpg",
	"5579596_151574987_n.jpg"}


// TODO: Change this to your current password.
var studentPassword string = "P6Hjqh"

// Implementing ReceiveHandler for student package.
type RcvHandler struct{}

//5leto global
var St= new(student.Student)

// to handle the msg after recive handler is caller
func handleMsg(from int, to int, username string,content string){
				P.Files[content] = from
				if from == connectedNodes[0]{
					return
				}
				error := St.SendMsg(connectedNodes[0],content)
				if error != nil {
					fmt.Println("Failed to SendMsg to node",connectedNodes[0],": ", error)
					return
					}	
				fmt.Println("File ",0,": ", content)
}

// Handle a message received.
func (rcvHand *RcvHandler) ReceiveHandler(from int, to int, username string,
	content string) {
	// DONOT CHANGE PARAMENTERS OR FUNCTION HEADER.
	// TODO: Implement handling a message received.
	go handleMsg(from,to,username,content)
}

func main() {
	P.Files = make(map[string]int)
	f_size := len(fileList)
	for i := 0 ; i < f_size ; i++ {
		P.Files[fileList[i]] = 5
	}
	fmt.Println(P.Files)
	S := 5
	time.Sleep(time.Second * time.Duration(S))

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

	// TODO: Broadcast your files to neighbours.

		fmt.Println("Intializing node 5\n");
	for j := 0; j < 3 ; j++ {
		error = St.SendMsg(connectedNodes[0],fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 1: ", error)
			return
		}
	}
	// TODO: It's expected to converge after N second
	// To be able to print a stable graph and shortest
	// path for file.
	N := 10
	time.Sleep(time.Second * time.Duration(N))
}
