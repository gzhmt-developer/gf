// Copyright 2019 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package gwheel

import (
    "gitee.com/johng/gf/g/container/glist"
    "gitee.com/johng/gf/g/container/gtype"
    "time"
)

// 循环任务管理对象
type Wheel struct {
    slots     []*glist.List   // 所有的循环任务项, 按照Slot Number进行分组
    number    int             // Slot Number
    closed    chan struct{}   // 停止事件
    ticks     *gtype.Int      // 当前时间轮已转动的刻度数量
    ticker    *time.Ticker    // 时间轮刻度间隔
    interval  int64           // 时间间隔(slot时间长度, 纳秒)
}

// 创建使用默认值的时间轮
func NewDefault() *Wheel {
    return New(gDEFAULT_SLOT_NUMBER, gDEFAULT_WHEEL_INTERVAL)
}

// 创建自定义的循环任务管理对象
func New(slot int, interval time.Duration) *Wheel {
    w := &Wheel {
        slots     : make([]*glist.List, slot),
        number    : slot,
        closed    : make(chan struct{}, 1),
        ticks     : gtype.NewInt(),
        ticker    : time.NewTicker(interval),
        interval  : interval.Nanoseconds(),
    }
    for i := 0; i < w.number; i++ {
        w.slots[i] = glist.New()
    }
    w.startLoop()
    return w
}

// 添加循环任务
func (w *Wheel) Add(interval time.Duration, job JobFunc) (*Entry, error) {
    return w.newEntry(interval, job, false, gDEFAULT_TIMES)
}

// 添加单例运行循环任务
func (w *Wheel) AddSingleton(interval time.Duration, job JobFunc) (*Entry, error) {
    return w.newEntry(interval, job, true, gDEFAULT_TIMES)
}

// 添加只运行一次的循环任务
func (w *Wheel) AddOnce(interval time.Duration, job JobFunc) (*Entry, error) {
    return w.newEntry(interval, job, false, 1)
}

// 添加运行指定次数的循环任务
func (w *Wheel) AddTimes(interval time.Duration, times int, job JobFunc) (*Entry, error) {
    return w.newEntry(interval, job, false, times)
}

// 延迟添加循环任务，delay参数单位为时间轮刻度
func (w *Wheel) DelayAdd(delay time.Duration, interval time.Duration, job JobFunc) {
    w.AddOnce(delay, func() {
        w.Add(interval, job)
    })
}

// 延迟添加单例循环任务，delay参数单位为时间轮刻度
func (w *Wheel) DelayAddSingleton(delay time.Duration, interval time.Duration, job JobFunc) {
    w.AddOnce(delay, func() {
        w.AddSingleton(interval, job)
    })
}

// 延迟添加只运行一次的循环任务，delay参数单位为时间轮刻度
func (w *Wheel) DelayAddOnce(delay time.Duration, interval time.Duration, job JobFunc) {
    w.AddOnce(delay, func() {
        w.AddOnce(interval, job)
    })
}

// 延迟添加只运行一次的循环任务，delay参数单位为时间轮刻度
func (w *Wheel) DelayAddTimes(delay time.Duration, interval time.Duration, times int, job JobFunc) {
    w.AddOnce(delay, func() {
        w.AddTimes(interval, times, job)
    })
}

// 任务数量
func (w *Wheel) Size() (size int) {
    for _, l := range w.slots {
        size += l.Len()
    }
    return
}

// 关闭循环任务
func (w *Wheel) Close() {
    w.ticker.Stop()
    w.closed <- struct{}{}
}
