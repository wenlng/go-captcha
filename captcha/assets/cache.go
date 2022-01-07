/**
 * @Author Awen
 * @Description
 * @Date 2021/12/18
 **/

package assets

import (
	"github.com/wenlng/go-captcha/captcha/assets/fonts"
	"github.com/wenlng/go-captcha/captcha/assets/images"
)

type AssetData struct {
	// 路径
	Path string
	// 内容
	Content []byte
}

var cache []*AssetData

var defaultAssetsImage = []string{
	"assets/images/1.jpg",
	"assets/images/2.jpg",
	"assets/images/3.jpg",
	"assets/images/4.jpg",
	"assets/images/5.jpg",
}

var defaultAssetsFont = []string{
	"assets/fonts/fzshengsksjw_cu.ttf",
}

/**
 * @Description: 获取默认资源
 * @param path
 * @return []byte
 * @return error
 */
func findFontsAsset(path string) ([]byte, error) {
	return fonts.FindAsset(path)
}

/**
 * @Description: 获取默认资源
 * @param path
 * @return []byte
 * @return error
 */
func findImagesAsset(path string) ([]byte, error) {
	return images.FindAsset(path)
}

/**
 * @Description: 内置默认资源
 * @return []string
 */
func DefaultBinFontList() []string {
	return defaultAssetsFont
}

/**
 * @Description: 内置默认资源
 * @return []string
 */
func DefaultBinImageList() []string {
	return defaultAssetsImage
}

// GetAssetCache is a function
/**
 * @Description: 获取缓存资源
 * @param path
 * @return []byte
 * @return error
 */
func GetAssetCache(path string) (ret []byte, erro error) {
	if len(cache) > 0 {
		for _, asset := range cache {
			if asset.Path == path{
				ret = asset.Content
				return
			}
		}
	}

	ret, erro = findFontsAsset(path)
	if len(ret) > 0{
		cache = append(cache, &AssetData{
			Path: path,
			Content: ret,
		})
		return
	}

	ret, erro = findImagesAsset(path)
	if len(ret) > 0{
		cache = append(cache, &AssetData{
			Path: path,
			Content: ret,
		})
		return
	}
	return
}

// HasAssetCache is a function
/**
 * @Description: 资源是否缓存
 * @param path
 * @return bool
 */
func HasAssetCache(path string) bool {
	if len(cache) > 0 {
		for _, asset := range cache {
			if asset.Path == path{
				return true
			}
		}
	}
	return false
}


// ClearAssetCache is a function
/**
 * @Description: 清除资源缓存
 * @param paths
 * @return bool
 */
func ClearAssetCache(paths []string) bool {
	if len(cache) > 0 {
		for _, path := range paths {
			for ak, asset := range cache {
				if asset.Path == path{
					cache = append(cache[:ak], cache[(ak+1):]...)
					break
				}
			}
		}
	}
	return true
}

// SetAssetCache is a function
/**
 * @Description: 设置缓存资源
 * @param path
 * @return error
 */
func SetAssetCache(path string, content []byte, force bool) bool {
	if len(cache) > 0 {
		for _, asset := range cache {
			if asset.Path == path && !force{
				return true
			}
		}
	}

	cache = append(cache, &AssetData{
		Path: path,
		Content: content,
	})
	return true
}