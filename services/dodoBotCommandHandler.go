package services

import (
	"context"
	"strconv"

	"fmt"
	"runtime/debug"
	"strings"
	"time"

	. "github.com/nft-rainbow/rainbow-app-service/appService-errors"
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

	router["绑定"] = func(req *CommandRequest) error {
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

	b, _ := models.FindBotServerByChannel(msgSource.channelId)
	if b == nil {
		return nil
	}

	methodAndArgs := strings.Split(command, "/")
	if len(methodAndArgs) == 1 {
		return nil
	}
	method := methodAndArgs[1]
	args := make([]string, 10) // avoid empty string
	copy(args, methodAndArgs[2:])
	logrus.WithField("method", method).WithField("args", args).Info("parse commands")

	if d.router[method] == nil {
		msg := GenCommandResponse(CrCommandUnSupport, CmdRespData{DodoSourceId: msgSource.userDodoSourceId})
		_, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
		return err
	}

	err := d.router[method](&CommandRequest{
		ChannelMsgSource: msgSource,
		args:             args,
	})
	if err != nil {
		logrus.WithField("method", method).WithField("args", args).WithField("error stack", fmt.Sprintf("%+v", errors.WithStack(err))).Info("failed run command")
		msg := GenCommandResponse(CrErrUnknown, CmdRespData{DodoSourceId: msgSource.userDodoSourceId})
		_, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
		return err
	}
	return nil
}

func (d *DodoBotCommander) Mint(msgSource ChannelMsgSource, pushInfoIdStr string, verbalSecret string) error {
	pushInfoIdInt, err := strconv.Atoi(pushInfoIdStr)
	if err != nil {
		errResp := GenCommandErrResponse(ERR_BUSINESS_ACTIVITY_NOT_EXIST, CmdRespData{DodoSourceId: msgSource.userDodoSourceId})
		_, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, errResp, msgSource.messageId)
		return err
	}
	pushInfoId := uint(pushInfoIdInt)

	// var activity *models.Activity
	cmdRespData := CmdRespData{
		DodoSourceId: msgSource.userDodoSourceId,
		PushInfoId:   pushInfoId,
	}

	err = func() error {
		var err error
		mainAddr, testAddr, err := GetBindAddress(msgSource.userDodoSourceId, enums.SOCIAL_TOOL_DODO)
		if err != nil {
			return err
		}

		pushInfo, err := models.FindPushInfoById(pushInfoId)
		if err != nil {
			return ERR_BUSINESS_ACTIVITY_NOT_EXIST
		}

		if pushInfo.ChannelId != msgSource.channelId {
			return ERR_BUSINESS_ACTIVITY_NOT_EXIST
		}

		activity := pushInfo.Activity
		if activity.Contract == nil {
			return ERR_BUSNISS_ACTIVITY_CONFIG_WRONG
		}

		msg := GenCommandResponse(CrReadyMint, cmdRespData)
		if _, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId); err != nil {
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

			msg := GenCommandResponse(CrMintSuccess, cmdRespData)
			_, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
			return err
		}
	}()

	if err != nil {
		errResp := GenCommandErrResponse(err, cmdRespData)
		_, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, errResp, msgSource.messageId)
		return err
	}
	return nil
}

func (d *DodoBotCommander) Bind(msgSource ChannelMsgSource, userAddress string) error {
	mainAddr, testAddr, err := BindCfxAddress(msgSource.userDodoSourceId, userAddress, enums.SOCIAL_TOOL_DODO)
	if err != nil {
		if err != ERR_BIND_ADDRESS_WRONG_FORMAT {
			err = ERR_BIND_ADDRESS_OTHER
		}
		errResp := GenCommandErrResponse(err, CmdRespData{DodoSourceId: msgSource.userDodoSourceId})
		_, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, errResp, msgSource.messageId)
		return err
	}

	cmdRespData := CmdRespData{
		MainnetAddress: mainAddr,
		TestnetAddress: testAddr,
		DodoSourceId:   msgSource.userDodoSourceId,
	}
	msg := GenCommandResponse(CrBindSuccess, cmdRespData)
	_, err = d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
	return err
}

func (d *DodoBotCommander) GetAddress(msgSource ChannelMsgSource) error {
	cmdRespData := CmdRespData{
		DodoSourceId: msgSource.userDodoSourceId,
	}

	mainAddr, testAddr, err := GetBindAddress(msgSource.userDodoSourceId, enums.SOCIAL_TOOL_DODO)
	if err != nil {
		errResp := GenCommandErrResponse(err, cmdRespData)
		_, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, errResp, msgSource.messageId)
		return err
	}

	cmdRespData.MainnetAddress = mainAddr
	cmdRespData.TestnetAddress = testAddr
	msg := GenCommandResponse(CrShowAddress, cmdRespData)
	_, err = d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
	return err
}

func (d *DodoBotCommander) GetTutorial(msgSource ChannelMsgSource) error {
	_, err := d.dodoBot.SendDirectMessage(context.Background(), msgSource.serverId, msgSource.userDodoSourceId, guide)
	return err
}

func (d *DodoBotCommander) GetAllCommands(msgSource ChannelMsgSource, language string) error {
	msg := GenCommandResponse(CrSeeOnDirectMessage, CmdRespData{DodoSourceId: msgSource.userDodoSourceId})
	if _, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId); err != nil {
		return err
	}

	_allCommand := CrAllCommandsZh
	if language == "en" {
		_allCommand = CrAllCommandsEn
	}

	_, err := d.dodoBot.SendDirectMessage(context.Background(), msgSource.serverId, msgSource.userDodoSourceId, _allCommand)
	return err
}

func (d *DodoBotCommander) CreateAccount(msgSource ChannelMsgSource) error {
	msg := GenCommandResponse(CrShowCreateAddressDoc, CmdRespData{DodoSourceId: msgSource.userDodoSourceId})
	_, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
	return err
}

func (d *DodoBotCommander) GetVerbalSecret(msgSource ChannelMsgSource, pushInfoIdStr string) error {
	pushInfoIdInt, err := strconv.Atoi(pushInfoIdStr)
	if err != nil {
		errResp := GenCommandErrResponse(ERR_BUSINESS_ACTIVITY_NOT_EXIST, CmdRespData{DodoSourceId: msgSource.userDodoSourceId})
		_, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, errResp, msgSource.messageId)
		return err
	}
	pushInfoId := uint(pushInfoIdInt)

	pushInfo, err := models.FindPushInfoById(pushInfoId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ERR_BUSINESS_ACTIVITY_NOT_EXIST
		}
		errMsg := GenCommandErrResponse(err, CmdRespData{DodoSourceId: msgSource.userDodoSourceId})
		_, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, errMsg, msgSource.messageId)
		return err
	}

	activity := pushInfo.Activity
	cmdRespData := CmdRespData{DodoSourceId: msgSource.userDodoSourceId,
		PushInfoId:   pushInfoId,
		VisperSecret: activity.Command,
	}

	if activity.Command == "" {
		msg := GenCommandResponse(CrnotNeedVisperSecret, cmdRespData)
		_, err := d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
		return err
	}

	msg := GenCommandResponse(CrShowVisperSecret, cmdRespData)
	if _, err := d.dodoBot.SendDirectMessage(context.Background(), msgSource.serverId, msgSource.userDodoSourceId, msg); err != nil {
		return err
	}
	msg = GenCommandResponse(CrSeeOnDirectMessage, CmdRespData{DodoSourceId: msgSource.userDodoSourceId})
	_, err = d.dodoBot.SendChannelMessage(context.Background(), msgSource.channelId, msg, msgSource.messageId)
	return err
}

func customNotFoundError(err error, msgOfNotFound string) error {
	if err == gorm.ErrRecordNotFound {
		return errors.New(msgOfNotFound)
	}
	return err
}
