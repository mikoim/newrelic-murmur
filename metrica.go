package main

type MetricaConnectedUsers struct {
	Client *MumbleClient
}

func NewMetricaConnectedUsers(client *MumbleClient) *MetricaConnectedUsers {
	return &MetricaConnectedUsers{
		Client: client,
	}
}

func (m *MetricaConnectedUsers) GetName() string {
	return "Connected Users"
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
	return "Maximum Users"
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