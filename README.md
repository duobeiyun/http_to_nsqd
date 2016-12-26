# http_to_nsq
Publish http request to nsqd with custom topic. (for record logs)

## Installation

```
$ go get github.com/0neSe7en/http_to_nsqd
```

## About

An util for collect logs via HTTP. Send logs to different urls, and `http_to_nsq` will publish logs to different topics.

## Usage

- `host`(optional) - HTTP Address to listen (default: 127.0.0.1:3000)
- `nsqd` - NSQD Address for publish (default: 127.0.0.1:4150)
- `urltopic` - URL and Topic mapping.
- `param` - Param in GET or POST (default: result)

## Example

Start `http_to_nsq` with following command.

```
$ http_to_nsqd --urltopic /log:log
$ curl "http://127.0.0.1:3000/log?result=%7B%22message%22:%20%22it%20is%20a%20log%20for%20test%22%7D"
// {"message":{"message": "it is a log for test"},"timestamp":"2016-12-26T06:13:21.043Z","ip":"127.0.0.1"}
// This log is published to the "log" topic.
```

# License

MIT