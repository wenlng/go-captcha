package main

import (
	"encoding/json"
	"fmt"
	"github.com/wenlng/goCaptcha/captcha"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Example: Get captcha data
	http.HandleFunc("/captcha-data", GetCaptchaData)
	// Example: Post check data
	http.HandleFunc("/check-data", CheckCaptcha)
	// Example: Demo
	http.HandleFunc("/demo", Demo)

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("ListenAndServe 0.0.0.0:8082")
}

// =========================================================

/**
 * @Description: Demo
 * @param w
 * @param r
 */
func Demo(w http.ResponseWriter, r *http.Request) {
	sessid := time.Now().UnixNano() / 1e6
	t, _ := template.ParseFiles(getPWD() + "/__example/demo.html")
	_ = t.Execute(w, map[string]interface{}{"sessid": sessid})
}

/**
 * @Description: Example
 * @param w
 * @param r
 */
func GetCaptchaData(w http.ResponseWriter, r *http.Request) {
	capt := captcha.GetCaptcha()

	//chars := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//_ = capt.SetRangChars(strings.Split(chars, ""))

	//chars := []string{"HE","CA","WO","NE","HT","IE","PG","GI","CH","CO","DA"}
	//_ = capt.SetRangChars(chars)

	//chars := []string{"你","好","呀","这","是","点","击","验","证","码","哟"}
	//_ = capt.SetRangChars(chars)

	/*



	 */

	capt.SetTextRangFontColors([]string{
		"#006600",
		"#005db9",
		"#aa002a",
		"#875400",
		"#6e3700",
		"#333333",
		"#660033",
	})

	capt.SetFont([]string{
		getPWD() + "/__example/resources/fonts/fzshengsksjw_cu.ttf",
		getPWD() + "/__example/resources/fonts/fzssksxl.ttf",
		getPWD() + "/__example/resources/fonts/hyrunyuan.ttf",
	})

	capt.SetBackground([]string{
		getPWD() + "/__example/resources/images/1.jpg",
		getPWD() + "/__example/resources/images/2.jpg",
		getPWD() + "/__example/resources/images/3.jpg",
		getPWD() + "/__example/resources/images/4.jpg",
		getPWD() + "/__example/resources/images/5.jpg",
	})

	//capt.SetThumbBackground([]string{
	//	getPWD() + "/__example/resources/images/thumb/r1.jpg",
	//	getPWD() + "/__example/resources/images/thumb/r2.jpg",
	//	getPWD() + "/__example/resources/images/thumb/r3.jpg",
	//	getPWD() + "/__example/resources/images/thumb/r4.jpg",
	//	getPWD() + "/__example/resources/images/thumb/r5.jpg",
	//})

	//capt.SetThumbBgCirclesNum(200)
	//capt.SetImageFontAlpha(0.5)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		bt, _ := json.Marshal(map[string]interface{}{
			"code": 1,
			"message": "GenCaptcha err",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}
	WriteCache(dots, key)
	bt, _ := json.Marshal(map[string]interface{}{
		"code": 0,
		"image_base64": b64,
		"thumb_base64": tb64,
		"captcha_key": key,
	})
	_, _ = fmt.Fprintf(w, string(bt))
}

/**
 * @Description: Verify where the user clicks on the image
 * @param w
 * @param r
 */
func CheckCaptcha(w http.ResponseWriter, r *http.Request) {
	code := 1
	_ = r.ParseForm()
	dots := r.Form.Get("dots")
	key := r.Form.Get("key")
	if dots == "" || key == "" {
		bt, _ := json.Marshal(map[string]interface{}{
			"code": code,
			"message": "dots or key param is empty",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}

	cacheData := ReadCache(key)
	if cacheData == "" {
		bt, _ := json.Marshal(map[string]interface{}{
			"code": code,
			"message": "illegal key",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}
	src := strings.Split(dots, ",")

	var dct map[int]captcha.CharDot
	if err := json.Unmarshal([]byte(cacheData), &dct); err != nil {
		bt, _ := json.Marshal(map[string]interface{}{
			"code": code,
			"message": "illegal key",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}

	chkRet := false
	if len(src) >= len(dct) * 2 {
		chkRet = true
		for i, dot := range dct {
			j := i * 2
			k := i * 2 + 1
			a, _ := strconv.Atoi(src[j])
			b, _ := strconv.Atoi(src[k])
			chkRet = CheckDist(a, b, dot.Dx, dot.Dy, dot.Width, dot.Height)
			if !chkRet {
				break
			}
		}
	}

	if chkRet && (len(dct) * 2) == len(src) {
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
func WriteCache(v interface{}, file string) {
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
func ReadCache(file string) string {
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
func CheckDist(sx, sy, dx, dy, width int, height int) bool {
	return sx >= dx &&
		sx <= dx + width &&
		sy <= dy &&
		sy >= dy - height
}

/**
 * @Description: Get cache dir path
 * @return string
 */
func getCacheDir() string  {
	return getPWD() + "/__example/.cache/"
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
