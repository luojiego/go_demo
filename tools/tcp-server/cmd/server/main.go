package main

import (
	"demo1/frame"
	"demo1/packet"
	"fmt"
	"net"
)

func handlePacket(framePayload []byte) (ackFramePayload []byte, err error) {
	var p packet.Packet
	p, err = packet.Decode(framePayload)
	if err != nil {
		fmt.Println("handlePacket: packet decode error: ", err)
		return
	}

	switch p.(type) {
	case *packet.Conn:
		c := p.(*packet.Conn)
		fmt.Printf("recv conn: id = %s, payload = %s\n", c.ID, c.Payload)
		connAck := &packet.ConnAck{
			ID:     c.ID,
			Result: 0,
		}
		ackFramePayload, err = packet.Encode(connAck)
		if err != nil {
			fmt.Println("handlePacket: packet encode err", err)
			return nil, err
		}
		return ackFramePayload, nil
	case *packet.Submit:
		submit := p.(*packet.Submit)
		fmt.Printf("recv submit: id = %s, payload = %s\n", submit.ID, submit.Payload)
		submitAck := &packet.SubmitAck{
			ID:     submit.ID,
			Result: 0,
		}
		ackFramePayload, err = packet.Encode(submitAck)
		if err != nil {
			fmt.Println("handlePacket: packet encode error: ", err)
			return nil, err
		}
		return ackFramePayload, nil
	default:
		return nil, fmt.Errorf("unknown packet type")
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	frameCodec := frame.NewMyFrameCodec()
	for {
		framePayload, err := frameCodec.Decode(c)
		if err != nil {
			fmt.Println("handleConn: frame decodec error: ", err)
			return
		}

		ackFramePayload, err := handlePacket(framePayload)
		if err != nil {
			fmt.Println("handleConn: handle packet error: ", err)
			return
		}

		err = frameCodec.Encode(c, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode error: ", err)
			return
		}

	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			break
		}
		go handleConn(c)
	}
}
