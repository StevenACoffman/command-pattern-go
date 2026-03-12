### Command pattern in Go, for building CLIs

From [Alternative Command Design Pattern in Go](https://cgiacomi.com/posts/2020-01-25-alternative-command-pattern-go/)

---

In my previous post regarding the command design pattern in Go I explored the way one would structure the code to implement this classic design pattern.

One thing that one would immediately notice about the implementation in Go is the lack of type inheritance and the fact that interfaces are not explicitly implemented either, this is of course due to how Go is designed. The code implementation of my previous post is based on the original implementation in C++ as can be seen in the book by the gang of four. So even though I believe the implementation in Go is correct it feels like it’s just not a perfect fit.

But what if there was a more Go oriented implementation, one that felt like it was really meant for Go? A more idiomatic implementation.

### Go mod
Well I believe I have found one, I did not come up with it, and at first I was surprised by it. I was looking into go mod and how the command is implemented. One great thing about Go, is that you have access to the source code, so all I did was go digging in the github repo.

While I was digging around I stumbled upon [this file](https://github.com/golang/go/blob/master/src/cmd/go/internal/base/base.go) which contains the base type for the go command. The code below is taken directly form the Go source code.
```go
// A Command is an implementation of a go command
// like go build or go fix.
type Command struct {
// Run runs the command.
// The args are the arguments after the command name.
Run func(cmd *Command, args []string)

    // UsageLine is the one-line usage message.
    // The words between "go" and the first flag or argument in the line are taken to be the command name.
    UsageLine string

    // Short is the short description shown in the 'go help' output.
    Short string

    // Long is the long message shown in the 'go help <this-command>' output.
    Long string

    // Flag is a set of flags specific to this command.
    Flag flag.FlagSet

    // CustomFlags indicates that the command will do its own
    // flag parsing.
    CustomFlags bool

    // Commands lists the available commands and help topics.
    // The order here is the order in which they are printed by 'go help'.
    // Note that subcommands are in general best avoided.
    Commands []*Command
}
```
If we omit the last line `Commands []*Command` the code above looks exactly like what a standard command base type would in other languages. But the question is how do you use this code when you are using a language like Go that has no concept of inheritance?

### Function pointer
Well the answer is simple. Simply take a look at the first function defined in the struct.
```go
// Run runs the command.
// The args are the arguments after the command name.
Run func(cmd *Command, args []string)
```
The code takes a function pointer, which is interesting. So how can we generalize this struct so that we have a base type to demonstrate the alternative pattern?

We can start by editing the struct and remove all the properties we don’t need for this post. This is what a simple command base type could look like. Of course you are free to customize the struct as you see fit, but for the purposes of this blog post the following should do just fine.
```go
package base

type Command struct {
// Run runs the command. The args are the arguments after the command name.
Run func(cmd *Command, args []string)

    // UsageLine is the one-line usage message.
    UsageLine string

    // Short is the short description shown in the 'help' output.
    Short string

    // Long is the long message shown in the 'go help <this-command>' output.
    Long string
}
```
Sample command
Now that we have removed some of the properties we don’t need, let’s go ahead and flesh out the first concrete Command we plan to demo.
```go
//first.go
package commands

import (
"fmt"
"base"
)

func NewFirstCommand() *base.Command {
cmd := &base.Command{
UsageLine:   "cmd1",
Short:       "command1",
Long:        "Sample command, the first one",
Run:         runMyFirstCommand,
}

    return cmd
}

func runMyFirstCommand(cmd *base.Command, args []string) {
fmt.Println("Hello from command1")
}
```
As you can see the code above uses the demo function NewFirstCommand() as a factory function that sets up all the plumbing for us. Our new factory sets the UsageLine to cmd1 and since the property is a string it would be a good candidate as a key in a map[string]*Command data structure.

The most important role of our factory is that it wires the function that actually performs the necessary work runMyFirstCommand into the instance of our command. In effect we have composed our parts together instead of inheriting them, like you would in other languages.

We could then use our `map[string]*Command` in our app to hold a reference to all our commands. This would allow us, at run time, to use the key and look up the concrete command and invoke its Run function to perform the actions of this command.

For now though lets just look at how we would instantiate and invoke our command from the main() function.
```
//main.go
package main

import (
"fmt",
"commands"
)

func main() {
cmd := commands.NewFirstCommand()
cmd.Run(cmd, nil)
}
```
The snippet above doesn’t pass any parameters to the Run function, but this is the intended behaviour for our small demo. If you feel comfortable looking at the Go source code you can take a look at how the Go team handles arguments and argument validation in their commands.

Although this version of the Command pattern is way different from the original one in the book by the ‘Gang of four` it is idiomatic Go. The cool thing about it is that is leverages the lack of type inheritance to create a perfectly valid command struct. And in a sense, like I previously mentioned, it uses ‘composition’ when declaring the different commands.

Of course the code shown above, is extremely simplistic but I am sure you can see all the possibilities.

Multiple Commands
In reality your application would most likely use some kind of data structure, like we saw with the map[string]*Command above. This would allow you to register all your commands, either statically or dynamically, at runtime and use the UsageLine property as the key for the appropriate command. Like so.
```go
//main_revised.go
package main

import (
"fmt",
"commands"
)

func main() {
arg := os.Args[1]

    cmd1 := NewFirstCommand()
    cmd2 := NewSecondCommand()

    var m = make(map[string]*base.Command)
    m[cmd1.UsageLine] = cmd1
    m[cmd2.UsageLine] = cmd2

    var cmd = m[arg]
    cmd.Run(cmd, nil)
}
```
This would make it easier to build CLI style apps where the user runs your application with the command name and any possible arguments. This is of course completely up to you how you would go about it, but at least I hope you appreciate how ingenious and creative the Go team is.

I personally have had a fun time finding and then understanding how the Go team had modelled this design pattern. It goes without saying that there are possibly many other ways to tackle this, each one with its own advantages and disadvantages.

For me it was important to understand this because it is definitely more idiomatic that the previous post I wrote, as such it allows me to learn and gain more of that mental flexibility required for this profession.

I hope you find this useful. The code for this post is available
