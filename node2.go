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

// var to know the sender 
var sender int;

// TODO: Change this to your current password.
var studentPassword string = "P6Hjqh"

// Implementing ReceiveHandler for student package.
type RcvHandler struct{}

// Handle a message received.
func (rcvHand *RcvHandler) ReceiveHandler(from int, to int, username string,
	content string) {
	// DONOT CHANGE PARAMENTERS OR FUNCTION HEADER.
	// TODO: Implement handling a message received.
	
	sender=from;
	fmt.Println(from, " ", to, username," ", content)
	fmt.Printf("%+v\n", rcvHand)
	//fmt.Println("rcvHand struct: ",rcvHand)
}

func main() {
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

	//sending msgs

	if sender == 1 {
	for j := 0; j < 3 ; j++ {
    	//gg:=strconv.Itoa(M[sorted[j]])
		error = student.SendMsg(3,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 3: ", error)
			return
		}
		error = student.SendMsg(4,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 4: ", error)
			return
		} 
		
	}
	

	} else if sender  == 3 {

		for j := 0; j < 3 ; j++ {
    	//gg:=strconv.Itoa(M[sorted[j]])
		error = student.SendMsg(1,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 1: ", error)
			return
		}
		error = student.SendMsg(4,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 4: ", error)
			return
		} 
		
		}

	} else if sender == 4 {

		for j := 0; j < 3 ; j++ {
    	//gg:=strconv.Itoa(M[sorted[j]])
		error = student.SendMsg(1,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 1: ", error)
			return
		}
		error = student.SendMsg(3,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 3: ", error)
			return
		} 
		
	}
	}

	// TODO: It's expected to converge after N second
	// To be able to print a stable graph and shortest
	// path for file.
	N := 10
	time.Sleep(time.Second * time.Duration(N))
}
