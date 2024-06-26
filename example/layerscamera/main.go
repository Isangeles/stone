/*
 * main.go
 *
 * Copyright 2019-2024 Dariusz Sikora <ds@isangeles.dev>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
 * MA 02110-1301, USA.
 *
 *
 */

// Example for identifying map layer on specific position
// while moving camera.
package main

import (
	"fmt"
	"math"

	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/gopxl/pixel/text"

	"github.com/isangeles/stone"
)

var (
	cameraPos  pixel.Vec
	cameraZoom = 1.0
	areaMap    *stone.Map
)

// Main function.
func main() {
	// Run Pixel graphic.
	pixelgl.Run(run)
}

// All window code fired from there.
func run() {
	// Create Pixel window.
	cfg := pixelgl.WindowConfig{
		Title:  "Stone map layers and camera example",
		Bounds: pixel.R(0, 0, 1600, 900),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(fmt.Errorf("Unable to create pixel window: %v", err))
	}
	// Create map from TMX file.
	m, err := stone.NewMap("res/map.tmx")
	if err != nil {
		panic(fmt.Errorf("Unable to create map: %v", err))
	}
	areaMap = m
	// Creat info text.
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	infoText := text.New(pixel.V(0, 0), atlas)
	// Main loop.
	for !win.Closed() {
		// Clear window.
		win.Clear(colornames.Black)
		// Draw map.
		areaMap.Draw(win, pixel.IM.Moved(cameraPos).Scaled(cameraPos, cameraZoom))
		// We need to convert mouse position to map position.
		mousePos := convCameraPos(win.MousePosition())
		// Retrieve layer for current mouse position.
		layer := areaMap.PositionLayer(mousePos)
		// Set layer info.
		infoText.Clear()
		if layer != nil {
			fmt.Fprintf(infoText, layer.Name())
		}
		// Draw info text.
		textPos := pixel.V(0, 0)
		infoText.Draw(win, pixel.IM.Moved(textPos))
		// Key events.
		keyMouseEvents(win)
		// Update.
		win.Update()
	}
}

// convCameraPos translates specified camera
// position to area position.
func convCameraPos(pos pixel.Vec) pixel.Vec {
	areaPos := pixel.V(pos.X+cameraPos.X, pos.Y+cameraPos.Y)
	areaPos.X /= cameraZoom
	areaPos.Y /= cameraZoom
	return areaPos
}

// keyMouseEvents handles window key and mouse events.
func keyMouseEvents(win *pixelgl.Window) {
	// Moves camera one tile right/left/up/down on WSAD or arrow keys event.
	if win.JustPressed(pixelgl.KeyW) || win.JustPressed(pixelgl.KeyUp) {
		cameraPos.Y += areaMap.TileSize().Y
	}
	if win.JustPressed(pixelgl.KeyD) || win.JustPressed(pixelgl.KeyRight) {
		cameraPos.X += areaMap.TileSize().X
	}
	if win.JustPressed(pixelgl.KeyS) || win.JustPressed(pixelgl.KeyDown) {
		cameraPos.Y -= areaMap.TileSize().Y
	}
	if win.JustPressed(pixelgl.KeyA) || win.JustPressed(pixelgl.KeyLeft) {
		cameraPos.X -= areaMap.TileSize().X
	}
	// Zoom on mouse scroll.
	zoomSpeed := 1.1
	cameraZoom *= math.Pow(zoomSpeed, win.MouseScroll().Y)
}
