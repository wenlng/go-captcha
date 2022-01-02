package images

/**
 * @Description: 获取资源
 * @param path
 * @return []byte
 * @return error
 */
func FindAsset(path string) ([]byte, error) {
	return Asset(path)
}
