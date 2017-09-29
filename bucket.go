package throttler

import (
	"time"
)

type bucket struct {
	tokens    int
	timestamp int64
}

type Rate struct {
	quantity int
	seconds  int64
}

type bucketThrottler struct {
	bucketStore map[string]bucket
	rate        Rate
}

func NewBucketThrottler(rate Rate) Throttler {
	return &bucketThrottler{
		bucketStore: make(map[string]bucket),
		rate:        rate,
	}
}

func (bt *bucketThrottler) ThrottleKey(key string) bool {
	currentTimestamp := time.Now()
	keyBuket, ok := bt.bucketStore[key]
	tokens := 0
	if ok {
		tokens = keyBuket.tokens - (bt.rate.quantity * (currentTimestamp - keyBuket.timestamp) / (bt.rate.seconds * time.Second))
		if tokens < 0 {
			tokens = 0
		}
	}
	if tokens >= bt.rate.quantity {
		return true
	}
	tokens++
	bt.bucketStore[key] = bucket{tokens, currentTimestamp}
	return false
}

func (bt *bucketThrottler) ResetKey(key string, value int64) error {
	return nil
}
