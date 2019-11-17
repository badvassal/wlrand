# wlrand

## Description

`wlrand` is a randomizer for the 1988 DOS game Wasteland.  The current version only randomizes transitions, i.e., the tiles that transport the player from one location to another.  `wlrand` produces a "new game" in the sense that even an experienced player must explore the randomized game and discover how to reach the various locations in the wastes.

## WARNING

`wlrand` permanently and irreversibly modifies your Wasteland `GAME1` and `GAME2` files.  It is imperative that you create backups of these files before running `wlrand`!

## Quick Start

The below instructions assume wasteland is "installed" at `/usr/local/share/games/dos-games/wasteland`.  Please adjust accordingly if this assumption is inaccurate.

1. Back up your game files:
```
cp /usr/local/share/games/dos-games/wasteland/GAME1 cp /usr/local/share/games/dos-games/wasteland/backup-GAME1 
cp /usr/local/share/games/dos-games/wasteland/GAME2 cp /usr/local/share/games/dos-games/wasteland/backup-GAME2 
```

2. Run the randomizer:
```
wlrand -p /usr/local/share/games/dos-games/wasteland
```

Wasteland should now be randomized.  Verify by entering any non-shop location in a city (e.g., Highpool community center).

## Building

`wlrand` is a Go program.  To build it, obtain the official [Go compiler](https://golang.org/dl/) and run the following command from the `wlrand` directory:
```
go build
```

This produces a `wlrand` executable in the current directory.

## De-randomize

There is currently no way to reverse the effect of `wlrand`.  The only way to get back to the original state is to restore backups of the `GAME1` and `GAME2` files.

## Bugs

Please report bugs using the wlrand issue tracker: <https://github.com/badvassal/wlrand/issues>

## To do

* Automatically back up `GAMEx` files; add a restore command.
* Randomize loot bags.
* Randomize shops.
* Randomize NPCs?
* Produce a winnable game when `--world` is specified.
* Make the Spade's Casino basement escapable without a shovel or explosives.
* Consider making the Base Cochise levels easily escapable.

## Acknowledgements

Thank you to the authors of [Wasteland: The Definitive Deconstruction](https://wasteland.gamepedia.com/Category:Wasteland:_The_Definitive_Deconstruction) without which this program would not be possible.
