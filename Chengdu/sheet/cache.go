package sheet

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/mobile/sprite"
)

var (
	//ErrNonSupported standard for a file format which can't supported now.
	ErrNonSupported = errors.New("not supported format")

	//ErrIllegalParam standard for illegal parameter(s)
	ErrIllegalParam = errors.New("illegal parameter(s)")
)

func NewCache(eng sprite.Engine) *Cache {
	c := &Cache{
		sheets:      make(map[string]*Sheet),
		pendingTask: make(chan func(), 256),
		eng:         eng,
	}

	go c.run()

	return c
}

type Cache struct {
	sheets map[string]*Sheet

	//pending loading task.
	pendingTask chan func()

	eng sprite.Engine

	locks sync.RWMutex
}

func (c *Cache) String() string {
	return fmt.Sprintf("<CCTextureCache | Number of atlas = %u>", len(c.sheets))
}

//Exit exit the cache.
func (c *Cache) Exit() {
	close(c.pendingTask)

}

//SnapshotTextures get a snapshot.
func (c *Cache) SnapshotSheet() map[string]*Sheet {

	ret := make(map[string]*Sheet)

	c.locks.RLock()
	defer c.locks.RUnlock()

	for k, v := range c.sheets {
		ret[k] = v
	}

	return ret

}

func (c *Cache) AddSheet(meta, sheet string) *Sheet {

	chanSheet := make(chan *Sheet)
	c.AddSheetAsync(meta, sheet, func(s *Sheet) {
		chanSheet <- s
	})

	return <-chanSheet

}

func (c *Cache) AddSheetAsync(meta, sheet string, cb func(*Sheet)) {

	tex := c.Sheet(sheet)
	if tex != nil {

		if cb != nil {
			cb(tex)
		}

		return
	}

	c.pendingTask <- func() {
		tex := c.addSheet(meta, sheet)
		if cb != nil {
			cb(tex)
		}

	}

}

func (c *Cache) addSheet(meta, sheet string) *Sheet {
	s, ok := c.sheets[sheet]
	if ok {
		return s
	}

	s = newSheet(c.eng, meta, sheet)

	if s != nil {
		c.locks.Lock()
		c.sheets[sheet] = s
		c.locks.Unlock()
	}

	return s

}

func (c *Cache) SubTex(sheet, tex string) *sprite.SubTex {
	if sheet == "" || tex == "" {
		return nil
	}

	a := c.Sheet(sheet)
	if a == nil {
		return nil
	}

	return a.SubTex(tex)
}

func (c *Cache) Sheet(name string) *Sheet {
	if name == "" {
		return nil
	}

	c.locks.RLock()
	s, ok := c.sheets[name]
	c.locks.RUnlock()

	if !ok {
		return nil
	}

	return s

}

func (c *Cache) Clear() {
	c.locks.Lock()
	c.sheets = map[string]*Sheet{}
	c.locks.Unlock()
}

func (c *Cache) run() {
	for {
		select {
		case cb, ok := <-c.pendingTask:
			{

				if !ok {
					return
				}
				cb()

			}

		default:
			time.Sleep(1e9)
		}

	} // end of for {
}

func (c *Cache) GC() {

}

// RemoveAtlas deletes a texture from the cache given a texture
func (c *Cache) RemoveSheet(s *Sheet) {

	var (
		k string
		v *Sheet
	)

	c.locks.Lock()
	defer c.locks.Unlock()

	for k, v = range c.sheets {
		if v == s {
			delete(c.sheets, k)
			break
		}
	}

}

//RemoveAtlasByName deletes a atlas from the cache given a its key name
func (c *Cache) RemoveSheetByName(key string) {
	c.locks.Lock()
	defer c.locks.Unlock()

	delete(c.sheets, key)
}

func (c *Cache) Dump() {
	/*
		count := 0
		totalBytes := 0

		buf := &bytes.Buffer{}

		for k, v := range c.textures {
			bpp := v.BitsPerPixelForCurrentFormat()
			bits := v.PixelsWide * v.PixelsHigh * bpp / 8
			totalBytes += bits
			count++

			fmt.Fprintln(buf, "ChengDu:  \"%s\" id=%d %d x %d @ %d bpp => %d KB",
				k,
				v.Name,
				v.PixelsWide,
				v.PixelsHigh,
				bpp,
				bits/1024)
		}

		fmt.Fprintln(buf, "ChengDu: TextureCache dumpDebugInfo: %d textures, for %d KB (%.2f MB)", count, totalBytes/1024, float32(totalBytes)/(1024.0*1024.0))

		log.Println(buf.String())

	*/

}
