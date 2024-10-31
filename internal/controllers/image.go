package controllers

import (
	"image"
	"image/png"
	"os"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
}

type RGBA struct {
	R uint8 `json:"r"` // 红色分量
	G uint8 `json:"g"` // 绿色分量
	B uint8 `json:"b"` // 蓝色分量
	A uint8 `json:"a"` // 透明度分量
}

// Image 定义了一张图片，包含宽度、高度和像素数据
type Image struct {
	Width  int      `json:"width"`  // 图片的宽度
	Height int      `json:"height"` // 图片的高度
	Pixels [][]RGBA `json:"pixels"` // 像素数组
}

func decodeImage(_file string) (*Image, error) {
	file, err := os.Open(_file)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 解码图像
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	im := &Image{
		Pixels: make([][]RGBA, 0),
	}

	// 转换为 RGBA
	rgbaImg := image.NewRGBA(img.Bounds())
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			c := img.At(x, y) // 获取每个像素的颜色
			rgbaImg.Set(x, y, c)
		}
	}

	im.Width = rgbaImg.Rect.Max.X - rgbaImg.Rect.Min.X
	im.Height = rgbaImg.Rect.Max.Y - rgbaImg.Rect.Min.Y

	for y := 0; y < rgbaImg.Bounds().Dy(); y++ {
		row := make([]RGBA, 0)

		for x := 0; x < rgbaImg.Bounds().Dx(); x++ {
			r, g, b, a := rgbaImg.At(x, y).RGBA()
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			a8 := uint8(a >> 8)

			p := RGBA{
				R: r8,
				G: g8,
				B: b8,
				A: a8,
			}
			row = append(row, p)
		}
		im.Pixels = append(im.Pixels, row)
	}
	return im, nil
}

func (sc *ImageController) Decode(c *gin.Context) {
	param := &BasicField{}
	err := c.BindJSON(&param)
	if err != nil {
		ReturnFailResponse(c, 400, "no param entry")
		return
	}

	im, err := decodeImage(param.Entry)
	if err != nil {
		ReturnFailResponse(c, 400, err.Error())
	}
	ReturnSuccessResponse(c, im)
}
