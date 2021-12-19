package pprof

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"net/http"
	_ "net/http/pprof"
	"os"
)

/**
 * @Description: 二进制编码
 * @param img
 * @return []byte
 */
func binaryEncoding(img image.Image) (ret []byte) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err.Error())
	}
	ret = buf.Bytes()
	buf.Reset()
	return
}

/**
 * @Description: Get pwd dir path
 * @return string
 */
func getRoot() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

func Handler(w http.ResponseWriter, r *http.Request) {
	imgBg, err := os.Open(getRoot() + "/tests/.cache/index.png")
	if err != nil {
		panic(err)
	}
	defer imgBg.Close()
	//image.Decode(imgBg)
	img, err := png.Decode(imgBg)
	if err != nil {
		panic(err)
	}
	//str := binaryEncoding(img)
	//fmt.Println(len(str))
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>> ", &img)
}
