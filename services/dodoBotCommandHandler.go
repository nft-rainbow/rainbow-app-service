package services

import (
	"context"

	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/nft-rainbow/rainbow-app-service/models"
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
	channelId        string
	userDodoSourceId string
	args             []string
}

const (
	allCommands = "\n" +
		"/教程\t获取教程链接\n" +
		"/查指令\t查所有指令\n" +
		"/铸造/活动ID/TokenID\t铸造nft\n" +
		"/查口令/活动ID\t查活动口令\n" +
		"/创建地址\t创建conflux地址并绑定到您的dodo账户\n" +
		"/查地址\t查询已绑定的conflux地址\n" +
		"/绑定地址/conflux地址\t绑定指定conflux地址到您的dodo账户\n"
)

func NewDodoBotCommander(_dodoBot *DodoBot) *DodoBotCommander {
	d := &DodoBotCommander{dodoBot: _dodoBot}
	router := make(map[string]func(req *CommandRequest) error)
	router["教程"] = func(req *CommandRequest) error {
		return d.GetTutorial(req.channelId, req.userDodoSourceId)
	}
	router["查指令"] = func(req *CommandRequest) error {
		return d.GetAllCommands(req.channelId, req.userDodoSourceId)
	}
	router["铸造"] = func(req *CommandRequest) error {
		return d.Mint(req.channelId, req.userDodoSourceId, req.args[0], req.args[1])
	}
	router["查口令"] = func(req *CommandRequest) error {
		return d.GetVerbalSecret(req.channelId, req.userDodoSourceId, req.args[0])
	}
	router["创建地址"] = func(req *CommandRequest) error {
		return d.CreateAccount(req.channelId, req.userDodoSourceId)
	}
	router["查地址"] = func(req *CommandRequest) error {
		return d.GetAddress(req.channelId, req.userDodoSourceId)
	}
	router["绑定地址"] = func(req *CommandRequest) error {
		return d.Bind(req.channelId, req.userDodoSourceId, req.args[0])
	}
	d.router = router
	return d
}

func (d *DodoBotCommander) ExcuteCommand(channelId string, userDodoSourceId string, command string) error {
	defer func() {
		if err := recover(); err != nil {
			logrus.WithField("command", command).WithField("err", err).WithField("stack", string(debug.Stack())).Info("exception when run command")
			msg := fmt.Sprintf("<@!%s> %s", userDodoSourceId, "发生内部错误")
			d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
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
		msg := fmt.Sprintf("<@!%s> %s", userDodoSourceId, fmt.Sprintf("不支持指令 %s", method))
		return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
	}

	err := d.router[method](&CommandRequest{
		channelId:        channelId,
		userDodoSourceId: userDodoSourceId,
		args:             args,
	})
	if err != nil {
		logrus.WithField("method", method).WithField("args", args).WithField("error stack", fmt.Sprintf("%+v", errors.WithStack(err))).Info("failed run command")
		fmt.Printf("failed run command\n%+v\n", errors.WithStack(err))
		msg := fmt.Sprintf("<@!%s> %s %s", userDodoSourceId, err.Error(), debug.Stack())
		return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
	}
	return nil
}

// TODO: verify channel match activity
func (d *DodoBotCommander) Mint(channelId string, userDodoSourceId string, activityId string, verbalSecret string) error {
	mainAddr, testAddr, err := GetBindAddress(userDodoSourceId, models.SOCIAL_TOOL_DODO)
	if err != nil {
		return err
	}

	activity, err := models.FindActivityByCode(activityId)
	if err != nil {
		return err
	}

	if _, err = models.FindPushInfo(models.PushInfo{ActivityId: activity.ID, ChannelId: channelId}); err != nil {
		return customNotFoundError(err, "the activity not support on this channel")
	}

	msg := fmt.Sprintf("<@!%s> Start to mint NFT. Please wait patiently.", userDodoSourceId)
	if err := d.dodoBot.SendChannelMessage(context.Background(), channelId, msg); err != nil {
		return err
	}

	toAddr := mainAddr
	if activity.Contract.ChainId == int32(utils.CONFLUX_TEST_ID) {
		toAddr = testAddr
	}

	res, err := GetActivityService().HandlePOAPH5Mint(&MintReq{
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

		msg = fmt.Sprintf("<@!%s> Mint NFT successfully. The correspending transaction hash is %v", userDodoSourceId, resp.Hash)
		return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
	}
}

func (d *DodoBotCommander) Bind(channelId string, userDodoSourceId string, userAddress string) error {
	mainAddr, testAddr, err := BindCfxAddress(userDodoSourceId, userAddress, models.SOCIAL_TOOL_DODO)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("<@!%s> Success bind conflux address!\nTestnet\t%s\nMainnet\t%s\n", userDodoSourceId, mainAddr, testAddr)
	return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
}

func (d *DodoBotCommander) GetAddress(channelId string, userDodoSourceId string) error {
	mainAddr, testAddr, err := GetBindAddress(userDodoSourceId, models.SOCIAL_TOOL_DODO)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("<@!%s>\nTestnet\t%s\nMainnet\t%s\n", userDodoSourceId, mainAddr, testAddr)
	return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
}

func (d *DodoBotCommander) GetTutorial(channelId string, userDodoSourceId string) error {
	msg := fmt.Sprintf("<@!%s> %s", userDodoSourceId, guide)
	return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
}

func (d *DodoBotCommander) GetAllCommands(channelId string, userDodoSourceId string) error {
	msg := fmt.Sprintf("<@!%s> %s", userDodoSourceId, allCommands)
	return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
}

func (d *DodoBotCommander) CreateAccount(channelId string, userDodoSourceId string) error {
	msg := fmt.Sprintf("<@!%s> %s", userDodoSourceId, anywebH5)
	return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
}

func (d *DodoBotCommander) GetVerbalSecret(channelId string, userDodoSourceId string, activityId string) error {
	config, err := models.FindActivityByCode(activityId)
	if err != nil {
		return err
	}
	if config.Command == "" {
		msg := fmt.Sprintf("<@!%s> The command is not needed in this activity", userDodoSourceId)
		return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
	}
	return d.dodoBot.SendChannelMessage(context.Background(), channelId, config.Command)
}

func customNotFoundError(err error, msgOfNotFound string) error {
	if err == gorm.ErrRecordNotFound {
		return errors.New(msgOfNotFound)
	}
	return err
}
