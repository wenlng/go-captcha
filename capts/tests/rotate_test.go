package tests

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"testing"

	"github.com/wenlng/go-captcha/capts/rotate"
)

var rotateCapt rotate.Captcha

func init() {
	rotateCapt = rotate.New()

	bgImage, err := loadPng("../.cache/bg.png")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage1, err := loadPng("../.cache/bg1.png")
	if err != nil {
		log.Fatalln(err)
	}

	rotateCapt.SetResources(
		rotate.WithImages([]image.Image{
			bgImage,
			bgImage1,
		}),
	)
}

func TestRotateDirectionCaptcha(t *testing.T) {
	captData, err := rotateCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	blockData := captData.GetData()
	if blockData == nil {
		log.Fatalln(">>>>> generate err")
	}

	block, _ := json.Marshal(blockData)
	fmt.Println(string(block))
	fmt.Println(captData.GetMasterImage().ToBase64())
	fmt.Println(captData.GetThumbImage().ToBase64())

	err = captData.GetMasterImage().SaveToFile("../.cache/master.png")
	if err != nil {
		fmt.Println(err)
	}
	err = captData.GetThumbImage().SaveToFile("../.cache/thumb.png")
	if err != nil {
		fmt.Println(err)
	}
}
