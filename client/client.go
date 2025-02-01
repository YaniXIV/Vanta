package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"os"
	"sync"
 
)

//const origin string = "http://localhost/"
//const url string = "ws://localhost:1444/ws"

const defaultIp string = "localhost"

// var ip string
const port string = "1444"

var name string

var pubKey, privKey = KeyGen()
// Hardcoding ip and port for now. *Change Later*

type msg struct {
	Usrname string `json:"usrname"`
	Text    string `json:"Text"`
}

func main() {
	initWebsocketClient()
	fmt.Println("End of Program reached")

}

func initWebsocketClient() {
	fmt.Println("<Client Side> Starting Client!")
	//fmt.Printf("ws://%s:%s/wq", ip, port)
	//fmt.Printf("http://%s/", ip)
	var ip string
	fmt.Println("Please Enter target Ip: ")
	fmt.Scanln(&ip)
	if ip == "" {
		ip = defaultIp
	}
	fmt.Println("Please Enter Username: ")
	fmt.Scanln(&name)

	//ws, err := websocket.DialConfig()
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s:%s/ws", ip, port), "", fmt.Sprintf("https://%s/", ip))
	if err != nil {
		log.Fatal(err)
	}
  /*
  replace first messaqge with the public key! 
	if _, err = ws.Write([]byte("<Client Side> connected to Server!")); err != nil {
		log.Fatal(err)
	}
 */
  if _, err = ws.Write(pubKey.Bytes());err != nil{
    log.Fatal(err)
  }
	// Use WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Start sendMessage and listenForMessages goroutines
	wg.Add(2) // We are starting two goroutines
	go func() {
		defer wg.Done() // Decrements the counter when the goroutine finishes
		sendMessage(ws)
	}()
	go func() {
		defer wg.Done()
		receiveMessage(ws)
	}()
	// Wait for all goroutines to finish
	wg.Wait()
}

func listenForMessages(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			fmt.Println("<Client Side> Error reading", err)
		}
		fmt.Println("<Client Side> Message from Server:", string(buf[:n]))
	}
}

func sendMessage(ws *websocket.Conn) {
	var usrin string
	fmt.Println("Attempting to send message")

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("SendMessage: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input.\nTry again")
			continue
		}

		usrin = input[:len(input)-1]
		if usrin == "q" {
			break
		}
		fmt.Println("usrin is ", usrin)
		b, err := prepareMsg(usrin)
		if err != nil {
			fmt.Println("Msg failed to send: ", err)
			continue
		}
		if _, err = ws.Write(b); err != nil {
			fmt.Println("Error with writing to server", err)
		}
	}

}

func prepareMsg(m string) ([]byte, error) {
	//fmt.Println(m)
	c := msg{name, m}
	b, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func receiveMessage(ws *websocket.Conn) []byte {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection broke!")
				continue
			}
			fmt.Println("Read error", err)
		}
		msg := buf[:n]
		fmt.Println(string(msg))
	}
}

/*
what is the current task
figuring out encryption. 
^^ I am trying to get the server to contain 
*/
