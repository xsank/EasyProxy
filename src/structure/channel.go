package structure

import "net"

const ChannelPairNum = 2

type Channel struct {
	SrcConn net.Conn
	DstConn net.Conn
}

func (channel *Channel) SrcUrl() string {
	return channel.SrcConn.RemoteAddr().String()
}

func (channel *Channel) DstUrl() string {
	return channel.DstConn.RemoteAddr().String()
}

func (channel *Channel) Close() {
	channel.SrcConn.Close()
	channel.DstConn.Close()
}
