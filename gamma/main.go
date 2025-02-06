package main

import (
	"gamma/renderer"
	"gamma/scene"
)

const (
	IMG_WIDTH  = 1920
	IMG_HEIGHT = 1080
	IMG_NAME   = "output.png"
	IMG_FORMAT = renderer.PNG
)

func main() {
	r := renderer.NewRenderer(IMG_WIDTH, IMG_HEIGHT)
	s := scene.NewScene()

	r.SetScene(s)
	r.Render()
	r.Export(IMG_NAME, IMG_FORMAT)
}
