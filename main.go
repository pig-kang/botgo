package main

import (
	"botgo/dto"
	"botgo/log"
	"botgo/openapi"
	"botgo/token"
	"botgo/websocket"
	"context"
	"fmt"
	"strings"
	"time"
)

var (
	testGuildID   = "1012082334931374501" // 频道id
	testChannelID = "3791527"             // 子频道id
	ctx           context.Context
	api           openapi.OpenAPI
)

func main() {
	ctx = context.Background()

	// 加载 appid 和 token
	botToken := token.BotToken(101999646, "Hf0DqdZcaEkI03805EYng5BdkeW3klXe")
	// 初始化沙箱环境
	api = NewSandboxOpenAPI(botToken).WithTimeout(3 * time.Second)

	wsInfo, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Error(err)
	}

	// 初始化数据库
	DbErr := initDB()
	if DbErr != nil {
		fmt.Printf("init DB failed,err%v\n", DbErr)
	}

	intent := websocket.RegisterHandlers(
		atMessage(),      // at 机器人事件
		guildMemberAdd(), //机器人新加入频道事件监听
	)

	// 启动 session manager 进行 ws 连接的管理，如果接口返回需要启动多个 shard 的连接，这里也会自动启动多个
	NewSessionManager().Start(wsInfo, botToken, &intent)
}

// 机器人新加入频道事件监听
func guildMemberAdd() websocket.GuildMemberEventHandler {
	return func(event *dto.WSPayload, data *dto.WSGuildMemberData) error {
		api.PostMessage(ctx, testChannelID, &dto.MessageToCreate{
			Content: "哈哈哈哈哈，我! 梗指南！现在正式宣布！我入住了该频道！你们！识相点！多多@本大爷，分分钟让你们不用网上冲浪也能成为一个梗达人" +
				"，懂？！现在来了解我！ (’∇’)シ┳━┳",
		})
		return nil
	}
}

//  @ 机器人事件监听
func atMessage() websocket.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		var ChannelID = data.ChannelID
		// 来个梗
		var isGetMeme = strings.Contains(data.Content, "/来个梗")
		if isGetMeme {
			return getMeme(ChannelID)
		}

		// 二次元经典梗
		var isAcgMeme = strings.Contains(data.Content, "/二次元")
		if isAcgMeme {
			return getAcgMeme(ChannelID)
		}

		// 不明意义的缩写
		var isAbbreviationMeme = strings.Contains(data.Content, "/缩写解释")
		if isAbbreviationMeme {
			return getAbbreviationMeme(ChannelID)
		}

		// 查询梗
		if strings.Contains(data.Content, "/查询梗") {
			var memeName = strings.SplitAfter(data.Content, "/查询梗")[1]
			return getMemeByName(strings.TrimSpace(memeName), ChannelID)
		}

		// @机器人台词
		api.PostMessage(ctx, ChannelID, &dto.MessageToCreate{
			Content: "哈喽，各位频道成员们，我是一个木得感情的科普机器人┗( ▔, ▔ )┛" +
				"我将提供：科普新梗网络热词，了解热词，场景热词等等网上冲浪小词语~\n" +
				"如何召唤我：#输入“@梗指南” 或者 #选择想执行的指令，快尝试召唤我叭 |ू･ω･` )",
		})
		return nil
	}
}

// 来个梗 指令处理
func getMeme(ChannelID string) error {
	meme := query()
	api.PostMessage(ctx, ChannelID, &dto.MessageToCreate{
		Content: meme.Name + "\n" + meme.Description,
	})
	return nil
}

// 二次元经典梗 指令处理
func getAcgMeme(ChannelID string) error {
	meme := queryTagMeme(1)
	api.PostMessage(ctx, ChannelID, &dto.MessageToCreate{
		Content: meme.Name + "\n" + meme.Description,
	})
	return nil
}

// 不明意义的缩写 指令处理
func getAbbreviationMeme(ChannelID string) error {
	meme := queryTagMeme(2)
	api.PostMessage(ctx, ChannelID, &dto.MessageToCreate{
		Content: meme.Name + "\n" + meme.Description,
	})
	return nil
}

// 查询梗 指令处理
func getMemeByName(memeName string, ChannelID string) error {
	meme := queryMemeByName(memeName)

	for _, value := range meme {
		api.PostMessage(ctx, ChannelID, &dto.MessageToCreate{
			Content: value.Name + "\n" + value.Description,
		})
	}

	return nil
}
