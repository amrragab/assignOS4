// TODO: In this file you need to build a graph of nodes
// and each node contains a set of files and print all files
// you get, graph and shortest path for specified file.

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
var masterAddr string = "10.0.0.4:46321"
var connectedNodes = []int{2}
var fileList = []string{
	"1939620_437577509_n.jpg",
	"5978610_937577509_n.jpg",
	"6436120_737577509_n.jpg"}

// TODO: Change this to your current password.
var studentPassword string = "P6Hjqh"

// Implementing ReceiveHandler for student package.
type RcvHandler struct{}

// Handle a message received.
func (rcvHand *RcvHandler) ReceiveHandler(from int, to int, username string,
	content string) {
	// DONOT CHANGE PARAMENTERS OR FUNCTION HEADER.
	// TODO: Implement handling a message received.
	fmt.Println(from, " ", to, username," ", content)
}

func main() {
	S := 10
	time.Sleep(time.Second * time.Duration(S))

	// Setup connection to master of current node.
	student := new(student.Student)
	error := student.Connect(masterAddr, studentPassword)
	if error != nil {
		fmt.Println("Failed to connect to master node:", error)
		return
	}

	// Link implementation of ReceiveHandler to student.
	rcv := new(RcvHandler)
	go student.Receive(rcv)
	// End of Setup.

	// TODO: Broadcast your files to neighbours.
	error = student.SendMsg(2,fileList[0])
	if error != nil {
		fmt.Println("Failed to SendMsg to master node:", error)
		return
	} 
	// TODO: It's expected to converge after N second
	// To be able to print a stable graph and shortest
	// path for file.
	N := 10
	time.Sleep(time.Second * time.Duration(N))

	// TODO: Print results in output file.
}
