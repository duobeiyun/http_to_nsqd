package main

import "time"

var Gif = []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x1, 0x0, 0x1, 0x0, 0x80, 0x0, 0x0, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x2c, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x1, 0x0, 0x0, 0x2, 0x2, 0x44, 0x1, 0x0, 0x3b}

func Process(l string, ip string) string {
	return "{\"message\":" + l + ",\"timestamp\":\"" + time.Now().UTC().Format("2006-01-02T15:04:05.000Z") + "\",\"ip\":\"" + ip + "\"}"
}
