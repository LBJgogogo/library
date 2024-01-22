package Seal

import (
	"fmt"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func GenerateSeal(name string) (image.Image, error) {
	imageSize := 300
	borderCircleWidth := 3
	borderCircleHeight := 140
	borderInnerCircleWidth := 1
	borderInnerCircleHeight := 135
	innerCircleWidth := 2
	innerCircleHeight := 105

	// 计算文字位置
	fontSize := 25
	centerY := imageSize / 2
	nameWidth := getTextWidth(name+"专用章", fontSize)
	nameMarginSize := (imageSize - nameWidth) / 2
	namePoint := freetype.Pt(nameMarginSize, centerY+fontSize/2)

	// 创建图片
	rect := image.Rect(0, 0, imageSize, imageSize)
	img := image.NewRGBA(rect)

	// 设置背景颜色
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)

	// 绘制外边线圆
	drawCircle(img, borderCircleWidth, borderCircleHeight, color.White)

	// 绘制内边线圆
	drawCircle(img, borderInnerCircleWidth, borderInnerCircleHeight, color.White)

	// 绘制内环线圆
	drawCircle(img, innerCircleWidth, innerCircleHeight, color.White)

	// 绘制印章名字
	context := freetype.NewContext()
	context.SetDPI(72)
	fontData := goregular.TTF
	font, err := truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	context.SetFont(font)
	context.SetFontSize(float64(fontSize))
	context.SetClip(img.Bounds())
	context.SetDst(img)
	context.SetSrc(image.NewUniform(color.Black))
	_, err = context.DrawString(name+"专用章", namePoint)
	if err != nil {
		fmt.Println("Error drawing text:", err)
	}

	// 绘制其他文字
	drawText(img, "权威认证", 22, 5, centerY, color.Black)
	drawText(img, "个人签章专用", 25, 0, centerY, color.Black)

	return img, nil
}
func getTextWidth(text string, fontSize int) int {
	context := freetype.NewContext()
	context.SetDPI(72)
	fontData := goregular.TTF
	font, err := truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	context.SetFont(font)
	context.SetFontSize(float64(fontSize))
	widths := make([]int, len(text))
	for i, _ := range text {
		bounds := font.Bounds(fixed.I(fontSize))
		widths[i] = (bounds.Max.X - bounds.Min.X).Ceil()
	}
	sum := 0
	for _, width := range widths {
		sum += width
	}
	return sum
}

func drawCircle(img *image.RGBA, lineWidth, radius int, c color.Color) {
	centerX, centerY := img.Bounds().Dx()/2, img.Bounds().Dy()/2
	for y := -radius; y < radius; y++ {
		for x := -radius; x < radius; x++ {
			if x*x+y*y <= radius*radius {
				img.Set(centerX+x, centerY+y, c)
			}
		}
	}
}

func drawText(img *image.RGBA, text string, fontSize, marginSize, centerY int, c color.Color) {
	context := freetype.NewContext()
	context.SetDPI(72)
	fontData := goregular.TTF
	font, err := truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	context.SetFont(font)
	context.SetFontSize(float64(fontSize))
	context.SetClip(img.Bounds())
	context.SetDst(img)
	context.SetSrc(image.NewUniform(c))

	// 计算文字位置
	point := freetype.Pt(marginSize, centerY+fontSize/2)

	_, err = context.DrawString(text, point)
	if err != nil {
		fmt.Println("Error drawing text:", err)
	}
}
