package impl

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

/**
 * @Author  jackie.lqj
 * @Date  2022/5/18 19:38
 * @Description
 */

type Connection struct {
	wsConn *websocket.Conn
	// 读取websocket的channel
	inChan chan []byte
	// 给websocket写消息的channel
	outChan chan []byte
	// 关闭标志
	closeChan chan byte
	mutex     sync.Mutex
	// closeChan 状态
	isClosed bool
}

// InitConnection 初始化长连接
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}
	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()
	return
}

// ReadMessage 读取websocket消息
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// WriteMessage 发送消息到websocket
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// Close 关闭连接
func (conn *Connection) Close() {
	// 线程安全的Close,可重入
	_ = conn.wsConn.Close()

	// 只执行一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	defer conn.Close()
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			return
		}
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			return
		}
	}
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	defer conn.Close()
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			return
		}
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			return
		}
	}
}
