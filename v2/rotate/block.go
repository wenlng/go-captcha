/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package rotate

// Block ..
type Block struct {
	// Deprecated: As of 2.1.0, it will be removed, please use [[CaptchaInstance].GetOptions().GetImageSize()].
	ParentWidth int `json:"parent_width"`
	// Deprecated: As of 2.1.0, it will be removed, please use [[CaptchaInstance].GetOptions().GetImageSize()].
	ParentHeight int `json:"parent_height"`
	Width        int `json:"width"`
	Height       int `json:"height"`
	Angle        int `json:"angle"`
}

// DrawBlock .
type DrawBlock struct {
	Block  *Block
	X      int
	Y      int
	Width  int
	Height int
	Angle  int
}
