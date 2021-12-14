/**
 * @Author Awen
 * @Description
 * @Date 2021/7/16
 **/

package captcha

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

// Md5ToString is a function
/**
 * @Description: Md5
 * @param str
 * @return string
 */
func Md5ToString(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GenUniqueId is a function
/**
 * @Description: GenUniqueId
 * @param str
 * @return string
 */
var num int64
func GenUniqueId() string {
	t := time.Now()
	s := t.Format("20060102150405")
	m := t.UnixNano() / 1e6 - t.UnixNano() / 1e9 * 1e3
	ms := Sup(m, 3)
	p := os.Getpid() % 1000
	ps := Sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := Sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

// Sup is a function
/**
 * @Description: 对长度不足n的数字前面补0
 * @param int64
 * @param int
 * @return string
 */
func Sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}