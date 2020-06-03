# events

This is simply an event subscription and does not support multiple listeners subscribing to one event

# Installation

```shell
go get -u github.com/whiteCcinn/events
```

# Feature

- Thread safety
- Golang Version > 1.9 （Because sync.map is used to ensure thread safety）