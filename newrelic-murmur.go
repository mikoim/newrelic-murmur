package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"github.com/yvasiyarov/newrelic_platform_go"
	"github.com/layeh/gumble/gumble"
)

const (
	AGENT_NAME = "Murmur"
	AGENT_GUID = "com.github.mikoim.newrelic.murmur"
	AGENT_VERSION = "0.0.2"
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

type MetricaConnectedUsers struct {
	Client *MumbleClient
}

func NewMetricaConnectedUsers(client *MumbleClient) *MetricaConnectedUsers {
	return &MetricaConnectedUsers{
		Client: client,
	}
}

func (m *MetricaConnectedUsers) GetName() string {
	return "Connected users"
}

func (m *MetricaConnectedUsers) GetUnits() string {
	return "users"
}

func (m *MetricaConnectedUsers) GetValue() (float64, error) {
	resp, err := m.Client.GetPingResponse()
	if err != nil {
		return 0, err
	}

	return float64(resp.ConnectedUsers), err
}

type MetricaMaximumBitrate struct {
	Client *MumbleClient
}

func NewMetricaMaximumBitrate(client *MumbleClient) *MetricaMaximumBitrate {
	return &MetricaMaximumBitrate{
		Client: client,
	}
}

func (m *MetricaMaximumBitrate) GetName() string {
	return "Maximum Bitrate"
}

func (m *MetricaMaximumBitrate) GetUnits() string {
	return "bps"
}

func (m *MetricaMaximumBitrate) GetValue() (float64, error) {
	resp, err := m.Client.GetPingResponse()
	if err != nil {
		return 0, err
	}

	return float64(resp.MaximumBitrate), err
}

type MetricaMaximumUsers struct {
	Client *MumbleClient
}

func NewMetricaMaximumUsers(client *MumbleClient) *MetricaMaximumUsers {
	return &MetricaMaximumUsers{
		Client: client,
	}
}

func (m *MetricaMaximumUsers) GetName() string {
	return "Maximum users"
}

func (m *MetricaMaximumUsers) GetUnits() string {
	return "users"
}

func (m *MetricaMaximumUsers) GetValue() (float64, error) {
	resp, err := m.Client.GetPingResponse()
	if err != nil {
		return 0, err
	}

	return float64(resp.MaximumUsers), err
}

func main() {
	var (
		host = flag.String("host", "localhost", "Murmur host")
		port = flag.Int("port", 64738, "Murmur port")
		license = flag.String("license", "", "New Relic license key")
		interval = flag.Int("interval", 60, "Poll interval (seconds)")
		timeout = flag.Int("timeout", 1000, "Timeout (milliseconds)")
		verbose = flag.Bool("verbose", false, "Verbose")
	)

	flag.Parse()
	if *license == "" {
		log.Fatal("New Relic license key is required.")
	}

	plugin := newrelic_platform_go.NewNewrelicPlugin(AGENT_VERSION, *license, *interval)
	component := newrelic_platform_go.NewPluginComponent(AGENT_NAME, AGENT_GUID, *verbose)
	plugin.AddComponent(component)

	client := NewMumbleClient(*host, *port, *timeout, *interval)
	component.AddMetrica(NewMetricaConnectedUsers(client))
	component.AddMetrica(NewMetricaMaximumBitrate(client))
	component.AddMetrica(NewMetricaMaximumUsers(client))

	plugin.Verbose = *verbose
	plugin.Run()
}