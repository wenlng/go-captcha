<div align="center">
<img width="120" style="padding-top: 50px; margin: 0;" src="https://github.com/wenlng/git-assets/blob/master/go-captcha/gocaptcha_logo.svg?raw=true"/>
<h1 style="margin: 0; padding: 0">GoCaptcha</h1>
<p>Golang 行为验证码</p>
<a href="https://goreportcard.com/report/github.com/wenlng/go-captcha"><img src="https://goreportcard.com/badge/github.com/wenlng/go-captcha"/></a>
<a href="https://godoc.org/github.com/wenlng/go-captcha"><img src="https://godoc.org/github.com/wenlng/go-captcha?status.svg"/></a>
<a href="https://github.com/wenlng/go-captcha/releases"><img src="https://img.shields.io/github/v/release/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/wenlng/go-captcha/blob/v2/LICENSE"><img src="https://img.shields.io/badge/License-Apache2.0-green.svg"/></a>
<a href="https://github.com/wenlng/go-captcha"><img src="https://img.shields.io/github/stars/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/wenlng/go-captcha"><img src="https://img.shields.io/github/last-commit/wenlng/go-captcha.svg"/></a>
</div>

<br/>

> [English](README.md) | 中文 

<p style="text-align: center">
<a style="font-weight: bold" href="https://github.com/wenlng/go-captcha">Go Captcha</a> 是行为式验证码，支持文本/图形点选、滑动/拖拽、旋转等验证模式。
</p>

<p style="text-align: center"> ⭐️ 如果能帮助到你，请随手给点一个star</p>
<p style="text-align: center">QQ交流群：178498936</p>


<br/>

<div align="center">
<img src="https://github.com/wenlng/git-assets/blob/master/go-captcha/go-captcha-v2.jpg?raw=true" alt="Poster">
</div>

<br/>
<hr/>
<br/>

## 项目索引

| 名称                                                                 | 说明                                  |
|--------------------------------------------------------------------|-------------------------------------|
| [document](http://gocaptcha.wencodes.com)                          | GoCaptcha 文档                        |
| [online demo](http://gocaptcha.wencodes.com/demo/)                 | GoCaptcha 在线演示                      |
| [go-captcha-example](https://github.com/wenlng/go-captcha-example) | Golang + 前端 + APP实例                 |
| [go-captcha-assets](https://github.com/wenlng/go-captcha-assets)   | Golang 内嵌素材资源                       |
| [go-captcha](https://github.com/wenlng/go-captcha)                 | Golang 验证码                          |
| [go-captcha-jslib](https://github.com/wenlng/go-captcha-jslib)     | Javascript 验证码                      |
| [go-captcha-vue](https://github.com/wenlng/go-captcha-vue)         | Vue 验证码                             |
| [go-captcha-react](https://github.com/wenlng/go-captcha-react)     | React 验证码                           |
| [go-captcha-angular](https://github.com/wenlng/go-captcha-angular) | Angular 验证码                         |
| [go-captcha-svelte](https://github.com/wenlng/go-captcha-svelte)   | Svelte 验证码                          |
| [go-captcha-solid](https://github.com/wenlng/go-captcha-solid)     | Solid 验证码                           |
| [go-captcha-uni](https://github.com/wenlng/go-captcha-uni)         | UniApp 验证码，兼容 Android、IOS、小程序、快应用等  |
| ...                                                                |                                     |

<br/>

## 设置Go代理
- Window
```shell
$ set GO111MODULE=on
$ set GOPROXY=https://goproxy.io,direct

### The Golang 1.13+ can be executed directly
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.io,direct
```
- Linux or Mac
```shell
$ export GO111MODULE=on
$ export GOPROXY=https://goproxy.io,direct

### or
$ echo "export GO111MODULE=on" >> ~/.profile
$ echo "export GOPROXY=https://goproxy.cn,direct" >> ~/.profile
$ source ~/.profile
```

## 安装模块
```shell
$ go get -u github.com/wenlng/go-captcha/v2@latest
```

## 引入模块
```go
package main

import "github.com/wenlng/go-captcha/v2"

func main(){
   // ....
}
```

<br />

## 🖖点选式
### 快速使用
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

	// 可以使用预置的素材资源：https://github.com/wenlng/go-captcha-assets
	fontN, err := loadFont("../resources/fzshengsksjw_cu.ttf")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage, err := loadPng("../resources/bg.png")
	if err != nil {
		log.Fatalln(err)
	}

	builder.SetResources(
		click.WithChars([]string{"这", "是", "随", "机", "的", "文", "本", "种", "子", "呀"}),
		click.WithFonts([]*truetype.Font{
			fontN,
		}),
		click.WithBackgrounds([]image.Image{
			bgImage,
		}),
	)

	textCapt = builder.Make()
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

### 创建实例
- builder.Make()  中文文本、字母数字混合点选
- builder.MakeWithShape()  图形点选

### 配置选项
> click.NewBuilder(click.WithXxx(), ...) 或 builder.SetOptions(click.WithXxx(), ...)

| Options                                    | Desc                                                  |
|--------------------------------------------|-------------------------------------------------------|
| 主图                                         |
| click.WithImageSize(option.Size)           | 设置主图尺寸，默认 300x220                                     |
| click.WithRangeLen(option.RangeVal)        | 设置随机内容长度范围                                            |
| click.WithRangeAnglePos([]option.RangeVal) | 设置随机角度范围                                              |
| click.WithRangeSize(option.RangeVal)       | 设置随机内容大小范围                                            |
| click.WithRangeColors([]string)            | 设置随机颜色                                                |
| click.WithDisplayShadow(bool)              | 设置是否显示阴影                                              |
| click.WithShadowColor(string)              | 设置阴影颜色                                                |
| click.WithShadowPoint(option.Point)        | 设置阴影偏移位置                                              |
| click.WithImageAlpha(float32)              | 设置主图透明度                                               |
| click.WithUseShapeOriginalColor(bool)      | 设置是否使用图形原始颜色，"图形点选"有效                                 |
| 缩略图                                        |
| click.WithThumbImageSize(option.Size)      | 设置缩略尺寸，默认 150x40                                      |
| click.WithRangeVerifyLen(option.RangeVal)  | 设置校验内容的随机长度范围                                         |
| click.WithDisabledRangeVerifyLen(bool)     | 禁用校验内容的随机长度，与主图内容的长度保持一致                              |
| click.WithRangeThumbSize(option.RangeVal)  | 设置随机缩略内容随机大小范围                                        |
| click.WithRangeThumbColors([]string)       | 设置缩略随机颜色范围                                            |
| click.WithRangeThumbBgColors([]string)     | 设置缩略随机背景颜色范围                                          |
| click.WithIsThumbNonDeformAbility(bool)    | 设置缩略图内容不变形，不受背景影响                                     |
| click.WithThumbBgDistort(int)              | 设置缩略图背景扭曲 option.DistortLevel1 至 option.DistortLevel5 |
| click.WithThumbBgCirclesNum(int)           | 设置缩略图绘制小圆点数量                                          |
| click.WithThumbBgSlimLineNum(int)          | 设置缩略图绘制线条数量                                           |


### 设置资源
> builder.SetResources(click.WithXxx(), ...)

| Options                                   | Desc      |
|-------------------------------------------|-----------|
| click.WithChars([]string)                 | 设置文本种子    |
| click.WithShapes(map[string]image.Image)  | 设置图形种子    |
| click.WithFonts([]*truetype.Font)         | 设置字体      |
| click.WithBackgrounds([]image.Image)      | 设置主图背景    |
| click.WithThumbBackgrounds([]image.Image) | 设置缩略图背景   |


### 验证码数据
> captData, err := capt.Generate()

| Method                                   | Desc      |
|------------------------------------------|-----------|
| GetData() map[int]*Dot                   | 获取当前校验的信息 |
| GetMasterImage() imagedata.JPEGImageData | 获取主图      |
| GetThumbImage() imagedata.PNGImageData   | 获取缩略图     |


<br />

## 🖖 滑动/拖拽式
### 快速使用
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

	// 可以使用预置的素材资源：https://github.com/wenlng/go-captcha-assets
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


### 创建实例
- builder.Make() 滑动式
- builder.MakeWithRegion()  区域内拖拽式


### 配置选项
> slide.NewBuilder(slide.WithXxx(), ...) 或 builder.SetOptions(slide.WithXxx(), ...)

| Options                                                        | Desc              |
|----------------------------------------------------------------|-------------------|
| slide.WithImageSize(*option.Size)                              | 设置主图尺寸，默认 300x220 |
| slide.WithImageAlpha(float32)                                  | 设置主图透明度           |
| slide.WithRangeGraphSize(val option.RangeVal)                  | 设置图形随机尺寸范围        |
| slide.WithRangeGraphAnglePos([]option.RangeVal)                | 设置图形随机角度范围        |
| slide.WithGenGraphNumber(val int)                              | 设置图形个数            |
| slide.WithEnableGraphVerticalRandom(val bool)                  | 设置图形水平方向是否随机排序    |
| slide.WithRangeDeadZoneDirections(val []DeadZoneDirectionType) | 设置贴图盲区            |


### 设置资源
> builder.SetResources(slide.WithXxx(), ...)

| Options                                       | Desc     |
|-----------------------------------------------|----------|
| slide.WithBackgrounds([]image.Image)          | 设置主图背景   |
| slide.WithGraphImages(images []*GraphImage)   | 设置贴图的图形  |


### 验证码数据
> captData, err := capt.Generate()

| Method                                   | Desc        |
|------------------------------------------|-------------|
| GetData() *Block                         | 获取当前校验的信息   |
| GetMasterImage() imagedata.JPEGImageData | 获取主图        |
| GetTileImage() imagedata.PNGImageData    | 获取缩略图       |

<br />

## 🖖  旋转式
### 快速使用
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

	// 可以使用预置的素材资源：https://github.com/wenlng/go-captcha-assets
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


### 创建实例
- builder.Make() 旋转式


### 配置选项
> rotate.NewBuilder(rotate.WithXxx(), ...) 或 builder.SetOptions(rotate.WithXxx(), ...)

| Options                                          | Desc              |
|--------------------------------------------------|-------------------|
| rotate.WithImageSquareSize(val int)              | 设置主图大小，默认 220x220 |
| rotate.WithRangeAnglePos(vals []option.RangeVal) | 设置校验随机角度范围        |
| rotate.WithRangeThumbImageSquareSize(val []int)  | 设置缩略图大小           |
| rotate.WithThumbImageAlpha(val float32)          | 设置缩略图透明度          |


### 设置资源
> builder.SetResources(rotate.WithXxx(), ...)

| Options                                    | Desc       |
|--------------------------------------------|------------|
| rotate.WithBackgrounds([]image.Image)      | 设置主图图片     |


### 验证码数据
> captData, err := capt.Generate()

| Method                                   | Desc        |
|------------------------------------------|-------------|
| GetData() *Block                         | 获取当前校验的信息   |
| GetMasterImage() imagedata.JPEGImageData | 获取主图        |
| GetTileImage() imagedata.PNGImageData    | 获取缩略图       |

<br/>

## 验证码图像

### JPEGImageData

| Method                                                    | Desc                                                                                                                       |
|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------|
| Get() image.Image                                         | <span style='padding: 0 10px'></span>获取原图像                                                                                 |
| ToBytes() ([]byte, error)                                 | <span style='padding: 0 10px'></span>转为字节数组                                                                                |
| ToBytesWithQuality(imageQuality int) ([]byte, error)      | 指定清晰度转为字节数组                                                                                                                |
| ToBase64() (string, error)                                | <span style='padding: 0 10px'></span>转为 Base64 字符串，带 <span style='color:#ed4630;'>"data:image/jpeg;base64,"</span> 前缀      |
| ToBase64Data() (string, error)                            | <span style='padding: 0 10px'></span>转为 Base64 字符串                                                                         |
| ToBase64WithQuality(imageQuality int)  (string, error)    | <span style='padding: 0 10px'></span>指定清晰度转为 Base64 字符串，带 <span style='color:#ed4630;'>"data:image/jpeg;base64,"</span> 前缀 |
| ToBase64DataWithQuality(imageQuality int) (string, error) | <span style='padding: 0 10px'></span>指定清晰度转为 Base64 字符串                                                                    |
| SaveToFile(filepath string, quality int) error            | <span style='padding: 0 10px'></span>保存 JPEG 到文件                                                                           |


### PNGImageData

| Method                                    | Desc                                                                                                                 |
|-------------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| Get() image.Image                         | <span style='padding: 0 10px'></span>获取原图像                                                                           |
| ToBytes() ([]byte, error)                 | <span style='padding: 0 10px'></span>转为字节数组                                                                          |
| ToBase64() (string, error)                | <span style='padding: 0 10px'></span>转为 Base64 字符串，带 <span style='color:#ed4630;'>"data:image/png;base64,"</span> 前缀 |
| ToBase64Data() (string, error)            | <span style='padding: 0 10px'></span>转为 Base64 字符串                                                                   |
| SaveToFile(filepath string) error         | <span style='padding: 0 10px'></span>保存 到文件                                                                          |


<br/>

## 验证模块
- [x] 文字点选
- [x] 图形点选
- [x] 滑动
- [x] 拖拽
- [x] 旋转
- [ ] 抛物线
- [ ] 看图选择
- [ ] 物品辨认
- [ ] ...

## 扩展&增强
- [x] 基本验证
- [ ] Docker部署镜像
- [ ] 行为检测增强
- [ ] 其他因素增强
- [ ] 多任务生成模式
- [ ] ...

## 语言支持
- [x] Golang
- [ ] NodeJs
- [ ] Rust
- [ ] Python
- [ ] Java
- [ ] PHP
- [ ] ...

## Web
- [x] JavaScript
- [x] Vue
- [x] React
- [x] Angular
- [x] Svelte
- [x] Solid
- [ ] ...

## App
- [x] UniApp
- [ ] WX-Applet
- [ ] React Native App
- [ ] Flutter App
- [ ] Android App
- [ ] IOS App
- ...

## Deployment
- [ ] Docker Image
- ...

<br/>

## 👍赞助一下

<p>如果觉得项目有帮助，可以请作者喝杯咖啡 🍹</p>
<div>
<a href="http://witkeycode.com/sponsor" target="_blank"><img src="http://47.104.180.148/payment-code/wxpay.png" alt="Buy Me A Coffee" style="width: 217px !important;" ></a>
<a href="http://witkeycode.com/sponsor" target="_blank"><img src="http://47.104.180.148/payment-code/alipay.png" alt="Buy Me A Coffee" style="width: 217px !important;" ></a>
</div>

<br/>

## LICENSE
Go Captcha source code is licensed under the Apache Licence, Version 2.0 [http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)

