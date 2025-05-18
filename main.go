package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

type MessageHandler func(*websocket.Conn, json.RawMessage)

/*var messageRouter = map[string]MessageHandler {
	"movePiece":
}*/

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func reverseString(s string) string {
	runes := []rune(s) // Convert to rune slice to handle Unicode properly
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func websocketHandler(responseWriter http.ResponseWriter, request *http.Request) {
	connection, err := upgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		fmt.Println("Error on http to websocket connection upgrade: ", err)
		return
	}

	go handleWebsocketConnection(connection)
}

func handleWebsocketConnection(connection *websocket.Conn) {
	defer connection.Close()
	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			fmt.Println("Error on reading websocket message: ", err)
			break
		}
		flippedMessage := reverseString(string(message))
		err = connection.WriteMessage(websocket.TextMessage, []byte(flippedMessage))
		if err != nil {
			fmt.Println("Error on writing websocket message: ", err)
			break
		}
	}
}

func handl() {

}

func main() {
	godotenv.Load(".env")
	http.HandleFunc("/websocketTest", websocketHandler)
	httpPort := os.Getenv("SERVER_PORT")
	fmt.Printf("Initializing http server on port %s", httpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil)
	if err != nil {
		fmt.Println("Error on http server initialization: ", err)
	}

}
