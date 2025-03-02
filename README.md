<div align="center">
<img width="120" style="padding-top: 50px; margin: 0;" src="https://github.com/wenlng/git-assets/blob/master/go-captcha/gocaptcha_logo.svg?raw=true"/>
<h1 style="margin: 0; padding: 0">GoCaptcha</h1>
<p>Behavior Captcha Of Golang</p>
<a href="https://goreportcard.com/report/github.com/wenlng/go-captcha"><img src="https://goreportcard.com/badge/github.com/wenlng/go-captcha"/></a>
<a href="https://godoc.org/github.com/wenlng/go-captcha"><img src="https://godoc.org/github.com/wenlng/go-captcha?status.svg"/></a>
<a href="https://github.com/wenlng/go-captcha/releases"><img src="https://img.shields.io/github/v/release/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/wenlng/go-captcha/blob/v2/LICENSE"><img src="https://img.shields.io/badge/License-Apache2.0-green.svg"/></a>
<a href="https://github.com/wenlng/go-captcha"><img src="https://img.shields.io/github/stars/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/wenlng/go-captcha"><img src="https://img.shields.io/github/last-commit/wenlng/go-captcha.svg"/></a>
</div>

<br/>

> English | [中文](README_zh.md)

<p style="text-align: center"><a href="https://github.com/wenlng/go-captcha">Go Captcha</a> is a behavior CAPTCHA, which implements click mode, slider mode, drag-drop mode and rotation mode.</p>

<p style="text-align: center"> ⭐️ If it helps you, please give a star.</p>

<div align="center"> 
<img src="https://github.com/wenlng/git-assets/blob/master/go-captcha/go-captcha-v2.jpg?raw=true" alt="Poster">
</div>

<br/>
<hr/>
<br/>

## Projects Index

| Name                                                               | Desc                       |
|--------------------------------------------------------------------|----------------------------|
| [document](http://gocaptcha.wencodes.com)                          | GoCaptcha Document         |
| [online demo](http://gocaptcha.wencodes.com/demo/)                 | GoCaptcha Online Demo      |
| [go-captcha-example](https://github.com/wenlng/go-captcha-example) | Golang + Web + APP Example |
| [go-captcha-assets](https://github.com/wenlng/go-captcha-assets)   | Golang Asset File          |
| [go-captcha](https://github.com/wenlng/go-captcha)                 | Golang Captcha             |
| [go-captcha-jslib](https://github.com/wenlng/go-captcha-jslib)     | Javascript Captcha         |
| [go-captcha-vue](https://github.com/wenlng/go-captcha-vue)         | Vue Captcha                |
| [go-captcha-react](https://github.com/wenlng/go-captcha-react)     | React Captcha              |
| [go-captcha-angular](https://github.com/wenlng/go-captcha-angular) | Angular Captcha            |
| [go-captcha-svelte](https://github.com/wenlng/go-captcha-svelte)   | Svelte Captcha             |
| [go-captcha-solid](https://github.com/wenlng/go-captcha-solid)     | Solid Captcha              |
| [go-captcha-uni](https://github.com/wenlng/go-captcha-uni)         | UniApp Captcha             |
| ...                                                                |                            |


<br/>

## Install Module
```shell
$ go get -u github.com/wenlng/go-captcha/v2@latest
```

## Import Module
```go
package main

import "github.com/wenlng/go-captcha/v2"

func main(){
   // ...
}
```

<br />

## 🖖 Click Mode
### Quick Use
```go
package main

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/click"
	"github.com/wenlng/go-captcha/v2/base/codec"
)

var textCapt click.Captcha

func init() {
	builder := click.NewBuilder(
		click.WithRangeLen(option.RangeVal{Min: 4, Max: 6}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 4}),
	)

	// You can use preset material resources：https://github.com/wenlng/go-captcha-assets
	fontN, err := loadFont("../resources/fzshengsksjw_cu.ttf")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage, err := loadPng("../resources/bg.png")
	if err != nil {
		log.Fatalln(err)
	}

	builder.SetResources(
		click.WithChars([]string{
			"1A",
			"5E",
			"3d",
			"0p",
			"78",
			"DL",
			"CB",
			"9M",
			// ...
		}),
		click.WithFonts([]*truetype.Font{
			fontN,
		}),
		click.WithBackgrounds([]image.Image{
			bgImage,
		}),
	)

	textCapt= builder.Make()
}

func loadPng(p string) (image.Image, error) {
	imgBytes, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return codec.DecodeByteToPng(imgBytes)
}

func loadFont(p string) (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}
	return freetype.ParseFont(fontBytes)
}


func main() {
	captData, err := textCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	dotData := captData.GetData()
	if dotData == nil {
		log.Fatalln(">>>>> generate err")
	}

	dots, _ := json.Marshal(dotData)
	fmt.Println(">>>>> ", string(dots))

	var mBase64, tBase64 string
	mBase64, err = captData.GetMasterImage().ToBase64()
	if err != nil {
		fmt.Println(err)
	}
	tBase64, err = captData.GetThumbImage().ToBase64()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(">>>>> ", mBase64)
	fmt.Println(">>>>> ", tBase64)
	
	//err = captData.GetMasterImage().SaveToFile("../resources/master.jpg", option.QualityNone)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = captData.GetThumbImage().SaveToFile("../resources/thumb.png")
	//if err != nil {
	//	fmt.Println(err)
	//}
}
```

### Make Instance
- builder.Make()
- builder.MakeWithShape()

### Configuration Options
> click.NewBuilder(click.WithXxx(), ...) OR builder.SetOptions(click.WithXxx(), ...)

| Options                                      | Desc                                        |
|----------------------------------------------|---------------------------------------------|
| 主图                                           |                                             |
| click.WithImageSize(option.Size)             | default 300x220                             |
| click.WithRangeLen(option.RangeVal)          |                                             |
| click.WithRangeAnglePos([]option.RangeVal)   |                                             |
| click.WithRangeSize(option.RangeVal)         |                                             |
| click.WithRangeColors([]string)              |                                             |
| click.WithDisplayShadow(bool)                |                                             |
| click.WithShadowColor(string)                |                                             |
| click.WithShadowPoint(option.Point)          |                                             |
| click.WithImageAlpha(float32)                |                                             |
| click.WithUseShapeOriginalColor(bool)        |                                             |
| 缩略图                                          |
| click.WithThumbImageSize(option.Size)        | default 150x40                              |
| click.WithRangeVerifyLen(option.RangeVal)    |                                             |
| click.WithDisabledRangeVerifyLen(bool)       |                                             |
| click.WithRangeThumbSize(option.RangeVal)    |                                             |
| click.WithRangeThumbColors([]string)         |                                             |
| click.WithRangeThumbBgColors([]string)       |                                             |
| click.WithIsThumbNonDeformAbility(bool)      |                                             |
| click.WithThumbBgDistort(int)                | option.DistortLevel1 ~ option.DistortLevel5 |
| click.WithThumbBgCirclesNum(int)             |                                             |
| click.WithThumbBgSlimLineNum(int)            |                                             |


### Set Resources
> builder.SetResources(click.WithXxx(), ...)

| Options                                   | Desc |
|-------------------------------------------|------|
| click.WithChars([]string)                 |      |
| click.WithShapes(map[string]image.Image)  |      |
| click.WithFonts([]*truetype.Font)         |      |
| click.WithBackgrounds([]image.Image)      |      |
| click.WithThumbBackgrounds([]image.Image) |      |

### Captcha Data
> captData, err := capt.Generate()

| Method                                   | Desc |
|------------------------------------------|------|
| GetData() map[int]*Dot                   |      |
| GetMasterImage() imagedata.JPEGImageData |      |
| GetThumbImage() imagedata.PNGImageData   |      |

<br />

## 🖖 Slider  Or Drag-Drop Mode
### Quick Use
```go
package main

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"io/ioutil"

	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/slide"
	"github.com/wenlng/go-captcha/v2/base/codec"
)

var slideTileCapt slide.Captcha

func init() {
	builder := slide.NewBuilder()

	// You can use preset material resources：https://github.com/wenlng/go-captcha-assets
	bgImage, err := loadPng("../resources/bg.png")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage1, err := loadPng("../resources/bg1.png")
	if err != nil {
		log.Fatalln(err)
	}

	graphs := getSlideTileGraphArr()

	builder.SetResources(
		slide.WithGraphImages(graphs),
		slide.WithBackgrounds([]image.Image{
			bgImage,
			bgImage1,
		}),
	)

	slideTileCapt = builder.Make()
}

func getSlideTileGraphArr() []*slide.GraphImage {
	tileImage1, err := loadPng("../resources/tile-1.png")
	if err != nil {
		log.Fatalln(err)
	}

	tileShadowImage1, err := loadPng("../resources/tile-shadow-1.png")
	if err != nil {
		log.Fatalln(err)
	}
	tileMaskImage1, err := loadPng("../resources/tile-mask-1.png")
	if err != nil {
		log.Fatalln(err)
	}

	return []*slide.GraphImage{
		{
			OverlayImage: tileImage1,
			ShadowImage:  tileShadowImage1,
			MaskImage:    tileMaskImage1,
		},
	}
}

func main() {
	captData, err := slideTileCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	blockData := captData.GetData()
	if blockData == nil {
		log.Fatalln(">>>>> generate err")
	}

	block, _ := json.Marshal(blockData)
	fmt.Println(">>>>>", string(block))

	var mBase64, tBase64 string
	mBase64, err = captData.GetMasterImage().ToBase64()
	if err != nil {
		fmt.Println(err)
	}
	tBase64, err = captData.GetTileImage().ToBase64()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(">>>>> ", mBase64)
	fmt.Println(">>>>> ", tBase64)
	
	//err = captData.GetMasterImage().SaveToFile("../resources/master.jpg", option.QualityNone)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = captData.GetTileImage().SaveToFile("../resources/thumb.png")
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func loadPng(p string) (image.Image, error) {
	imgBytes, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return codec.DecodeByteToPng(imgBytes)
}
```


### Make Instance
- builder.Make()
- builder.MakeWithRegion() 


### Configuration Options
> slide.NewBuilder(slide.WithXxx(), ...) OR builder.SetOptions(slide.WithXxx(), ...)

| Options                                                        | Desc            |
|----------------------------------------------------------------|-----------------|
| slide.WithImageSize(*option.Size)                              | default 300x220 |
| slide.WithImageAlpha(float32)                                  |                 |
| slide.WithRangeGraphSize(val option.RangeVal)                  |                 |
| slide.WithRangeGraphAnglePos([]option.RangeVal)                |                 |
| slide.WithGenGraphNumber(val int)                              |                 |
| slide.WithEnableGraphVerticalRandom(val bool)                  |                 |
| slide.WithRangeDeadZoneDirections(val []DeadZoneDirectionType) |                 |


### Set Resources
> builder.SetResources(slide.WithXxx(), ...)

| Options                                       | Desc |
|-----------------------------------------------|------|
| slide.WithBackgrounds([]image.Image)          |      |
| slide.WithGraphImages(images []*GraphImage)   |      |

### Captcha Data

> captData, err := capt.Generate()

| Method                                   | Desc |
|------------------------------------------|------|
| GetData() *Block                         |      |
| GetMasterImage() imagedata.JPEGImageData |      |
| GetTileImage() imagedata.PNGImageData    |      |


<br />

## 🖖 Rotation Mode
### Quick Use
```go
package main

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"io/ioutil"

	"github.com/wenlng/go-captcha/v2/rotate"
	"github.com/wenlng/go-captcha/v2/base/codec"
)

var rotateCapt rotate.Captcha

func init() {
	builder := rotate.NewBuilder()

	// You can use preset material resources：https://github.com/wenlng/go-captcha-assets
	bgImage, err := loadPng("../resources/bg.png")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage1, err := loadPng("../resources/bg1.png")
	if err != nil {
		log.Fatalln(err)
	}

	builder.SetResources(
		rotate.WithImages([]image.Image{
			bgImage,
			bgImage1,
		}),
	)

	rotateCapt = builder.Make()
}

func main() {
	captData, err := rotateCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	blockData := captData.GetData()
	if blockData == nil {
		log.Fatalln(">>>>> generate err")
	}

	block, _ := json.Marshal(blockData)
	fmt.Println(">>>>>", string(block))

	var mBase64, tBase64 string
	mBase64, err = captData.GetMasterImage().ToBase64()
	if err != nil {
		fmt.Println(err)
	}
	tBase64, err = captData.GetThumbImage().ToBase64()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(">>>>> ", mBase64)
	fmt.Println(">>>>> ", tBase64)
	
	//err = captData.GetMasterImage().SaveToFile("../resources/master.png")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = captData.GetThumbImage().SaveToFile("../resources/thumb.png")
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func loadPng(p string) (image.Image, error) {
	imgBytes, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return codec.DecodeByteToPng(imgBytes)
}
```


### Make Instance
- builder.Make()


### Configuration Options
> rotate.NewBuilder(rotate.WithXxx(), ...) OR builder.SetOptions(rotate.WithXxx(), ...)

| Options                                          | Desc            |
|--------------------------------------------------|-----------------|
| rotate.WithImageSquareSize(val int)              | default 220x220 |
| rotate.WithRangeAnglePos(vals []option.RangeVal) |                 |
| rotate.WithRangeThumbImageSquareSize(val []int)  |                 |
| rotate.WithThumbImageAlpha(val float32)          |                 |


### Set Resources
> builder.SetResources(rotate.WithXxx(), ...)

| Options                                    | Desc |
|--------------------------------------------|------|
| rotate.WithBackgrounds([]image.Image)      |      |

### Captcha Data
> captData, err := capt.Generate()

| Method                                   | Desc |
|------------------------------------------|------|
| GetData() *Block                         |      |
| GetMasterImage() imagedata.JPEGImageData |      |
| GetTileImage() imagedata.PNGImageData    |      |

<br/>

## Captcha Image Data
### Object Method Of JPEGImageData

| Method                                                     | Desc |
|------------------------------------------------------------|------|
| Get() image.Image                                          |      |
| ToBytes() ([]byte, error)                                  |      |
| ToBytesWithQuality(imageQuality int) ([]byte, error)       |      |
| ToBase64() (string, error)                                 |      |
| ToBase64Data() (string, error)                             |      |
| ToBase64WithQuality(imageQuality int)  (string, error)     |      |
| ToBase64DataWithQuality(imageQuality int) (string, error)  |      |
| SaveToFile(filepath string, quality int) error             |      |



### Object Method Of PNGImageData

| Method                                    | Desc |
|-------------------------------------------|------|
| Get() image.Image                         |      |
| ToBytes() ([]byte, error)                 |      |
| ToBase64() (string, error)                |      |
| ToBase64Data() (string, error)            |      |
| SaveToFile(filepath string) error         |      |


<br/>

## Language Support
- [x] Golang
- [ ] NodeJs
- [ ] Rust
- [ ] Python
- [ ] Java
- [ ] PHP
- [ ] ...

## Install Package
- [x] JavaScript
- [x] Vue 
- [x] React
- [x] Angular
- [x] Svelte
- [x] Solid
- [x] UniApp
- [ ] WX-Applet
- [ ] React Native App
- [ ] Flutter App
- [ ] Android App
- [ ] IOS App
- [ ] ... 

<br/>

## LICENSE
Go Captcha source code is licensed under the Apache Licence, Version 2.0 [http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)
