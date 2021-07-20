/**
 * @Author Awen
 * @Description
 * @Date 2021/7/20
 **/

package main

import (
	"fmt"
	"github.com/wenlng/go-captcha/captcha"
	"testing"
)

func TestImageSize(t *testing.T) {
	capt := getCaptcha()

	capt.SetImageSize(&captcha.Size{Width: 300, Height: 300})

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestSetThumbSize(t *testing.T) {
	capt := getCaptcha()

	capt.SetThumbSize(&captcha.Size{Width: 300, Height: 300})

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestChars(t *testing.T) {
	capt := getCaptcha()
	//chars := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//capt.SetRangChars(strings.Split(chars, ""))
	chars := []string{"HE","CA","WO","NE","HT","IE","PG","GI","CH","CO","DA"}
	_ = capt.SetRangChars(chars)
	//chars := []string{"你","好","呀","这","是","点","击","验","证","码","哟"}
	//capt.SetRangChars(chars)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestColors(t *testing.T) {
	capt := getCaptcha()
	capt.SetTextRangFontColors([]string{
		"#1d3f84",
		"#3a6a1e",
		"#712217",
		"#885500",
		"#392585",
	})

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestAlpha(t *testing.T) {
	capt := getCaptcha()

	capt.SetImageFontAlpha(0.5)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestImageFontDistort(t *testing.T) {
	capt := getCaptcha()

	capt.SetImageFontDistort(captcha.ThumbBackgroundDistortLevel2)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestRangAnglePos(t *testing.T) {
	capt := getCaptcha()

	rang := []*captcha.RangeVal{
		{1, 15},
		{15, 30},
		{30, 45},
		{315, 330},
		{330, 345},
		{345, 359},
	}
	capt.SetTextRangAnglePos(rang)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestThumbBackground(t *testing.T) {
	capt := getCaptcha()

	capt.SetThumbBackground([]string{
		getPWD() + "/__example/resources/images/thumb/r1.jpg",
		getPWD() + "/__example/resources/images/thumb/r2.jpg",
		getPWD() + "/__example/resources/images/thumb/r3.jpg",
		getPWD() + "/__example/resources/images/thumb/r4.jpg",
		getPWD() + "/__example/resources/images/thumb/r5.jpg",
	})

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestThumbBgCircles(t *testing.T) {
	capt := getCaptcha()

	capt.SetThumbBgCirclesNum(200)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}