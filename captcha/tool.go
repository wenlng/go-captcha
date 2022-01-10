/**
 * @Author Awen
 * @Description Captcha Tool
 * @Date 2021/7/18
 * @Email wengaolng@gmail.com
 **/

package captcha

import (
	"crypto/rand"
	"errors"
	"image/color"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/**
 * @Description: 转16进制
 * @param t
 * @return string
 */
func t2x(t int64) string {
	result := strconv.FormatInt(t, 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

// RgbToHex is a function
/**
 * @Description: RBG颜色转十六进制颜色
 * @param red
 * @param green
 * @param blue
 * @return CaptchaColorHEX
 */
func RgbToHex(red int64, green int64, blue int64) string {
	r := t2x(red)
	g := t2x(green)
	b := t2x(blue)
	return r + g + b
}

// HexToRgb is a function
/**
 * @Description: 十六进制转RBG颜色
 * @param hex
 * @return int64
 * @return int64
 * @return int64
 */
func HexToRgb(hex string) (int64, int64, int64) {
	r, _ := strconv.ParseInt(hex[:2], 16, 10)
	g, _ := strconv.ParseInt(hex[2:4], 16, 18)
	b, _ := strconv.ParseInt(hex[4:], 16, 10)
	return r, g, b
}

// ParseHexColor is a function
/**
 * @Description: 把十六进制颜色转 color.RGBA
 * @param v
 * @return out
 * @return err
 */
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	if s[0] != '#' {
		return c, errors.New("hex color must start with '#'")
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errors.New("hexToByte component invalid")
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1]) << 4 + hexToByte(s[2])
		c.G = hexToByte(s[3]) << 4 + hexToByte(s[4])
		c.B = hexToByte(s[5]) << 4 + hexToByte(s[6])

	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errors.New("hexToByte component invalid")
	}
	return
}

// PathExists is a function
/**
 * @Description: 检测文件是否存在
 * @param path
 * @return bool
 * @return error
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Child is a function
/**
 * @Description: 获取路径
 * @param path
 * @return []string
 */
func Child(path string) []string {
	fullPath, _ := filepath.Abs(path)
	listStr := make([]string, 0)
	_ = filepath.Walk(fullPath, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".ttf") {
			listStr = append(listStr, path)
		}
		return err
	})
	return listStr
}

// InArrayWithStr is a function
/**
 * @Description: InArrayS 如果 s 在 items 中,返回 true；否则，返回 false。
 * @param items
 * @param s
 * @return bool
 */
func InArrayWithStr(items []string, s string) bool {
	for _, item := range items {
		if item == s {
			return true
		}
	}
	return false
}

// RandInt is a function
/**
 * @Description: 生成区间[-m, n]的安全随机数
 * @param min
 * @param max
 * @return int
 */
func RandInt(min, max int) int {
	if min > max {
		return max
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(int64(max+1+i64Min)))

		return int(result.Int64() - int64(i64Min))
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	return int(int64(min) + result.Int64())
}

// RandFloat is a function
/**
 * @Description: 随机浮点数
 * @param min
 * @param max
 * @return float64
 */
func RandFloat(min, max int) float64 {
	return float64(RandInt(min, max))
}

// IsChineseChar is a function
/**
 * @Description: 判断是否是中文
 * @param str
 * @return bool
 */
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

// LenChineseChar is a function
/**
 * @Description: 计算中文及字母长度，例如：“你好hello” = 7
 * @param str
 * @return int
 */
func LenChineseChar(str string) int {
	return utf8.RuneCountInString(str)
}