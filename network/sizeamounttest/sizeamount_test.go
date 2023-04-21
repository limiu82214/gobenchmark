package sizeamounttest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

const testMsg string = "abcdefghijklmnopqrstuvwxyz"
const repeatCount = 40000

func BenchmarkWebSocketAmount(b *testing.B) {
	// 啟動 WebSocket 伺服器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}))
	defer server.Close()

	// 建立 WebSocket 客戶端
	client := &websocket.Dialer{}
	wsURL := fmt.Sprintf("ws://%s", server.Listener.Addr())

	conn, _, err := client.Dial(wsURL, nil)
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	// 傳送訊息到伺服器並等待回應
	for j := 0; j < repeatCount; j++ {
		for i := 0; i < b.N; i++ {
			err = conn.WriteMessage(websocket.TextMessage, []byte(testMsg))
			if err != nil {
				b.Fatal(err)
			}

			_, message, err := conn.ReadMessage()
			if err != nil {
				b.Fatal(err)
			}

			if string(message) != testMsg {
				b.Fatalf("Unexpected response: %s", message)
			}
		}
	}
}

func BenchmarkWebSocketSize(b *testing.B) {
	// 啟動 WebSocket 伺服器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}))
	defer server.Close()

	// 建立 WebSocket 客戶端
	client := &websocket.Dialer{}
	wsURL := fmt.Sprintf("ws://%s", server.Listener.Addr())

	conn, _, err := client.Dial(wsURL, nil)
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	// 傳送訊息到伺服器並等待回應
	msg := strings.Repeat(testMsg, repeatCount)
	for i := 0; i < b.N; i++ {
		err = conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			b.Fatal(err)
		}

		_, message, err := conn.ReadMessage()
		if err != nil {
			b.Fatal(err)
		}

		if string(message) != msg {
			b.Fatalf("Unexpected response: %s", message)
		}
	}
}
