package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	g.Log().SetFlags(glog.F_TIME_TIME | glog.F_FILE_SHORT)
	g.Log().Print("time and short line number")
	g.Log().SetFlags(glog.F_TIME_MILLI | glog.F_FILE_LONG)
	g.Log().Print("time with millisecond and long line number")
	g.Log().SetFlags(glog.F_TIME_STD | glog.F_FILE_LONG)
	g.Log().Print("standard time format and long line number")
}
