## Introduction
Stone is simple library that allows rendering TMX maps with [Pixel](https://github.com/faiface/pixel) library.

Originally created as part of [Mural](https://github.com/Isangeles/mural) GUI.

## Usage
Get Stone:
```
$ go get -u github.com/isangeles/stone
```

Create map:
```
tmxMap, err := stone.NewMap("path/to/map.tmx")
if err != nil {
   panic(fmt.Errorf("fail to create map: %v", err))
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

If you looking for things to do, then check TODO file or contact me(dev@isangeles.pl).

When you find something to do, create new branch for your feature.
After you finish, open pull request to merge your changes with master branch.

## Contact
* Isangeles <<dev@isangeles.pl>>

## License
Copyright 2019 Dariusz Sikora <<dev@isangeles.pl>>

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
