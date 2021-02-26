package helper

import (
	_ "errors"
	"fmt"
	"runtime"
	"time"
)

type Callback func(data interface{})

type ChanQueue struct {
	len   uint32
	queue chan interface{}
	//consumer_count uint32
}

func (l ChanQueue) Create(len uint32) ChanQueue {
	qqq := ChanQueue{}
	qqq.len = len
	qqq.queue = make(chan interface{}, len)
	return qqq
}

func (l *ChanQueue) Push(data interface{}, outtime time.Duration) (err error) {

	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		if errr := recover(); errr != nil {

			switch errr.(type) {
			case runtime.Error: // 运行时错误
				err = errr.(runtime.Error)
				fmt.Println("runtime error:", errr)
			default: // 非运行时错误

			}

			fmt.Println(err)
		}
	}()

	if outtime < 0 {
		select {
		case l.queue <- data:
			err = nil
		default:
			panic("push 失败")
		}
	} else {
		select {
		case l.queue <- data:
			err = nil
		case <-time.After(outtime * time.Millisecond):
			panic("push 超时")
		}
	}
	return
}

func (l *ChanQueue) consumer(callback func(data interface{})) {
	for x := range l.queue {

		callback(x)
	}
}

func (l *ChanQueue) Start(callback func(data interface{}), consumer_count uint32) {
	go func() {
		var i uint32 = 0
		for true {
			go func() {
				defer func() {
					if errr := recover(); errr != nil {
						fmt.Println(errr)
					}
				}()
				l.consumer(callback)
			}()
			i++
			if i > consumer_count {
				return
			}

		}
	}()

}
