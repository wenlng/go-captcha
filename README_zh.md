<div align="center">
<img width="120" style="padding-top: 50px; margin: 0;" src="http://47.104.180.148/go-captcha/gocaptcha_logo.svg"/>
<h1 style="margin: 0; padding: 0">Go Captcha</h1>
<p>行为式安全验证码</p>
<a href="https://goreportcard.com/report/github.com/wenlng/go-captcha"><img src="https://goreportcard.com/badge/github.com/wenlng/go-captcha"/></a>
<a href="https://godoc.org/github.com/wenlng/go-captcha"><img src="https://godoc.org/github.com/wenlng/go-captcha?status.svg"/></a>
<a href="https://github.com/wenlng/go-captcha/releases"><img src="https://img.shields.io/github/v/release/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/wenlng/go-captcha/blob/master/LICENSE"><img src="https://img.shields.io/github/license/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/wenlng/go-captcha"><img src="https://img.shields.io/github/stars/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/wenlng/go-captcha"><img src="https://img.shields.io/github/last-commit/wenlng/go-captcha.svg"/></a>
</div>

<br/>

> [English](README.md) | 中文

<p>
<a href="https://github.com/wenlng/go-captcha">go-captcha</a> 是一个简洁易用、交互友好、高安全性的 "行为式验证码" Go库 ，采用 “验证码展示-采集用户行为-验证行为数据” 为流程，用户无需键盘手动输入，极大优化传统验证码用户体验不佳的问题，支持PC端及移动端，支持前端Vue、React...等项目。
</p>

<p> ⭐️ 如果能帮助到你，记得随手给点一个star。</p>

- Github：[https://github.com/wenlng/go-captcha](https://github.com/wenlng/go-captcha)
- 后端Go实例：[https://github.com/wenlng/go-captcha-example](https://github.com/wenlng/go-captcha-example)
- 前端Vue实例：[https://github.com/wenlng/go-captcha-vue](https://github.com/wenlng/go-captcha-vue)
- 前端React实例：[https://github.com/wenlng/go-captcha-react](https://github.com/wenlng/go-captcha-react)
- 在线演示：[http://47.104.180.148:8081/go_captcha_demo](http://47.104.180.148:8081/go_captcha_demo)

<br/>

<div align="center"> 
    <img src="http://47.104.180.148/go-captcha/go-captcha.jpg?v=9" alt="Reward Support">
    <br/>
    <br/>
    <img src="http://47.104.180.148/go-captcha/go-captcha-01.png?v=7" alt="Reward Support">
    <br/>
    <br/>
</div>

### 设置Go代理
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

## 依赖golang官方标准库
```shell
$ go get -u github.com/golang/freetype
$ go get -u golang.org/x/crypto
$ go get -u golang.org/x/image
```

## 安装模块
```shell
$ go get -u github.com/wenlng/go-captcha/captcha
```

## 引入模块
```go
package main

import "github.com/wenlng/go-captcha/captcha"

func main(){
   // ....
}
```

## 快速使用
```go
package main
import (
    "fmt"
    "github.com/wenlng/go-captcha/captcha"
)

func main(){
    // Captcha Single Instances
    capt := captcha.GetCaptcha()

    // 生成验证码
    dots, b64, tb64, key, err := capt.Generate()
    if err != nil {
        panic(err)
        return
    }
    
    // 主图base64
    fmt.Println(len(b64))
    
    // 缩略图base64
    fmt.Println(len(tb64))
    
    // 唯一key
    fmt.Println(key)
    
    // 文本位置验证数据
    fmt.Println(dots)
}

```

## 验证码实例
- 创建实例或者获取单例模式的实例
```go
package main
import (
    "fmt"
    "github.com/wenlng/go-captcha/captcha"
)

func main(){
	// 创建验证码实例
    // capt := captcha.NewCaptcha() 
    
    // 单例模式的验证码实例
    capt := captcha.GetCaptcha()

    // ====================================================
    fmt.Println(capt)

}
```

## 验证码配置
v1.2.3版本后大图默认尺寸为：300×240px，小图默认尺寸为：150×40px。

#### 文本相关配置
默认情况下内置了定制的字体。如果设置了其他中文的文字，则可能需要设置字体文件。
```go
package main
import (
    "fmt"
    "github.com/wenlng/go-captcha/captcha"
)

func main(){
    capt := captcha.GetCaptcha()
    
    // ====================================================
    // Method: SetRangChars (chars []string) error;
    // Desc: 设置验证码文本随机种子
    // ====================================================
    // 字符
    //chars := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    //_ = capt.SetRangChars(strings.Split(chars, ""))
    
    // 双字母
    //chars := []string{"HE","CA","WO","NE","HT","IE","PG","GI","CH","CO","DA"}
    //_ = capt.SetRangChars(chars)

    // 汉字
    chars := []string{"你","好","呀","这","是","点","击","验","证","码","哟"}
    _ = capt.SetRangChars(chars)

    // ====================================================
    fmt.Println(capt)
}
```

#### 设置字体文件
你可以拷贝实例仓库中 "https://github.com/wenlng/go-captcha-example/tree/main/resources" 的图片资源和字体文件到你的项目中使用。
```go
package main
import (
    "fmt"
    "os"
    "github.com/wenlng/go-captcha/captcha"
)

func main(){
    capt := captcha.GetCaptcha()
    
    path, _ := os.Getwd()
    // ====================================================
    // Method: SetFont(fonts []string);
    // Desc: 设置验证码字体，自动仅读取一次并加载到内存中缓存，如需重置可清除缓存
    // ====================================================
    capt.SetFont([]string{
        path + "/__example/resources/fonts/fzshengsksjw_cu.ttf",
    })
}
```

#### 大图相关配置
```go
package main
import (
    "fmt"
    "os"
    "golang.org/x/image/font"
    "github.com/wenlng/go-captcha/captcha"
)

func main(){
    capt := captcha.GetCaptcha()
    
    path, _ := os.Getwd()    
    // ====================================================
    // Method: SetBackground(color []string);
    // Desc: 设置验证码背景图，自动仅读取一次并加载到内存中缓存，如需重置可清除缓存
    // ====================================================
    capt.SetBackground([]string{
        path + "/__example/resources/images/1.jpg",
        path + "/__example/resources/images/2.jpg",
    })

    // ====================================================
    // Method: SetImageSize(size Size);
    // Desc: 设置验证码主图的尺寸
    // ====================================================
    capt.SetImageSize(captcha.Size{300, 300})

    // ====================================================
    // Method: SetImageQuality(val int);
    // Desc: 设置验证码主图清晰度，压缩级别范围 QualityCompressLevel1 - 5，QualityCompressNone无压缩，默认为最低压缩级别
    // ====================================================
    capt.SetImageQuality(captcha.QualityCompressNone)

    // ====================================================
    // Method: SetFontHinting(val font.Hinting);
    // Desc: 设置字体Hinting值 (HintingNone,HintingVertical,HintingFull)
    // ====================================================
    capt.SetFontHinting(font.HintingFull)

    // ====================================================
    // Method: SetTextRangLen(val captcha.RangeVal);
    // Desc: 设置验证码文本显示的总数随机范围
    // ====================================================
    capt.SetTextRangLen(captcha.RangeVal{6, 7})

    // ====================================================
    // Method: SetRangFontSize(val captcha.RangeVal);
    // Desc: 设置验证码文本的随机大小
    // ====================================================
    capt.SetRangFontSize(captcha.RangeVal{32, 42})
    
    // ====================================================
    // Method: SetTextRangFontColors(colors []string);
    // Desc: 设置验证码文本的随机十六进制颜色
    // ====================================================
    capt.SetTextRangFontColors([]string{
        "#1d3f84",
        "#3a6a1e",
    })

    // ====================================================
    // Method: SetImageFontAlpha(val float64);
    // Desc:设置验证码字体的透明度
    // ====================================================
    capt.SetImageFontAlpha(0.5)

    // ====================================================
    // Method: SetTextShadow(val bool);
    // Desc:设置字体阴影
    // ====================================================
    capt.SetTextShadow(true)

    // ====================================================
    // Method: SetTextShadowColor(val string);
    // Desc:设置字体阴影颜色
    // ====================================================
    capt.SetTextShadowColor("#101010")

    // ====================================================
    // Method: SetTextShadowPoint(val captcha.Point);
    // Desc:设置字体阴影偏移位置
    // ====================================================
    capt.SetTextShadowPoint(captcha.Point{1, 1})

    // ====================================================
    // Method: SetTextRangAnglePos(pos []captcha.RangeVal);
    // Desc:设置验证码文本的旋转角度
    // ====================================================
    capt.SetTextRangAnglePos([]captcha.RangeVal{
        {1, 15},
        {15, 30},
        {30, 45},
        {315, 330},
        {330, 345},
        {345, 359},
    })

    // ====================================================
    // Method: SetImageFontDistort(val int);
    // Desc:设置验证码字体的扭曲程度
    // ====================================================
    capt.SetImageFontDistort(captcha.DistortLevel2)

}
```

#### 小图相关配置
```go
package main
import (
    "fmt"
    "os"
    "github.com/wenlng/go-captcha/captcha"
)

func main(){
    capt := captcha.GetCaptcha()
    
    path, _ := os.Getwd()

    // ====================================================
    // Method: SetThumbSize(size Size);
    // Desc: 设置缩略图的尺寸
    // ====================================================
    capt.SetThumbSize(captcha.Size{150, 40})

    // ====================================================
    // Method: SetRangCheckTextLen(val captcha.RangeVal);
    // Desc:设置缩略图校验文本的随机长度范围
    // ====================================================
    capt.SetRangCheckTextLen(captcha.RangeVal{2, 4})

    // ====================================================
    // Method: SetRangCheckFontSize(val captcha.RangeVal);
    // Desc:设置缩略图校验文本的随机大小
    // ====================================================
    capt.SetRangCheckFontSize(captcha.RangeVal{24, 30})
    
    // ====================================================
    // Method: SetThumbTextRangFontColors(colors []string);
    // Desc: 设置缩略图文本的随机十六进制颜色
    // ====================================================
    capt.SetThumbTextRangFontColors([]string{
        "#1d3f84",
        "#3a6a1e",
    })
  
    // ====================================================
    // Method: SetThumbBgColors(colors []string);
    // Desc: 设置缩略图的背景随机十六进制颜色
    // ====================================================
    capt.SetThumbBgColors([]string{
        "#1d3f84",
        "#3a6a1e",
    })

    // ====================================================
    // Method: SetThumbBackground(colors []string);
    // Desc:设置缩略图的随机图像背景，自动仅读取一次并加载到内存中缓存，如需重置可清除缓存
    // ====================================================
    capt.SetThumbBackground([]string{
        path + "/__example/resources/images/r1.jpg",
        path + "/__example/resources/images/r2.jpg",
    })

    // ====================================================
    // Method: SetThumbBgDistort(val int);
    // Desc:设置缩略图背景的扭曲程度
    // ====================================================
    capt.SetThumbBgDistort(captcha.DistortLevel2)

    // ====================================================
    // Method: SetThumbFontDistort(val int);
    // Desc:设置缩略图字体的扭曲程度
    // ====================================================
    capt.SetThumbFontDistort(captcha.DistortLevel2)

    // ====================================================
    // Method: SetThumbBgCirclesNum(val int);
    // Desc:设置缩略图背景的圈点数
    // ====================================================
    capt.SetThumbBgCirclesNum(20)

    // ====================================================
    // Method: SetThumbBgSlimLineNum(val int);
    // Desc:设置缩略图背景的线条数
    // ====================================================
    capt.SetThumbBgSlimLineNum(3)
    
}
```

## 其它
```go
package main
import (
    "fmt"
    "os"
    "github.com/wenlng/go-captcha/captcha"
)

func main(){    
    path, _ := os.Getwd()    
    // ====================================================
    // Method: ClearAssetCacheWithPath(paths []string) bool;
    // Desc: 根据路径清除应资源的缓存
    // ====================================================
    captcha.ClearAssetCacheWithPaths([]string{
    	path + "/__example/resources/images/1.jpg",
    	path + "/__example/resources/fonts/fonts.ttf",
    }) 

    // ====================================================
    // Method: captcha.CheckPointDist(sx, sy, dx, dy, width, height int64) bool;
    // Desc: 校验点的位置
    // ====================================================
    captcha.CheckPointDist(12, 22, 10, 15, 30, 30)    

    // ====================================================
    // Method: captcha.CheckPointDistWithPadding(sx, sy, dx, dy, width, height, padding int64) bool;
    // Desc: 校验点的位置,在原有的区域上添加额外边距进行扩张计算区域,不推荐设置过大的padding
    // 例如：文本的宽和高为30，校验范围x为10-40，y为15-45，此时扩充5像素后校验范围宽和高为40，则校验范围x为5-45，位置y为10-50
    // ====================================================
    captcha.CheckPointDistWithPadding(12, 22, 10, 15, 30, 30, 5) 
}
```

## 生成验证码数据
```go
package main
import (
    "fmt"
    "os"
    "github.com/wenlng/go-captcha/captcha"
)

func main(){
    capt := captcha.GetCaptcha()
   
    // generate ...
    // ====================================================
    // dots:  文本字符的位置信息
    //  - {"0":{"Index":0,"Dx":198,"Dy":77,"Size":41,"Width":54,"Height":41,"Text":"SH","Angle":6,"Color":"#885500"} ...}
    // imageBase64:  Verify the clicked image
    // thumbImageBase64: Thumb displayed
    // key: Only Key
    // ====================================================
    dots, imageBase64, thumbImageBase64, key, err := capt.Generate()
    if err != nil {
        panic(err)
        return
    }
    fmt.Println(len(imageBase64))
    fmt.Println(len(thumbImageBase64))
    fmt.Println(key)
    fmt.Println(dots)
}
```

## 前端在数据请求或提交验证数据时的格式： 
```
// Example: 获取验证码数据
    Respose Data = {
        "code": 0,
        "image_base64": "...",
        "thumb_base64": "...",
        "captcha_key": "xxxxxx",
    }     

// Example: 提交校验数据 
    Request Data = {
        dots: "x1,y1,x2,y2,...."
        key: "xxxxxx"
    }
```
<br/>

> 请作者喝咖啡：[http://witkeycode.com/sponsor](http://witkeycode.com/sponsor)

<br/>

## LICENSE
    MIT
