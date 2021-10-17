package origin

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/cbsinteractive/bakery/config"
	"github.com/cbsinteractive/bakery/logging"
)

// RemoteDeviceConfig origin is for filters that are set server side
type RemoteDeviceConfig struct {
	url *url.URL
}

// configreDeviceConfig will decode the base64 encoded path for fetching the manifest
func configureRemoteDeviceConfig(ctx context.Context, c config.Config, encodedPath string) (*RemoteDeviceConfig, error) {
	logging.UpdateCtx(ctx, logging.Params{"origin": "remote-device-config"})

	// todo(km) trim the DeviceConfig string from path
	urlBytes, err := base64.RawURLEncoding.DecodeString(encodedPath)
	if err != nil {
		return &RemoteDeviceConfig{}, fmt.Errorf("configuring RemoteDeviceConfig origin: %w", err)
	}

	u, err := url.Parse(string(urlBytes))
	if err != nil {
		return &RemoteDeviceConfig{}, err
	}

	return &RemoteDeviceConfig{url: u}, nil
}

func (r *RemoteDeviceConfig) GetPlaybackURL() string {
	return r.url.String()
}

func (r *RemoteDeviceConfig) FetchOriginContent(ctx context.Context, c config.Client) (OriginContentInfo, error) {
	return fetch(ctx, c, r.GetPlaybackURL())
}
