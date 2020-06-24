package redis

import "time"

type Client struct {
}

func (c *Client) SetNX(k, v string, dur time.Duration) (bool, error) {
	return true, nil
}
