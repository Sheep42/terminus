# terminus
A simple terminal game engine in Go (still WIP)

Built using [tcell](https://github.com/gdamore/tcell)

Terminus is a hobby project and it comes with no warranty or guarantee. I built it for myself, in order to create a flexible, easy-to-use, cross-platform engine for building games for the command line.

If you like it and find it useful, please feel free to use it. Keep in mind though that it was built for fun, and not necessarily to be highly performant or revolutionary in any way.

## Installing
    
    $ go get github.com/Sheep42/terminus

That's all you need to get started :)

## Running the Examples

Change your directory to the example that you want to run. Then run **go build**. Finally, run the executable, it will have the same name as the example directory. 

    $ cd terminus/examples/collision/
    $ go build
    $ ./collision

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


### Game

    Coming Soon

### Scene

    Coming Soon

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
