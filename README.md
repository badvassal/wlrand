# wlrand

https://github.com/badvassal/wlrand

## Description

`wlrand` is a randomizer for the 1988 DOS game Wasteland.  The current version randomizes two aspects of Wasteland:

1. Transitions, i.e., the tiles that transport the player from one location to another
2. NPCs

`wlrand` produces a "new game" in the sense that even an experienced player must explore the game to discover its various locations.

## Quick Start

The below instructions assume wasteland is "installed" at `/usr/local/share/games/wasteland`.  Please adjust accordingly if this assumption is inaccurate.

Run the `wlrand rand` command:
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
3. Inventory

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

### Skill Class

The first step in NPC randomization is skill class selection.  A skill class has the following properties:
* a weight for each of the 35 skills in the game.
* minimum IQ requirement.
* set of armor types the character can start with.
* range of "armor points" per experience level.
* range of starting cash per experience level.

A skill class is selected at random from a predefined set.

### Attributes

`wlrand` uses the following procedure when randomizing an NPC's attributes.

1. Randomly select an "attribute class" from a predefined set.  An attribute class assigns weights to each of the seven attributes.  
2. Calculate the number of attribute points that the NPC starts with.  This number is calculated from a configurable range plus two points for each experience level beyond level one.  Use the following procedure to distribute the points among the seven attributes:
    1. Increase all attributes from 0 to 6, consuming the necessary number of points.
    2. Boost IQ to the character's skill class's minimum IQ requirement, consuming the necessary number of points.  Each time IQ is increased, its corresponding weight is multiplied by a reduction factor, making it less likely to receive points in the next step.
    3. Distribute the remaining points among the seven attributes, favoring those with greater weights.  Each time an attribute is increased, its corresponding weight is multiplied by a reduction factor, making it less likely to receive additional points.


The `--npc-attr-min` and `--npc-attr-max` command line options can be used to specify the range of attribute points.  `wlrand` selects a random number from within this range for each NPC.

### Skills

After calculating attributes, `wlrand` assigns skills.  The program uses the following procedure to distribute skill points:

1. Start with X skill points, where X is equal to the NPC's IQ.
2. Distribute skill points, favoring those skills that the skill class assigns a greater weight to.  The program follows the in-game rules when distributing skill points:
    1. Every skill has a minimum IQ requirement.
    2. The cost of a skill doubles each time it is improved.
3. Each time a skill is improved, its weight is reduced.
4. A skill will not be improved beyond a configurable maximum level.
5. If `wlrand` is unable to distribute any points, these leftovers are preserved as spendable skill points (`skp`).

The `--npc-skill-min` and `--npc-skill-max` command line options can be used to specify a range of extra skill points.  `wlrand` selects a random number from within this range for each NPC.  The selected value is the number of extra skill points the NPC gets.

The `--npc-learn-level-max` command line option specifies the maximum skill level that can be attained during this phase.  The default value is 2.

### Mastery

This phase is meant to simulate the natural improving of skills through use.  The number of "mastery points" is based on the NPC's experience level.  These points are used to improve the skills that were acquired in the "skills" phase.  It costs X mastery points to improve a skill to level X.  For example, it would take four mastery points to improve the Doctor skill from level three to four.  During this phase, skills are selected at random with no regard to the NPC's skill class.

The `--npc-mastery-min` and `--npc-mastery-max` command line options specify the range of mastery points that each NPC gets per level beyond level one.

### Inventory

An NPC's inventory is randomized to contain the following three types of items:
* Weapons
* Armor
* Accessories

#### Weapons

Weapons are selected at random based on the character's skill set.  First, `wlrand` isolates the character's top three weapon skills (or fewer if the character does not have three).  These skills are labeled "primary", "secondary", and "tertiary".  Next, the program randomly calculates a number of "weapon points" for each of these skills.  If the number of points is greater than 0, the character gets a corresponding weapon.  The greater the number of points, the better the weapon given.  A character is guaranteed to get a weapon for his primary skill, but may not get a secondary weapon, and is even less likely to get a tertiary weapon.

If a weapon is non-reusable (e.g., AT weapons), the character gets a random assortment of weapons from the appropriate category.  If the weapon uses clips, the character gets a random number of clips in addition to the weapon.

The `--npc-weapon-ppl-min` and `--npc-weapon-ppl-max` command line options specify the range of weapon points per level that each NPC gets.

The `--npc-weapon-count-min` and `--npc-weapon-count-max` command line options specify the range of counts of non-reusable items that an NPC gets.

The `--npc-weapon-clips-min` and `--npc-weapon-clips-max` command line options specify the range of clips of non-reusable items that an NPC gets.

#### Armor

Armor is selected at random from the skill class's set of allowed armor types.  The higher the experience level, the better the armor (on average).

#### Accessories

Accessories include the following items:
* Canteen
* Matches
* Crowbar
* Geiger counter
* Ropes
* Shovel
* TNT

All NPCs are equally likely to get each of the above items.

### Etc

* Experience level is randomly selected from a range (1 to 10 by default).  This range can be configured with the `--npc-level-min` and `--npc-level-max` command line options.
* Max/con is calculated using this (made up) formula: `20 + 2d8 + 2 * (explvl - 1)`.
* Starting cash is based on two things: 1) the NPC's skill class, and 2) the NPC's experience level.
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
