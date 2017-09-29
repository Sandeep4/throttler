package throttler

import (
	"time"
)

type bucket struct {
	tokens    int64
	timestamp int64
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
	currentTimestamp := int64(time.Now().UnixNano())
	keyBucket, ok := bt.bucketStore[key]
	tokens := int64(0)
	if ok {
		tokens = keyBucket.tokens - (bt.rate.Quantity * (currentTimestamp - keyBucket.timestamp) / (bt.rate.Seconds * int64(time.Second)))
		if tokens < 0 {
			tokens = 0
		}
	}
	if tokens >= bt.rate.Quantity {
		return true
	}
	tokens++
	bt.bucketStore[key] = bucket{tokens, currentTimestamp}
	return false
}

func (bt *bucketThrottler) ResetKey(key string, value int64) error {
	return nil
}
