/**
 * @Author Awen
 * @Description
 * @Date 2021/12/18
 **/

package assets

type AssetData struct {
	// 路径
	Path string
	// 内容
	Content []byte
}

var cache []*AssetData

/**
 * @Description: 获取默认资源
 * @param path
 * @return []byte
 * @return error
 */
func findAsset(path string) ([]byte, error) {
	return Asset(path)
}

/**
 * @Description: 内置默认资源
 * @return []string
 */
func DefaultBinFontList() []string {
	return []string{
		"assets/fonts/fzshengsksjw_cu.ttf",
		"assets/fonts/hyrunyuan.ttf",
	}
}

/**
 * @Description: 内置默认资源
 * @return []string
 */
func DefaultBinImageList() []string {
	return []string{
		"assets/images/1.jpg",
		"assets/images/2.jpg",
		"assets/images/3.jpg",
		"assets/images/4.jpg",
		"assets/images/5.jpg",
	}
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

	ret, erro = findAsset(path)
	if len(ret) > 0{
		cache = append(cache, &AssetData{
			Path: path,
			Content: ret,
		})
	}
	return
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