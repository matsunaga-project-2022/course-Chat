package valueobject

import "os"

// Image は画像を表すValueObjectです
type Image []byte

// FileToImage はファイルからValueObjectへ変換する関数です
func FileToImage(file os.File) (Image, error) {
	panic("implement me")
}
