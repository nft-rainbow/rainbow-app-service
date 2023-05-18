package services

import (
	"bytes"
	"text/template"

	. "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/pkg/errors"
)

const (
	CrReadyMint   = "<@!{{.DodoSourceId}}> ⭕️准备铸造活动 {{.PushInfoId}} NFT, 请耐心等待......"
	CrMintSuccess = "<@!{{.DodoSourceId}}> ⭕活动 {{.PushInfoId}} 铸造成功！请到区块链浏览器或绑定的钱包查看"

	CrBindSuccess          = "<@!{{.DodoSourceId}}> ⭕️绑定成功。\n主网: \t{{.MainnetAddress}}\n测试网: \t{{.TestnetAddress}}\n"
	CrShowVisperSecret     = "活动 {{.PushInfoId}} 领取口令为: {{.VisperSecret}}"
	CrnotNeedVisperSecret  = "活动 {{.PushInfoId}} 无需领取口令"
	CrShowCreateAddressDoc = "<@!{{.DodoSourceId}}> 查看如何创建钱包地址: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	CrShowAddress          = "<@!{{.DodoSourceId}}> 您绑定的地址为: \n主网: \t{{.MainnetAddress}}\n测试网: \t{{.TestnetAddress}}"
	CrSeeOnDirectMessage   = "<@!{{.DodoSourceId}}> 请查看RainbowBot的私信"

	CrAllCommandsZh = "指令集:\n" +
		"1./帮助——查询指令集\n" +
		"2./铸造/ID——铸造社区活动NFT「例如: /铸造/1234」\n" +
		"3./查口令/ID——查询社区活动口令「例如: /查口令/1234」\n" +
		"4./铸造/ID/口令——使用口令铸造社区活动NFT「例如: /铸造/1234/123456」\n" +
		"5./绑定/cfx:123456——绑定钱包地址\n" +
		"6./查地址——查询绑定的钱包地址\n" +
		"7./创建地址——学习如何创建钱包地址\n"

	CrAllCommandsEn = "commands:\n" +
		"/help\t查所有指令\n" +
		"/tutorial\t获取教程链接\n" +
		"/mint/activity_id/token_id\t铸造nft\n" +
		"/command/activity_id\t查活动口令\n" +
		"/create_address\t创建conflux地址并绑定到您的dodo账户\n" +
		"/address\t查询已绑定的conflux地址\n" +
		"/bind_address/conflux地址\t绑定指定conflux地址到您的dodo账户\n"
)

var (
	CrPushJsonTemplate = `{
		"card": {
			"type": "card",
			"components": [
				{
					"type": "header",
					"text": {
						"type": "dodo-md",
						"content": "{{.Content}}"
					}
				},
				{
					"type": "section",
					"text": {
						"type": "dodo-md",
						"content": "**ID: {{.PushInfoID}}**\n**活动名称：{{.ActivityName}}**\n**开始日期：{{.StartTime}}**\n**结束日期：{{.EndTime}}**"
					}
				},
				{
					"type": "countdown",
					"title": "活动开始计时:",
					"style": "hour",
					"endTime": {{.StartTimeInMillisec}}
				},
				{
					"type": "image",
					"src": "{{.ActivityImage}}"
				},
				{
					"type": "button-group",
					"elements": [
						{
							"type": "button",
							"interactCustomId": "claim",
							"click": {
								"value": "{{.ClaimLink}}",
								"action": "link_url"
							},
							"color": "green",
							"name": "去领取"
						}
					]
				},
				{
					"type": "remark",
					"elements": [{
							"type": "image",
							"src": "https://console.nftrainbow.cn/nftrainbow-logo-icon.png"
						}, {
							"type": "dodo-md",
							"content": "NFTRainbowBot帮助您使用POAP和专属NFT奖励以及管理您的社区。"
						}
					]
				}
			],
			"theme": "green",
			"title": "NFT Rainbow"
		}
	}`
)

var (
	CrErrResponses = map[RainbowAppServiceError]string{
		// ❌未绑定钱包地址！
		ERR_BUSINESS_NOT_BIND_WALLET: "<@!{{.DodoSourceId}}> ❌未绑定钱包地址！",
		// ❌此活动不存在。
		ERR_BUSINESS_ACTIVITY_NOT_EXIST: "<@!{{.DodoSourceId}}>❌此活动不支持或不存在。",
		// ❌活动1234铸造失败。活动未开始。
		ERR_BUSINESS_TIME_EARLY: "<@!{{.DodoSourceId}}>❌活动{{.PushInfoId}}铸造失败。活动未开始。",
		// ❌活动1234铸造失败。活动已过期。
		ERR_BUSINESS_TIME_EXPIRED: "<@!{{.DodoSourceId}}>❌活动{{.PushInfoId}}铸造失败。活动已过期。",
		// ❌活动1234铸造失败。NFT已领取完。
		ERR_BUSINESS_ACTIVITY_MAX_AMOUNT_ARRIVED: "<@!{{.DodoSourceId}}>❌活动{{.PushInfoId}}铸造失败。NFT已领取完。",
		// ❌活动1234铸造失败。超过活动领取限制。
		ERR_BUSINESS_PERSONAL_MAX_AMOUNT_ARRIVED: "<@!{{.DodoSourceId}}>❌活动{{.PushInfoId}}铸造失败。超过活动领取限制。",
		// ❌活动1234铸造失败。需要领取口令。使用帮助指令查看如何获取领取口令。
		ERR_BUSINESS_MISS_VISPER: "<@!{{.DodoSourceId}}> ❌活动{{.PushInfoId}}铸造失败。需要领取口令。使用帮助指令查看如何获取领取口令。",
		// ❌活动1234铸造失败。需要领取口令。使用帮助指令查看如何获取领取口令。
		ERR_BUSINESS_VISPER_WRONG: "<@!{{.DodoSourceId}}> ❌活动{{.PushInfoId}}铸造失败。领取口令错误。",
		// ❌活动1234铸造失败。无领取资格。
		ERR_BUSINESS_NO_MINT_PERMISSIION: "<@!{{.DodoSourceId}}> ❌活动{{.PushInfoId}}铸造失败。无领取资格。",

		ERR_BUSNISS_ACTIVITY_CONFIG_WRONG: "<@!{{.DodoSourceId}}> ❌活动{{.PushInfoId}}配置错误。请联系管理员。",
		// ❌绑定失败，地址格式不正确。
		ERR_BIND_ADDRESS_WRONG_FORMAT: "<@!{{.DodoSourceId}}> ❌绑定失败，地址格式不正确。",
		// ❌绑定失败，发生未知错误，请重新绑定。
		ERR_BIND_ADDRESS_OTHER: "<@!{{.DodoSourceId}}> ❌绑定失败，发生未知错误，请重新绑定。",
	}
	CrErrUnknown       = "<@!{{.DodoSourceId}}> ❌发生未知错误，请重试。"
	CrCommandUnSupport = "<@!{{.DodoSourceId}}> ❌不支持的指令！"
)

type CmdRespData struct {
	DodoSourceId   string
	PushInfoId     uint
	MainnetAddress string
	TestnetAddress string
	VisperSecret   string
}

func GenCommandErrResponse(err error, data CmdRespData) string {
	var _template string
	e, ok := errors.Cause(err).(RainbowAppServiceError)
	if ok && CrErrResponses[e] != "" {
		_template = CrErrResponses[e]
	} else {
		_template = CrErrUnknown
	}
	return GenCommandResponse(_template, data)
}

func GenCommandResponse(crTemplateStr string, data CmdRespData) string {
	return ExcuteTemplate(crTemplateStr, data)
}

func ExcuteTemplate(_template string, data interface{}) string {
	t := template.Must(template.New("").Parse(_template))
	buf := bytes.NewBuffer(nil)
	t.Execute(buf, data)
	return buf.String()
}
