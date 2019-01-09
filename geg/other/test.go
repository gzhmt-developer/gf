package main

import (
    "gitee.com/johng/gf/g/container/gtype"
    "gitee.com/johng/gf/g/os/glog"
    "gitee.com/johng/gf/g/os/gwheel"
    "time"
)

func main() {
    v := gtype.NewInt()
    w := gwheel.New(1, 10*time.Millisecond)
    glog.Println("start")
    for i := 0; i < 100000; i++ {
        w.AddOnce(time.Second, func() {
           v.Add(1)
        })
    }
    glog.Println("end")
    time.Sleep(1020*time.Millisecond)
    glog.Println(v.Val())
    //gwheel.AddSingleton(time.Second, func() {
    //    fmt.Println(time.Now().String())
    //})
    //select { }
}
