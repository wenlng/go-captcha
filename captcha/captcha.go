/**
 * @Author Awen
 * @Description Captcha
 * @Date 2021/7/18
 * @Email wengaolng@gmail.com
 **/

package captcha

import (
	"encoding/json"
	"fmt"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"io/ioutil"
	"math"
	"math/rand"
	"strings"
	"time"
)

// CharDot is a type
/**
 * @Description: 图片点数据
 */
type CharDot struct {
	// 顺序索引
	Index int
	// x,y位置
	Dx int
	Dy int
	// 字体大小
	Size int
	// 字体宽
	Width int
	// 字体高
	Height int
	// 字符文本
	Text string
	// 字体角度
	Angle int
	// 颜色
	Color string
	// 颜色2
	Color2 string
}

// Captcha is a type
/**
 * @Description: 点选验证码
 */
type Captcha struct {
	// 字符集合，用于随机字符串
	chars *[]string
	// 难度码配置
	config *Config
	// 验证画图
	captchaDraw *Draw
}

var clickCaptcha *Captcha

// NewCaptcha is a function
/**
 * @Description: 创建点选验证码
 * @return *Captcha
 */
func NewCaptcha() *Captcha {
	return &Captcha{
		chars:       GetCaptchaDefaultChars(),
		config:      GetCaptchaDefaultConfig(),
		captchaDraw: &Draw{},
	}
}

// GetCaptcha is a function
/**
 * @Description: 获取点选验证码
 * @return *Captcha
 */
func GetCaptcha() *Captcha {
	if clickCaptcha == nil {
		clickCaptcha = NewCaptcha()
	}

	return clickCaptcha
}

// SetRangChars is a function
/**
 * @Description: 设置随机字符串，每个单词不能超出2个字符，超出会影响位置的验证
 * @receiver cc
 * @param chars
 * @return error
 */
func (cc *Captcha) SetRangChars(chars []string) error {
	var err error
	if len(chars) > 0 {
		for _, char := range chars {
			if IsChineseChar(char) {
				if LenChineseChar(char) > 1 {
					err = fmt.Errorf("Captcha SetRangChars Error: The chinese char [%s] must be equal to 1", char)
					break
				}
			} else if LenChineseChar(char) > 2 {
				err = fmt.Errorf("Captcha SetRangChars Error: The char [%s] must be less than or equal to 2", char)
				break
			}
		}
	}

	if err != nil {
		return err
	}

	cc.chars = &chars
	return nil
}

// =============================================
// Captcha Set Config
// =============================================

// SetBackground is a function
/**
 * @Description: 设置随机背景图片
 * @receiver cc
 * @param images
 * @param args	true|false 是否强制刷新缓存
 */
func (cc *Captcha) SetBackground(images []string, args ...bool) {
	for _, path := range images {
		if has, err := PathExists(path); !has || err != nil {
			panic(fmt.Errorf("CaptchaConfig Error: The [%s] file does not exist", path))
		}

		hasCache := hasAssetCache(path)
		if !hasCache || (hasCache && len(args) > 0 && args[0]) {
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			setAssetCache(path, bytes, len(args) > 0 && args[0])
		}
	}

	cc.config.rangBackground = images
}

// SetFont is a function
/**
 * @Description: 设置随机字体
 * @receiver cc
 * @param fonts
 * @param args	true|false 是否强制刷新缓存
 */
func (cc *Captcha) SetFont(fonts []string, args ...bool) {
	for _, path := range fonts {
		if has, err := PathExists(path); !has || err != nil {
			panic(fmt.Errorf("CaptchaConfig Error: The [%s] file does not exist", path))
		}
		hasCache := hasAssetCache(path)
		if !hasCache || (hasCache && len(args) > 0 && args[0]) {
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			setAssetCache(path, bytes, len(args) > 0 && args[0])
		}
	}

	cc.config.rangFont = fonts
}

// SetImageSize is a function
/**
 * @Description: 设置验证码尺寸
 * @receiver cc
 * @param size
 */
func (cc *Captcha) SetImageSize(size Size) {
	cc.config.imageSize = size
}

// SetImageQuality is a function
/**
 * @Description: 设置验证码清晰度1-100为压缩图，999为原图
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetImageQuality(val int) {
	if val == QualityCompressNone || (val <= QualityCompressLevel1 && val >= QualityCompressLevel5) {
		cc.config.imageQuality = val
	}
}

// SetThumbSize is a function
/**
 * @Description: 设置缩略图尺寸
 * @receiver cc
 * @param size
 */
func (cc *Captcha) SetThumbSize(size Size) {
	cc.config.thumbnailSize = size
}

// SetRangFontSize is a function
/**
 * @Description: 设置随机字体大小
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetRangFontSize(val RangeVal) {
	cc.config.rangFontSize = val
}

// SetTextRangLen is a function
/**
 * @Description: 设置字符随机长度范围
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetTextRangLen(val RangeVal) {
	cc.config.rangTextLen = val
}

// SetTextRangFontColors is a function
/**
 * @Description: 设置文本随机颜色
 * @receiver cc
 * @param colors
 */
func (cc *Captcha) SetTextRangFontColors(colors []string) {
	cc.config.rangFontColors = colors
}


// SetThumbTextRangFontColors is a function
/**
 * @Description: 设置缩略图文本随机颜色
 * @receiver cc
 * @param colors
 */
func (cc *Captcha) SetThumbTextRangFontColors(colors []string) {
	cc.config.rangThumbFontColors = colors
}

// SetFontDPI is a function
/**
 * @Description: 设置分辨率，72为标准
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetFontDPI(val int) {
	cc.config.fontDPI = val
}

// SetFontDPI is a function
/**
 * @Description: 设置分辨率，72为标准
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetFontHinting(val font.Hinting) {
	cc.config.fontHinting = val
}

// SetImageFontAlpha is a function
/**
 * @Description: 设置验证码文本透明度
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetImageFontAlpha(val float64) {
	cc.config.imageFontAlpha = val
}

// SetTextShadow is a function
/**
 * @Description: 设置验证码文本阴影显示
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetTextShadow(val bool) {
	cc.config.showTextShadow = val
}

// SetTextShadowColor is a function
/**
 * @Description: 设置验证码文本阴影颜色
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetTextShadowColor(val string) {
	cc.config.textShadowColor = val
}

// SetTextShadowPoint is a function
/**
 * @Description: 设置验证码文本阴影位置
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetTextShadowPoint(val Point) {
	cc.config.textShadowPoint = val
}

// SetImageFontDistort is a function
/**
 * @Description: 设置验证码文本扭曲程度
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetImageFontDistort(val int) {
	if val >= DistortNone || val <= DistortLevel5 {
		cc.config.imageFontDistort = val
	}
}

// SetTextRangAnglePos is a function
/**
 * @Description: 设置文本角度位置
 * @receiver cc
 * @param pos
 */
func (cc *Captcha) SetTextRangAnglePos(pos []RangeVal) {
	cc.config.rangTexAnglePos = pos
}

// SetRangCheckTextLen is a function
/**
 * @Description: 设置随机验证文本长度
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetRangCheckTextLen(val RangeVal) {
	// 检测验证文本范围最大值是否小于随机字符串的最小范围
	if cc.config.rangCheckTextLen.Max > cc.config.rangTextLen.Min {
		panic(fmt.Errorf("CaptchaConfig Error: RangCheckTextLen.max must be less than or equal to RangTextLen.min"))
	}

	cc.config.rangCheckTextLen = val
}

// SetRangCheckFontSize is a function
/**
 * @Description: 设置随机验证文本大小
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetRangCheckFontSize(val RangeVal) {
	cc.config.rangCheckFontSize = val
}

// SetThumbBgColors is a function
/**
 * @Description: 设置缩略图随机背景颜色
 * @receiver cc
 * @param colors
 */
func (cc *Captcha) SetThumbBgColors(colors []string) {
	cc.config.rangThumbBgColors = colors
}

// SetThumbBackground is a function
/**
 * @Description: 设置随机背景图片
 * @receiver cc
 * @param images
 * @param args	true|false 是否强制刷新缓存
 */
func (cc *Captcha) SetThumbBackground(images []string, args ...bool) {
	for _, path := range images {
		if has, err := PathExists(path); !has || err != nil {
			panic(fmt.Errorf("CaptchaConfig Error: The [%s] file does not exist", path))
		}
		hasCache := hasAssetCache(path)
		if !hasCache || (hasCache && len(args) > 0 && args[0]) {
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			setAssetCache(path, bytes, len(args) > 0 && args[0])
		}
	}

	cc.config.rangThumbBackground = images
}

// SetThumbBgDistort is a function
/**
 * @Description: 设置缩略图扭曲程度
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetThumbBgDistort(val int) {
	if val >= DistortNone || val <= DistortLevel5 {
		cc.config.thumbBgDistort = val
	}
}

// SetThumbFontDistort is a function
/**
 * @Description: 设置缩略图字体扭曲程度
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetThumbFontDistort(val int) {
	if val >= DistortNone || val <= DistortLevel5 {
		cc.config.thumbFontDistort = val
	}
}

// SetThumbBgCirclesNum is a function
/**
 * @Description: 设置缩略图小圆点数量
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetThumbBgCirclesNum(val int) {
	cc.config.thumbBgCirclesNum = val
}

// SetThumbBgSlimLineNum is a function
/**
 * @Description: 缩略图线条数量
 * @receiver cc
 * @param val
 */
func (cc *Captcha) SetThumbBgSlimLineNum(val int) {
	cc.config.thumbBgSlimLineNum = val
}

// =============================================
// Captcha Call API
// =============================================
/**
 * @Description: 根据路径清除资源缓存
 * @param paths
 * @return bool
 */
func (cc *Captcha) ClearAssetCacheWithPaths(paths []string) bool {
	return clearAssetCache(paths)
}

// =============================================

/**
 * @Description: 检测配置是否完整和合法，字体和图片背景必须设置
 * @receiver cc
 * @return error
 */
func (cc *Captcha) checkConfig() error {
	// 验证颜色总和是否超出255个
	if len(cc.config.rangFontColors) >= 255 {
		return fmt.Errorf("CaptchaConfig Error: len(rangFontColors) must be less than or equal to 255")
	}

	// 验证颜色总和是否超出255个
	if len(cc.config.rangThumbFontColors) + len(cc.config.rangThumbBgColors) >= 255 {
		return fmt.Errorf("CaptchaConfig Error: len(rangThumbBgColors + RangThumbBgColors) must be less than or equal to 255")
	}

	return nil
}

// Generate is a function
/**
 * @Description: 			根据设置的尺寸生成验证码图片
 * @return CaptchaCharDot	位置信息
 * @return string			主图Base64
 * @return string			验证码KEY
 * @return string			缩略图Base64
 * @return error
 */
func (cc *Captcha) Generate() (map[int]CharDot, string, string, string, error) {
	dots, ib64, tb64, key, err := cc.GenerateWithSize(cc.config.imageSize, cc.config.thumbnailSize)
	return dots, ib64, tb64, key, err
}

// GenerateWithSize is a function
/**
 * @Description: 			生成验证码图片
 * @param imageSize			主图尺寸
 * @param thumbnailSize		缩略图尺寸
 * @return CaptchaCharDot	位置信息
 * @return string			主图Base64
 * @return string			验证码KEY
 * @return string			缩略图Base64
 * @return error
 */
func (cc *Captcha) GenerateWithSize(imageSize Size, thumbnailSize Size) (map[int]CharDot, string, string, string, error) {
	err := cc.checkConfig()
	length := RandInt(cc.config.rangTextLen.Min, cc.config.rangTextLen.Max)
	chars := cc.genRandChar(length)
	if chars == "" {
		return nil, "", "", "", fmt.Errorf("genCaptchaImage Error: No character generation")
	}

	var allDots, thumbDots, checkDots map[int]CharDot
	var imageBase64, tImageBase64 string
	var checkChars string

	allDots = cc.genDots(imageSize, cc.config.rangFontSize, chars, 10)
	// checkChars = "A:B:C"
	checkDots, checkChars = cc.rangeCheckDots(allDots)
	thumbDots = cc.genDots(thumbnailSize, cc.config.rangCheckFontSize, checkChars, 0)
	if err != nil {
		return nil, "", "", "", err
	}
	imageBase64, err = cc.genCaptchaImage(imageSize, allDots)
	if err != nil {
		return nil, "", "", "", err
	}
	tImageBase64, err = cc.genCaptchaThumbImage(thumbnailSize, thumbDots)
	if err != nil {
		return nil, "", "", "", err
	}

	str, _ := json.Marshal(checkDots)
	key, _ := cc.genCaptchaKey(string(str))
	return checkDots, imageBase64, tImageBase64, key, err
}

// EncodeB64string is a function
/**
 * @Description: base64编码
 * @receiver cc
 * @param img
 * @return string
 */
func (cc *Captcha) EncodeB64stringWithJpeg(img image.Image) string {
	if cc.config.imageQuality <= QualityCompressLevel1 && cc.config.imageQuality >= QualityCompressLevel1{
		return EncodeB64stringWithJpeg(img, cc.config.imageQuality)
	}
	return EncodeB64stringWithPng(img)
}


// EncodeB64string is a function
/**
 * @Description: base64编码
 * @receiver cc
 * @param img
 * @return string
 */
func (cc *Captcha) EncodeB64stringWithPng(img image.Image) string {
	return EncodeB64stringWithPng(img)
}

/**
 * @Description: 生成唯一KEY
 * @receiver cc
 * @param str
 * @return string
 * @return error
 */
func (cc *Captcha) genCaptchaKey(str string) (string, error) {
	t := GenUniqueId()
	keyStr := Md5ToString(str + t)
	return keyStr, nil
}

/**
 * @Description: 生成字符在图片上的点
 * @receiver cc
 * @param imageSize
 * @param fontSize
 * @param chars
 * @param padding
 * @return []*CaptchaCharDot
 */
func (cc *Captcha) genDots(imageSize Size, fontSize RangeVal, chars string, padding int) map[int]CharDot {
	dots := make(map[int]CharDot) // 各个文字点位置
	width := imageSize.Width
	height := imageSize.Height
	if padding > 0 {
		width -= padding
		height -= padding
	}

	strs := strings.Split(chars, ":")
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		// 随机角度
		randAngle := cc.getRandAngle()
		// 随机颜色
		randColor := cc.getRandColor(cc.config.rangFontColors)
		randColor2 := cc.getRandColor(cc.config.rangThumbFontColors)

		// 随机文字大小
		randFontSize := RandInt(fontSize.Min, fontSize.Max)
		fontHeight := randFontSize
		fontWidth := randFontSize

		if LenChineseChar(str) > 1 {
			fontWidth = randFontSize * LenChineseChar(str)

			if randAngle > 0 {
				surplus := fontWidth - randFontSize
				ra := randAngle % 90
				pr := float64(surplus) / 90
				h := math.Max(float64(ra)*pr, 1)
				fontHeight = fontHeight + int(h)
			}
		}

		//_w := (width - randFontSize) / len(str)
		_w := width / len(strs)
		rd := math.Abs(float64(_w) - float64(fontWidth))
		x := (i * _w) + RandInt(0, int(math.Max(rd, 1)))
		x = int(math.Min(math.Max(float64(x), 10), float64(width - 10 - (padding * 2))))
		y := RandInt(10, height + fontHeight)
		y = int(math.Min(math.Max(float64(y), float64(fontHeight + 10)), float64(height + (fontHeight / 2) - (padding * 2))))
		text := fmt.Sprintf("%s", str)

		dot := CharDot{i, x, y, randFontSize, fontWidth, fontHeight, text, randAngle, randColor, randColor2}
		dots[i] = dot
	}

	return dots
}

/**
 * @Description: 随机检测点
 * @receiver cc
 * @param dots
 * @return map[int]CaptchaCharDot
 * @return string
 */
func (cc *Captcha) rangeCheckDots(dots map[int]CharDot) (map[int]CharDot, string) {
	rand.Seed(time.Now().UnixNano())
	rs := rand.Perm(len(dots))
	chkDots := make(map[int]CharDot)
	count := RandInt(cc.config.rangCheckTextLen.Min, cc.config.rangCheckTextLen.Max)
	var chars []string
	for i, value := range rs {
		if i >= count {
			continue
		}
		dot := dots[value]
		dot.Index = i
		chkDots[i] = dot
		chars = append(chars, chkDots[i].Text)
	}
	return chkDots, strings.Join(chars, ":")
}

/**
 * @Description: 验证码画图
 * @receiver cc
 * @param size
 * @param dots
 * @return string
 * @return error
 */
func (cc *Captcha) genCaptchaImage(size Size, dots map[int]CharDot) (base64 string, erro error) {
	var drawDots []DrawDot
	for _, dot := range dots {
		drawDot := DrawDot{
			Dx:      dot.Dx,
			Dy:      dot.Dy,
			FontDPI: cc.config.fontDPI,
			Text:    dot.Text,
			Angle:   dot.Angle,
			Color:   dot.Color,
			Size:    dot.Size,
			Width:   dot.Width,
			Height:  dot.Height,
			Font:   cc.genRandWithString(cc.config.rangFont),
		}

		drawDots = append(drawDots, drawDot)
	}

	img, err := cc.captchaDraw.Draw(DrawCanvas{
		Width:             	size.Width,
		Height:            	size.Height,
		Background:        	cc.genRandWithString(cc.config.rangBackground),
		BackgroundDistort: 	cc.getRandDistortWithLevel(cc.config.imageFontDistort),
		TextAlpha:         	cc.config.imageFontAlpha,
		FontHinting: 	   	cc.config.fontHinting,
		CaptchaDrawDot:    	drawDots,

		ShowTextShadow: 	cc.config.showTextShadow,
		TextShadowColor: 	cc.config.textShadowColor,
		TextShadowPoint: 	cc.config.textShadowPoint,
	})
	if err != nil {
		erro = err
		return
	}

	// 转 base64
	base64 = cc.EncodeB64stringWithJpeg(img)
	return
}

/**
 * @Description: 验证码缩略画图
 * @receiver cc
 * @param size
 * @param dots
 * @return string
 * @return error
 */
func (cc *Captcha) genCaptchaThumbImage(size Size, dots map[int]CharDot) (string, error) {
	var drawDots []DrawDot

	fontWidth := size.Width / len(dots)
	for i, dot := range dots {
		Dx := int(math.Max(float64(fontWidth * i + fontWidth / dot.Width), 8))
		Dy := size.Height / 2 + dot.Size/2 - rand.Intn(size.Height / 16 * len(dot.Text))

		drawDot := DrawDot{
			Dx:      Dx,
			Dy:      Dy,
			FontDPI: cc.config.fontDPI,
			Text:    dot.Text,
			Angle:   dot.Angle,
			Color:   dot.Color2,
			Size:    dot.Size,
			Width:   dot.Width,
			Height:  dot.Height,
			Font:    cc.genRandWithString(cc.config.rangFont),
		}
		drawDots = append(drawDots, drawDot)
	}

	params := DrawCanvas{
		Width:                 size.Width,
		Height:                size.Height,
		CaptchaDrawDot:        drawDots,
		BackgroundDistort:     cc.getRandDistortWithLevel(cc.config.thumbFontDistort),
		BackgroundCirclesNum:  cc.config.thumbBgCirclesNum,
		BackgroundSlimLineNum: cc.config.thumbBgSlimLineNum,
	}

	if len(cc.config.rangThumbBackground) > 0 {
		params.Background = cc.genRandWithString(cc.config.rangThumbBackground)
	}

	var colorA []color.Color
	for _, cStr := range cc.config.rangThumbFontColors {
		co, _ := ParseHexColor(cStr)
		colorA = append(colorA, co)
	}

	var colorB []color.Color
	for _, co := range cc.config.rangThumbBgColors {
		rc, _ := ParseHexColor(co)
		colorB = append(colorB, rc)
	}

	img, err := cc.captchaDraw.DrawWithPalette(params, colorA, colorB)
	if err != nil {
		return "", err
	}

	// 转 base64
	dist := cc.EncodeB64stringWithPng(img)
	return dist, err
}

/**
 * @Description: 根据级别获取扭曲程序
 * @receiver cc
 * @param level
 * @return int
 */
func (cc *Captcha) getRandDistortWithLevel(level int) int {
	if level == 1 {
		return RandInt(240, 320)
	} else if level == 2 {
		return RandInt(180, 240)
	} else if level == 3 {
		return RandInt(120, 180)
	} else if level == 4 {
		return RandInt(100, 160)
	} else if level == 5 {
		return RandInt(80, 140)
	}
	return 0
}

/**
 * @Description: 获取随机角度
 * @receiver cc
 * @return int
 */
func (cc *Captcha) getRandAngle() int {
	angles := cc.config.rangTexAnglePos
	anglesLen := len(angles)
	index := RandInt(0, anglesLen)
	if index >= anglesLen {
		index = anglesLen - 1
	}

	angle := angles[index]
	res := RandInt(angle.Min, angle.Max)

	return res
}

/**
 * @Description: 随机获取颜色
 * @param colors
 * @return string
 */
func (cc *Captcha) getRandColor(colors []string) string {
	colorLen := len(colors)
	index := RandInt(0, colorLen)
	if index >= colorLen {
		index = colorLen - 1
	}

	return colors[index]
}

/**
 * @Description: 随机生成中文字符串
 * @param length
 * @return string
 */
func (cc *Captcha) genRandChar(length int) string {
	var strA []string
	for len(strA) < length {
		char := cc.randChar()
		if !InArrayWithStr(strA, char) {
			strA = append(strA, char)
		}
	}

	return strings.Join(strA, ":")
}

/**
 * @Description: 随机获取值
 * @param strs
 * @return string
 */
func (cc *Captcha) genRandWithString(strs []string) string {
	strLen := len(strs)
	if strLen == 0 {
		return ""
	}

	index := RandInt(0, strLen)
	if index >= strLen {
		index = strLen - 1
	}

	return strs[index]
}

/**
 * @Description: 随机一个字符
 * @return string
 */
func (cc *Captcha) randChar() string {
	chars := *cc.chars
	k := rand.Intn(len(chars))
	return chars[k]
}
