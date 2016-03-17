package main

import (
	"flag"
	"log"
	"github.com/yvasiyarov/newrelic_platform_go"
)

const (
	AGENT_NAME = "Murmur"
	AGENT_GUID = "com.github.mikoim.newrelic.murmur"
	AGENT_VERSION = "0.0.2"
)

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