
package main


import (
    "fmt"
    "github.com/lewgun/zturn/Chengdu/cmd/creator/config"
    "github.com/lewgun/zturn/Chengdu/cmd/creator/builder"
)


var (
    tips = `Welcome to Chengdu creator.
With this tool. follow the directives.
You can create the skeleton of your fantastic game.
Enjoy!`
)

func main() {
    fmt.Println(tips)

    c := config.New()
    c.Config()

    b := builder.New(c)
    if err := b.Build(); err != nil {
        b.Clean()
    }


    fmt.Println(b)



}



