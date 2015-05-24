// TODO: In this file you need to build a graph of nodes
// and each node contains a set of files and print all files
// you get, graph and shortest path for specified file.

package main

import (
	"./core/student"
	"fmt"
	"time"
	"strings"
	"strconv"
	"os"
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
    
    ///////////// Floyd Marshall algo intilization arrays
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
	//fmt.Println(from," ",to," ",content)
	lines := strings.Split(content," ")
	P.Files[lines[0]] = int(lines[1][0]) - '0'
	//fmt.Println("after", len(P.Files))
	sentstr := fmt.Sprintf("%s1",content)
	go constg(lines[1])
	for c := 0; c < len(connectedNodes); c++ {
		if( ! strings.Contains(lines[1],string(connectedNodes[c]))) {
			error := St.SendMsg(connectedNodes[c],sentstr)
			x := 2
			time.Sleep(time.Second * time.Duration(x))
			if error != nil {
				fmt.Println("Failed to SendMsg to node",connectedNodes[c],": ", error)
				return
			}	
		}
	}
}

func handleNode(from int, content string){
	tmp := strings.Count(content,"x")
	if tmp == 0 {
		content = content[5:]
		fmt.Println("damn ",content)
		edges := strings.Split(content,":")
		e_sz := len(edges)
		for i := 0 ; i < e_sz ; i++{
			ee := len(edges[i])
			for j := 0 ; j < ee ; j++{
				y := int(edges[i][j]) - '0'
				P.arr[i+1][y] = 1
				P.arr[y][i+1] = 1
				P.next[i+1][y] = y
				P.next[y][i+1] = i+1
			}
		}
		return
	}
	content = strings.Replace(content, "x1", "2", -1)
	content = "node " + content
	fmt.Println("alone ",content)
	for _,c := range connectedNodes{

			St.SendMsg(c,content)
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
			go handleNode(from,lines[1])
		}else{
			_, found := P.Files[lines[0]]
			if( ! found ){
					go handleMsg(from,to,username,content)
			} 
		}
	}
}



///// path reconstruction function to get the shortest path -> Floyd Marshall Algo.
func Path(u int, v int) []int{
	var path []int 
	if (P.next[u][v]  == -1){
			return nil
			}
			path = append(path,u)
			for ; u != v ; {
				u = P.next[u][v]
				path = append(path,u)
			}
			return path
}

func main() {

//// intialize the map that will have all files and thier node

	P.Files = make(map[string]int)
	f_size := len(fileList)
	for i := 0 ; i < f_size ; i++ {
		P.Files[fileList[i]] = 1
	}

	///////////// Floyd Marshall algo intilization arrays

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
	c_sz := len(connectedNodes)
	f_sz := len(fileList)
	// send nodes to other connected nodes
	node := "node 2:x2:x3:x4:x5"
	for c := 0; c < c_sz; c++{
		error := St.SendMsg(connectedNodes[c],node)
		x := 5
			time.Sleep(time.Second * time.Duration(x))
			if error != nil {
				fmt.Println("Failed to SendMsg to node 1: ", error)
				return
			}
	}
	//fmt.Println("Intializing node 1\n");
	for j := 0; j < f_sz ; j++ {
		for c := 0; c < c_sz; c++ {
			sentstr := fmt.Sprintf("%s 1",fileList[j])
			error = St.SendMsg(connectedNodes[c],sentstr)
			x := 5
			time.Sleep(time.Second * time.Duration(x))
			if error != nil {
				fmt.Println("Failed to SendMsg to node 1: ", error)
				return
			}
		}
	}
	// TODO: It's expected to converge after N second
	// To be able to print a stable graph and shortest
	// path for file.
	N :=20
	time.Sleep(time.Second * time.Duration(N))
	
	//fmt.Println("final results")
	//fmt.Println(P.Files)
	//fmt.Println(len(P.Files))
	//fmt.Println("node 1 done ")

/////////////////////////////////////////////// write into output file ////////////////////////////////
	out1,_ := os.Create("output")
	defer out1.Close()    
/*	out1.WriteString("--------------------- List of all files in all nodes ------------------------ \n\n")
    for key,value := range P.Files {
    	    b := strconv.Itoa(value)
		    out1.WriteString(key+" ---> "+b+"\n")
	}*/   

	for i := 1; i < 6 ; i++ {
		fmt.Println(P.arr[i])
		fmt.Println("\n")	
	}	



////////// Floyd Marshall algorithm on array to find the shortest path 
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


out1.WriteString("\n\n--------------------- Graph ------------------------ \n\n")	

out1.WriteString("\n Graph shows the connection between each node and other nodes\n")	

for i := 1; i < 6 ; i++ {
		u := strconv.Itoa(i)
		S2 :=(u + " --> ")
	//	out1.WriteString(u + " --> ")
		for j := 1; j < 6 ; j++ {
			u2 := strconv.Itoa(j)
			if P.arr[i][j] == 1{
				//out1.WriteString(u2 + " ,")
				S2 =S2 + (u2 + " ,")
			}
		}
		if last := len(S2) - 1; last >= 0 && S2[last] == ',' {
        	S2 = S2[:last]
   		 }
		out1.WriteString(S2 + "\n")
}	

out1.WriteString("\n\n---------------------------------------------------- \n\n")


//// Print  the node on which the file was found
fileFound := P.Files["5834591_818124870_n.jpg"]

fileFoundOut := strconv.Itoa(fileFound)
out1.WriteString("\n\nFile -- 5834591_818124870_n.jpg -- found in node: "+ fileFoundOut +"\n\n")

//////// call path function that find the shortesst path betweeen nodes 1 and the node on which the file was found (fileFound)

printPath:= Path(1,fileFound)

out1.WriteString("\n\n--------------------- Shortest Path ------------------------ \n\n")	

for i := 0; i < len(printPath) ; i++ {
	g := strconv.Itoa(printPath[i])
	if (i != len(printPath)-1){
		out1.WriteString(g + " --> ")
	}else{
		out1.WriteString(g + "\n\n")
	}
}	

fmt.Println("after floyed\n")	
for i := 1; i < 6 ; i++ {
		fmt.Println(P.arr[i])
		fmt.Println("\n")	
	}	
fmt.Println(printPath)
	// TODO: Print results in output file.
}
