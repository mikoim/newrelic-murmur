package main

import (
	"time"
	"github.com/layeh/gumble/gumble"
	"fmt"
)

type MumbleClient struct {
	Host          string
	Port          int
	Timeout       time.Duration
	Cache         *gumble.PingResponse
	CacheError    error
	CacheModified time.Time
	CacheDuration time.Duration
}

func NewMumbleClient(host string, port int, timeout int, cacheDuration int) *MumbleClient {
	return &MumbleClient{
		Host: host,
		Port: port,
		Timeout: time.Millisecond * time.Duration(timeout),
		CacheDuration: time.Second * time.Duration(cacheDuration),
	}
}

func (m *MumbleClient) GetPingResponse() (*gumble.PingResponse, error) {
	if time.Now().Sub(m.CacheModified) > m.CacheDuration {
		m.Cache, m.CacheError = gumble.Ping(fmt.Sprintf("%s:%d", m.Host, m.Port), m.Timeout)
		m.CacheModified = time.Now()
	}

	return m.Cache, m.CacheError
}
