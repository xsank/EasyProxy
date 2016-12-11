package structure

import (
	"sync"
	"errors"
	"github.com/xsank/EasyProxy/src/util"
)

type ChannelManager struct {
	channels []Channel
	mapSrc   map[string]*Channel
	mapDst   map[string]*Channel
	mutex    *sync.Mutex
}

func (channelManager *ChannelManager) Init() {
	channelManager.channels = make([]Channel, 0)
	channelManager.mapSrc = make(map[string]*Channel)
	channelManager.mapDst = make(map[string]*Channel)
	channelManager.mutex = new(sync.Mutex)
}

func (channelManager *ChannelManager) PutChannel(channel *Channel) {
	channelManager.mutex.Lock()
	channelManager.channels = append(channelManager.channels, *channel)
	channelManager.mapSrc[channel.SrcUrl()] = channel
	channelManager.mapDst[channel.DstUrl()] = channel
	defer channelManager.mutex.Unlock()
}

func (channelManager *ChannelManager) DeleteChannel(channel *Channel) {
	channelManager.mutex.Lock()
	index := util.SliceIndex(channelManager.channels, *channel)
	if index >= 0 {
		channelManager.channels = append(channelManager.channels[:index], channelManager.channels[index + 1:]...)
		deleteMap(channelManager.mapSrc, channel.SrcUrl())
		deleteMap(channelManager.mapDst, channel.DstUrl())
	}
	defer channelManager.mutex.Unlock()
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

func deleteMap(_map map[string]*Channel, url string) {
	_, ok := _map[url]
	if ok {
		delete(_map, url)
	}
}

func (channelManager *ChannelManager) Clean() {
	for _, channel := range channelManager.channels {
		deleteMap(channelManager.mapSrc, channel.SrcUrl())
		deleteMap(channelManager.mapDst, channel.DstUrl())
		channel.Close()
	}
	channelManager.channels = channelManager.channels[:0]
}