package web

import (
	"net/http"
	"github.com/xsank/EasyProxy/src/proxy"
)

const (
	StatisticsUrl = "/statistic"
)

func Statistic(writer http.ResponseWriter, request *http.Request) {
	proxy.Record()
	Render(writer, "statistic", StatisticHtml, proxy.StatisticData())
}