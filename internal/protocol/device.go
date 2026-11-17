package protocol

import (
	"fmt"

	"github.com/beilunyang/openhongguo/sdk/model"
)

// UserAgent is the app's Cronet User-Agent for d, in the captured client's
// format (2026-10-24), which names the model code rather than a market name.
func UserAgent(d model.Device) string {
	return UserAgentFor(d, VersionCode)
}

// UserAgentFor is UserAgent with an explicit version_code, so a client that
// announces a newer build also announces it here. An empty versionCode keeps
// the compiled-in one.
func UserAgentFor(d model.Device, versionCode string) string {
	d = d.OrDefault()
	if versionCode == "" {
		versionCode = VersionCode
	}
	return fmt.Sprintf("%s/%s (Linux; U; Android %s; zh_CN; %s; Build/%s; Cronet/TTNetVersion:8d40f833 2026-03-03 QuicVersion:21ac1950 2025-11-18)",
		Package, versionCode, d.OSVersion, d.Model, d.Build)
}
