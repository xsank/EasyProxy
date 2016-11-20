package net

import "net"

type ChannelManager struct {
	mapSrc map[string]Channel
	mapDst map[string]Channel
}

func (channelManager *ChannelManager) updateDst(urlOld string, urlNew string, con net.Conn) {
	channelManager.mapDst[urlOld].updateDst(urlNew, con)
}
