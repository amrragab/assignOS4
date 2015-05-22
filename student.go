// Student Framework.

package student

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os/user"
	"strconv"
	"time"
)

var timeStamp int = 0
var BUFFER_SIZE int = 1024 * 32

type StudentComm interface {
	ReceiveHandler(from int, to int, username string, content string)
}

type Student struct {
	UserName string
	UserId   int
	Conn     net.Conn
	Status   bool
}

type Message struct {
	From     int
	To       int
	UserName string
	Content  string
}

type ConnMessage struct {
	Type     string
	Id       int
	UserName string
	Password string
}

// Closes connection for request.
func (student *Student) Close() {
	student.Conn.Close()
}

// Send message to current connection.
func (student *Student) Send(msg []byte) {
	timeStamp += 200000
	if timeStamp > 20000000 {
		timeStamp = 0
	}
	time.Sleep(time.Duration(timeStamp))
	student.Conn.Write(msg)
}

// Connect to specific node.
func (student *Student) Connect(address string, password string) error {
	tcpAddr, error := net.ResolveTCPAddr("tcp", address)
	if error != nil {
		return error
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	} else {
		usr, _ := user.Current()
		userId, _ := strconv.Atoi(usr.Uid)
		msg := &ConnMessage{"user", userId, usr.Username, password}
		buf, error := json.Marshal(msg)
		if error != nil {
			return error
		}
		student.Conn = conn
		student.UserName = usr.Username
		student.UserId = userId
		student.Send(buf)
		student.Status = true
		fmt.Println("Connected to master successfully")
	}
	return nil
}

// Send message for current connection.
func (student *Student) SendMsg(to int, content string) error {
	msg := &Message{-1, to, student.UserName, content}
	buf, error := json.Marshal(msg)
	if error != nil {
		return error
	}
	student.Send(buf)
	return nil
}

// Receive message for current connection.
func (student *Student) Receive(stdComm StudentComm) error {
	buffer := make([]byte, BUFFER_SIZE)
	for {
		time.Sleep(time.Duration(student.random(9, 15)))
		student.Conn.SetReadDeadline(time.Now().Add(150 * time.Second))
		bufferSize, error := student.Conn.Read(buffer)
		if error != nil || error == io.EOF {
			student.Conn.Close()
			student.Status = false
			return error
		}
		// Decode received message
		var recvMessage Message
		err := json.Unmarshal(buffer[0:bufferSize], &recvMessage)
		if err != nil {
			fmt.Println(string(buffer))
			continue
		}
		stdComm.ReceiveHandler(recvMessage.From, recvMessage.To,
			recvMessage.UserName, recvMessage.Content)

		// Cleaning the buffer
		for i := 0; i < BUFFER_SIZE; i++ {
			buffer[i] = 0x00
		}
	}
}

// Generates random integer.
func (student *Student) random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
