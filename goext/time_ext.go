/**
 * Auth :   liubo
 * Date :   2018/10/31 9:39
 * Comment: 受系统时钟影响的timer
 */

package goext

import (
	"time"
)

// 误差1秒
func Sleep(d time.Duration) {
	// 超过1分钟的，采用时钟比较
	feature := time.Now().Add(d).UnixNano()
	for true {
		now := time.Now().UnixNano()
		if now - (feature) >= 0 {
			break
		}

		// 分梯度，为了性能
		if d < time.Second {
			time.Sleep(time.Millisecond)
		} else if d < time.Minute {
			time.Sleep(time.Millisecond * 30)
		} else {
			time.Sleep(time.Second)
		}
	}
}

// timer
type Timer struct {
	C chan time.Time
	stop bool
}
func NewTimer(d time.Duration) *Timer {
	c := make(chan time.Time, 1)
	t := &Timer{ C: c }
	go func(tt *Timer) {
		Sleep(d)
		if !tt.stop {
			tt.C <- time.Now()
		}
	}(t)
	return t
}
func (self *Timer) Stop() {
	self.stop = true
}


