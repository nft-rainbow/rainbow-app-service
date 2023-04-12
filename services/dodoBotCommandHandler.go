package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/sirupsen/logrus"
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

func NewDodoBotCommander(_dodoBot *DodoBot) *DodoBotCommander {
	d := &DodoBotCommander{dodoBot: dodoBot}
	router := make(map[string]func(req *CommandRequest) error)
	router["教程"] = func(req *CommandRequest) error {
		return d.GetTutorial(req.channelId, req.userDodoSourceId)
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
	router["绑定"] = func(req *CommandRequest) error {
		return d.Bind(req.channelId, req.userDodoSourceId, req.args[0])
	}
	d.router = router
	return d
}

func (d *DodoBotCommander) ExcuteCommand(channelId string, userDodoSourceId string, command string) error {
	methodAndArgs := strings.Split(command, "/")
	if len(methodAndArgs) == 1 {
		return nil
	}
	method := methodAndArgs[1]
	args := make([]string, 10) // avoid empty string
	copy(args, methodAndArgs[2:])
	logrus.WithField("method", method).WithField("args", args).Info("parse commands")

	err := d.router[method](&CommandRequest{
		channelId:        channelId,
		userDodoSourceId: userDodoSourceId,
		args:             args,
	})
	if err != nil {
		msg := fmt.Sprintf("<@!%s> %s", userDodoSourceId, err.Error())
		return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
	}
	return nil
}

func (d *DodoBotCommander) Mint(channelId string, userDodoSourceId string, activityId string, verbalSecret string) error {
	bind, err := models.FindBindingCFXAddressById(userDodoSourceId, utils.DoDo)
	if err != nil {
		return err
	}

	config, err := models.FindPOAPActivityConfigById(activityId)
	if err != nil {
		return err
	}

	if err := config.CheckActivityValid(); err != nil {
		return err
	}

	msg := fmt.Sprintf("<@!%s> Start to mint NFT. Please wait patiently.", userDodoSourceId)
	if err := d.dodoBot.SendChannelMessage(context.Background(), channelId, msg); err != nil {
		return err
	}

	res, err := HandlePOAPH5Mint(&POAPRequest{
		ActivityID:  config.ActivityID,
		UserAddress: bind.CFXAddress,
		Command:     verbalSecret,
	})
	if err != nil {
		return err
	}

	for {
		resp, _ := models.FindPOAPResultById(config.ActivityID, int(res.ID))
		if resp.Hash == "" {
			time.Sleep(time.Second)
			continue
		}

		msg = fmt.Sprintf("<@!%s> Mint NFT successfully. The correspding transaction hash is %v", userDodoSourceId, resp.Hash)
		return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
	}
}

func (d *DodoBotCommander) Bind(channelId string, userDodoSourceId string, userAddress string) error {
	err := HandleBindCfxAddress(userDodoSourceId, userAddress, "dodo")
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("<@!%s> success!", userDodoSourceId)
	return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
}

func (d *DodoBotCommander) GetAddress(channelId string, userDodoSourceId string) error {
	resp, err := GetDoDoBindCFXAddress(userDodoSourceId)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("<@!%s> %s", userDodoSourceId, resp)
	return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
}

func (d *DodoBotCommander) GetTutorial(channelId string, userDodoSourceId string) error {
	msg := fmt.Sprintf("<@!%s> %s", userDodoSourceId, guide)
	return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
}

func (d *DodoBotCommander) CreateAccount(channelId string, userDodoSourceId string) error {
	msg := fmt.Sprintf("<@!%s> %s", userDodoSourceId, anywebH5)
	return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
}

func (d *DodoBotCommander) GetVerbalSecret(channelId string, userDodoSourceId string, activityId string) error {
	config, err := models.FindPOAPActivityConfigById(activityId)
	if err != nil {
		return err
	}
	if config.Command == "" {
		msg := fmt.Sprintf("<@!%s> The command is not needed in this activity", userDodoSourceId)
		return d.dodoBot.SendChannelMessage(context.Background(), channelId, msg)
	}
	return d.dodoBot.SendChannelMessage(context.Background(), channelId, config.Command)
}
