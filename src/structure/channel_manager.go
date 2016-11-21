package structure

import "net"

type ChannelManager struct {
	mapSrc map[string]*Channel
	mapDst map[string]*Channel
}

func (channelManager *ChannelManager) Init() {
	channelManager.mapSrc = make(map[string]*Channel)
	channelManager.mapDst = make(map[string]*Channel)
}

func (channelManager *ChannelManager) UpdateDst(urlOld string, newConn net.Conn) {
	channelManager.mapDst[urlOld].UpdateDst(newConn)
}

func (channelManager *ChannelManager) PutChannelPair(srcConn net.Conn, dstConn net.Conn) {
	channel := &Channel{srcConn:srcConn, dstConn:dstConn}
	channelManager.mapSrc[srcConn.RemoteAddr().String()] = channel
	channelManager.mapDst[dstConn.LocalAddr().String()] = channel
}

func (channelManager *ChannelManager) CloseBySrcConn(srcConn net.Conn) {
	channelManager.mapSrc[srcConn.RemoteAddr().String()].Close()
}

func (channelManager *ChannelManager) CloseByDstConn(dstConn net.Conn) {
	channelManager.mapSrc[dstConn.LocalAddr().String()].Close()
}