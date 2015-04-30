package texture

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"bytes"
	"github.com/lewgun/mobile/gl/glutil"
	"sync"
	"time"
)

type ImageFormat uint8

const (
	ImageFormatUnknown ImageFormat = iota
	ImageFormatJPG
	ImageFormatPNG
	ImageFormatTIFF
	ImageFormatWebP
)

var SharedCache *cache

func init() {
	SharedCache = &cache{
		textures:    make(map[string]*Texture),
		pendingTask: make(chan func(), 256),
	}

	go SharedCache.run()
}

func ImageFormat(filename string) ImageFormat {

	switch strings.ToLower(filepath.Ext(filename)) {
	case "jpg", "jpeg":
		return ImageFormatJPG

	case "png":
		return ImageFormatPNG

	case "tiff":
		return ImageFormatTIFF

	case "webp":
		return ImageFormatWebP

	default:

	}
	return ImageFormatUnknown

}

type cache struct {
	//all textures key: abspath value: *Texture.
	textures map[string]*Texture

	//pending loading task.
	pendingTask chan func()

	locks sync.RWMutex
}

func (c *cache) String() string {
	return fmt.Sprintf("<CCTextureCache | Number of textures = %u>", len(c.textures))
}

//Exit exit the cache.
func (c *cache) Exit() {
	close(c.pendingTask)

}

//SnapshotTextures get a snapshot.
func (c *cache) SnapshotTextures() map[string]*Texture {

	ret := make(map[string]*Texture)

	c.locks.RLock()
	defer c.locks.RUnlock()

	for k, v := range c.textures {
		ret[k] = v
	}

	return ret

}

/*
AddImage returns a Texture2D object given an file image
 If the file image was not previously loaded, it will create a new CCTexture2D
  object and it will return it. It will use the filename as a key.
 Otherwise it will return a reference of a previously loaded image.
 Supported image extensions: .png, .bmp, .tiff, .jpeg, .pvr, .gif
*/
func (c *cache) AddImage(file string) *Texture {

	chanTex := make(*Texture, 1)
	c.AddImageAsync(file, func(tex *Texture) {
		chanTex <- tex
	})

	return <-chanTex

}

/*
  AddImageAsync returns a Texture2D object given a file image
     If the file image was not previously loaded, it will create a new CCTexture2D object and it will return it.
     Otherwise it will load a texture in a new thread, and when the image is loaded, the callback will be called with the Texture2D as a parameter.
    // The callback will be called from the main thread, so it is safe to create any cocos2d object from the callback.
     Supported image extensions: .png, .jpg
*/

func (c *cache) AddImageAsync(file string, cb func(*Texture)) {

	//pathKey = CCFileUtils::sharedFileUtils()->fullPathForFilename(file);
	absPath := ""

	tex := c.TextureForKey(absPath)
	if tex != nil {

		if cb != nil {
			cb(tex)
		}

		return
	}

	c.pendingTask <- func() {
		tex := c.loadImage(file)
		if cb != nil {
			cb(tex)
		}

	}

}

func (c *cache) addNormalImage(file string) *Texture {

}
func (c *cache) loadImage(file string) *Texture {
	absPath := file

	tex := c.TextureForKey(absPath)
	if tex != nil {
		return tex
	}

	ext := strings.ToLower(filepath.Ext(absPath))

	switch ext {
	case "pvr":
		tex = c.addPVRImage(absPath)
	case "pkm":
		tex = c.addETCImage(absPath)

	default:
		tex = c.addNormalImage(file)
	}

	if tex != nil {
		c.locks.Lock()
		c.textures[absPath] = tex
		c.locks.Unlock()
	}

	return tex

}

func (c *cache) TextureForKey(key string) *Texture {
	if key == "" {
		return nil
	}

	c.locks.RLock()
	tex, ok := c.textures[key]
	c.locks.RUnlock()

	if !ok {
		return nil
	}

	return tex

}

/** Reload texture from the image file
 * If the file image hasn't loaded before, load it.
 * Otherwise the texture will be reloaded from the file image.
 * The "filenName" parameter is the related/absolute path of the file image.
 * Return true if the reloading is succeed, otherwise return false.
 */
func (c *cache) ReloadTexture(file string) bool {

}

/*
 RemoveAllTextures Purges the map of loaded textures.
 Call this method if you receive the "Memory Warning"
 In the short term: it will free some resources preventing your app from being killed
 In the medium term: it will allocate more resources
 In the long term: it will be the same
*/
func (c *cache) RemoveAllTextures() {
	c.locks.Lock()
	c.textures = map[string]*Texture{}
	c.locks.Unlock()
}

func (c *cache) run() {
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

/** Removes unused textures
 * Textures that have a retain count of 1 will be deleted
 * It is convenient to call this method after when starting a new Scene
 * @since v0.8
 */
func (c *cache) RemoveUnusedTextures() {

}

/*
 RemoveTexture deletes a texture from the cache given a texture
*/
func (c *cache) RemoveTexture(tex *Texture) {

	var (
		k string
		v *Texture
	)

	c.locks.Lock()
	defer c.locks.Unlock()

	for k, v = range c.textures {
		if v == tex {
			delete(c.textures, k)
			break
		}
	}

}

/*
RemoveTextureForKey deletes a texture from the cache given a its key name
*/
func (c *cache) RemoveTextureForKey(key string) {
	c.locks.Lock()
	defer c.locks.Unlock()

	delete(c.textures, key)
}

/*
 DumpCachedTextureInfo output to log the current contents of this CCTextureCache
 This will attempt to calculate the size of each texture, and the total texture memory in use
*/
func (c *cache) DumpCachedTextureInfo() {

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

}

/*
 AddPVRImage returns a Texture2D object given an PVR filename
 If the file image was not previously loaded, it will create a new CCTexture2D
  object and it will return it. Otherwise it will return a reference of a previously loaded image
*/
func (c *cache) AddPVRImage(filename string) *Texture {
	var tex *Texture
	tex = c.TextureForKey(filename)
	if tex != nil {
		return tex
	}

	abs := filename

	tex = &Texture{}

	if tex.initWithPVRFile(abs) {
		return tex
	}

	return nil

}

/*
 AddPVRImage returns a Texture2D object given an ETC filename
  If the file image was not previously loaded, it will create a new CCTexture2D
   object and it will return it. Otherwise it will return a reference of a previously loaded image
*/
func (c *cache) AddETCImage(filename string) {
	var tex *Texture
	tex = c.TextureForKey(filename)
	if tex != nil {
		return tex
	}

	abs := filename

	tex = &Texture{}

	if tex.initWithETCFile(abs) {
		return tex
	}

	return nil

}

///** Reload all textures
//It's only useful when the value of CC_ENABLE_CACHE_TEXTURE_DATA is 1
//*/
//static void reloadAllTextures();

//#if CC_ENABLE_CACHE_TEXTURE_DATA
//
//class VolatileTexture
//{
//typedef enum {
//kInvalid = 0,
//kImageFile,
//kImageData,
//kString,
//kImage,
//}ccCachedImageType;
//
//public:
//VolatileTexture(CCTexture2D *t);
//~VolatileTexture();
//
//static void addImageTexture(CCTexture2D *tt, const char* imageFileName, CCImage::EImageFormat format);
//static void addStringTexture(CCTexture2D *tt, const char* text, const CCSize& dimensions, CCTextAlignment alignment,
//CCVerticalTextAlignment vAlignment, const char *fontName, float fontSize);
//static void addDataTexture(CCTexture2D *tt, void* data, CCTexture2DPixelFormat pixelFormat, const CCSize& contentSize);
//static void addCCImage(CCTexture2D *tt, CCImage *image);
//
//static void setTexParameters(CCTexture2D *t, ccTexParams *texParams);
//static void removeTexture(CCTexture2D *t);
//static void reloadAllTextures();
//
//public:
//static std::list<VolatileTexture*> textures;
//static bool isReloading;
//
//private:
//// find VolatileTexture by CCTexture2D*
//// if not found, create a new one
//static VolatileTexture* findVolotileTexture(CCTexture2D *tt);
//
//protected:
//CCTexture2D *texture;
//
//CCImage *uiImage;
//
//ccCachedImageType m_eCashedImageType;
//
//void *m_pTextureData;
//CCSize m_TextureSize;
//CCTexture2DPixelFormat m_PixelFormat;
//
//std::string m_strFileName;
//CCImage::EImageFormat m_FmtImage;
//
//ccTexParams     m_texParams;
//CCSize          m_size;
//CCTextAlignment m_alignment;
//CCVerticalTextAlignment m_vAlignment;
//std::string     m_strFontName;
//std::string     m_strText;
//float           m_fFontSize;
//};
//
//#endif
