package main

import (
	"encoding/json"
	"fmt"
	"github.com/wenlng/go-captcha/captcha"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
	"time"
)

// Please use command line mode to start
// go run main.go
func main() {
	// Example: Get captcha data
	http.HandleFunc("/go_captcha_data", getCaptchaData)
	// Example: Post check data
	http.HandleFunc("/go_captcha_check_data", checkCaptcha)
	// Example: demo
	http.HandleFunc("/demo", demo)

	log.Println("ListenAndServe 0.0.0.0:8002")
	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// =========================================================

// demo is a function
/**
 * @Description: demo
 * @param w
 * @param r
 */
func demo(w http.ResponseWriter, r *http.Request) {
	//sessid := time.Now().UnixNano() / 1e6
	t, _ := template.ParseFiles(getPWD() + "/demo.html")
	_ = t.Execute(w, map[string]interface{}{})
}

/**
 * @Description: Example
 * @param w
 * @param r
 */
func getCaptchaData(w http.ResponseWriter, r *http.Request) {
	capt := captcha.GetCaptcha()

	//capt.SetImageQuality(captcha.QualityCompressLevel1)

	//chars := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//_ = capt.SetRangChars(strings.Split(chars, ""))

	//chars := []string{"HE","CA","WO","NE","HT","IE","PG","GI","CH","CO","DA"}
	//_ = capt.SetRangChars(chars)

	//chars := []string{"你","好","呀","这","是","点","击","验","证","码","哟"}
	//_ = capt.SetRangChars(chars)

	//capt.SetTextRangFontColors([]string{
	//	"#fdefac",
	//	"#8abcff",
	//	"#ffa37a",
	//	"#fcb3ff",
	//	"#b4fed4",
	//	"#cbfaa9",
	//})
	//
	//capt.SetThumbTextRangFontColors([]string{
	//	"#006600",
	//	"#005db9",
	//	"#aa002a",
	//	"#875400",
	//	"#6e3700",
	//	"#660033",
	//})

	// capt.SetFont([]string{
	// 	getPWD() + "/resources/fonts/fzshengsksjw_cu.ttf",
	// 	getPWD() + "/resources/fonts/hyrunyuan.ttf",
	// })

	//capt.SetBackground([]string{
	//	getPWD() + "/resources/images/1.jpg",
	//	getPWD() + "/resources/images/2.jpg",
	//	getPWD() + "/resources/images/3.jpg",
	//	getPWD() + "/resources/images/4.jpg",
	//	getPWD() + "/resources/images/5.jpg",
	//})

	//capt.SetThumbBackground([]string{
	//	getPWD() + "/resources/images/thumb/r1.jpg",
	//	getPWD() + "/resources/images/thumb/r2.jpg",
	//	getPWD() + "/resources/images/thumb/r3.jpg",
	//	getPWD() + "/resources/images/thumb/r4.jpg",
	//	getPWD() + "/resources/images/thumb/r5.jpg",
	//})

	//capt.SetThumbBgCirclesNum(200)
	//capt.SetImageFontAlpha(0.5)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		bt, _ := json.Marshal(map[string]interface{}{
			"code":    1,
			"message": "GenCaptcha err",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}
	writeCache(dots, key)
	bt, _ := json.Marshal(map[string]interface{}{
		"code":         0,
		"image_base64": b64,
		"thumb_base64": tb64,
		"captcha_key":  key,
	})
	_, _ = fmt.Fprintf(w, string(bt))
}

/**
 * @Description: Verify where the user clicks on the image
 * @param w
 * @param r
 */
func checkCaptcha(w http.ResponseWriter, r *http.Request) {
	code := 1
	_ = r.ParseForm()
	dots := r.Form.Get("dots")
	key := r.Form.Get("key")
	if dots == "" || key == "" {
		bt, _ := json.Marshal(map[string]interface{}{
			"code":    code,
			"message": "dots or key param is empty",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}

	cacheData := readCache(key)
	if cacheData == "" {
		bt, _ := json.Marshal(map[string]interface{}{
			"code":    code,
			"message": "illegal key",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}
	src := strings.Split(dots, ",")

	var dct map[int]captcha.CharDot
	if err := json.Unmarshal([]byte(cacheData), &dct); err != nil {
		bt, _ := json.Marshal(map[string]interface{}{
			"code":    code,
			"message": "illegal key",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}

	chkRet := false
	if len(src) >= len(dct)*2 {
		chkRet = true
		for i, dot := range dct {
			j := i * 2
			k := i*2 + 1
			sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[j]), 64)
			sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[k]), 64)
			// 检测点位置
			chkRet = captcha.CheckPointDist(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height))
			if !chkRet {
				break
			}
		}
	}

	if chkRet && (len(dct)*2) == len(src) {
		code = 0
	}

	bt, _ := json.Marshal(map[string]interface{}{
		"code": code,
	})
	_, _ = fmt.Fprintf(w, string(bt))
	return
}

/**
 * @Description: Write Cache，“Redis” is recommended
 * @param v
 * @param file
 */
func writeCache(v interface{}, file string) {
	bt, _ := json.Marshal(v)
	month := time.Now().Month().String()
	cacheDir := getCacheDir() + month + "/"
	_ = os.MkdirAll(cacheDir, 0660)
	file = cacheDir + file + ".json"
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer logFile.Close()
	_, _ = io.WriteString(logFile, string(bt))
}

/**
 * @Description: Read Cache，“Redis” is recommended
 * @param file
 * @return string
 */
func readCache(file string) string {
	month := time.Now().Month().String()
	cacheDir := getCacheDir() + month + "/"
	file = cacheDir + file + ".json"

	if !checkFileIsExist(file) {
		return ""
	}

	bt, err := ioutil.ReadFile(file)
	err = os.Remove(file)
	if err == nil {
		return string(bt)
	}
	return ""
}

/**
 * @Description: Calculate the distance between two points
 * @param sx
 * @param sy
 * @param dx
 * @param dy
 * @param width
 * @param height
 * @return bool
 */
func checkDist(sx, sy, dx, dy, width, height int64) bool {
	return sx >= dx &&
		sx <= dx+width &&
		sy <= dy &&
		sy >= dy-height
}

/**
 * @Description: Get cache dir path
 * @return string
 */
func getCacheDir() string {
	return getPWD() + "/.cache/"
}

/**
 * @Description: Get pwd dir path
 * @return string
 */
func getPWD() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

/**
 * @Description: Check file exist
 * @param filename
 * @return bool
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
