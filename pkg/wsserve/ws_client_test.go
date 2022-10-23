/*
 * @Author: mknight(tianyh)
 * @Mail: 824338670@qq.com
 * @Date: 2022-09-16 14:18:14
 * @LastEditTime: 2022-09-16 14:21:52
 * @FilePath: \trcell\pkg\wsserve\ws_client_test.go
 */
package wsserve

import (
	"log"
	"net/url"
	"testing"
	"time"
	"trcell/pkg/loghlp"

	"github.com/gorilla/websocket"
)

func TestWSClient(t *testing.T) {
	u := url.URL{Scheme: "ws", Host: "localhost:7200", Path: "/ws"}
	loghlp.Infof("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
			// case <-interrupt:
			// 	log.Println("interrupt")

			// 	// Cleanly close the connection by sending a close message and then
			// 	// waiting (with timeout) for the server to close the connection.
			// 	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			// 	if err != nil {
			// 		log.Println("write close:", err)
			// 		return
			// 	}
			// 	select {
			// 	case <-done:
			// 	case <-time.After(time.Second):
			// 	}
			// 	return
		}
	}
}
