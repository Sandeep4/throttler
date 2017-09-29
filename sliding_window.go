package throttler

import (
	"strconv"
	"time"
)

type windowThrottler struct {
	callCountStore map[string]int64
	rate           Rate
}

func NewWindowThrottler(rate Rate) Throttler {
	return &windowThrottler{
		callCountStore: make(map[string]int64),
		rate:           rate,
	}
}

func (wt *windowThrottler) ThrottleKey(key string) bool {
	currentTime := time.Now().Unix()
	var i, totalCalls, callCount int64
	totalCalls = 0
	for i = 0; i < wt.rate.Seconds; i++ {
		windowKey := key + "_" + strconv.FormatInt(currentTime-i, 10)
		callCount, _ = wt.callCountStore[windowKey]
		totalCalls = totalCalls + callCount
	}
	if totalCalls >= wt.rate.Quantity {
		return true
	} else {
		windowKey := key + "_" + strconv.FormatInt(currentTime, 10)
		callCount, _ = wt.callCountStore[windowKey]
		wt.callCountStore[windowKey] = callCount + 1
	}
	return false
}

func (wt *windowThrottler) ResetKey(key string, value int64) error {
	return nil
}
