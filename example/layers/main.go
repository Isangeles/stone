/*
 * main.go
 *
 * Copyright 2019-2023 Dariusz Sikora <ds@isangeles.dev>
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

// Example for identifying map layer on specific position.
package main

import (
	"fmt"

	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"

	"github.com/isangeles/stone"
)

// Main function.
func main() {
	// Run Pixel graphic.
	pixelgl.Run(run)
}

// All window code fired from there.
func run() {
	// Create Pixel window configuration.
	cfg := pixelgl.WindowConfig{
		Title:  "Stone map layers example",
		Bounds: pixel.R(0, 0, 1600, 900),
	}
	// Create window.
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(fmt.Errorf("Unable to create pixel window: %v", err))
	}
	// Create map from TMX file.
	tmxMap, err := stone.NewMap("res/map.tmx")
	if err != nil {
		panic(fmt.Errorf("Unable to create map: %v", err))
	}
	// Creat info text.
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	infoText := text.New(pixel.V(0, 0), atlas)
	// Main loop.
	for !win.Closed() {
		// Clear window.
		win.Clear(colornames.Black)
		// Draw map.
		pos := pixel.V(0, 0) // e.g. camera pos
		tmxMap.Draw(win, pixel.IM.Moved(pos))
		// Retrieve layer for current mouse position.
		layer := tmxMap.PositionLayer(win.MousePosition())
		// Set layer info.
		infoText.Clear()
		if layer != nil {
			fmt.Fprintf(infoText, layer.Name())
		}
		// Draw info text.
		textPos := pixel.V(0, 0)
		infoText.Draw(win, pixel.IM.Moved(textPos))
		// Update.
		win.Update()
	}
}
