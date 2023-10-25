package boomer

import (
	"sync"
	"time"
)

type UserStartFunc func(*User) error
type UserStopFunc func(*User)
type WaitTimeFunc func() time.Duration

type User struct {
	tasks []*Task

	// This mutex protect keys map
	mu   sync.RWMutex
	keys map[string]interface{}

	startFunc    UserStartFunc
	stopFunc     UserStopFunc
	waitTimeFunc WaitTimeFunc
}

type UserConfig struct {
	Tasks     []*Task
	StartFunc UserStartFunc
	StopFunc  UserStopFunc
	WaitTime  WaitTimeFunc
}

func NewUser(config *UserConfig) *User {
	return &User{
		tasks:        config.Tasks,
		keys:         make(map[string]interface{}),
		startFunc:    config.StartFunc,
		stopFunc:     config.StopFunc,
		waitTimeFunc: config.WaitTime,
	}
}

func (c *User) Set(key string, value interface{}) {
	c.mu.Lock()
	if c.keys == nil {
		c.keys = make(map[string]interface{})
	}

	c.keys[key] = value
	c.mu.Unlock()
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exists it returns (nil, false)
func (c *User) Get(key string) (value interface{}, exists bool) {
	c.mu.RLock()
	value, exists = c.keys[key]
	c.mu.RUnlock()
	return
}
