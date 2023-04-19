package services

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

func drawPoster(templatePath string, fontPath string,
	activityId string, activityPicUrl string,
	name, description string, startTime, endTime int) (*bytes.Buffer, error) {
	// now := time.Now()

	var dc *gg.Context
	paintSig := make(chan interface{}, 2)

	drawBackground := func() error {
		templateImg, err := gg.LoadImage(templatePath)
		if err != nil {
			return err
		}

		dc = gg.NewContext(templateImg.Bounds().Dx(), templateImg.Bounds().Dy())
		// fmt.Printf("0 %v\n", time.Since(now))
		dc.DrawImage(templateImg, 0, 0)
		// fmt.Printf("1 %v\n", time.Since(now))
		for i := 0; i < 2; i++ {
			paintSig <- struct{}{}
		}
		// fmt.Println("loadTemplate done")
		return nil
	}

	drawHeadPic := func() error {
		resp, err := http.Get(activityPicUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		imgData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		img, err := imaging.Decode(bytes.NewReader(imgData))
		if err != nil {
			return err
		}
		img = imaging.Fill(img, 1260, 1260, 0, imaging.ResampleFilter{})
		<-paintSig
		dc.DrawImage(img, 120, 200)
		// fmt.Printf("2 %v\n", time.Since(now))
		// fmt.Println("drawBackground done")
		return nil
	}

	drawTexts := func() error {
		<-paintSig
		// 增加文字
		err := dc.LoadFontFace(fontPath, 88)
		if err != nil {
			return err
		}
		dc.SetHexColor("#05001F")
		dc.DrawStringAnchored(name, 120, 1580, 0, 0)

		// fmt.Printf("3 %v\n", time.Since(now))
		// err = dc.LoadFontFace("./assets/fonts/PingFang.ttf", 64)
		err = dc.LoadFontFace(fontPath, 48)
		if err != nil {
			return err
		}

		lines := []string{""}
		var lineLen float64
		for _, r := range description {
			w, _ := dc.MeasureString(string(r))
			if lineLen+w > 1260 {
				if len(lines) == 3 {
					lastLine := lines[len(lines)-1]
					lines[len(lines)-1] = lastLine[:len(lastLine)-5] + "..."
					break
				}
				lines = append(lines, "")
				lineLen = 0
			}
			lines[len(lines)-1] += string(r)
			lineLen += w
		}

		paintHeight := float64(1732)
		addPaintHeight := func(delta int) float64 {
			_paintHeight := paintHeight
			paintHeight += float64(delta)
			return _paintHeight
		}

		dc.SetHexColor("#696679")
		for _, line := range lines {
			dc.DrawString(line, 120, addPaintHeight(96))
		}

		var start, end string
		if startTime == -1 {
			start = "不限时"
		} else {
			start = time.Unix(int64(startTime), 0).Format("2006-01-02")
		}
		if endTime == -1 {
			end = "不限时"
		} else {
			end = time.Unix(int64(endTime), 0).Format("2006-01-02")
		}

		err = dc.LoadFontFace(fontPath, 64)
		if err != nil {
			return err
		}
		paintHeight = 2084
		dc.DrawStringAnchored(fmt.Sprintf("开始时间：%v", start), 120, addPaintHeight(96), 0, 0)
		dc.DrawStringAnchored(fmt.Sprintf("结束时间：%v", end), 120, addPaintHeight(96), 0, 0)
		// fmt.Printf("4 %v\n", time.Since(now))

		// QR Code Generate
		targetUrl := generateActivityURLById(activityId)
		qrCode, err := qrcode.New(targetUrl, qrcode.Low)
		if err != nil {
			return err
		}

		paintHeight = 2245
		qrImg := qrCode.Image(268)
		dc.DrawImage(qrImg, 1112, int(paintHeight))
		// fmt.Printf("5 %v\n", time.Since(now))
		// fmt.Println("drawTexts done")
		return nil
	}

	group := new(errgroup.Group)
	group.Go(drawBackground)
	group.Go(drawHeadPic)
	group.Go(drawTexts)
	err := group.Wait()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// encode to png
	buf := new(bytes.Buffer)
	if err := dc.EncodePNG(buf); err != nil {
		return nil, errors.WithStack(err)
	}
	// fmt.Printf("6 %v\n", time.Since(now))

	return buf, nil
}

func generateActivityPoster(config *models.ActivityReq, activityId string) (string, error) {
	// if err := config.CheckActivityValid(); err != nil {
	// 	return "", err
	// }

	buf, err := drawPoster("./assets/images/activityPoster.png",
		"./assets/fonts/PingFang.ttf",
		activityId,
		config.ActivityPictureURL,
		config.Name,
		config.Description,
		int(config.StartedTime),
		int(config.EndedTime),
	)
	if err != nil {
		return "", errors.WithStack(err)
	}

	bucket, err := getOSSBucket(viper.GetString("oss.bucketName"))
	if err != nil {
		return "", errors.WithStack(err)
	}
	if err := bucket.PutObject(path.Join(viper.GetString("posterDir.activity"), activityId+".png"), buf); err != nil {
		return "", errors.WithStack(err)
	}

	url := generateAcvitivyPosterUrl(activityId)
	logrus.WithField("url", url).Info("write activity poster img")
	return url, nil
}

func generateResultPoster(result *models.POAPResult, name string) error {
	templateImg, err := gg.LoadImage("./assets/images/resultPoster.png")

	dc := gg.NewContext(templateImg.Bounds().Dx(), templateImg.Bounds().Dy())
	dc.DrawImage(templateImg, 0, 0)

	resp, err := http.Get(generateActivityUrlByFileUrl(result.FileURL, result.ActivityCode))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	imgData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	img, err := imaging.Decode(bytes.NewReader(imgData))
	if err != nil {
		return err
	}
	img = imaging.Fit(img, 1260, 1260, imaging.Lanczos)
	dc.DrawImage(img, 120, 200)

	// QR Code Generate
	targetUrl := generateActivityURLById(result.ActivityCode)
	qrCode, _ := qrcode.New(targetUrl, qrcode.Low)
	qrImg := qrCode.Image(268)
	dc.DrawImage(qrImg, 1112, 2212)

	// 增加文字
	err = dc.LoadFontFace("./assets/fonts/PingFang.ttf", 88)
	if err != nil {
		return err
	}
	dc.SetHexColor("#05001F")
	dc.DrawStringAnchored(name, 120, 1708, 0, 0)

	err = dc.LoadFontFace("./assets/fonts/PingFang.ttf", 64)
	if err != nil {
		return err
	}
	dc.SetHexColor("#696679")
	x := 120.00
	dc.DrawString("由「", x, 1580)
	w, _ := dc.MeasureString("由「")
	x += w
	dc.SetHexColor("#6953EF")

	dc.DrawString(fmt.Sprintf("%v", utils.SimpleAddress(result.Address)), x, 1580)
	w, _ = dc.MeasureString(fmt.Sprintf("%v", utils.SimpleAddress(result.Address)))
	x += w
	dc.SetHexColor("#696679")
	dc.DrawString("」拥有", x, 1580)

	drawTimeStringWithColor(dc, "：", fmt.Sprintf("徽章编号：%v", result.TokenID), 120, 1908, "#6953EF")
	drawTimeStringWithColor(dc, "：", fmt.Sprintf("领取时间：%v", result.CreatedAt.Format("2006-01-02")), 120, 2036, "#05001F")
	buf := new(bytes.Buffer)
	dc.EncodePNG(buf)

	bucket, err := getOSSBucket(viper.GetString("oss.bucketName"))
	if err != nil {
		return err
	}

	if err := bucket.PutObject(path.Join(viper.GetString("posterDir.result"), result.ActivityCode, result.Address, strconv.Itoa(int(result.ID))+".png"), buf); err != nil {
		return err
	}

	return nil
}

func AddLogoAndUpload(url, name, activity string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	imgData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return err
	}

	logoFile, err := os.Open("./assets/images/logo.png")
	if err != nil {
		return err
	}
	defer logoFile.Close()

	logo, _, err := image.Decode(logoFile)
	if err != nil {
		return err
	}

	withLogo, err := addLogo(img, logo)
	if err != nil {
		return err
	}

	bucket, err := getOSSBucket(viper.GetString("oss.bucketName"))
	if err != nil {
		return err
	}

	if err := bucket.PutObject(path.Join(viper.GetString("imagesDir.minted"), activity, name), bytes.NewReader(withLogo)); err != nil {
		return err
	}

	return nil
}
