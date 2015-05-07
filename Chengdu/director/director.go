package director

import (
	"github.com/lewgun/stack"
	"github.com/lewgun/zturn/Chengdu/element"
)

var (
	Chief *Chief
)

func init () {

}


type Chief struct {

}

type Executive struct {

	current *element.Scene
	scenes stack.Stacker
}

func ( e *Executive) Play() {

}

func ( e *Executive) Pause() {

}

func ( e *Executive) Stop() {

}

func ( e *Executive) *Scene {

}

func (e *Executive) Push(s *Scene) {

}

func (e *Executive) Pop() {

}

func NewExecutive() *Executive {

}

