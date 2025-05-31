package main

import (
	"bytes"
	"encoding/binary"
)

type Response struct {
	MessageSize int32
	Header      ResponseHeader
	Body        ResponseBody
}

// TODO: maybe use the builder pattern here?
func NewResponseWithoutBody(size, correlationId int32) *Response {
	return &Response{
		MessageSize: size,
		Header:      NewResponseHeaderV0(correlationId),
	}
}

func NewResponseForApiVersions(size, correlationId int32, errorCode int16) *Response {
	resp := NewResponseWithoutBody(size, correlationId)
	resp.Body.ApiVersions = NewApiVersionsResponseV4(errorCode)
	return resp
}

func (r *Response) Encode() []byte {
	var buffer bytes.Buffer

	// message_size
	messageSizeBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(messageSizeBuffer, uint32(r.MessageSize))
	buffer.Write(messageSizeBuffer)

	// correlation_id
	correlationIdBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(correlationIdBuffer, uint32(r.Header.CorrelationId()))
	buffer.Write(correlationIdBuffer)

	// body
	buffer.Write(r.Body.Encode())

	return buffer.Bytes()
}

type ResponseHeader interface {
	CorrelationId() int32
}

type ResponseBody struct {
	ApiVersions *ApiVersionsResponse
}

func (rb *ResponseBody) Encode() []byte {
	return rb.ApiVersions.Encode()
}

// ResponseHeaderV0 - see https://kafka.apache.org/protocol.html#protocol_messages.
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
