package scene


import (
	"golang.org/x/mobile/sprite"
)

type Scene struct {
	*sprite.Node
	Tag string
	ZOrder int
}

func NewScene( e sprite.Engine, tag string ) *Scene {
	if e == nil {
		return nil
	}
	s := &Scene{
		Node: &Sprite.Node{},
		Tag: tag,
	}
	e.Register(s)

	return s
}

func ( s *Scene) AddChild( n *Sprite.Node, tag string ) error {

	return nil
}

func (s *Scene) RemoveChild( l *Scene) error {
	return nil
}

func (s *Scene) RemoveChildByTag( tag string ) error {
	return nil
}

func (s *Scene) ChildWithTag( tag string ) *Scene {
	return nil
}

func (s *Scene) Render() {

}


func (s *Scene) tagExists(tag string ) bool {
	return nil
}





