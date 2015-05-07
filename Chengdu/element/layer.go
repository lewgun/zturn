package scene



import (
"golang.org/x/mobile/sprite"
)

type Layer struct {
*sprite.Node
Tag string
}

func NewLayer( e sprite.Engine, tag string ) *Layer {
if e == nil {
return nil
}

r := &Layer{
Node: &Sprite.Node{},
Tag: tag,
}
e.Register(r)

return r
}

func ( s *Layer) AddLayer( l *Layer, tag string ) error {

return nil
}

func (s *Layer) RemoveLayer( l *Layer) error {
return nil
}

func (s *Layer) RemoveLayerByTag( tag string ) error {
return nil
}

func (s *Layer) LayerWithTag( tag string ) *Layer {
return nil
}

func (s *Layer) Render() {

}


func (s *Layer) tagExists(tag string ) bool {
return nil
}






