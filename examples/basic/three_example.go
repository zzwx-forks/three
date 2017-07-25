package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lngramos/three"
)

var (
	scene    three.Scene
	camera   three.PerspectiveCamera
	renderer three.WebGLRenderer
	mesh     three.Mesh
)

func main() {
	document := js.Global.Get("document")
	windowWidth := js.Global.Get("innerWidth").Float()
	windowHeight := js.Global.Get("innerHeight").Float()
	devicePixelRatio := js.Global.Get("devicePixelRatio").Float()

	camera = three.NewPerspectiveCamera(70, windowWidth/windowHeight, 1, 1000)
	camera.Position.Set(0, 0, 400)

	scene = three.NewScene()

	renderer = three.NewWebGLRenderer()
	renderer.SetPixelRatio(devicePixelRatio)
	renderer.SetSize(windowWidth, windowHeight, true)
	document.Get("body").Call("appendChild", renderer.Get("domElement"))

	// Create cube
	geometry := three.NewBoxGeometry(100, 100, 100)
	material := three.NewMeshBasicMaterial(three.MeshBasicMaterialParameters{Color: three.NewColor(255, 0, 0)})
	mesh = three.NewMesh(geometry, material)

	scene.Add(mesh)

	animate()
}

func animate() {
	js.Global.Call("requestAnimationFrame", animate)

	pos := mesh.Object.Get("rotation")
	pos.Set("x", pos.Get("x").Float()+float64(0.01))
	pos.Set("y", pos.Get("y").Float()+float64(0.01))

	renderer.Render(scene, camera)
}
