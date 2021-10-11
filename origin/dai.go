package origin

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/cbsinteractive/bakery/config"
	"github.com/cbsinteractive/bakery/logging"
)

type DeviceConfig struct {
	url *url.URL
	// todo(km) extend with device specific args needed for making the request to DeviceConfig
}

// configreDeviceConfig will decode the base64 encoded path and set the device capabilities
// as configured in the request via query params
func configureDeviceConfig(ctx context.Context, encodedPath string) (*DeviceConfig, error) {
	// todo(km) trim QUERY ARGS that define device capabilities and extend the DeviceConfig struct
	// to maintain state of the origin request
	logging.UpdateCtx(ctx, logging.Params{"origin": "device-config"})

	// todo(km) trim the DeviceConfig string from path
	urlBytes, err := base64.RawURLEncoding.DecodeString(encodedPath)
	if err != nil {
		return &DeviceConfig{}, fmt.Errorf("configuring DeviceConfig origin: %w", err)
	}

	u, err := url.Parse(string(urlBytes))
	if err != nil {
		return &DeviceConfig{}, err
	}

	return &DeviceConfig{url: u}, nil
}

func (d *DeviceConfig) GetPlaybackURL() string {
	return d.url.String()
}

func (d *DeviceConfig) FetchOriginContent(ctx context.Context, c config.Client) (OriginContentInfo, error) {
	return fetch(ctx, c, d.GetPlaybackURL())
}
