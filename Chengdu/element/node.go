package scene

import (
	"golang.org/x/mobile/sprite"
)


type Node struct {
	*sprite.Node
	Tag string
	ZOrder int
}

func NewNode( tag string, ) {

}