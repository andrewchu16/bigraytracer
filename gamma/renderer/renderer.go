package renderer

import (
	"errors"
	"fmt"
	"gamma/geometry"
	"gamma/scene"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

type Renderer struct {
	imgWidth       int
	imgHeight      int
	viewportWidth  float64
	viewportHeight float64
	focalLength    float64

	scene       *scene.Scene
	pixelBuffer [][]geometry.Vec3
	rendered    bool
}

func NewRenderer(imgWidth, imgHeight int) Renderer {
	var viewportHeight float64 = 2.0
	var viewportWidth float64 = 2.0 * float64(imgWidth) / float64(imgHeight)

	var focalLength float64 = 1.0

	var buffer [][]geometry.Vec3
	for i := 0; i < imgHeight; i++ {
		buffer = append(buffer, make([]geometry.Vec3, imgWidth))
	}

	return Renderer{
		imgWidth:       imgWidth,
		imgHeight:      imgHeight,
		viewportWidth:  viewportWidth,
		viewportHeight: viewportHeight,
		focalLength:    focalLength,
		pixelBuffer:    buffer,
		rendered:       false,
	}
}

func (r *Renderer) SetScene(scene *scene.Scene) {
	r.scene = scene
}

func (r *Renderer) Render() {
	// do the rendering...
	r.rendered = true
}

func (r *Renderer) Resize(imgWidth, imgHeight int) {
	r.imgWidth = imgWidth
	r.imgHeight = imgHeight
	r.viewportWidth = 2.0 * float64(imgWidth) / float64(imgHeight)

	var buffer [][]geometry.Vec3
	for i := 0; i < imgHeight; i++ {
		buffer = append(buffer, make([]geometry.Vec3, imgWidth))
	}
	r.pixelBuffer = buffer
	r.rendered = false
}

func (r *Renderer) createImageData() (*image.RGBA, error) {
	if !r.rendered {
		return nil, errors.New("cannot create image data before rendering")
	}

	img := image.NewRGBA(image.Rect(0, 0, r.imgWidth, r.imgHeight))

	// Convert buffer to image
	for y := range r.imgHeight {
		for x := range r.imgWidth {
			// Convert normalized color to 8-bit color
			red := uint8(r.pixelBuffer[y][x].X * 255)
			green := uint8(r.pixelBuffer[y][x].Y * 255)
			blue := uint8(r.pixelBuffer[y][x].Z * 255)

			img.Set(x, y, color.RGBA{red, green, blue, 255}) // Alpha is always 255
		}
	}

	return img, nil
}

// Export the rendered image to the specified filename and format
func (r *Renderer) Export(filename string, format SupportedImageFormats) error {
	img, err := r.createImageData()

	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	done := make(chan error, 1)

	go func(done chan error) {
		defer file.Close()

		var encodeErr error
		switch format {
		case PNG:
			encodeErr = png.Encode(file, img)
		case JPEG:
			encodeErr = jpeg.Encode(file, img, nil)
		default:
			encodeErr = fmt.Errorf("unsupported image format. %v", format)
		}

		if encodeErr != nil {
			fmt.Fprintln(os.Stderr, "Error encoding image:", encodeErr)
		}

		done <- encodeErr
	}(done)

	err = <-done

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error encoding image:", err)
		return err
	}

	return nil
}
