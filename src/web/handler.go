package web

import (
	"net/http"
	"fmt"
)

const(
	StatisticsUrl="/statistic"
)

func Statistic(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer,"hello world")
}