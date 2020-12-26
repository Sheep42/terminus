# terminus
A simple terminal game engine in Go (still WIP)

Built using [tcell](https://github.com/gdamore/tcell)

Terminus is a hobby project and it comes with no warranty or guarantee. I built it for myself, in order to create a flexible, easy-to-use, cross-platform engine for building games for the command line.

If you like it and find it useful, please feel free to use it. Keep in mind though that it was built for fun, and not necessarily to be highly performant or revolutionary in any way.

![Image of snake game](examples/images/snake.gif?raw=true "Snake Example")

![Image of scenes example](examples/images/scenes.gif?raw=true "Scenes Example")

![Image of collision example](examples/images/collision.gif?raw=true "Collision Example")

# Contents

- [Installing](#installing)
- [Examples](#running-the-examples)
- [Understanding the Engine](#understanding-the-engine)
- [Types](#game)
    - [Game](#game)
    - [Scene](#scene)
    - [Entity](#entity)
    - [EntityGroup](#entitygroup)
    - [Text](#text-1)
    - [StateManager](#statemanager)
    - [State](#state)

## Installing

```bash
$ go get github.com/Sheep42/terminus
```

That's all you need to get started :)

## Running the Examples

Change your directory to the example that you want to run. Then run **go build**. Finally, run the executable, it will have the same name as the example directory. 

```bash
$ cd terminus/examples/collision/
$ go build
$ ./collision
```

### What Are the Examples?

In order of appearance:

### Collision

This is a simple demonstration of how collision can be implemented.

`Moveable` extends `Entity` through composition, and  when input is detected, we can use `Entity`'s built-in `CheckDir` function to detect a collision with any `Entity` in `collidables` before setting the position. This is one of several possible approaches, `Entity`'s built-in utility functions will be explained in more detail below.

### Entity Groups

This is a simple demonstration of how to use an `EntityGroup`. In this case I extend `EntityGroup` in order to override `Update` and move the group as a whole. I also demonstrate moving a single `Entity` within the group.

### Hello World

This example customizes nothing, and simply creates a `Scene` to which the "Hello World" is added.

### Logging

This example demonstrates how you can use logging for print debugging during development, outputting useful information when issues arise, or throwing fatal errors and exiting gameplay.

***This example is designed to crash***. Errors are first printed, and the final error is an intentional fatal error. Errors are logged to the `terminus.log` which will be created inside of the directory where you run the executable.

### Moveable Entity

This example showcases how to implement custom `Entity` logic through a simple `Moveable` entity. It is a simplified version of the Collision example. This is simply an `Entity` which has been extended in order to listen for input, and change its position according to the key pressed. There is some additional logic in place for screen wrapping.

### Scenes

This example showcases how to implement custom `Scene` logic. This is a simple demonstration with different colored screens and text which can be toggled between using the 'z' and 'x' keys.

### Snake

This example is a simple implementation of Snake. This is meant to showcase all of the aspects that might go into a full game built with Terminus. 

While Snake is a relatively simple example, I did manage to make use of `State`s and extended `Scene`'s functionality. 

The game keeps and displays score, increases the snake's speed as the score goes up, presents the snake and food in different colors, and resets when pressing enter from the Game Over state. 

### States

This example is a clone of the Collision example, but it has been expanded to include a pause state which can be toggled by pressing 'p'.

This is meant to be a very simple demonstration of the use of `State` and `StateManager` within the context of a game.

### Text

This example showcases some simple examples of how `Text` can be manipulated and extended in your game.

## Understanding the Engine

### General

I tried to design Terminus to be simple to use without being too restrictive. The running theme you will notice is that things are heavily Interface-based and you'll be using a lot of composition in order to customize different elements. 

It is easy to create some awkward designs this way, so it is worth thinking through an approach before committing too hard to it.

Outside of that, there really isn't that much to learn, and you should be able to be making a game relatively quickly.

### Simple Example

The most basic example is the Hello World example included in the examples directory. The full source is contained in main.go, and is duplicated below.

```go
package main

// import terminus - I abbreviate as 't'
import (
    t "github.com/Sheep42/terminus"
)

func main() {

    // Create the Game
    g := t.NewGame()

    // Create a Scene
    s := t.NewSceneCustom(g, t.Black, t.Gray)

    // Add some text

    // override scene color
    s.Add(t.NewText(2, 2, "Press ESC to quit", t.White, t.Black))

    // Inherit scene color
    s.Add(t.NewText(5, 5, "Hello World"))

    // g.Init takes a slice of IScenes
    ss := []t.IScene{s}

    // Init the Game
    g.Init(ss)

    // Start the Game
    g.Start()

}
```

---

## Game

`Game` is the main component of the engine. There is no interface to allow for extension of `Game` itself.

A `Game` must always be created and started in order to use terminus. The most basic game would contain the following lines in `main`.

```go
game := t.NewGame() // Create a Game

scene := t.NewScene(game) // Create a Scene

// Init() requires a slice containing all of the 
// Game's Scenes
scenes := []t.IScene{scene}

game.Init(scenes) // Run Game's Init function

game.Start() // Start the Game
```

Scenes are stored as a slice in `Game`, and referenced by an internal index which always points to the current active scene. The first `Game` scene by default is always the one in `scenes[0]`. 

#### **Functions**

---

#### `NewGame`

Constructor function: create a new `Game` &ndash; Generally you should only need to call this once inside of `main`.

```go
game := t.NewGame()
```

#### `Init`

**Params**

* `scenes IScene[]` &ndash; A slice of `Scene`s to load for use in the `Game` 

Initialize the `Game`, set up the Logger, call the `Setup` function on every `Scene` in `scenes`, and call the `Init` function of the first `Scene` &ndash; Generally you should only need to call this once inside of `main`.

**This function must be invoked before `game.Start()`**

    game.Init(scenes) // Assume scenes is of type IScene[]

#### `Start`

Run the actual game loop. This calls the `Update` and `Draw` functions of the currently active scene, and listens for input changes &ndash; Generally you should only need to call this once inside of `main`.

**This function should always be at the end of `main`**

```go
game.Start()
```

#### `NextScene`

Increment the Scene index by one and run the `Init` function of the new scene after doing so.

Stop if the index would exceed the last index of `scenes`

```go
game.NextScene()
```

#### `PrevScene`

Derement the Scene index by one and run the `Init` function of the new scene after doing so.

Stop if the index would be 0

```go
game.PrevScene()
```

#### `SetScene`

**Params**

* `index int` &ndash; The scene index to switch to.

Set the Scene index to a specific number and run the `Init` function of the new scene after doing so.

Fall back to 0 if index exceeds the number of `Scenes` or if index is negative

```go
game.SetScene(3) 
```

#### `ExitKey`

**Return**

* `exitKey tcell.Key`

Fetch the `Game`'s current exit key

```go
key := game.ExitKey()
```

#### `SetExitKey`

**Params**

* `exitKey tcell.Key` &ndash; For convenience, I have migrated some major Keys from tcell to terminus in the form of constants, but there are many different keys you can choose from in tcell

Set the `Game`'s exit key

**Default exit key value is ESC**

```go
game.SetExitKey(tcell.KeyCtrlC)
```

#### `GetFPS`

**Return**

* `fps float64`

Fetch the `Game`'s target FPS.

```go
fps := game.GetFPS() 
```

#### `SetFPS`

**Params**

* `fps float64` &ndash; The target FPS number

Set the `Game`'s target FPS.

**Default FPS is 60**

**Call this before `game.Init`**

```go
game.SetFPS(30)
game.Init()
```

#### `GetLogger`

**Return**

* `logger *log.Logger`

Fetch the `Game`'s Logger instance &ndash; Once you have it, use it like any log.Logger

```go
l := game.GetLogger()
l.Println("Logger Out")
```

#### `SetLogFileName`

**Params**

* `filename string` &ndash; The log file name 

Set a custom log file name for
logger output. Default value is 'terminus.log'

**Call this before `game.Init`**

```go
game.SetLogFileName("custom_log.log")
game.Init()
```

#### `Input`

**Return**

* `input *tcell.EventKey` &ndash; The engine constantly listens for input, if there is none the return value will be `nil`

Fetch the current `Game`'s input data.  

```go
i := game.Input()
```

#### `ScreenSize`

**Return** 

* `width int, height int`

Fetch the current screen size

```go
w, h := game.ScreenSize()
```

#### `CurrentScene`

**Return** 

* `scene *Scene`

Fetch the current `Scene`

```go
scene := game.CurrentScene()
```

---


## Scene

`Scene`s are used to render content to a `Game` screen, and a `Scene` is usually the first thing that you will add to a game.

`Scene` can be extended via composition, in order to override `Setup`, `Init`, `Update`, or `Draw` with custom logic.

#### **Functions**

---

#### `NewScene`

**Params**

* `game *Game`

**Return**

* `scene *Scene`

Creates a new `Scene` to be used in a `Game`. Uses default screen colors.

```go
s := t.NewScene(g)
```

#### `NewSceneCustom`

**Params**

* `game *Game` 
* `foreground tcell.Color` 
* `background tcell.Color`

**Return**

* `scene *Scene`

Creates a new `Scene` with custom foreground and background colors. Foreground affects `Entities`, background is the screen background color.

```go
s := t.NewSceneCustom(g, t.Black, t.Gray)
```

#### `Setup`

Fires **only once** during `Game`'s `Init` function for every `Scene` in `scenes`. Keep in mind this happens at the launch of the `Game`, so this will run before any `Scene`s are rendered, and it runs for all `Scene`s at once.

This is a good place to do scene one-time tasks, such as adding `Entities`, or property initializations that only need to happen once.

This function can be overridden in order to customize your `Scene`.

#### `Init`

Fires just before the `Scene` is first rendered. 

This runs for the first `Scene` in `scenes` when the `Game`'s `Init` function runs. It is fired after `Setup` is completed for all `Scenes`.

You should do `Scene` setup actions here that cannot be done in `Setup` or must be run each time the `Scene` is re-entered. For example, reset `Scene` data, or reload removed `Entities`.

In the **Scenes** example, I used `Init` to center the `Scene` text, because screen size is not yet available in `Setup`.

This function can be overridden in order to customize your `Scene`.

#### `Update`

**Params**

* `delta float64` &ndash; The time elapsed since the last pass through the game loop.

Fires on each pass of the game loop. You can use `delta` to implement timers.

This is where the meat of your custom `Scene` logic should go. You should add any custom interactivity logic, movement, etc to your overridden `Update` function.

If you are using a `StateManager` in your `Scene` you will most likely only be calling `StateManager`'s `Update` inside of here. 

This function can be overridden in order to customize your `Scene`.


#### `Draw`

`Draw` is fired after the `Scene` updates, on each pass through
the game loop.

A design decision worth noting is that my implementation of `Draw` only redraws the screen when `scene.redraw` has been flagged. Several actions in the engine flag a scene for a redraw. This is to work around screen flicker on each update in Windows terminals.

This function can be overridden in order to customize your `Scene`. **However**, if you don't call `myScene.Scene.Draw()` in your overridden function, you will need to render all `Scene` children and hanlde scene refreshing on your own.

#### `Entities`

**Return**

* `entities []IEntity`

Get the `Scene`'s slice of `Entities`. 

This will return a slice of type `IEntity`. In order to get access to the underlying `Entity`, you can call `iEntity.GetEntity()`.

```go
s := t.NewSceneCustom(g, t.Black, t.Gray)

// ...

for _, iEntity := range s.Entities() {

    e := iEntity.GetEntity()
    e.SetPosition( e.GetX(), e.GetY() + 1)

}
```

#### `GetScene`

**Return**

* `scene *Scene`

Returns the `Scene`. This function is exposed via `IScene`, and can be used to get the actual `Scene` for manipulation given a generalized  `IScene`. 

```go
g := NewGame()

ss := []IScene{
    t.NewCustomScene(g),
    t.NewOtherCustomScene(g),
}

// ...

for _, iScene := ss {

    s := iScene.GetScene()
    s.Add(t.NewEntity(2, 2))

}
```

#### `Add`

**Params**

* `entity IEntity`

Attach the specified `Entity` to the `Scene`. Once an `Entity` is added to a `Scene`, that Entity will be rendered by the `Scene`'s `Draw` function.

A single `Entity` **can** be added to multiple `Scene`s

**This function flags the `Scene` for redraw**

#### `Remove`

**Params**

* `entity IEntity`

Remove the specified `Entity` from the `Scene`. Once an `Entity` is removed from a `Scene`, that Entity will no longer be rendered by the `Scene`'s `Draw` function.

**This function flags the `Scene` for redraw**

#### `Game`

**Return**

* `game *Game`

Returns the `Scene`'s game. Useful if you need to reference `Game` in the context of a `Scene`.

```go
g := scene.Game()
```

#### `SetRedraw`

**Params**

* `redraw bool`

Allows you to tell the `Scene` to redraw (true) or not (false) on the next frame.

```go
// Do something...

// Force a scene redraw
scene.SetRedraw( true )
```

#### **Custom Scenes**

---

Below is a very simple skeleton of a custom `Scene` through composition. You will find more detailed real-world examples by reading through the examples included in the package. This `Scene` adds text to the screen and centers it.

```go
package main

import (
    t "github.com/Sheep42/terminus"
    "github.com/gdamore/tcell"
)

type CustomScene struct {
    *t.Scene
    title *t.Text
}

func NewCustomScene(g *t.Game, fg, bg tcell.Color, title string) *CustomScene {

    cs := &CustomScene{
        // NewSceneCustom is like NewScene, but allows
        // custom foreground and background colors
        t.NewSceneCustom(g, fg, bg),
        t.NewText(0, 0, title),
    }

    return cs

}

func (cs *CustomScene) Setup() {

    cs.Scene.Setup() // super
    cs.Add(cs.title)

}

func (cs *CustomScene) Init() {

    cs.Scene.Init() // super

    game := cs.Game()

    screenWidth, screenHeight := game.ScreenSize()
    textWidth, textHeight := cs.title.GetDimensions()

    cs.title.SetPosition(screenWidth/2-textWidth/2, screenHeight/2-textHeight/2)

}

func (cs *CustomScene) Update(delta float64) {

    cs.Scene.Update() // super

}
```

---


## Entity

An `Entity` is used to represent any general object that will exist in a `Scene`. 

`Entities` must have an x position and a y position, and may have a sprite. `Text` and `EntityGroup`s are extensions of `Entity`, and can be manipulated in much the same way. 

`Entities` can be added to a `Scene` or removed from a `Scene`. An `Entity` will be rendered to the screen once added to a `Scene` during that `Scene`'s `Draw`. 

#### **Functions**

---

#### `NewEntity`

Takes an x position and a y position and creates an `Entity` without a sprite. This can be used if you plan to extend `Entity` and it does not make sense for the root `Entity` to have a sprite, such as with `EntityGroup`. This can also be used if you plan on assigning a sprite at a later time.

**Params**

* `x int`
* `y int`

```go
e := NewEntity(5, 5)
```

#### `NewSpriteEntity`

Takes an x position, a y position, and a `rune` to be used as a visual representation, and creates an `Entity`. 

Colors are optional - foreground & background required if used.

**Params**

* `x int`
* `y int`
* `sprite rune`
* `fg tcell.Color` (optional)
* `bg tcell.Color` (optional)

```go
spriteE := t.NewSpriteEntity(5, 5, '#')
colorE := t.NewSpriteEntity(5, 5, '#',  tcell.ColorBlack, tcell.ColorGray)
```

#### `Init`

Fires duting `game.Init`, if the `Entity` has been added to a `Scene` at that point.

If you add an `Entity` to a `Scene` later, you must call `Init` manually. 

By default `Init` does nothing and is not needed, but it can be overridden for your own custom implementations.

#### `Update`

**Params**

* `delta float64` &ndash; The time elapsed since the last pass through the game loop.

Fires after the `Scene` `Update` on each pass through the game loop. 

This can be overridden in order to customize an `Entity`.

#### `Draw`

Fires during scene redraw, and is responsible for rendering the `Entity`.

Draw and can be overridden in order to extend or replace functionality. But be careful, overridding this without calling `customEntity.Entity.Draw()` means that you will need to handle rendering the `Entity` on your own.

#### `GetEntity`

**Return**

* `entity *Entity`

Returns the `Entity`. Generally used to get the actual `Entity` from an `IEntity`.

```go
entities := []IEntity{
    t.NewEntity(5, 5),
    t.NewSpriteEntity(3, 3, '@'),
}

for _, e := range entities {

    e.GetEntity().SetScene(scene)

}
```

#### `SetScene`

**Params**

* `scene *Scene`

Sets the `Entity`'s `Scene` and `Game`. 

`Scene.Add()` uses this function to update the `Entity`'s `Scene` value when it is added.

#### `GetScene`

**Return**

* `scene *Scene`

Gets the `Scene` that the `Entity` is associated with.

#### `GetGame`

**Return**

* `game *Game`

Gets the Game the the Entity is associated with.

#### `GetX`

**Return**

* `x int`

Gets the x position of the `Entity`

#### `GetY`

**Return**

* `y int`

Gets the y position of the `Entity`

#### `SetPosition`

**Params**

* `x int`
* `y int`

Sets the `Entity`'s x and y position simultaneously.

```go
e.SetPosition(2, 5)
```

**This function flags the `Scene` for redraw**

#### `GetPosition`

**Return**

* `x int, y int`

Gets the `Entity`'s x and y position simultaneously.

```go
x, y := e.GetPosition()
```

#### `SetSprite`

**Params**

* `sprite rune`

Sets the `Entity`'s sprite. 

**This function flags the `Scene` for redraw**

#### `GetSprite`

**Return**

* `sprite rune`

Returns the rune that visually represents the `Entity`.

#### `SetColor`

**Params**

* `foreground tcell.Color`
* `background tcell.Color`

Changes the `Entity`'s style foreground and background colors.

**This function flags the `Scene` for redraw**

#### `Overlaps`

**Params**

* `target IEntity`

**Return**

* `overlaps bool`

Checks if the `Entity` currently overlaps the target `Entity`. Overlaps is simply a check if two `Entities` occupy the same coordinates.

```go
// Checks if e is currently overlapping e2
overlaps := e.Overlaps(e2)
```

#### `OverlapsPoint`

**Params**

* `x int` - The x position to check
* `y int` - The y position to check

**Return**

* `overlaps bool`

Checks if the `Entity` overlaps the specified screen point.

```go
// Checks if e is currently overlapping (2, 5)
overlaps := e.OverlapsPoint(2, 5)
```

#### `CheckDir`

**Params**

* `axis rune` - 'x' or 'y'
* `distance int` - distance of 0 is the same as overlapping.
* `point int` - The point to check on the specified axis 

**Return**

* `isDistanceAway bool`

Checks if the `Entity` is the specified distance away from the target point.

```go

eX, eY := e.GetPosition()
e2X, e2Y := e2.GetPosition()

changeX := 1
changeY := 0

// check if changing position of e by changeX, changeY would result in a collision
collided := e.CheckDir('x', changeX, e2X) && e.CheckDir('y', changeY, e2Y)

```

#### `IsLeftOf`

**Params**

* `target IEntity`

**Return**

* `isLeft bool`

Checks if the `Entity` is directly to the left of the target `Entity`

Note: This function checks if `Entity` is exactly 1 unit in the specified direction.

```go
// checks if e is directly left of e2
isLeft := e.IsLeftOf(e2)
```

#### `IsRightOf`

**Params**

* `target IEntity`

**Return**

* `isRight bool`

Checks if the `Entity` is directly to the right of the target `Entity`

Note: This function checks if `Entity` is exactly 1 unit in the specified direction.

#### `IsAbove`

**Params**

* `target IEntity`

**Return**

* `isAbove bool`

Checks if the `Entity` is directly above the target `Entity`

Note: This function checks if `Entity` is exactly 1 unit in the specified direction.

#### `IsBelow`

**Params**

* `target IEntity`

**Return**

* `isBelow bool`

Checks if the `Entity` is directly below the target `Entity`

Note: This function checks if `Entity` is exactly 1 unit in the specified direction.

---

#### Custom Entities

`Entities` can be extended through composition. Below is the simplest version of the `Moveable` defined in several of the example projects. `Moveable` extends `Entity`'s `Update` to process input and move the `Entity` on screen, including screen wrapping.


```go
package main

import (
	t "github.com/Sheep42/terminus"
)

type Moveable struct {
	*t.Entity
}

func NewMoveable(x, y int, sprite rune) *Moveable {

	m := &Moveable{
		t.NewSpriteEntity(x, y, sprite),
	}

	return m
}

func (m *Moveable) Update(delta float64) {

	// super
	m.Entity.Update(delta)

	game := m.GetGame()
	input := game.Input()

	// Screen Wrap
	gw, gh := game.ScreenSize()

	if m.GetX() >= gw {
		m.SetPosition(0, m.GetY())
	} else if m.GetX() < 0 {
		m.SetPosition(gw-1, m.GetY())
	}

	if m.GetY() >= gh {
		m.SetPosition(m.GetX(), 0)
	} else if m.GetY() < 0 {
		m.SetPosition(m.GetX(), gh-1)
	}

	// Moveable movement
	if nil != input {

		if t.KeyLeft == input.Key() {

			m.SetPosition(m.GetX()-1, m.GetY())

		} else if t.KeyRight == input.Key() {

			m.SetPosition(m.GetX()+1, m.GetY())

		} else if t.KeyUp == input.Key() {

			m.SetPosition(m.GetX(), m.GetY()-1)

		} else if t.KeyDown == input.Key() {

			m.SetPosition(m.GetX(), m.GetY()+1)

		}

	}

}

```

---


## EntityGroup

`EntityGroup`s are a simple extension of `Entity` which allow for grouping of many `Entities` into the context of a single `Entity`. 

`EntityGroups` have some special properties:

* `Entities` added to an `EntityGroup` are postioned relative to the `EntityGroup`, not the screen. This means, if `e := NewEntity(0, 0)` and `eg := NewEntityGroup(5, 5, 10, 10, []IEntity{e})`,  `e`'s screen position would be (5, 5), while it's relative position would be (0, 0).

* `Entities` whose coordinates exist outside of the `EntityGroup`'s bounds will not be rendered to the screen. However, the `Entities` will still exist and can still be manipulated.

* `EntityGroup`s will move as a single `Entity`, moving all `Entities` contained within. So, moving an `EnitityGroup` 1 unit to the right will move all of that `EntityGroup`'s children 1 unit to the right as well. Individual `Entities` can be targeted and moved within the `EntityGroup` as well, if needed.

* `Entities` within an `EntityGroup` will inherit their color from the `EntityGroup`. At the moment, you cannot set individual `Entity` colors in an `EntityGroup`.

#### **Functions**

---

As `EntityGroup` is an extension of `Entity` it shares all `Entity`'s functions. You can reference above for details on the functions that carry over. 

I'll detail `EntityGroup` specific or overridden functions below.

`NewEntityGroup`

**Params**

* `x int`
* `y int`
* `width int`
* `height int`
* `entities []IEntity`
* `fg tcell.Color` - optional 
* `bg tcell.Color` - optional

**Return**

* `entityGroup *EntityGroup`

Creates a new `EntityGroup` with coordinates (`x`, `y`) of specified `width` and `height`. `entities` should contain the `Entities` to be grouped.

```go
e := t.NewSpriteEntity(0, 0, '#')
e2 := t.NewCustomEntity(0, 0, '#')

eg := t.NewEntityGroup(5, 5, 10, 10, []IEntity{e, e2})
eg2 := t.NewEntityGroup(20, 20, 10, 10, []IEntity{e, e2}, t.Red, t.Blue) 
```

`Init`

Invokes `eg.Entity.Init()`. Can be overridden for custom functionality.

`Update`

**Params**

* `delta float`

Invokes `eg.Entity.Update()`. Can be overridden for custom functionality.

`Draw`

Does not invoke `Entity`'s `Draw` function. 

Loops through `entities` and renders the `Entities` in the group to the screen. This can be overridden for custom functionality, but doing so without calling `EntityGroup.Draw()` means you will need to handle rendering on your own.

`SetScene`

**Params**

* `scene *Scene`

Sets the `Scene` for the `EntityGroup`, but also invokes `SetScene` for each child `Entity`, passing in `scene`. 

`GetEntity`

**Return**

* `entity *Entity`

Gets the `Entity` that represents the `EntityGroup` as a whole. (The positioning `Entity`)

`GetEntities`

**Return**

* `entities []IEntity`

Gets the slice of `Entities` within the `EntityGroup`. Since `entities` is of type `[]IEntity`, you'll need to call `GetEntity()` to operate on specific types of `Entities`. 

```go
entities := eg.GetEntities()

for _, iEntity := entities {
    
    e := iEntity.GetEntity()
    
    e.SetPosition(2, 2)

}
```

`SetWidth`

**Params**

* `width int`

Sets the `EntityGroup`'s width

**This function flags the `Scene` for redraw**

`SetHeight`

**Params**

* `height int`

Sets the `EntityGroup`'s height

**This function flags the `Scene` for redraw**

`GetDimensions`

**Return**

* `width int, height int`

Returns the width and height of the `EntityGroup`

`SetEntities`

**Params**

* `entities []IEntity`

Sets the `EntityGroup`'s list of entities

**This function flags the `Scene` for redraw**

---

## Text

`Text` is an extension of `EntityGroup`. As such, `Text` inherits all of `Entity`'s functionality, as well as `EntityGroup`'s.

Below are the overridden or unique functions.

#### **Functions**

---

`NewText`

**Params**

* `x int`
* `y int`
* `text string`
* `fg tcell.Color` - optional
* `bg tcell.Color` - optional

Takes an x position, y position, and textvalue and creates a new `Text` on the screen. 

If colors are passed, fg & bg are required.

```go
t.NewText(5, 5, "Hello World")
t.NewText(0, 0, "Press ESC to quit", t.Blue, t.Black)
```

`Update`

**Params**

* `delta float`

Invokes `text.EntityGroup.Update()`. Can be overridden for custom functionality.

`ToEntities`

**Params**

* `text string`

**Return**

* `entities []IEntity`

Converts the given string of text into a slice of `IEntities`.

`SetText`

**Params**

* `newText string`

Sets the `text` value of the `Text`.

`GetText`

**Return**

* `text string`

Gets the string value of `text`.

`GetEntityGroup`

**Return**

* `entityGroup *EntityGroup`

Gets the underlying `EntityGroup` behind the `Text`.

---

## StateManager

`StateManager` is a simple state machine that should suffice for most simple games as is. However, it can be extended via composition if desired.

A `StateManager` can be used with any derivative of `Scene` or `Entity`.

Generally, when using a `StateManager`, you'll want to override `Scene` or `Entity` `Update` function, and call `sm.Update(delta)` from there. See the states or snake examples for more examples of basic state management.

#### Functions 

---

`NewStateManager`

**Params**

* `defaultState IState`

**Return**

* `stateManager *StateManager`

Creates a new `StateManager` and sets the default `State` to `defaultState`.

`ChangeState`

**Params**

* `state IState`

Changes the current `State` of the `StateManager` to `state`.

`BackToDefault`

Changes the current `State` of the `StateManager` back to the default `State`.

`BackToPrevious`

Changes the current `State` of the `StateManager` back to the previous `State`.

`Update`

**Params**

* `delta float64`

On the first pass, sets the `StateManager`'s current `State` to the default `State`. 

On every pass, invokes current `State`'s `Tick` function.

---

## State

`State` is an abstract struct meant to be extended for use with any `Scene` or `Entity`, and to be managed by a `StateManager`.

See the states or snake examples for more examples of custom `State`s.

#### Functions

---

`NewState`

**Return**

* `state *State`

Creates a new State to be used by a `StateManager`.

`OnEnter`

Fired every time a `State` is entered. (After a call to `stateManager.ChangeState()`)

`OnExit`

Fired every time a `State` is exited. (After a call to `stateManager.ChangeState()`)

`Tick`

**Params**

* `delta float64`

Fired on every pass through `stateManager.Update()`, when the `State` is `StateManager`'s current `State`.