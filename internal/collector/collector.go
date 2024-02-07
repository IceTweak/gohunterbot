package collector

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

var DefaultProxyList = []string{
	"socks5://kupfwikj:xm3642nd5169@38.154.227.167:5868",
	"socks5://kupfwikj:xm3642nd5169@185.199.229.156:7492",
	"socks5://kupfwikj:xm3642nd5169@185.199.228.220:7300",
	"socks5://kupfwikj:xm3642nd5169@185.199.231.45:8382",
	"socks5://kupfwikj:xm3642nd5169@188.74.210.207:6286",
	"socks5://kupfwikj:xm3642nd5169@188.74.183.10:8279",
	"socks5://kupfwikj:xm3642nd5169@188.74.210.21:6100",
	"socks5://kupfwikj:xm3642nd5169@45.155.68.129:8133",
	"socks5://kupfwikj:xm3642nd5169@154.95.36.199:6893",
	"socks5://kupfwikj:xm3642nd5169@45.94.47.66:8110",
}

type CollectorOption func(c *colly.Collector) error

func defaultCollector() *colly.Collector {
	c := colly.NewCollector(colly.CheckHead())
	c.SetRequestTimeout(120 * time.Second)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
		Delay:       2 * time.Second,
		RandomDelay: 5 * time.Second,
	})
	return c
}

func NewCollector(options ...CollectorOption) (*colly.Collector, error) {
	c := defaultCollector()
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, fmt.Errorf("collector: options: %w", err)
		}
	}
	return c, nil
}
