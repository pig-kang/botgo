// Package botgo 是一个QQ频道机器人 sdk 的 golang 实现
package main

import (
	"botgo/openapi"
	v1 "botgo/openapi/v1"
	"botgo/token"
	"botgo/websocket/client"
)

func init() {
	v1.Setup()     // 注册 v1 接口
	client.Setup() // 注册 websocket client 实现
}

// NewSessionManager 获得 session manager 实例
func NewSessionManager() SessionManager {
	return defaultSessionManager
}

// NewSandboxOpenAPI 创建测试环境的 openapi 实例
func NewSandboxOpenAPI(token *token.Token) openapi.OpenAPI {
	return openapi.DefaultImpl.Setup(token, true)
}
