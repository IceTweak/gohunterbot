# GoHunterBot

HeadHunter bot that raises resumes in searches after a certain time interval.

## Description

Bot have ordinary construction with 3 components in it:

```go
type Bot struct {
	ResourceUrl   string
	BorwserCaps   selenium.Capabilities
	Refreshing    time.Duration
}
```
- **ResourceUrl** - link to page that bot will visiting;
- **BorwserCaps** - it's external library struct from [tebeka/selenium](https://pkg.go.dev/github.com/tebeka/selenium). Capabilities configures both the WebDriver process and the target browsers, with standard and browser-specific options;
- **Refreshing** - repeating time interval after which the bot will perform its target actions using the **Collector**.

## How to use

Below are described ways to launch this bot locally with the necessary configuration.

1. Clone this repo by typing:

	```bash
	git clone https://github.com/IceTweak/gohunterbot.git
	cd gohunterbot
	```

2. Build Docker Image using Dockerfile

	```bash
	docker build --tag hh-bot .
	```
3. Check builded image via:

	```bash
	docker images
	```
	should print out same thing:

	```bash
	REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
	hh-bot       latest    30cab23d6fa8   2 minutes ago   990MB
	``` 
4. Run image using docker compose:
	
	`docker compose up` (run command in the same folder as compose.yaml file).