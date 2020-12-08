# terminus
A simple terminal game engine in Go (still WIP)

Built using [tcell](https://github.com/gdamore/tcell)

Terminus is a hobby project and it comes with no warranty or guarantee. I built it for myself, in order to create a flexible, easy-to-use, cross-platform engine for building games for the command line.

If you like it and find it useful, please feel free to use it. Keep in mind though that it was built for fun, and not necessarily to be highly performant or revolutionary in any way.

## Installing

Start by installing tcell: 

    go get https://github.com/gdamore/tcell

Then by installing terminus: 
    
    go get https://github.com/Sheep42/terminus

That's all you need to get started :)

## Running the Examples

Change your directory to the example that you want to run. Then run **go build**. Finally, run the executable, it will have the same name as the example directory. 

    $ cd terminus/examples/collision/
    $ go build
    $ ./collision

<!-- ## Understanding the Engine -->
<!-- TODO: Content -->