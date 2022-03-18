package openapi

import (
	"botgo/dto"
	"botgo/token"
	"golang.org/x/net/context"
	"time"
)

// OpenAPI openapi 完整实现
type OpenAPI interface {
	Base
	UserAPI
	WebsocketAPI
	MessageAPI
}

// Base 基础能力接口
type Base interface {
	Setup(token *token.Token, inSandbox bool) OpenAPI

	// WithTimeout 设置请求接口超时时间
	WithTimeout(duration time.Duration) OpenAPI
}

// UserAPI 用户相关接口
type UserAPI interface {
	Me(ctx context.Context) (*dto.User, error)
	MeGuilds(ctx context.Context, pager *dto.GuildPager) ([]*dto.Guild, error)
}

// WebsocketAPI websocket 接入地址
type WebsocketAPI interface {
	WS(ctx context.Context, params map[string]string, body string) (*dto.WebsocketAP, error)
}

// MessageAPI 消息相关接口
type MessageAPI interface {
	PostMessage(ctx context.Context, channelID string, msg *dto.MessageToCreate) (*dto.Message, error)
}
