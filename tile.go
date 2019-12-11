/*
 * tile.go
 *
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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

package stone

import (
	"github.com/faiface/pixel"
)

// Struct for map tile.
type Tile struct {
	*pixel.Sprite
	bounds pixel.Rect
}

// newTile creates new map tile with specified sprite
// and position.
func newTile(spr *pixel.Sprite, pos pixel.Vec) *Tile {
	t := new(Tile)
	t.Sprite = spr
	t.bounds = pixel.R(pos.X, pos.Y, pos.X + t.Sprite.Frame().W(),
		pos.Y + t.Sprite.Frame().H())
	return t
}

// Position returns tile position.
func (t *Tile) Position() pixel.Vec {
	return t.bounds.Min
}

// Bounds returns tile size bounds.
func (t *Tile) Bounds() pixel.Rect {
	return t.bounds
}
