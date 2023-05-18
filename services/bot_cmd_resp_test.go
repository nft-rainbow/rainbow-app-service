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
		StartTime:     time.Now(),
		EndTime:       time.Now(),
		ClaimLink:     "http://claim.link",
		ActivityImage: "https://i.seadn.io/gcs/files/0d9dd708e670842a624e585024fa3540.png?auto=format&dpr=1&w=1000",
	}
	pushDataForTemplate := struct {
		PushData
		Roles          string
		StartTime      string
		EndTime        string
		StartTimeInSec int64
	}{
		PushData:       pushData,
		Roles:          "",
		StartTime:      pushData.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:        pushData.EndTime.Format("2006-01-02 15:04:05"),
		StartTimeInSec: pushData.StartTime.Unix(),
	}

	exectued := ExcuteTemplate(CrPushJsonTemplate, pushDataForTemplate)
	fmt.Println(exectued)

	var message model.CardMessage
	err := json.Unmarshal([]byte(exectued), &message)
	assert.NoError(t, err)
}
