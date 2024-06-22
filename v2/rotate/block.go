/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package rotate

type Block struct {
	ParentWidth  int `json:"parent_width"`
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
