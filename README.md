# GoHunterBot

HeadHunter bot that raises resumes in searches after a certain time interval.

## Description

Bot have ordinary construction with 3 components in it:

```go
type Bot struct {
	ResourceUrl string
	Collector   *colly.Collector
	Refreshing  time.Duration
}
```
- **ResourceUrl** - link to page that bot will visiting;
- **Collector** - it's external library struct from [gocolly repo](https://github.com/gocolly/colly). Used for scraping web page of resource and made some actions in it;
- **Refreshing** - repeating time interval after which the bot will perform its target actions using the **Collector**.

## How to use

Below are described ways to launch this bot locally with the necessary configuration.

### Docker Compose

In development...