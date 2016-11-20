package net

import "net"

type Channel struct {
	urlSrc  string
	connSrc net.Conn
	urlDst  string
	connDst net.Conn
}

func (channel Channel) updateDst(url string, con net.Conn) {
	channel.connDst.Close()
	channel.connDst = con
	channel.urlDst = url
}
