package core

import (
	"Vanta/client/models"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

const retryLimit int = 3

func ListenForMessages(ws *websocket.Conn) models.DataPayload {

	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			log.Println("Error Reading Data", err)
		}
		var Data models.DataPayload
		err = json.Unmarshal(buf[:n], &Data)
		switch Data.Type {
		case models.TextMessage:
			HandleMessage(&Data)
		case models.Ping:
			HandlePing(1, ws)

		case models.KeyExchange:
			HandleKeyExchange(&Data)
		default:
			log.Fatal("Failed to read data type.", Data.Type)

		}

	}
}

func HandleMessage(d *models.DataPayload) {
	s, ok := d.Data.(string)
	if !ok {
		log.Printf("Data of type %T, is not of type string.", d.Data)
		return
	}
	fmt.Printf("%s: %s\n", d.Username, s)
}

// HandlePing This function should just immediately send a ping back to the server.
func HandlePing(retries int, ws *websocket.Conn) {
	if retries > retryLimit {
		return
	}
	ping := models.DataPayload{
		Type:     models.Ping,
		Username: name,
		Data:     "Ping",
	}
	p, err := json.Marshal(ping)
	if err != nil {
		log.Println("Failed to ping server!\nretrying!")
		HandlePing(retries+1, ws)
	}
	if _, writeErr := ws.Write(p); err != nil {
		log.Println("Error writing to server!", writeErr)
	}

}

// HandleKeyExchange Work on this function, I need to work on correctly handling the cases for each key type.
func HandleKeyExchange(d *models.DataPayload) {
	s, ok := d.Data.(models.Key)
	if !ok {
		log.Printf("Data of type %T, is not of type Key.", s)
		return

	}
	switch s.KeyType {
	case models.IdentityKey:
		fmt.Println("Do something Later here!")
	case models.EphemeralKey:
		fmt.Println("Do something Later here!")
	case models.PreKeys:
		fmt.Println("Do something Later here!")
	}
}
