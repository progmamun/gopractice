package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

func Checksum(b []byte) []byte {
	var sum uint64
	for len(b) >= 5 {
		for i := range b[:5] {
			v := uint64(b[i])
			for j := 0; j < i; j++ {
				v = v * 256
			}
			sum += v
		}
		b = b[5:]
	}
	for _, v := range b {
		sum += uint64(v)
	}
	s := make([]byte, 8)
	binary.LittleEndian.PutUint64(s, sum)
	return s[:4]
}

var ErrLength = errors.New("message too long")

func CreateMessage(content []byte) ([]byte, error) {
	if len(content) > 65535 {
		return nil, ErrLength
	}
	data := make([]byte, 0, len(content)+14)
	data = append(data, Sequence...)
	data = append(data, byte(len(content)/256), byte(len(content)%256))
	data = append(data, Checksum(content)...)
	data = append(data, content...)
	data = append(data, Sequence...)
	return data, nil
}

func MessageContent(b []byte) ([]byte, error) {
	n := len(b)
	if n < 14 {
		return nil, fmt.Errorf("Too short")
	}
	if open := b[:4]; !bytes.Equal(open, Sequence) {
		return nil, fmt.Errorf("Wrong opening sequence %x", open)
	}
	if length := int(b[4])*256 + int(b[5]); n-14 != length {
		return nil, fmt.Errorf("Wrong length: %d (expected %d)", length, n-14)
	}
	if close := b[n-4 : n]; !bytes.Equal(close, Sequence) {
		return nil, fmt.Errorf("Wrong closing sequence %x", close)
	}
	content := b[10 : n-4]
	if !bytes.Equal(Checksum(content), b[6:10]) {
		return nil, fmt.Errorf("Wrong checksum")
	}
	return content, nil
}

// Send
data, err := common.CreateMessage(msg)
if err != nil {
log.Println("->", addr, "Encode error:", err)
continue
}
if _, err := conn.WriteTo(data, addr); err != nil {
log.Println("->", addr, "Send error:", err)
}

//Receive
n, addr, err := conn.ReadFromUDP(b)
if err != nil {
log.Println("<-", addr, "Message error:", err)
continue
}
msg, err := common.MessageContent(b[:n])
if err != nil {
log.Println("<-", addr, "Decode error:", err)
continue
}
log.Printf("<- %q from %s", msg, addr)