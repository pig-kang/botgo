// Package token 用于调用 openapi，websocket 的 token 对象。
package token

import (
	"fmt"
	_ "fmt"
)

// Type token 类型
type Type string

// TokenType
const (
	TypeBot    Type = "Bot"
	TypeNormal Type = "Bearer"
)

// Token 用于调用接口的 token 结构
type Token struct {
	AppID       uint64
	AccessToken string
	Type        Type
}

// GetString 获取授权头字符串
func (t *Token) GetString() string {
	if t.Type == TypeNormal {
		return t.AccessToken
	}
	return fmt.Sprintf("%v.%s", t.AppID, t.AccessToken)
}

// BotToken 机器人身份的 token
func BotToken(appID uint64, accessToken string) *Token {
	return &Token{
		AppID:       appID,
		AccessToken: accessToken,
		Type:        TypeBot,
	}
}

// New 创建一个新的 Token
func New(tokenType Type) *Token {
	return &Token{
		Type: tokenType,
	}
}
