package roku

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type DeviceInfo struct {
	XMLName                     xml.Name `xml:"device-info"`
	UDN                         string   `xml:"udn"`
	SerialNumber                string   `xml:"serial-number"`
	DeviceID                    string   `xml:"device-id"`
	VendorName                  string   `xml:"vendor-name"`
	ModelNumber                 string   `xml:"model-number"`
	ModelName                   string   `xml:"model-name"`
	ModelRegion                 string   `xml:"model-region"`
	IsTV                        bool     `xml:"is-tv"`
	IsStick                     bool     `xml:"is-stick"`
	ScreenSize                  int      `xml:"screen-size"`
	PanelID                     string   `xml:"panel-id"`
	TunerType                   string   `xml:"tuner-type"`
	SupportsEthernet            string   `xml:"supports-ethernet"`
	WiFiMac                     string   `xml:"wifi-mac"`
	WiFiDriver                  string   `xml:"wifi-driver"`
	EthernetMac                 string   `xml:"ethernet-mac"`
	NetworkType                 string   `xml:"network-type"`
	NetworkName                 string   `xml:"network-name"`
	FriendlyDeviceName          string   `xml:"friendly-device-name"`
	FriendlyModelName           string   `xml:"friendly-model-name"`
	DefaultDeviceName           string   `xml:"default-device-name"`
	UserDeviceName              string   `xml:"user-device-name"`
	SoftwareVersion             string   `xml:"software-version"`
	SoftwareBuild               string   `xml:"software-build"`
	SecureDevice                bool     `xml:"secure-device"`
	Language                    string   `xml:"language"`
	Country                     string   `xml:"country"`
	Locale                      string   `xml:"locale"`
	TimeZoneAuto                bool     `xml:"time-zone-auto"`
	TimeZone                    string   `xml:"time-zone"`
	TimeZoneName                string   `xml:"time-zone-name"`
	TimeZoneTZ                  string   `xml:"time-zone-tz"`
	TimeZoneOffset              int      `xml:"time-zone-offset"`
	ClockFormat                 string   `xml:"clock-format"`
	Uptime                      int64    `xml:"uptime"`
	PowerMode                   string   `xml:"power-mode"`
	SupportsSuspend             bool     `xml:"supports-suspend"`
	SupportsFindRemote          bool     `xml:"supports-find-remote"`
	SupportsAudioGuide          bool     `xml:"supports-audio-guide"`
	DeveloperEnabled            bool     `xml:"developer-enabled"`
	KeyedDeveloperID            string   `xml:"keyed-developer-id"`
	SearchEnabled               bool     `xml:"search-enabled"`
	SearchChannelsEnabled       bool     `xml:"search-channels-enabled"`
	VoiceSearchEnabled          bool     `xml:"voice-search-enabled"`
	NotificationsEnabled        bool     `xml:"notifications-enabled"`
	NotificationsFirstUse       bool     `xml:"notifications-first-use"`
	SupportsPrivateListening    bool     `xml:"supports-private-listening"`
	SupportsPrivateListeningDTV bool     `xml:"supports-private-listening-dtv"`
	SupportsWarmStandby         bool     `xml:"supports-warm-standby"`
	HeadphonesConnected         bool     `xml:"headphones-connected"`
	ExpertPQEnabled             float64  `xml:"expert-pq-enabled"`
	SupportsECSTextEdit         bool     `xml:"supports-ecs-textedit"`
	SupportsECSMicrophone       bool     `xml:"supports-ecs-microphone"`
	SupportsWakeupOnWLAN        bool     `xml:"supports-wakeup-on-wlan"`
	HasPlayOnRoku               bool     `xml:"has-play-on-roku"`
	HasMobileScreenSaver        bool     `xml:"has-mobile-screensaver"`
	SupportURL                  string   `xml:"support-url"`
}

// Get the Roku Device info
func (r Roku) GetDeviceInfo() (DeviceInfo, error) {
	res, err := http.Get(fmt.Sprintf("%squery/device-info", r.APIAddress))
	var deviceInfo DeviceInfo
	if err != nil {
		return deviceInfo, err
	}
	defer res.Body.Close()

	err = xml.NewDecoder(res.Body).Decode(&deviceInfo)
	return deviceInfo, err
}
