/**
 * @Author Awen
 * @Description
 * @Date 2021/7/20
 **/

package main

import (
	"fmt"
	"github.com/wenlng/go-captcha/captcha"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestDrawTextImg(t *testing.T) {
	draw := GetDraw()

	drawDots := &captcha.DrawDot{
		Dx:      0,
		Dy:      0,
		FontDPI: 72,
		Text:    "你好",
		Angle:   45,
		Size:    20,
		Color:   "#841524",
		Width:   20,
		Height:  20,
		Font:    getPWD() + "/__example/resources/fonts/fzshengsksjw_cu.ttf",
	}

	canvas, ap, _ := draw.DrawTextImg(drawDots, &captcha.DrawCanvas{
		Width:      20,
		Height:     20,
		Background: getPWD() + "/__example/resources/images/1.jpg",
	})

	nW := canvas.Bounds().Max.X
	nH := canvas.Bounds().Max.Y
	minX := ap.MinX
	maxX := ap.MaxX
	minY := ap.MinY
	maxY := ap.MaxY
	width := maxX - minX
	height := maxY - minY

	co, _ := captcha.ParseHexColor("#841524")
	var coArr = []color.RGBA{
		co,
	}
	canvas2 := draw.CreateCanvasWithPalette(&captcha.DrawCanvas{
		Width:  width,
		Height: height,
	}, coArr)

	// 开始裁剪
	for x := 0; x < nW; x++ {
		for y := 0; y < nH; y++ {
			co := canvas.At(x, y)
			if _, _, _, a := co.RGBA(); a > 0 {
				canvas2.Set(x, y, canvas.At(x, y))
			}
		}
	}

	file := getPWD() + "/tests/.cache/" + fmt.Sprintf("%v", captcha.RandInt(1, 200)) + "textImg.png"
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer logFile.Close()
	err := png.Encode(logFile, canvas)
	if err != nil {
		panic(err)
	}
}
