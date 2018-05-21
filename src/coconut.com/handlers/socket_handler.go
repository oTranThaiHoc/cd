package handlers

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"github.com/Tomasen/realip"
	"coconut.com/worker"
)

var (
	upgrader = websocket.Upgrader{}
	Conn = make(map[string]*websocket.Conn)
)

var SocketHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	clientIp := realip.FromRequest(r)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	existingConn, ok := Conn[clientIp]
	if ok {
		// existing connection, close it
		existingConn.Close()
	}
	log.Printf("new connection: %v\n", clientIp)
	serveConnection(c)
})

func NotifyJobDone(job worker.Job) {
	log.Printf("notify job done: %v\n", job.Title)
	conn, ok := Conn[job.ClientIp]
	if !ok {
		log.Printf("notify job done failed, client not found: %v\n", job.ClientIp)
		return
	}
	conn.WriteMessage(0, []byte("reload"))
}

func serveConnection(c *websocket.Conn) {

}