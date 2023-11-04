package tools

import (
	"fmt"
	"net"
	"time"
)

func Ping(host string, count int, timeout time.Duration) {
	addr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	conn, err := net.DialIP("ip4:icmp", nil, addr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	for i := 1; i <= count; i++ {
		msg := make([]byte, 64)
		msg[0] = 8
		msg[1] = 0
		msg[2] = 0
		msg[3] = 0
		msg[4] = 0
		msg[5] = 0
		msg[6] = 0
		msg[7] = byte(i)
		checksum := checkSum(msg)
		msg[2] = byte(checksum >> 8)
		msg[3] = byte(checksum & 0xFF)

		startTime := time.Now()

		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		reply := make([]byte, 64)
		conn.SetReadDeadline(time.Now().Add(timeout))
		_, err = conn.Read(reply)
		if err != nil {
			fmt.Println("Request timed out")
		} else {
			duration := time.Since(startTime)
			fmt.Printf("Reply from %s: time=%v\n", addr.String(), duration)
		}

		time.Sleep(1 * time.Second)
	}
}

func checkSum(msg []byte) uint16 {
	sum := uint32(0)
	for i := 0; i < len(msg); i += 2 {
		sum += uint32(msg[i+1]) | (uint32(msg[i]) << 8)
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	return uint16(^sum)
}
