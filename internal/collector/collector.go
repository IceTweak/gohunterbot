package collector

import (
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
)

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

func GetDefaultProxyList() []string {
	user := os.Getenv("PROXY_USER")
	passwd := os.Getenv("PROXY_PASSWD")
	return []string{
		fmt.Sprintf("socks5://%s:%s@38.154.227.167:5868", user, passwd),
		fmt.Sprintf("socks5://%s:%s@185.199.229.156:7492", user, passwd),
		fmt.Sprintf("socks5://%s:%s@185.199.228.220:7300", user, passwd),
		fmt.Sprintf("socks5://%s:%s@185.199.231.45:8382", user, passwd),
		fmt.Sprintf("socks5://%s:%s@188.74.210.207:6286", user, passwd),
		fmt.Sprintf("socks5://%s:%s@188.74.183.10:8279", user, passwd),
		fmt.Sprintf("socks5://%s:%s@188.74.210.21:6100", user, passwd),
		fmt.Sprintf("socks5://%s:%s@45.155.68.129:8133", user, passwd),
		fmt.Sprintf("socks5://%s:%s@154.95.36.199:6893", user, passwd),
		fmt.Sprintf("socks5://%s:%s@45.94.47.66:8110", user, passwd),
	}
}
