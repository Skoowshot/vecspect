package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Skoowshot/vecspect/domain"
)

func (b *TelegramBot) GetUpdates(offset int) (*domain.Updates, error) {
	ep := b.GetUpdatesEndpoint(offset, b.pollTimeout)
	res, err := http.Get(ep)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	var updates domain.Updates
	if err := decodeJson(res.Body, &updates); err != nil {
		return nil, err
	}

	return &updates, nil
}

func decodeJson(r io.Reader, result interface{}) error {
	return json.NewDecoder(r).Decode(result)
}
