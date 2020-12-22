# terminus
A simple terminal game engine in Go (still WIP)

Built using [tcell](https://github.com/gdamore/tcell)

Terminus is a hobby project and it comes with no warranty or guarantee. I built it for myself, in order to create a flexible, easy-to-use, cross-platform engine for building games for the command line.

If you like it and find it useful, please feel free to use it. Keep in mind though that it was built for fun, and not necessarily to be highly performant or revolutionary in any way.

![Image of snake game](examples/images/snake.gif?raw=true "Snake Example")

![Image of scenes example](examples/images/scenes.gif?raw=true "Scenes Example")

![Image of collision example](examples/images/collision.gif?raw=true "Collision Example")

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

### Game

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

* `width, height (int, int)`

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

### Scene

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
        NewCustomScene(g),
        NewOtherCustomScene(g),
    }

    // ...

    for _, iScene := ss {

        s := iScene.GetScene()
        s.Add(NewEntity(2,2))

    }
```

#### `Add`

**Params**

* `entity IEntity`

Attach the specified `Entity` to the `Scene`. Once an `Entity` is added to a `Scene`, that Entity will be rendered by the `Scene`'s `Draw` function.

#### `Remove`

**Params**

#### `Game`

#### `Entities`

#### `GetScene`

#### `SetRedraw`

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


### Entity

    Coming Soon

### EntityGroup

    Coming Soon

### Text

    Coming Soon

### StateManager

    Coming Soon

### State

    Coming Soon
