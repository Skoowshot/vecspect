package telegram

import (
	"net/url"
	"strconv"
)

func (b *TelegramBot) GetUpdatesEndpoint(offset int, timeout int) string {
	u, _ := url.Parse("https://api.telegram.org/bot" + b.token + "/getUpdates")
	q := u.Query()
	
	q.Set("offset", strconv.Itoa(offset))
	q.Set("timeout", strconv.Itoa(timeout))

	u.RawQuery = q.Encode()
	return u.String()
}
