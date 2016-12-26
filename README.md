# http_to_nsq
Publish http request to nsqd with custom topic. (for record logs)

## Installation

```
$ go get github.com/duobeiyun/http_to_nsqd
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

```
Copyright 2016 Duobei. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
```
