/**
 * @Author Awen
 * @Description
 * @Date 2021/7/16
 **/

package captcha

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

/**
 * @Description: 生成KEY
 * @param str str
 * @return Hash
 * @return err
 */
func GenerateKey(str string) (string, error) {
	secret := "HW85SDdRhu1Y45av"
	pwd := []byte(str + secret)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	_hash := string(hash)
	return _hash, nil
}

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
