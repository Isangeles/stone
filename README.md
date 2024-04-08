## Introduction
Stone is simple library that allows rendering [Tiled](https://www.mapeditor.org/) maps with [Pixel](https://github.com/gopxl/pixel) library.

Originally created as part of [Mural](https://github.com/Isangeles/mural) GUI.

## Usage
First, make sure you have dependencies required by [Pixel](https://github.com/gopxl/pixel).

Get sources from git:
```
go get -u github.com/isangeles/stone
```

Create map:
```
tmxMap, err := stone.NewMap("path/to/map.tmx")
if err != nil {
   panic(fmt.Errorf("Unable to create map: %v", err))
}
```

Draw map in Pixel window:
```
for !win.Closed() {
    // ...
    pos = pixel.V(0, 0) // e.g. camera pos
    tmxMap.Draw(win, pixel.IM.Moved(pos))
}
```

Check [example](https://github.com/Isangeles/stone/tree/master/example) package for more examples.

## Contributing
You are welcome to contribute to project development.

If you looking for things to do, then check [TODO file](https://github.com/Isangeles/stone/blob/master/TODO) or contact maintainer(ds@isangeles.dev).

When you find something to do, create new branch for your feature.
After you finish, open pull request to merge your changes with master branch.

## Contact
* Isangeles <<ds@isangeles.dev>>

## License
Copyright 2019-2024 Dariusz Sikora <<ds@isangeles.dev>>

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
MA 02110-1301, USA.
