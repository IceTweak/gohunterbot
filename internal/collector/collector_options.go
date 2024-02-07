package collector

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
)

func WithProxyFunc(proxies []string) CollectorOption {
	return func(c *colly.Collector) error {
		// Rotate proxies
		rp, err := proxy.RoundRobinProxySwitcher(proxies...)
		if err != nil {
			return fmt.Errorf("collector: %w", err)
		}
		// Setting proxy function for collector
		c.SetProxyFunc(rp)
		return nil
	}
}
