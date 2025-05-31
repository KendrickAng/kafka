package main

import "encoding/binary"

// ApiVersionsResponse - see https://kafka.apache.org/protocol.html#The_Messages_ApiVersions.
type ApiVersionsResponse struct {
	Version    int
	ResponseV4 *ApiVersionsResponseV4
}

func (a *ApiVersionsResponse) Encode() []byte {
	errorCodeBuffer := make([]byte, 8)
	binary.BigEndian.PutUint16(errorCodeBuffer, uint16(a.ResponseV4.ErrorCode()))
	return errorCodeBuffer
}

func NewApiVersionsResponseV4(errorCode int16) *ApiVersionsResponse {
	return &ApiVersionsResponse{
		Version: 4,
		ResponseV4: &ApiVersionsResponseV4{
			errorCode: errorCode,
		},
	}
}

type ApiVersionsResponseV4 struct {
	errorCode int16
}

func (v4 *ApiVersionsResponseV4) ErrorCode() int16 {
	return v4.errorCode
}
