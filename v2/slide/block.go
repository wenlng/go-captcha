/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package slide

import (
	"image"
)

// Block defines the block data for the slide CAPTCHA
type Block struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
	Angle  int `json:"angle"`
	// Deprecated: As of 2.1.0, it will be removed, please use [Block.DX].
	TileX int `json:"tile_x"`
	// Deprecated: As of 2.1.0, it will be removed, please use [Block.DY].
	TileY int `json:"tile_y"`
	// Display x,y
	DX int `json:"dx"`
	DY int `json:"dy"`
}

// DrawBlock defines the parameters for drawing slide CAPTCHA blocks
type DrawBlock struct {
	Block  *Block
	X      int
	Y      int
	Image  image.Image
	Width  int
	Height int
	Angle  int
}
