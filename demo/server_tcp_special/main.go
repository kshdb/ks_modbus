package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	modbus "github.com/kshdb/ks_modbus"
)

func main() {
	srv := modbus.NewTCPServerSpecial().
		SetOnConnectHandler(func(c *modbus.TCPServerSpecial) error {
			_, err := c.UnderlyingConn().Write([]byte("hello world"))
			return err
		}).
		SetConnectionLostHandler(func(c *modbus.TCPServerSpecial) {
			log.Println("connect lost")
		}).
		SetKeepAlive(true, time.Second*20, func(c *modbus.TCPServerSpecial) {
			_, _ = c.UnderlyingConn().Write([]byte("keep alive"))
		})
	if err := srv.AddRemoteServer("127.0.0.1:3001"); err != nil {
		panic(err)
	}
	srv.LogMode(true)
	srv.AddNodes(
		modbus.NewNodeRegister(
			1,
			0, 10, 0, 10,
			0, 10, 0, 10),
		modbus.NewNodeRegister(
			2,
			0, 10, 0, 10,
			0, 10, 0, 10),
		modbus.NewNodeRegister(
			3,
			0, 10, 0, 10,
			0, 10, 0, 10),
	)

	if err := srv.Start(); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":6060", nil); err != nil {
		panic(err)
	}
}
