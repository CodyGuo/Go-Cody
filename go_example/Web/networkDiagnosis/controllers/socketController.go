package controllers

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 8192

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Time to wait before force close on connection.
	closeGracePeriod = 10 * time.Second
)

const (
	ping = iota
	tcpdump
	traceroute
	nslookup
)

func pumpStdin(ws *websocket.Conn, w io.Writer) {
	defer ws.Close()
	ws.SetReadLimit(maxMessageSize)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var ipInfo2 map[string]interface{}
		if err := websocket.ReadJSON(ws, &ipInfo2); err != nil {
			return
		}

		if ipInfo2["stop"].(bool) {
			log.Printf("received -----> stop {%v}\n", ipInfo2["stop"])
			ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		}
	}
}

func gbkDecode(f io.Reader) *mahonia.Reader {
	decoder := mahonia.NewDecoder("gbk")
	return decoder.NewReader(f)
}

func pumpStdout(ws *websocket.Conn, r io.Reader, done chan struct{}) {
	defer func() {
	}()
	r = gbkDecode(r)
	s := bufio.NewScanner(r)
	for s.Scan() {
		ws.SetWriteDeadline(time.Now().Add(writeWait))
		if err := ws.WriteMessage(websocket.TextMessage, s.Bytes()); err != nil {
			ws.Close()
			break
		}
	}
	if s.Err() != nil {
		log.Println("scan ----->", s.Err())
	}
	close(done)

	ws.SetWriteDeadline(time.Now().Add(writeWait))
	ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	time.Sleep(closeGracePeriod)
	ws.Close()
}

func wsPing(ws *websocket.Conn, done chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				log.Println("wsPing ----->", err)
			}
		case <-done:
			return
		}
	}
}

func internalError(ws *websocket.Conn, msg string, err error) {
	log.Println(msg, err)
	ws.WriteMessage(websocket.TextMessage, []byte("Internal server error."))
}

var upgrader = websocket.Upgrader{}

type WsPing struct {
	beego.Controller
}

func (wp *WsPing) Get() {
	w, r := wp.Ctx.ResponseWriter, wp.Ctx.Request
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade ----->", err)
		return
	}
	defer ws.Close()

	var ipInfo map[string]interface{}
	if err := websocket.ReadJSON(ws, &ipInfo); err != nil {
		return
	}
	log.Println("received ----->", ipInfo)

	typeTool, _ := strconv.Atoi(fmt.Sprint(ipInfo["type"]))
	args := fmt.Sprint(ipInfo["args"])
	cmds := strings.Split(args, " ")
	cmd := exec.Command(types[typeTool], cmds...)
	inw, _ := cmd.StdinPipe()
	outr, _ := cmd.StdoutPipe()
	cmd.Start()

	stdoutDone := make(chan struct{})
	go pumpStdout(ws, outr, stdoutDone)
	go wsPing(ws, stdoutDone)

	pumpStdin(ws, inw)

	// Some commands will exit when stdin is closed.
	inw.Close()

	cmd.Process.Kill()

	// Other commands need a bonk on the head.
	if err := cmd.Process.Signal(os.Interrupt); err != nil {
		log.Println("inter error ----->", err)
	}

	select {
	case <-stdoutDone:
		log.Println("done ----->", types[typeTool])
	case <-time.After(time.Second):
		if err := cmd.Process.Signal(os.Kill); err != nil {
			log.Println("timeout ----->", err)
		}
		<-stdoutDone
	}

	cmd.Wait()

}
