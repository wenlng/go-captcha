package tests

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"testing"

	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/slide"
)

var slideTileCapt slide.Captcha

func init() {
	slideTileCapt = slide.New(
		//slide.WithGenGraphNumber(2),
		//slide.WithEnableGraphVerticalRandom(true),
	)

	bgImage, err := loadPng("../.cache/bg.png")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage1, err := loadPng("../.cache/bg1.png")
	if err != nil {
		log.Fatalln(err)
	}

	graphs := getSlideTileGraphArr()

	slideTileCapt.SetResources(
		slide.WithGraphImages(graphs),
		slide.WithBackgrounds([]image.Image{
			bgImage,
			bgImage1,
		}),
		//slide.WithThumbBackgrounds([]image.Image{
		//	img1,
		//}),
	)
}

func getSlideTileGraphArr() []*slide.GraphImage {
	tileImage1, err := loadPng("../.cache/tile-1.png")
	if err != nil {
		log.Fatalln(err)
	}

	tileShadowImage1, err := loadPng("../.cache/tile-shadow-1.png")
	if err != nil {
		log.Fatalln(err)
	}
	tileMaskImage1, err := loadPng("../.cache/tile-mask-1.png")
	if err != nil {
		log.Fatalln(err)
	}

	tileImage2, err := loadPng("../.cache/tile-2.png")
	if err != nil {
		log.Fatalln(err)
	}

	tileShadowImage2, err := loadPng("../.cache/tile-shadow-2.png")
	if err != nil {
		log.Fatalln(err)
	}

	tileMaskImage2, err := loadPng("../.cache/tile-mask-2.png")
	if err != nil {
		log.Fatalln(err)
	}

	return []*slide.GraphImage{
		{
			OverlayImage: tileImage1,
			ShadowImage:  tileShadowImage1,
			MaskImage:    tileMaskImage1,
		},
		{
			OverlayImage: tileImage2,
			ShadowImage:  tileShadowImage2,
			MaskImage:    tileMaskImage2,
		},
	}
}

func TestSlideTileCaptcha(t *testing.T) {
	captData, err := slideTileCapt.Generate()
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
	fmt.Println(captData.GetTileImage().ToBase64())

	err = captData.GetMasterImage().SaveToFile("../.cache/master.jpg", option.QualityNone)
	if err != nil {
		fmt.Println(err)
	}
	err = captData.GetTileImage().SaveToFile("../.cache/thumb.png")
	if err != nil {
		fmt.Println(err)
	}
}
