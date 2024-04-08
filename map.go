/*
 * map.go
 *
 * Copyright 2018-2024 Dariusz Sikora <ds@isangeles.dev>
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

// Stone allows rendering TMX maps.
package stone

import (
	"fmt"
	"path/filepath"

	"github.com/salviati/go-tmx/tmx"

	"github.com/gopxl/pixel"
)

// Struct for graphical representation of TMX map.
type Map struct {
	tmxMap      *tmx.Map
	tilesets    map[string]pixel.Picture
	tileBatches map[pixel.Picture]*pixel.Batch
	tilesize    pixel.Vec
	mapsize     pixel.Vec
	tilescount  pixel.Vec
	layers      []*Layer
}

// NewMap creates new map from .tmx file with specified path.
func NewMap(path string) (*Map, error) {
	tmxMap, err := tmxMap(path)
	if err != nil {
		return nil, fmt.Errorf("unable to retive TMX map: %v", err)
	}
	m := new(Map)
	m.tmxMap = tmxMap
	m.tilesize = pixel.V(float64(m.tmxMap.TileWidth),
		float64(m.tmxMap.TileHeight))
	m.tilescount = pixel.V(float64(m.tmxMap.Width),
		float64(m.tmxMap.Height))
	m.mapsize = pixel.V(float64(int(m.tilesize.X*m.tilescount.X)),
		float64(int(m.tilesize.Y*m.tilescount.Y)))
	m.tilesets = make(map[string]pixel.Picture)
	m.tileBatches = make(map[pixel.Picture]*pixel.Batch)
	mapDir := filepath.Dir(path)
	// Tilesets.
	for _, ts := range m.tmxMap.Tilesets {
		tsPath := filepath.FromSlash(mapDir + "/" + ts.Image.Source)
		tsPic, err := picture(tsPath)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve tilset source: %v",
				ts.Name)
		}
		m.tilesets[ts.Name] = tsPic
		m.tileBatches[tsPic] = pixel.NewBatch(&pixel.TrianglesData{}, tsPic)
	}
	// Map layers.
	for _, l := range m.tmxMap.Layers {
		layer, err := newLayer(m, l)
		if err != nil {
			return nil, fmt.Errorf("unable to create layer: %s: %v",
				l.Name, err)
		}
		m.layers = append(m.layers, layer)
	}
	return m, nil
}

// DrawSize use specified matrix and size to draw map on target.
// Draws part of the map in specified size starting from position
// specified in given matrix.
func (m *Map) DrawPart(tar pixel.Target, matrix pixel.Matrix, size pixel.Vec) {
	drawArea := pixel.R(matrix[4], matrix[5], matrix[4]+size.X,
		matrix[5]+size.Y)
	// Clear all tilesets draw batches.
	for _, batch := range m.tileBatches {
		batch.Clear()
	}
	// Draw layers tiles to tilesets batechs.
	for _, l := range m.layers {
		for _, t := range l.tiles {
			tilePos := pixel.V(t.Position().X*matrix[0], t.Position().Y*matrix[0])
			if drawArea.Contains(tilePos) {
				batch := m.tileBatches[t.Picture()]
				if batch == nil {
					continue
				}
				tileDrawPos := mapDrawPos(t.Position(), matrix)
				t.Draw(batch, pixel.IM.Scaled(pixel.V(0, 0),
					matrix[0]).Moved(tileDrawPos))
			}
		}
	}
	// Draw bateches with layer tiles.
	drawn := make(map[pixel.Picture]*pixel.Batch)
	for _, l := range m.layers {
		for _, t := range l.tiles {
			batch := m.tileBatches[t.Picture()]
			if batch == nil || drawn[t.Picture()] != nil {
				continue
			}
			batch.Draw(tar)
			drawn[t.Picture()] = batch
		}
	}
}

// Draw use specified matrix to draw map on target.
// Draws whole map starting from position specified in given matrix.
func (m *Map) Draw(tar pixel.Target, matrix pixel.Matrix) {
	// Clear all tilesets draw batches.
	for _, batch := range m.tileBatches {
		batch.Clear()
	}
	// Draw layers tile to tileset batechs.
	for _, l := range m.layers {
		for _, t := range l.tiles {
			batch := m.tileBatches[t.Picture()]
			if batch == nil {
				continue
			}
			tileDrawPos := mapDrawPos(t.Position(), matrix)
			t.Draw(batch, pixel.IM.Scaled(pixel.V(0, 0),
				matrix[0]).Moved(tileDrawPos))
		}
	}
	// Draw bateches with layer tiles.
	drawn := make(map[pixel.Picture]*pixel.Batch)
	for _, l := range m.layers {
		for _, t := range l.tiles {
			batch := m.tileBatches[t.Picture()]
			if batch == nil || drawn[t.Picture()] != nil {
				continue
			}
			batch.Draw(tar)
			drawn[t.Picture()] = batch
		}
	}
}

// TileSize returns size of singe map tile.
func (m *Map) TileSize() pixel.Vec {
	return m.tilesize
}

// Size returns size of the map.
func (m *Map) Size() pixel.Vec {
	return m.mapsize
}

// Layers returns all map layers.
func (m *Map) Layers() []*Layer {
	return m.layers
}

// PositionLayer returns visible layer on specified
// position on map or nil if there is no tiles on
// this position.
func (m *Map) PositionLayer(p pixel.Vec) *Layer {
	var visibleLayer *Layer
	for _, l := range m.Layers() {
		for _, t := range l.Tiles() {
			if t.Bounds().Contains(p) {
				visibleLayer = l
			}
		}
	}
	return visibleLayer
}

// tileBounds returns bounds for tile with specified size and ID
// from specified tileset picture.
func (m *Map) tileBounds(tileset pixel.Picture, tileID tmx.ID) pixel.Rect {
	tilesetSize := roundTilesetSize(tileset.Bounds().Max, m.tilesize)
	tileCount := 0
	for h := tilesetSize.Y - m.tilesize.Y; h >= 0; h -= m.tilesize.Y {
		for w := 0.0; w+m.tilesize.X <= tilesetSize.X; w += m.tilesize.X {
			if tileCount == int(tileID) {
				tileBounds := pixel.R(w, h, w+m.tilesize.X,
					h+m.tilesize.Y)
				return tileBounds
			}
			tileCount++
		}
	}
	return pixel.R(0, 0, 0, 0)

}
