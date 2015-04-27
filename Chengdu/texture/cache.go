package texture

import (
	"github.com/lewgun/mobile/gl/glutil"
	"path/filepath"
	"strings"
	"fmt"
)



type ImageFormat uint8
const (
	ImageFormatUnknown ImageFormat = iota
	ImageFormatJPG
	ImageFormatPNG
	ImageFormatTIFF
	ImageFormatWebP
)


func ImageFormat( filename string ) ImageFormat {


	switch strings.ToLower(filepath.Ext(filename)) {
case "jpg", "jpeg":
return ImageFormatJPG

case "png":
return ImageFormatPNG

case "tiff":
return ImageFormatTIFF

case "webp"
return ImageFormatWebP

default:

}
return ImageFormatUnknown

}


/*
func loadImageData( fileName string ) {

	if ImageFormat(fileName) == ImageFormatUnknown {
		return
	}

	img := glutil.Image{}
	//if (pImage && !pImage->initWithImageFileThreadSafe(filename, imageType))

}

*/

type Cache interface {
	String() string
	SnapshotTextures() map[string] *Texture
}


type cache struct {
	textures map[string]*Texture
}



func (c *cache) addImageAsyncCallBack(dt float32) {

}

func (c *cache) String() string {

	return fmt.Sprintf("<CCTextureCache | Number of textures = %u>", len(c.textures))

}

func (c *cache) SnapshotTextures()  map[string] *Texture {

	ret := make(map[string]*Texture)

	for k,v := range c.textures {
		ret[k] = v
	}

	return ret

}

var g_sharedCache Cache

func SharedCache() Cache {

	if g_sharedCache == nil {
		g_sharedCache = &cache{
			textures: make(map[string]*Texture),
		}

	}
	return g_sharedCache

}

func PurgeSharedCache() {

}

/** Returns a Texture2D object given an file image
* If the file image was not previously loaded, it will create a new CCTexture2D
*  object and it will return it. It will use the filename as a key.
* Otherwise it will return a reference of a previously loaded image.
* Supported image extensions: .png, .bmp, .tiff, .jpeg, .pvr, .gif
 */
func (c *cache) addImage(fileImage string) Texture {

}

// /* Returns a Texture2D object given a file image
//    * If the file image was not previously loaded, it will create a new CCTexture2D object and it will return it.
//    * Otherwise it will load a texture in a new thread, and when the image is loaded, the callback will be called with the Texture2D as a parameter.
//    * The callback will be called from the main thread, so it is safe to create any cocos2d object from the callback.
//    * Supported image extensions: .png, .jpg
//    * @since v0.8
//    * @lua NA
//    */
//
//void addImageAsync(const char *path, CCObject *target, SEL_CallFuncO selector);

/* Returns a Texture2D object given an CGImageRef image
* If the image was not previously loaded, it will create a new CCTexture2D object and it will return it.
* Otherwise it will return a reference of a previously loaded image
* The "key" parameter will be used as the "key" for the cache.
* If "key" is nil, then a new texture will be created each time.
 */
func (c *cache) addUIImage(image *glutil.Image, key string) Texture {

}

func (c *cache) TextureForKey(key string) *Texture {

}

/** Reload texture from the image file
 * If the file image hasn't loaded before, load it.
 * Otherwise the texture will be reloaded from the file image.
 * The "filenName" parameter is the related/absolute path of the file image.
 * Return true if the reloading is succeed, otherwise return false.
 */
func (c *cache) reloadTexture(fileName string) bool {

}

/** Purges the dictionary of loaded textures.
* Call this method if you receive the "Memory Warning"
* In the short term: it will free some resources preventing your app from being killed
* In the medium term: it will allocate more resources
* In the long term: it will be the same
 */
func (c *cache) removeAllTextures() {

}

/** Removes unused textures
 * Textures that have a retain count of 1 will be deleted
 * It is convenient to call this method after when starting a new Scene
 * @since v0.8
 */
func (c *cache) removeUnusedTextures() {

}

/** Deletes a texture from the cache given a texture
 */
func (c *cache) removeTexture(texture *Texture) {

}

/** Deletes a texture from the cache given a its key name
  @since v0.99.4
*/
func (c *cache) removeTextureForKey(key string) {

}

/** Output to CCLOG the current contents of this CCTextureCache
* This will attempt to calculate the size of each texture, and the total texture memory in use
*
* @since v1.0
 */
func (c *cache) dumpCachedTextureInfo() {

}

/** Returns a Texture2D object given an PVR filename
* If the file image was not previously loaded, it will create a new CCTexture2D
*  object and it will return it. Otherwise it will return a reference of a previously loaded image
 */
func (c *cache) addPVRImage(filename string) {

}

/** Returns a Texture2D object given an ETC filename
 * If the file image was not previously loaded, it will create a new CCTexture2D
 *  object and it will return it. Otherwise it will return a reference of a previously loaded image
 *  @lua NA
 */
func (c *cache) addETCImage(filename string) {

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
