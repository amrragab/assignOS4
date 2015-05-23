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
var masterAddr string = "10.0.0.7:46321"
var connectedNodes = []int{2, 3, 5}
var fileList = []string{
	"3737941_134967132_n.jpg",
	"1935851_722579545_n.jpg",
	"1512714_286659690_n.jpg"}

var sendLoop int = 3
	
// var to know the sender 
var sender int = 0

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
	S := 10
	time.Sleep(time.Second * time.Duration(S))

	if sender == 0 {
	for j := 0; j < 3 ; j++ {
		fmt.Println("node 4 sender 0\n");
		error = student.SendMsg(2,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 2: ", error)
			return
		}
		error = student.SendMsg(3,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 3: ", error)
			return
		}
		error = student.SendMsg(5,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 5: ", error)
			return
		}
	time.Sleep(time.Second * time.Duration(sendLoop))
		
	}
	
	} else if sender == 2 {
	for j := 0; j < 3 ; j++ {
		fmt.Println("node 4 sender 2\n");
		error = student.SendMsg(3,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 3: ", error)
			return
		}
		error = student.SendMsg(5,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 5: ", error)
			return
		} 
	time.Sleep(time.Second * time.Duration(sendLoop))
		
	}
	

	} else if sender  == 3 {
		for j := 0; j < 3 ; j++ {
		fmt.Println("node 4 sender 3\n");
		error = student.SendMsg(2,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 2: ", error)
			return
		}
		error = student.SendMsg(5,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 5: ", error)
			return
		} 
		time.Sleep(time.Second * time.Duration(sendLoop))
	}

	} else if sender == 5 {
		for j := 0; j < 3 ; j++ {
			fmt.Println("node 4 sender 5\n");
		error = student.SendMsg(2,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 2: ", error)
			return
		}
		error = student.SendMsg(3,fileList[j])
		if error != nil {
			fmt.Println("Failed to SendMsg to node 3: ", error)
			return
		} 
		
		time.Sleep(time.Second * time.Duration(sendLoop))	
	}
	}
	// TODO: It's expected to converge after N second
	// To be able to print a stable graph and shortest
	// path for file.
	N := 50
	time.Sleep(time.Second * time.Duration(N))
}
