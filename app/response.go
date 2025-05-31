package main

import "encoding/binary"

type Response struct {
	MessageSize int32
	Header      ResponseHeader
}

type ResponseHeader interface {
	CorrelationId() int32
}

func NewResponse(size, correlationId int32) *Response {
	return &Response{
		MessageSize: size,
		Header:      NewResponseHeaderV0(correlationId),
	}
}


func (r *Response) Encode() []byte {
	buffer := make([]byte, 8)

	// message_size
	binary.BigEndian.PutUint32(buffer[:4], 0)

	// correlation_id
	correlationId := uint32(r.Header.CorrelationId())
	binary.BigEndian.PutUint32(buffer[4:], correlationId)

	return buffer
}

// ########################
// # Response Header (V0) #
// ########################
type ResponseHeaderV0 struct {
	correlationId int32
}

func NewResponseHeaderV0(correlationId int32) *ResponseHeaderV0 {
	return &ResponseHeaderV0{
		correlationId: correlationId,
	}
}

func (v0 *ResponseHeaderV0) CorrelationId() int32 {
	return v0.correlationId
}
