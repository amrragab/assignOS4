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
	AdjList arr[5]string
}
// Global Declarations.
var masterAddr string = "10.0.0.6:46321"
var connectedNodes = []int{2, 4}
var fileList = []string{
	"4294923_402918889_n.jpg",
	"1107509_945113888_n.jpg",
	"1092345_660561345_n.jpg"}
/*
var sendLoop int = 3

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
		for c := 0; c < 2; c++ {
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
}


// Handle a message received.
func (rcvHand *RcvHandler) ReceiveHandler(from int, to int, username string,
	content string) {
	// DONOT CHANGE PARAMENTERS OR FUNCTION HEADER.
	// TODO: Implement handling a message received.

	go handleMsg(from,to,username,content)
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

	// Link implementation of ReceiveHandler to student.
	rcv := new(RcvHandler)
	go St.Receive(rcv)
	// End of Setup.

	// TODO: Broadcast your files to neighbours.

	for j := 0; j < 3 ; j++ {
		fmt.Println("Intializing node 3\n");
		for c := 0; c < 2; c++ {
				error = St.SendMsg(connectedNodes[c],fileList[j])
				if error != nil {
					fmt.Println("Failed to SendMsg to node",connectedNodes[c],": ", error)
					return
				} 
			}		
	}
	


	// TODO: It's expected to converge after N second
	// To be able to print a stable graph and shortest
	// path for file.
	N := 10
	time.Sleep(time.Second * time.Duration(N))

	fmt.Println("node 2 done ")

}

