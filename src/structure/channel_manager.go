package structure

import "errors"

type ChannelManager struct {
	channels []Channel
	mapSrc   map[string]*Channel
	mapDst   map[string]*Channel
}

func (channelManager *ChannelManager) Init() {
	channelManager.channels = make([]Channel, 0)
	channelManager.mapSrc = make(map[string]*Channel)
	channelManager.mapDst = make(map[string]*Channel)
}

func (channelManager *ChannelManager) PutChannel(channel *Channel) {
	channelManager.channels = append(channelManager.channels, *channel)
	channelManager.mapSrc[channel.SrcUrl] = channel
	channelManager.mapDst[channel.DstUrl] = channel
}

func (channelManager *ChannelManager) DeleteChannel(channel *Channel) {
	index := 0
	for i, cnl := range channelManager.channels {
		if cnl == *channel {
			index = i
			break
		}
	}
	channelManager.channels = append(channelManager.channels[:index], channelManager.channels[index + 1:]...)
	channelManager.deleteMap(channelManager.mapSrc, channel.SrcUrl)
	channelManager.deleteMap(channelManager.mapDst, channel.DstUrl)
}

func (channelManager *ChannelManager) GetChannels() []Channel {
	return channelManager.channels
}

func (channelManager *ChannelManager) Check() (error, error) {
	var srcErr, dstErr error
	channelLenth := len(channelManager.channels)
	if len(channelManager.mapSrc) != channelLenth {
		srcErr = errors.New("client socket close maybe error")
	}
	if len(channelManager.mapDst) != channelLenth {
		dstErr = errors.New("server socket close maybe error")
	}
	return srcErr, dstErr
}

func (channelManager *ChannelManager) deleteMap(_map map[string]*Channel, url string) {
	_, ok := _map[url]
	if ok {
		delete(_map, url)
	}
}