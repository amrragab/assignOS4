// TODO: In this file you need to build a graph of nodes
// and each node contains a set of files and print all files
// you get, graph and shortest path for specified file.

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
	AdjList [6]string
	arr [6][6]int
	next [6][6]int

}

var P Param


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

//5leto global
var St= new(student.Student)

func constg(str string){
	str = fmt.Sprintf("%s1",str)
		size := len(str)
		for i := size - 1; i > 0 ; i--{
			x := int(str[i]) - '0'
			if ( ! strings.Contains(P.AdjList[x],string(str[i-1])) ){
				//P.AdjList[x] = append(P.AdjList[x],string(str[i-1]) )
				P.AdjList[x] = fmt.Sprintf("%s%s",P.AdjList[x],string(str[i-1]))
			}
			// ex let str = "321"
			y := int(str[i-1]) - '0'
			P.arr[x][y] = 1
			P.arr[y][x] = 1
			P.next[x][y] = y
			P.next[y][x] = x
		}


}
// to handle the msg after recive handler is caller
func handleMsg(from int, to int, username string,content string){
				fmt.Println(from," ",to," ",content)
				//if from == connectedNodes[0]{
				//	return	
				//}
				lines := strings.Split(content," ")
				P.Files[lines[0]] = int(lines[1][0]) - '0'
				fmt.Println("after", len(P.Files))
				sentstr := fmt.Sprintf("%s1",content)
				go constg(lines[1])
				if( ! strings.Contains(lines[1],string(connectedNodes[0]))) {
				error := St.SendMsg(connectedNodes[0],sentstr)
					x := 2
				time.Sleep(time.Second * time.Duration(x))
				if error != nil {
					fmt.Println("Failed to SendMsg to node",connectedNodes[0],": ", error)
					return
					}	
				}
				//fmt.Println("File ",0,": ", content)
}

// Handle a message received.
func (rcvHand *RcvHandler) ReceiveHandler(from int, to int, username string,
	content string) {
	//tmp_str := fmt.Sprintf("%s is not connected at node %d",username,to)
	lines := strings.Split(content, " ")
	_, found := P.Files[lines[0]]
	if( ! strings.Contains(content,"not connected") && ! found ){
			go handleMsg(from,to,username,content)
	}
	// DONOT CHANGE PARAMENTERS OR FUNCTION HEADER.
	// TODO: Implement handling a message received.
}

/*procedure Path(u, v)
   if next[u][v] = null then
       return []
   path = [u]
   while u ≠ v
       u ← next[u][v]
       path.append(u)
   return path*/
func Path(u int, v int) []int{
	if ( next[u][v]  == -1){
				return []
			}
			path = append(path,u)
			for ; u != v ; {
				u = next[u][v]
				path = append(path,u)
			}
			return path
}

func main() {


	P.Files = make(map[string]int)
	for i := 0 ; i < 6 ; i++{
		P.AdjList[i] = ""
	}
	f_size := len(fileList)
	for i := 0 ; i < f_size ; i++ {
		P.Files[fileList[i]] = 1
	}
	for i := 0; i < 6 ; i++ {
		for j := 0 ; j < 6 ; j++ {
			P.arr[i][j] = 1000
			P.next[i][j] = -1
		}
		P.arr[i][i] = 0
	}
	//fmt.Println("length", len(P.Files))
	
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
		fmt.Println("Intializing node 1\n");
	for j := 0; j < 3 ; j++ {
		sentstr := fmt.Sprintf("%s 1",fileList[j])
		error = St.SendMsg(connectedNodes[0],sentstr)
		x := 5
				time.Sleep(time.Second * time.Duration(x))
		if error != nil {
			fmt.Println("Failed to SendMsg to node 1: ", error)
			return
		}
	}
	// TODO: It's expected to converge after N second
	// To be able to print a stable graph and shortest
	// path for file.
	N :=20
	time.Sleep(time.Second * time.Duration(N))
	fmt.Println("final results")
	fmt.Println(P.Files)
	//fmt.Println(P.AdjList)
	sz := len(P.AdjList)
	for i := 0 ; i < sz ; i++{
		fmt.Println(i,":",P.AdjList[i])
	}

	fmt.Println(len(P.Files))
	fmt.Println("node 1 done ")
	for i := 1; i < 6 ; i++ {
		fmt.Println(P.arr[i])
		fmt.Println("\n")	
	}	
for k := 1; k < 6 ; k++ {
	for i := 1; i < 6 ; i++ {
		for j := 1; j < 6 ; j++ {
               if (P.arr[i][k] + P.arr[k][j]) < P.arr[i][j]{
               P.arr[i][j] = P.arr[i][k] + P.arr[k][j];
               P.next[i][j]=P.next[i][k]
               }
		}
	}
}
fmt.Println("after floyed\n")	
for i := 1; i < 6 ; i++ {
		fmt.Println(P.arr[i])
		fmt.Println("\n")	
	}	
	// TODO: Print results in output file.
}
