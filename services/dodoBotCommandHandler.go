package services

import (
	"context"

	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DodoBotCommander struct {
	dodoBot *DodoBot
	router  map[string]func(*CommandRequest) error
}

type CommandRequest struct {
	ChannelMsgSource
	args []string
}

const (
	allCommandsZh = "\n" +
		"/帮助\t查所有指令\n" +
		"/教程\t获取教程链接\n" +
		"/铸造/活动ID/TokenID\t铸造nft\n" +
		"/查口令/活动ID\t查活动口令\n" +
		"/创建地址\t创建conflux地址并绑定到您的dodo账户\n" +
		"/查地址\t查询已绑定的conflux地址\n" +
		"/绑定地址/conflux地址\t绑定指定conflux地址到您的dodo账户\n\n"

	allCommandsEn = "\n" +
		"/help\t查所有指令\n" +
		"/tutorial\t获取教程链接\n" +
		"/mint/activity_id/token_id\t铸造nft\n" +
		"/command/activity_id\t查活动口令\n" +
		"/create_address\t创建conflux地址并绑定到您的dodo账户\n" +
		"/address\t查询已绑定的conflux地址\n" +
		"/bind_address/conflux地址\t绑定指定conflux地址到您的dodo账户\n"
)

func NewDodoBotCommander(_dodoBot *DodoBot) *DodoBotCommander {
	d := &DodoBotCommander{dodoBot: _dodoBot}
	router := make(map[string]func(req *CommandRequest) error)

	router["帮助"] = func(req *CommandRequest) error {
		return d.GetAllCommands(req.ChannelMsgSource, "zh")
	}

	router["教程"] = func(req *CommandRequest) error {
		return d.GetTutorial(req.ChannelMsgSource)
	}

	router["铸造"] = func(req *CommandRequest) error {
		return d.Mint(req.ChannelMsgSource, req.args[0], req.args[1])
	}

	router["查口令"] = func(req *CommandRequest) error {
		return d.GetVerbalSecret(req.ChannelMsgSource, req.args[0])
	}

	router["创建地址"] = func(req *CommandRequest) error {
		return d.CreateAccount(req.ChannelMsgSource)
	}

	router["查地址"] = func(req *CommandRequest) error {
		return d.GetAddress(req.ChannelMsgSource)
	}

	router["绑定地址"] = func(req *CommandRequest) error {
		return d.Bind(req.ChannelMsgSource, req.args[0])
	}

	/*======================= en =======================*/
	router["help"] = func(req *CommandRequest) error {
		return d.GetAllCommands(req.ChannelMsgSource, "en")
	}
	router["tutorial"] = router["教程"]
	router["mint"] = router["铸造"]
	router["command"] = router["查口令"]
	router["create_address"] = router["创建地址"]
	router["address"] = router["查地址"]
	router["bind_address"] = router["绑定地址"]

	d.router = router
	return d
}

func (d *DodoBotCommander) ExcuteCommand(msgSource ChannelMsgSource, command string) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("exception when run command\n%s", debug.Stack())
			logrus.WithField("command", command).WithField("err", err).WithField("stack", string(debug.Stack())).Info("exception when run command")
			msg := fmt.Sprintf("<@!%s> %s", msgSource.userDodoSourceId, "发生内部错误")
			d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
		}
	}()

	methodAndArgs := strings.Split(command, "/")
	if len(methodAndArgs) == 1 {
		return nil
	}
	method := methodAndArgs[1]
	args := make([]string, 10) // avoid empty string
	copy(args, methodAndArgs[2:])
	logrus.WithField("method", method).WithField("args", args).Info("parse commands")

	if d.router[method] == nil {
		msg := fmt.Sprintf("<@!%s> %s", msgSource.userDodoSourceId, fmt.Sprintf("不支持的指令 %s", method))
		return d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
	}

	err := d.router[method](&CommandRequest{
		ChannelMsgSource: msgSource,
		args:             args,
	})
	if err != nil {
		logrus.WithField("method", method).WithField("args", args).WithField("error stack", fmt.Sprintf("%+v", errors.WithStack(err))).Info("failed run command")
		fmt.Printf("failed run command\n%+v\n", errors.WithStack(err))
		msg := fmt.Sprintf("<@!%s> %s", msgSource.userDodoSourceId, err.Error())
		return d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
	}
	return nil
}

func (d *DodoBotCommander) Mint(msgSource ChannelMsgSource, activityId string, verbalSecret string) error {
	mainAddr, testAddr, err := GetBindAddress(msgSource.userDodoSourceId, enums.SOCIAL_TOOL_DODO)
	if err != nil {
		return err
	}

	activity, err := models.FindActivityByCode(activityId)
	if err != nil {
		return err
	}

	if _, err = models.FindPushInfo(models.PushInfo{ActivityId: activity.ID, ChannelId: msgSource.channelId}); err != nil {
		return customNotFoundError(err, "the activity not support on this channel")
	}

	msg := fmt.Sprintf("<@!%s> 开始铸造, 大约需要30秒, 请耐心等待", msgSource.userDodoSourceId)
	if err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId); err != nil {
		return err
	}

	toAddr := mainAddr
	if activity.Contract.ChainId == int32(utils.CONFLUX_TEST_ID) {
		toAddr = testAddr
	}

	res, err := GetActivityService().H5Mint(&MintReq{
		ActivityID:  activity.ActivityCode,
		UserAddress: toAddr,
		Command:     verbalSecret,
	})
	if err != nil {
		return err
	}

	for {
		resp, _ := models.FindPOAPResultById(activity.ActivityCode, int(res.ID))
		if resp.Hash == "" {
			time.Sleep(time.Second)
			continue
		}

		msg = fmt.Sprintf("<@!%s> 铸造成功。 交易hash为 %v", msgSource.userDodoSourceId, resp.Hash)
		return d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
	}
}

func (d *DodoBotCommander) Bind(msgSource ChannelMsgSource, userAddress string) error {
	mainAddr, testAddr, err := BindCfxAddress(msgSource.userDodoSourceId, userAddress, enums.SOCIAL_TOOL_DODO)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("<@!%s> 您已成功绑定 conflux 地址！\nTestnet\t%s\nMainnet\t%s\n", msgSource.userDodoSourceId, mainAddr, testAddr)
	return d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
}

func (d *DodoBotCommander) GetAddress(msgSource ChannelMsgSource) error {
	mainAddr, testAddr, err := GetBindAddress(msgSource.userDodoSourceId, enums.SOCIAL_TOOL_DODO)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("<@!%s> 您的地址:\nTestnet\t%s\nMainnet\t%s\n", msgSource.userDodoSourceId, mainAddr, testAddr)
	return d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
}

func (d *DodoBotCommander) GetTutorial(msgSource ChannelMsgSource) error {
	// msg := fmt.Sprintf("<@!%s> %s", msgSource.userDodoSourceId, guide)
	return d.dodoBot.SendDirectMessage(context.Background(), msgSource.serverId, msgSource.userDodoSourceId, guide)
}

func (d *DodoBotCommander) GetAllCommands(msgSource ChannelMsgSource, language string) error {
	_allCommand := allCommandsZh
	if language == "en" {
		_allCommand = allCommandsEn
	}
	msg := fmt.Sprintf("<@!%s> %s", msgSource.userDodoSourceId, _allCommand)
	return d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
}

func (d *DodoBotCommander) CreateAccount(msgSource ChannelMsgSource) error {
	msg := fmt.Sprintf("<@!%s> %s", msgSource.userDodoSourceId, anywebH5)
	return d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
}

func (d *DodoBotCommander) GetVerbalSecret(msgSource ChannelMsgSource, activityCode string) error {
	config, err := models.FindActivityByCode(activityCode)
	if err != nil {
		return err
	}
	if config.Command == "" {
		msg := fmt.Sprintf("<@!%s> %s 不需要口令码", msgSource.userDodoSourceId, activityCode)
		return d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
	}

	msg := fmt.Sprintf("活动 %s 的口令码是 %s", activityCode, config.Command)
	return d.dodoBot.SendDirectMessage(context.Background(), msgSource.serverId, msgSource.userDodoSourceId, msg)
	// return d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, config.Command, msgSource.messageId)
}

func customNotFoundError(err error, msgOfNotFound string) error {
	if err == gorm.ErrRecordNotFound {
		return errors.New(msgOfNotFound)
	}
	return err
}
