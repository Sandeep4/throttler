package throttler

// Throttler ...
type Throttler interface {
	ThrottleKey(key string) bool
	ResetKey(key string, value int64) error
}
