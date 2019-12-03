# wlrand

https://github.com/badvassal/wlrand

## Description

`wlrand` is a randomizer for the 1988 DOS game Wasteland.  The current version randomizes two aspects of Wasteland:

1. Transitions, i.e., the tiles that transport the player from one location to another
2. NPCs

`wlrand` produces a "new game" in the sense that even an experienced player must explore the game to discover its various locations.

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

### Transitions

In its default mode, `wlrand` randomizes transitions that meet the following criteria:

* *Don't involve the world map.*  Transitions involving the world map, i.e., transitions into and out of cities, are left unmodified.  Specify `--world` to lift this restriction [\*].
* *Are not post-sewers.*  That is to say, locations within the Darwin base, Sleeper Base, Guardian Citadel, and Base Cochise are not considered during randomization.  The rationale for this restriction is 1) Several post-sewers transitions put a nascent party into an inescapable situation, either due to unavoidable combat or a lack of physical exit, and 2) The post-sewers part of Wasteland is much more linear than the parts before it and doesn't lend itself well to randomization.  Specify `--post-sewers` to lift this restriction [\*].
* *Are not shops.*  Stores, hospitals, and libraries are unaffected by `wlrand`.  This is simply a technical limitation.  Hopefully it will be removed in the future.
* *Are not intra-location transitions.*  That is to say, `wlrand` ignores transitions between locations that are logically connected.  For example, the transition from downtown west to downtown east is an intra-transition.  Another example is the transitions between the various floors of the Needles waste pit.  Specify `--auto-intra` and `--hard-intra` to lift this restriction.
* *Do not have the same parent.*  In other words, `wlrand` won't replace the Highpool-\>Cave transition with Highpool-\>Workshop.  Both of these transitions have the same parent (Highpool).  This restriction is just to make the game seem more "random".  Specify `--same-parent` to lift this restriction.

[\*] These options will almost certainly produce an unwinnable game.

### NPCs

All NPCs get randomized.  `wlrand` randomizes the following aspects of each NPC:

1. Attributes
2. Skill list

Inventories are currently untouched.

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

## NPC Details

### Attributes

`wlrand` uses the following procedure when randomizing an NPC's attributes.

1. Randomly select an "attribute class" from a predefined set.  An attribute class assigns weights to each of the seven attributes.  
2. Roll 4d6 for each attribute, discarding the die with the lowest value.  These sums are assigned according to the attribute class weights.  The highest value is assigned to the attribute with the greatest weight.
3. Distribute extra attribute points.  The NPC gets two extra points for each experience level greater than one.  An extra point can be applied to any of the seven attributes, but it is more likely to be assigned to attributes with greater weights.

The `--npc-attr-min` and `--npc-attr-max` command line options can be used to specify a range of extra attribute points.  `wlrand` selects a random number from within this range for each NPC.  The selected value is the number of extra attribute points the NPC gets.

### Skills

After calculating attributes, `wlrand` assigns skills.  The program uses the following procedure to distribute skill points:

1. Randomly select a "skill class" from a predefined set.  A skill class assigns weights to each of the 35 skills in the game.
2. Start with X skill points, where X is equal to the NPC's IQ.
3. Distribute skill points, favoring those skills that the skill class assigns a greater weight to.  The program follows the in-game rules when distributing skill points:
    1. Every skill has a minimum IQ requirement.
    2. The cost of a skill doubles each time it is improved.
4. If `wlrand` is unable to distribute any points, these leftovers are preserved as spendable skill points (`skp`).

The `--npc-skill-min` and `--npc-skill-max` command line options can be used to specify a range of extra skill points.  `wlrand` selects a random number from within this range for each NPC.  The selected value is the number of extra skill points the NPC gets.

### Mastery

This phase is meant to simulate the natural improving of skills through use.  The number of "mastery points" is based on the NPC's experience level.  These points are used to improve the skills that were acquired in the "skills" phase.  It costs X mastery points to improve a skill to level X.  For example, it would take four mastery points to improve the Doctor skill from level three to four.  During this phase, skills are selected at random with no regard to the NPC's skill class.

The `--npc-mastery-min` and `--npc-mastery-max` command line options specify the range of mastery points that each NPC gets per level beyond level one.

### Etc

* Experience level is randomly selected from a range (1 to 10 by default).  This range can be configured with the `--npc-level-min` and `--npc-level-max` command line options.
* Max/con is calculated using this (made up) formula: `20 + 2d8 + 2 * (explvl - 1)`.
* The `npc` flag is cleared for all NPCs.  This allows the player to control who they attack and their shooting style, and it prevents them from disobeying any orders.

## Bugs

Please report bugs using the wlrand issue tracker: <https://github.com/badvassal/wlrand/issues>

## To do

* Produce a winnable game when `--world` is specified.
* Make the Spade's Casino basement escapable without a shovel or explosives.
* Randomize shops.
* Randomize loot bags.
* Consider making the Base Cochise levels easily escapable.

## Acknowledgements

Thank you to the authors of [Wasteland: The Definitive Deconstruction](https://wasteland.gamepedia.com/Category:Wasteland:_The_Definitive_Deconstruction) without whom this program would not be possible.

## License

[MIT License](https://opensource.org/licenses/MIT)
