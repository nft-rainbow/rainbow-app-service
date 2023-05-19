package services

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/dodo-open/dodo-open-go/model"
	"github.com/stretchr/testify/assert"
)

func TestCrPushJsonTemplate(t *testing.T) {
	pushData := PushData{
		Content:       "push content",
		PushInfoID:    12,
		ActivityName:  "activity name",
		StartTime:     time.Now().AddDate(0, 0, 2),
		EndTime:       time.Now(),
		ClaimLink:     "http://claim.link",
		ActivityImage: "https://i.seadn.io/gcs/files/0d9dd708e670842a624e585024fa3540.png?auto=format&dpr=1&w=1000",
	}
	pushDataForTemplate := struct {
		PushData
		Roles               string
		StartTime           string
		EndTime             string
		StartTimeInMillisec int64
		CountdownStyle      string
	}{
		PushData:            pushData,
		Roles:               "",
		StartTime:           pushData.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:             pushData.EndTime.Format("2006-01-02 15:04:05"),
		StartTimeInMillisec: pushData.StartTime.UnixMilli(),
		CountdownStyle:      "hour",
	}

	countDown := time.Until(pushData.StartTime)
	if countDown > time.Hour*24 {
		pushDataForTemplate.CountdownStyle = "day"
	}

	exectued := ExcuteTemplate(CrPushJsonTemplate, pushDataForTemplate)
	fmt.Println(exectued)

	var message model.CardMessage
	err := json.Unmarshal([]byte(exectued), &message)
	assert.NoError(t, err)

	if pushData.StartTime.Before(time.Now()) {
		message.Card.Components = append(message.Card.Components[:2], message.Card.Components[3:]...)
	}

	j, _ := json.Marshal(message)
	fmt.Println(string(j))
}
