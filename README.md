# adhocore/chin

[![Go Report](https://goreportcard.com/badge/github.com/adhocore/chin)](https://goreportcard.com/report/github.com/adhocore/chin)
[![Lint](https://github.com/adhocore/chin/actions/workflows/lint-action.yml/badge.svg)](https://github.com/adhocore/chin/actions/workflows/lint-action.yml)
[![Donate 15](https://img.shields.io/badge/donate-paypal-blue.svg?style=flat-square&label=donate+15)](https://www.paypal.me/ji10/15usd)
[![Donate 25](https://img.shields.io/badge/donate-paypal-blue.svg?style=flat-square&label=donate+25)](https://www.paypal.me/ji10/25usd)
[![Donate 50](https://img.shields.io/badge/donate-paypal-blue.svg?style=flat-square&label=donate+50)](https://www.paypal.me/ji10/50usd)
[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Golang+tool+to+check+internet+speed+right+from+the+terminal&url=https://github.com/adhocore/fast&hashtags=golang,terminal,spinner,spin)


A GO lang command line tool to show a spinner as you wait for some long running jobs to finish.

> This is a simple project but carries a tremendous value to me [❤️].

## Usage

Install `chin`:
```sh
go get -u github.com/adhocore/chin
```

Use in Go code with `WaitGroup`:

```go
var wg sync.WaitGroup

s := chin.New().WithWait(&wg)
go s.Start()

// invoke some long running task
// (you can also call s.Stop() from that task)
longTask(&wg)

s.Stop()
wg.Wait()
```

> Refer [example](./examples/main.go) for more (there is also one without waitgroup).

> To run the examples: `go run examples/main.go`

## Screen

![CHIN](./assets/chin.gif)

---
### Other projects
My other golang projects you might find interesting and useful:

- [**gronx**](https://github.com/adhocore/gronx) - Lightweight, fast and dependency-free Cron expression parser (due checker), task scheduler and/or daemon for Golang (tested on v1.13 and above) and standalone usage.
- [**urlsh**](https://github.com/adhocore/urlsh) - URL shortener and bookmarker service with UI, API, Cache, Hits Counter and forwarder using postgres and redis in backend, bulma in frontend; has [web](https://urlssh.xyz) and cli client.
- [**fast**](https://github.com/adhocore/fast) - Check your internet speed with ease and comfort right from the terminal. (It uses `adhocore/chin`.)
- [**goic**](https://github.com/adhocore/goic) - Go Open ID Connect, is OpenID connect client library for Golang, supports the Authorization Code Flow of OpenID Connect specification.
