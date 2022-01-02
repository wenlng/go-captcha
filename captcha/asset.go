/**
 * @Author Awen
 * @Description
 * @Date 2021/12/18
 **/

package captcha

import "github.com/wenlng/go-captcha/captcha/assets"

/**
 * @Description: 获取缓存资源
 * @param path
 * @return []byte
 * @return error
 */
func getAssetCache(path string) (ret []byte, erro error) {
	return assets.GetAssetCache(path)
}

/**
 * @Description: 资源是否缓存
 * @param path
 * @return []byte
 * @return error
 */
func hasAssetCache(path string) bool {
	return assets.HasAssetCache(path)
}

/**
 * @Description: 清除资源缓存
 * @param paths
 * @return bool
 */
func clearAssetCache(paths []string) bool {
	return assets.ClearAssetCache(paths)
}

/**
 * @Description: 设置缓存资源
 * @param path
 * @return error
 */
func setAssetCache(path string, content []byte, force bool) bool {
	return assets.SetAssetCache(path, content, force)
}