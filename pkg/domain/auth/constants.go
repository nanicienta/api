// Package auth contains the logic for validating JWT tokens and handle authentication
package auth

// ContextUserIDKey is the key used to store the user ID in the context
const ContextUserIDKey = "kimosUserId"

// DeviceIdHeaderName is the header name for the device ID
const DeviceIdHeaderName = "X-Device-Id"

// AuthorizationIdHeaderName is the header name for the authorization ID
const AuthorizationIdHeaderName = "Authorization"

// AuthorizationHeaderPrefix is the prefix for the authorization header
const AuthorizationHeaderPrefix = "Bearer "
