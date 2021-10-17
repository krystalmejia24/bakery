package origin

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/cbsinteractive/bakery/config"
)

//Origin interface is implemented by DefaultOrigin and Propeller struct
type Origin interface {
	GetPlaybackURL() string
	FetchOriginContent(ctx context.Context, c config.Client) (OriginContentInfo, error)
}

//DefaultOrigin struct holds Origin and Path of DefaultOrigin
//Variant level DefaultOrigins will be base64 encoded absolute Urls
type DefaultOrigin struct {
	Host string
	URL  url.URL
}

//OriginContentInfo holds http response info from manifest request
type OriginContentInfo struct {
	Payload      string
	LastModified time.Time
	Status       int
}

const (
	PropellerOrigin          = "propeller"
	RemoteDeviceConfigOrigin = "rdc"
)

//Configure will return proper Origin interface
func Configure(ctx context.Context, c config.Config, path string) (Origin, error) {
	if strings.Contains(path, PropellerOrigin) {
		return configurePropeller(ctx, c, path)
	}

	if strings.Contains(path, RemoteDeviceConfigOrigin) {
		return configureRemoteDeviceConfig(ctx, c, path)
	}

	// check if path is base64 encoded url
	if strings.Count(path, "/") == 1 && (strings.HasSuffix(path, ".m3u8") || strings.HasSuffix(path, ".vtt")) {
		decodedPath, err := trimAndDecodePath(strings.TrimPrefix(path, "/"))
		if err != nil {
			return &DefaultOrigin{}, fmt.Errorf("decoding base64 url %q: %w", path, err)
		}
		path = decodedPath
	}

	return NewDefaultOrigin(c.OriginHost, path)
}

//NewDefaultOrigin returns a new Origin struct
//host is not required if path is absolute
func NewDefaultOrigin(host string, p string) (*DefaultOrigin, error) {
	u, err := url.Parse(p)
	if err != nil {
		return &DefaultOrigin{}, err
	}

	if len(host) == 0 && !u.IsAbs() {
		return &DefaultOrigin{}, errors.New("an origin host must be provided for relative paths")
	}

	return &DefaultOrigin{
		Host: host,
		URL:  *u,
	}, nil
}

//GetPlaybackURL will retrieve url
func (d *DefaultOrigin) GetPlaybackURL() string {
	if d.URL.IsAbs() {
		return d.URL.String()
	}

	return d.Host + d.URL.String()
}

//FetchOriginContent will grab DefaultOrigin contents of configured origin
func (d *DefaultOrigin) FetchOriginContent(ctx context.Context, c config.Client) (OriginContentInfo, error) {
	return fetch(ctx, c, d.GetPlaybackURL())
}

func fetch(ctx context.Context, client config.Client, originURL string) (OriginContentInfo, error) {
	req, err := http.NewRequest(http.MethodGet, originURL, nil)
	if err != nil {
		return OriginContentInfo{}, fmt.Errorf("generating request to fetch origin: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, client.Timeout)
	defer cancel()

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return OriginContentInfo{}, fmt.Errorf("fetching origin: %w", err)
	}
	defer resp.Body.Close()

	origin, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return OriginContentInfo{}, fmt.Errorf("reading origin response body: %w", err)
	}

	var lastModified time.Time
	if header := resp.Header.Get("Last-Modified"); header != "" {
		lastModified, err = http.ParseTime(header)
		if err != nil {
			return OriginContentInfo{}, err
		}
	}

	return OriginContentInfo{
		Payload:      string(origin),
		LastModified: lastModified,
		Status:       resp.StatusCode,
	}, nil
}

func trimAndDecodePath(encodedPath string) (string, error) {
	encodedPath = strings.TrimSuffix(encodedPath, path.Ext(encodedPath))
	url, err := base64.RawURLEncoding.DecodeString(encodedPath)
	if err != nil {
		return "", err
	}

	return string(url), nil
}
