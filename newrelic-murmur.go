package main

import (
	"flag"
	"log"
	"github.com/yvasiyarov/newrelic_platform_go"
	mumble "github.com/layeh/gumble/gumble"
	"fmt"
	"time"
)

const (
	AGENT_NAME = "Murmur"
	AGENT_GUID = "com.github.mikoim.newrelic.murmur"
	AGENT_VERSION = "0.0.1"
)

type NewRelicMurmur struct {
	Host    string
	Port    int
	Timeout time.Duration
}

func (m *NewRelicMurmur) GetName() string {
	return "Current users"
}

func (m *NewRelicMurmur) GetUnits() string {
	return "users"
}

func (m *NewRelicMurmur) GetValue() (float64, error) {
	resp, err := mumble.Ping(fmt.Sprintf("%s:%d", m.Host, m.Port), m.Timeout)
	if err != nil {
		return 0, err
	}

	return float64(resp.ConnectedUsers), err
}

func NewNewRelicMurmur(host string, port int, timeout int) *NewRelicMurmur {
	return &NewRelicMurmur{
		Host: host,
		Port: port,
		Timeout: time.Millisecond * time.Duration(timeout),
	}
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

	component.AddMetrica(NewNewRelicMurmur(*host, *port, *timeout))

	plugin.Verbose = *verbose
	plugin.Run()
}