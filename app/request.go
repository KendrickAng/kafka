package main

import "encoding/binary"

type Request struct {
	MessageSize int32
	Header      RequestHeader
}

type RequestHeader interface {
	RequestApiVersion() int16
	CorrelationId() int32
}

func NewRequest(size int32, headerAndBody []byte) *Request {
	return &Request{
		MessageSize: size,
		Header:      NewRequestHeaderV2(headerAndBody),
	}
}

// RequestHeaderV2 - see https://kafka.apache.org/protocol.html#protocol_messages.
type RequestHeaderV2 struct {
	RequestApiKey     int16
	requestApiVersion int16
	correlationId     int32

	// TODO: handle these later
	ClientId  string
	TagBuffer []string
}

func NewRequestHeaderV2(headerAndBody []byte) *RequestHeaderV2 {
	return &RequestHeaderV2{
		RequestApiKey:     int16(binary.BigEndian.Uint16(headerAndBody[0:2])),
		requestApiVersion: int16(binary.BigEndian.Uint16(headerAndBody[2:4])),
		correlationId:     int32(binary.BigEndian.Uint32(headerAndBody[4:8])),
	}
}

func (v2 *RequestHeaderV2) CorrelationId() int32 {
	return v2.correlationId
}

func (v2 *RequestHeaderV2) RequestApiVersion() int16 { return v2.requestApiVersion }
