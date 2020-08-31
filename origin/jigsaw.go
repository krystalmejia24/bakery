package origin

import (
	"context"
	"fmt"
	"strings"

	"github.com/cbsinteractive/bakery/config"
	"github.com/cbsinteractive/bakery/logging"
)

const (
	aliasKey     = "alias"
	protocolKey  = "protocol"
	playtokenKey = "playtoken"
)

//Jigsaw holds the URL and playtoken for a jigsaw origin
type Jigsaw struct {
	URL       string
	playtoken string
}

func NewJigsaw(ctx context.Context, c config.Config, path string) (*Jigsaw, error) {
	logging.UpdateCtx(ctx, logging.Params{"origin": "jigsaw"})

	// path = /jigsaw/alias/protocol/playtoken/master.m3u8
	parts := strings.Split(path, "/")
	if len(parts) <= 5 {
		return &Jigsaw{}, fmt.Errorf("Jigsaw origin: invalid url format:%v", path)
	}

	alias := parts[2]
	protocol := parts[3]
	playtoken := parts[4]

	host, found := c.Jigsaw.Alias[alias]
	if !found {
		return &Jigsaw{}, fmt.Errorf("Jigsaw origin: invalid host `%v` requested", alias)
	}

	return &Jigsaw{
		URL:       fmt.Sprintf("https://%v/%v/%v/master.m3u8", host, protocol, playtoken),
		playtoken: playtoken,
	}, nil
}

func (j *Jigsaw) GetPlaybackURL() string {
	return j.URL
}
func (j *Jigsaw) FetchManifest(ctx context.Context, c config.Client) (ManifestInfo, error) {
	return fetch(ctx, c, j.GetPlaybackURL())
}
