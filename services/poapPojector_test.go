package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawPoster(t *testing.T) {
	_, err := drawPoster("/Users/dayong/myspace/mywork/rainbow-app-service/assets/images/activityPoster.png",
		"/Users/dayong/myspace/mywork/rainbow-app-service/assets/fonts/PingFang.ttf",
		"12345", "https://nftrainbow.oss-cn-hangzhou.aliyuncs.com/events/changandao-event.jpeg",
		"问DAO长安WEB3数字艺术节",
		"问DAO长安WEB3数字艺术节由Tang DAO与西安文化科技创业城联合举办，活动融合展览、峰会、演出、周边集市等活动，打造沉浸式Web3体验，引爆流量，是国内首个具有破圈影响力的Web3艺术节。",
		-1, -1)
	assert.NoError(t, err)
	// fmt.Println(url)
}
