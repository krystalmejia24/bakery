package parsers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cbsinteractive/bakery/config"
)

type DeviceConfig struct {
	Description   string        `json:"description,omitempty"`
	Condition     Condition     `json:"condition,omitempty"`
	Configuration Configuration `json:"configuration,omitempty"`
}

type Condition struct {
	Model        string `json:"model,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
}

type Configuration struct {
	AdaptiveVR360             bool `json:"adaptiveVR360,omitempty"`
	MaxBitrate                int  `json:"maxBitRate,omitempty"`
	MinBuffer                 int  `json:"minBuffer,omitempty"`
	MaxBuffer                 int  `json:"maxBuffer,omitempty"`
	MaxClientSideAdBitrate    int  `json:"maxClientSideAdBitrate,omitempty"`
	EnableSoftwareDecoder     bool `json:"enableSoftwareDecoder,omitempty"`
	DisablePremiumAudio       bool `json:"disablePremiumAudio,omitempty"`
	DisablePremiumAudioForAds bool `json:"disablePremiumAudioForAds,omitempty"`
	EnableSmallSeekThumbnail  bool `json:"enableSmallSeekThumbnail,omitempty"`
}

// todo not all of these above apply to manifest filters, this is a lift and shift of the sample
// json provided by the player team for the sake of unmarshaling the data

func (mf *MediaFilters) getRemoteFilters(ctx context.Context, c config.Config, capabilites DeviceCapabilites) error {
	deviceConfigs := make([]DeviceConfig, 0)
	if err := fetchRemoteDeviceConfig(ctx, c, &deviceConfigs); err != nil {
		return err
	}

	// below is a lazy approach to applying filters for POC purposes. Ideally we would iterate
	// through the remote config and create a map[condition]Configuartion to easily apply
	// filters based on multiple conditions

	// using the device capabilites passed in the request, you iterate through
	// the device config to find a conidtion (model, manufacturer, etc.) match
	// for POC we will match the following "model": "(?i)AFTMM" and applying bitrate filter
	for _, dc := range deviceConfigs {
		if strings.Contains(dc.Condition.Model, capabilites.Model) {
			//todo check units for bitrate
			mf.Bitrate = &Bitrate{Min: 0, Max: dc.Configuration.MaxBitrate}
		}

		if strings.Contains(dc.Condition.Manufacturer, capabilites.Manufacturer) {
			//add filters
		}
	}

	return nil
}

func fetchRemoteDeviceConfig(ctx context.Context, c config.Config, dcs *[]DeviceConfig) error {
	req, err := http.NewRequest(http.MethodGet, c.RemoteDeviceConfig, nil)
	if err != nil {
		return fmt.Errorf("generating request to fetch remote device config: %w", err)
	}

	client := c.Client
	ctx, cancel := context.WithTimeout(ctx, client.Timeout)
	defer cancel()

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("fetching remote device config: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading device config json: %w", err)
	}

	if resp.StatusCode/100 > 3 {
		return fmt.Errorf("remote device config returning http status code %d", resp.StatusCode)
	}

	if err = json.Unmarshal(body, dcs); err != nil {
		return fmt.Errorf("unmarshaling device config: %w", err)
	}

	return nil
}
