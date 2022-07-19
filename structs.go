package main

type GpsData struct {
	*Pvt
	*Svin
}

type Pvt struct {
	fixType       string
	long          int32
	lat           int32
	height        int32
	msl           int32
	gnssFixOk     bool
	confirmedDate bool
	confirmedTime bool
	confirmedAvai bool
}

type Svin struct {
	accuracy   string
	dataPoints uint32
	valid      byte
}

type Config struct {
	Server  Server
	Control Control
	Options Options
}

type Server struct {
	TtyPorts   string   `json:"ttyPort"`
	Nav        []string `json:"nav"`
	Rxm        []string `json:"rxm"`
	Svin       []uint32 `json:"svin"`
	Rtcm       []string `json:"rtcm"`
	Nmea       []string `json:"nmea"`
	Baud       uint32   `json:"baud"`
	Mountpoint string   `json:"mountpoint"`
	Caster     string   `json:"caster"`
	CasterPort string   `json:"casterPort"`
	Str2Str    string   `json:"str2str"`
}

type Control struct {
	TtyPorts string   `json:"ttyPort"`
	Nav      []string `json:"nav"`
	Rxm      []string `json:"rxm"`
	Svin     []uint32 `json:"svin"`
	Rtcm     []string `json:"rtcm"`
	Nmea     []string `json:"nmea"`
	Baud     uint32   `json:"baud"`
}

type Options struct {
	Nav  []string `json:"nav_options"`
	Rxm  []string `json:"rxm_options"`
	Rtcm []string `json:"rtcm_options"`
	Nmea []string `json:"nmea_options"`
}

type startServerChannel struct {
	fix  bool
	stop bool
}
