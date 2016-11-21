package structure

import "net"

type Channel struct {
	srcConn net.Conn
	dstConn net.Conn
}

func (channel *Channel) UpdateDst(conn net.Conn) {
	channel.dstConn.Close()
	channel.dstConn = conn
}

func (channel *Channel) Close() {
	channel.srcConn.Close()
	channel.dstConn.Close()
}
