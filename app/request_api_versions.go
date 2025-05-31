package main

// ApiVersionsRequest - see https://kafka.apache.org/protocol.html#The_Messages_ApiVersions.
type ApiVersionsRequest struct{}

func NewApiVersionsRequest(version int16) ApiVersionsRequest {
	// For now, version is unused.
	return ApiVersionsRequest{}
}
