package handlers

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"coconut.com/worker"
	"time"
	"net"
	"strings"
)

var (
	upgrader = websocket.Upgrader{}
	Conn = make(map[string]*websocket.Conn)
)

var SocketHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	clientPort := r.RemoteAddr
	// get client port
	if strings.ContainsRune(r.RemoteAddr, ':') {
		_, clientPort, _ = net.SplitHostPort(r.RemoteAddr)
	}
	existingConn, ok := Conn[clientPort]
	if ok {
		// existing connection, close it
		existingConn.Close()
	}
	Conn[clientPort] = c
	log.Printf("new connection: %v\n", r.RemoteAddr)
	serveConnection(c, clientPort)
})

func NotifyJobDone(job worker.Job) {
	log.Printf("notify job done: %v, client ip: %v\n", job.Title, job.ClientIp)
	log.Printf("call connections: %v\n", Conn)
	clientPort := job.ClientIp
	// get client port
	if strings.ContainsRune(job.ClientIp, ':') {
		_, clientPort, _ = net.SplitHostPort(job.ClientIp)
	}
	c, ok := Conn[clientPort]
	if ok {
		log.Printf("notify to client: %v\n", job.ClientIp)
		c.WriteMessage(websocket.TextMessage, []byte("reload"))
	} else {
		// broadcast
		log.Println("notify broadcast")
		for _, c := range Conn {
			c.WriteMessage(websocket.TextMessage, []byte("reload"))
		}
	}
}

func serveConnection(c *websocket.Conn, port string) {
	go func() {
		for {
			_, _, err := c.ReadMessage()
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
				log.Printf("close connection: %v\n", c.RemoteAddr())
				c.Close()
				delete(Conn, port)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
