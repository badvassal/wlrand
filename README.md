# wlrand

https://github.com/badvassal/wlrand

## Description

`wlrand` is a randomizer for the 1988 DOS game Wasteland.  The current version only randomizes transitions, i.e., the tiles that transport the player from one location to another.  `wlrand` produces a "new game" in the sense that even an experienced player must explore the game to discover its various locations.

## Quick Start

The below instructions assume wasteland is "installed" at `/usr/local/share/games/wasteland`.  Please adjust accordingly if this assumption is inaccurate.

Run the `wlrand rand` comamnd:
```
wlrand rand -p /usr/local/share/games/wasteland
```

Wasteland should now be randomized.  Verify by entering any non-shop location in a city (e.g., Highpool community center).

To restore your game to its pre-randomized state, use `wlrand restore`:
```
wlrand restore -p /usr/local/share/games/wasteland
```

## Scope

In its default mode, `wlrand` randomizes transitions that meet the following criteria:

* *Don't involve the world map.*  Transitions involving the world map, i.e., transitions into and out of cities, are left unmodified.  Specify `--world` to lift this restriction [\*].
* *Are not post-sewers.*  That is to say, locations within the Darwin base, Sleeper Base, Guardian Citadel, and Base Cochise are not considered during randomization.  The rationale for this restriction is 1) Several post-sewers transitions put a nascent party into an inescapable situation, either due to unavoidable combat or a lack of physical exit, and 2) The post-sewers part of Wasteland is much more linear than the parts before it and doesn't lend itself well to randomization.  Specify `--post-sewers` to lift this restriction [\*].
* *Are not shops.*  Stores, hospitals, and libraries are unaffected by `wlrand`.  This is simply a technical limitation.  Hopefully it will be removed in the future.
* *Are not intra-location transitions.*  That is to say, `wlrand` ignores transitions between locations that are logically connected.  For example, the transition from downtown west to downtown east is an intra-transition.  Another example is the transitions between the various floors of the Needles waste pit.  Specify `--auto-intra` and `--hard-intra` to lift this restriction.
* *Do not have the same parent.*  In other words, `wlrand` won't replace the Highpool-\>Cave transition with Highpool-\>Workshop.  Both of these transitions have the same parent (Highpool).  This restriction is just to make the game seem more "random".  Specify `--same-parent` to lift this restriction.

[\*] These options will almost certainly produce an unwinnable game.

## Building

Building requires two tools:

* [Go compiler](https://golang.org/dl/)
* `make`

To build, run one of the following invocations:
```
make build GOOS=linux   # Linux
make build GOOS=darwin  # MacOS
make build GOOS=windows # Windows
```

This produces a `wlrand` executable in the current directory.

## Bugs

Please report bugs using the wlrand issue tracker: <https://github.com/badvassal/wlrand/issues>

## To do

* Randomize loot bags.
* Randomize shops.
* Randomize NPCs?
* Produce a winnable game when `--world` is specified.
* Make the Spade's Casino basement escapable without a shovel or explosives.
* Consider making the Base Cochise levels easily escapable.

## Acknowledgements

Thank you to the authors of [Wasteland: The Definitive Deconstruction](https://wasteland.gamepedia.com/Category:Wasteland:_The_Definitive_Deconstruction) without whom this program would not be possible.

## License

[MIT License](https://opensource.org/licenses/MIT)
