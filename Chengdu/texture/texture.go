/****************************************************************************
Copyright (c) 2010-2012 cocos2d-x.org
Copyright (C) 2008      Apple Inc. All Rights Reserved.
http://www.cocos2d-x.org
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
****************************************************************************/
package texture

import (
	"fmt"
	"github.com/lewgun/mobile/geom"
	"github.com/lewgun/mobile/gl"
	"github.com/lewgun/mobile/gl/glutil"
)

//Texture2DPixelFormat is stand for possible texture pixel formats
type Texture2DPixelFormat uint8

const (
	Texture2DPixelFormat_Unknown Texture2DPixelFormat = iota

	//! 32-bit texture: RGBA8888
	Texture2DPixelFormat_RGBA8888
	//! 24-bit texture: RGBA888
	Texture2DPixelFormat_RGB888
	//! 16-bit texture without Alpha channel
	Texture2DPixelFormat_RGB565
	//! 8-bit textures used as masks
	Texture2DPixelFormat_A8
	//! 8-bit intensity texture
	Texture2DPixelFormat_I8
	//! 16-bit textures used as masks
	Texture2DPixelFormat_AI88
	//! 16-bit textures: RGBA4444
	Texture2DPixelFormat_RGBA4444
	//! 16-bit textures: RGB5A1
	Texture2DPixelFormat_RGB5A1
	//! 4-bit PVRTC-compressed texture: PVRTC4
	Texture2DPixelFormat_PVRTC4
	//! 2-bit PVRTC-compressed texture: PVRTC2
	Texture2DPixelFormat_PVRTC2

	//! Default texture format: RGBA8888
	Texture2DPixelFormat_Default = Texture2DPixelFormat_RGBA8888
)

/**
Extension to set the Min / Mag filter
*/
type TexParams struct {
	MinFilter uint
	MagFilter uint
	WrapS     uint
	WrapT     uint
}

//CLASS INTERFACES:

/** @brief CCTexture2D class.
* This class allows to easily create OpenGL 2D textures from images, text or raw data.
* The created CCTexture2D object will always have power-of-two dimensions.
* Depending on how you create the CCTexture2D object, the actual image area of the texture might be smaller than the texture dimensions i.e. "contentSize" != (pixelsWide, pixelsHigh) and (maxS, maxT) != (1.0, 1.0).
* Be aware that the content of the generated textures will be upside-down!
 */

type Texture struct {

	// By default PVR images are treated as if they don't have the alpha channel premultiplied
	PVRHaveAlphaPremultiplied bool

	/** pixel format of the texture */
	PixelFormat Texture2DPixelFormat

	/** width in pixels */
	PixelsWide uint

	/** height in pixels */
	PixelsHigh uint

	/** texture name */
	Name uint

	/** texture max S */
	MaxS float32

	/** texture max T */
	MaxT float32

	/** content size */
	ContentSize CCSize

	/** whether or not the texture has their Alpha premultiplied */
	HasPremultipliedAlpha bool

	HasMipmaps bool

	ShaderProgram gl.Shader
}

func (t *Texture) String() string {

	return fmt.Sprintf("<CCTexture2D | Name = %u | Dimensions = %u x %u | Coordinates = (%.2f, %.2f)>",
		t.Name, t.PixelsWide, t.PixelsHigh, t.MaxS, t.MaxT)

}

func (t *Texture) releaseData(data []byte) {

}

/**
Drawing extensions to make it easy to draw basic quads using a CCTexture2D object.
These functions require GL_TEXTURE_2D and both GL_VERTEX_ARRAY and GL_TEXTURE_COORD_ARRAY client states to be enabled.
*/

//DrawAtPoint draws a texture at a given point */
func (t *Texture) DrawAtPoint(p geom.Point) {
}

//DrawInRect draws a texture inside a rect.
func (t *Texture) DrawInRect(rect geom.Rectangle) {

}

/**
Extensions to make it easy to create a CCTexture2D object from an image file.
Note that RGBA type textures will have their alpha premultiplied - use the blending mode (GL_ONE, GL_ONE_MINUS_SRC_ALPHA).
*/
/** Initializes a texture from a glutil.Image object */
func (t *Texture) initWithImage(img *glutil.Image) bool {
	if img == nil {
		//CCLOG("cocos2d: CCTexture2D. Can't create Texture. UIImage is nil");
		return false;
	}


	//imageWidth := img->getWidth();
	//imageHeight := img->getHeight();

	CCConfiguration *conf = CCConfiguration::sharedConfiguration();

 maxTextureSize := conf.getMaxTextureSize()
if imageWidth > maxTextureSize || imageHeight > maxTextureSize {

//CCLOG("cocos2d: WARNING: Image (%u x %u) is bigger than the supported %u x %u", imageWidth, imageHeight, maxTextureSize, maxTextureSize);
return false
}

// always load premultiplied images
return initPremultipliedATextureWithImage(uiImage, imageWidth, imageHeight);

}

//initWithData Initializes with a texture2d with data
func (t *Texture) initWithData(
	data []byte,
	pixelFormat Texture2DPixelFormat,
	pixelsWide uint,
	pixelsHigh uint,
	contentSize geom.Size) bool {

	var bitsPerPixel uint
	//Hack: bitsPerPixelForFormat returns wrong number for RGB_888 textures. See function.
	if pixelFormat == Texture2DPixelFormat_RGB888 {
		bitsPerPixel = 24
	} else {
		bitsPerPixel = bitsPerPixelForFormat(pixelFormat)
	}

	bytesPerRow := pixelsWide * bitsPerPixel / 8

	switch {
	case bytesPerRow%8 == 0:
		//glPixelStorei(GL_UNPACK_ALIGNMENT, 8)

	case bytesPerRow%4 == 0:
		//glPixelStorei(GL_UNPACK_ALIGNMENT, 4)
	case bytesPerRow%2 == 0:
		//glPixelStorei(GL_UNPACK_ALIGNMENT, 2)
	default:
		//glPixelStorei(GL_UNPACK_ALIGNMENT, 1)

	}

	//glGenTextures(1, &m_uName)
	//ccGLBindTexture2D(m_uName)
	//
	//glTexParameteri( GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR )
	//glTexParameteri( GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR )
	//glTexParameteri( GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_CLAMP_TO_EDGE )
	//glTexParameteri( GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_CLAMP_TO_EDGE )

	// Specify OpenGL texture image

	switch pixelFormat {
	case Texture2DPixelFormat_RGBA8888:
	//glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, (GLsizei)pixelsWide, (GLsizei)pixelsHigh, 0, GL_RGBA, GL_UNSIGNED_BYTE, data)

	case Texture2DPixelFormat_RGB888:
	//glTexImage2D(GL_TEXTURE_2D, 0, GL_RGB, (GLsizei)pixelsWide, (GLsizei)pixelsHigh, 0, GL_RGB, GL_UNSIGNED_BYTE, data)

	case Texture2DPixelFormat_RGBA4444:
	//glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, (GLsizei)pixelsWide, (GLsizei)pixelsHigh, 0, GL_RGBA, GL_UNSIGNED_SHORT_4_4_4_4, data)

	case Texture2DPixelFormat_RGB5A1:
	//glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, (GLsizei)pixelsWide, (GLsizei)pixelsHigh, 0, GL_RGBA, GL_UNSIGNED_SHORT_5_5_5_1, data)

	case Texture2DPixelFormat_RGB565:
	//glTexImage2D(GL_TEXTURE_2D, 0, GL_RGB, (GLsizei)pixelsWide, (GLsizei)pixelsHigh, 0, GL_RGB, GL_UNSIGNED_SHORT_5_6_5, data)

	case Texture2DPixelFormat_AI88:
	//glTexImage2D(GL_TEXTURE_2D, 0, GL_LUMINANCE_ALPHA, (GLsizei)pixelsWide, (GLsizei)pixelsHigh, 0, GL_LUMINANCE_ALPHA, GL_UNSIGNED_BYTE, data)

	case Texture2DPixelFormat_A8:
	//glTexImage2D(GL_TEXTURE_2D, 0, GL_ALPHA, (GLsizei)pixelsWide, (GLsizei)pixelsHigh, 0, GL_ALPHA, GL_UNSIGNED_BYTE, data)

	case Texture2DPixelFormat_I8:
	//glTexImage2D(GL_TEXTURE_2D, 0, GL_LUMINANCE, (GLsizei)pixelsWide, (GLsizei)pixelsHigh, 0, GL_LUMINANCE, GL_UNSIGNED_BYTE, data)

	default:
		//CCAssert(0, "NSInternalInconsistencyException")

	}

	t.ContentSize = contentSize
	t.PixelsWide = pixelsWide
	t.PixelsHigh = pixelsHigh
	t.PixelFormat = pixelFormat
	t.MaxS = contentSize.width / float32(pixelsWide)
	t.MaxT = contentSize.height / float32(pixelsHigh)

	t.HasPremultipliedAlpha = false
	t.HasMipmaps = false

	//setShaderProgram(CCShaderCache::sharedShaderCache()->programForKey(Shader_PositionTexture))

	return true

}

/** Initializes a texture from a PVR file */
func (t *Texture) initWithPVRFile(file string) bool {

}

/** Initializes a texture from a ETC file */
func (t *Texture) initWithETCFile(file string) bool {
}

//SetTexParameters sets the min filter, mag filter, wrap s and wrap t texture parameters.
//If the texture size is NPOT (non power of 2), then in can only use GL_CLAMP_TO_EDGE in GL_TEXTURE_WRAP_{S,T}.
func (t *Texture) SetTexParameters(texParams *TexParams) {
}

// SetAntiAliasTexParameters sets antialias texture parameters:
// - GL_TEXTURE_MIN_FILTER = GL_LINEAR
// - GL_TEXTURE_MAG_FILTER = GL_LINEAR
func (t *Texture) SetAntiAliasTexParameters() {
}

//SetAliasTexParameters sets alias texture parameters:
//- GL_TEXTURE_MIN_FILTER = GL_NEAREST
//- GL_TEXTURE_MAG_FILTER = GL_NEAREST
func (t *Texture) SetAliasTexParameters()

//GenerateMipmap Generates mipmap images for the texture.
//It only works if the texture size is POT (power of 2).
func (t *Texture) GenerateMipmap() {
}

//BitsPerPixelForFormat returns the bits-per-pixel of the in-memory OpenGL texture
func (t *Texture) BitsPerPixelForFormat() uint {}

//BitsPerPixelForFormat Helper functions that returns bits per pixels for a given format.
func (t *Texture) BitsPerPixelForFormat(format Texture2DPixelFormat) uint {}

/** content size */
func (t *Texture) ContentSizeInPixels() geom.Size {

}

func (t *Texture) initPremultipliedATextureWithImage(
	image *glutil.Image,
	pixelsWide uint,
	pixelsHigh uint) bool {

            tempData := image->getData();

			var (
inPixel32 *uint32
inPixel8 *uint8
outPixel16 *uint16
pixelFormat CCTexture2DPixelFormat
			)

bool                      hasAlpha := image->hasAlpha();
CCSize                    imageSize := CCSizeMake((float)(image->getWidth()), (float)(image->getHeight()));
CCTexture2DPixelFormat    pixelFormat;
size_t                    bpp = image->getBitsPerComponent();

// compute pixel format
if (hasAlpha)
{
pixelFormat = g_defaultAlphaPixelFormat;
}
else
{
if (bpp >= 8)
{
pixelFormat = kCCTexture2DPixelFormat_RGB888;
}
else
{
pixelFormat = kCCTexture2DPixelFormat_RGB565;
}

}

// Repack the pixel data into the right format
unsigned int length = width * height;

if (pixelFormat == kCCTexture2DPixelFormat_RGB565)
{
if (hasAlpha)
{
// Convert "RRRRRRRRRGGGGGGGGBBBBBBBBAAAAAAAA" to "RRRRRGGGGGGBBBBB"

tempData = new unsigned char[width * height * 2];
outPixel16 = (unsigned short*)tempData;
inPixel32 = (unsigned int*)image->getData();

for(unsigned int i = 0; i < length; ++i, ++inPixel32)
{
*outPixel16++ =
((((*inPixel32 >>  0) & 0xFF) >> 3) << 11) |  // R
((((*inPixel32 >>  8) & 0xFF) >> 2) << 5)  |  // G
((((*inPixel32 >> 16) & 0xFF) >> 3) << 0);    // B
}
}
else
{
// Convert "RRRRRRRRRGGGGGGGGBBBBBBBB" to "RRRRRGGGGGGBBBBB"

tempData = new unsigned char[width * height * 2];
outPixel16 = (unsigned short*)tempData;
inPixel8 = (unsigned char*)image->getData();

for(unsigned int i = 0; i < length; ++i)
{
*outPixel16++ =
(((*inPixel8++ & 0xFF) >> 3) << 11) |  // R
(((*inPixel8++ & 0xFF) >> 2) << 5)  |  // G
(((*inPixel8++ & 0xFF) >> 3) << 0);    // B
}
}
}
else if (pixelFormat == kCCTexture2DPixelFormat_RGBA4444)
{
// Convert "RRRRRRRRRGGGGGGGGBBBBBBBBAAAAAAAA" to "RRRRGGGGBBBBAAAA"

inPixel32 = (unsigned int*)image->getData();
tempData = new unsigned char[width * height * 2];
outPixel16 = (unsigned short*)tempData;

for(unsigned int i = 0; i < length; ++i, ++inPixel32)
{
*outPixel16++ =
((((*inPixel32 >> 0) & 0xFF) >> 4) << 12) | // R
((((*inPixel32 >> 8) & 0xFF) >> 4) <<  8) | // G
((((*inPixel32 >> 16) & 0xFF) >> 4) << 4) | // B
((((*inPixel32 >> 24) & 0xFF) >> 4) << 0);  // A
}
}
else if (pixelFormat == kCCTexture2DPixelFormat_RGB5A1)
{
// Convert "RRRRRRRRRGGGGGGGGBBBBBBBBAAAAAAAA" to "RRRRRGGGGGBBBBBA"
inPixel32 = (unsigned int*)image->getData();
tempData = new unsigned char[width * height * 2];
outPixel16 = (unsigned short*)tempData;

for(unsigned int i = 0; i < length; ++i, ++inPixel32)
{
*outPixel16++ =
((((*inPixel32 >> 0) & 0xFF) >> 3) << 11) | // R
((((*inPixel32 >> 8) & 0xFF) >> 3) <<  6) | // G
((((*inPixel32 >> 16) & 0xFF) >> 3) << 1) | // B
((((*inPixel32 >> 24) & 0xFF) >> 7) << 0);  // A
}
}
else if (pixelFormat == kCCTexture2DPixelFormat_A8)
{
// Convert "RRRRRRRRRGGGGGGGGBBBBBBBBAAAAAAAA" to "AAAAAAAA"
inPixel32 = (unsigned int*)image->getData();
tempData = new unsigned char[width * height];
unsigned char *outPixel8 = tempData;

for(unsigned int i = 0; i < length; ++i, ++inPixel32)
{
*outPixel8++ = (*inPixel32 >> 24) & 0xFF;  // A
}
}

if (hasAlpha && pixelFormat == kCCTexture2DPixelFormat_RGB888)
{
// Convert "RRRRRRRRRGGGGGGGGBBBBBBBBAAAAAAAA" to "RRRRRRRRGGGGGGGGBBBBBBBB"
inPixel32 = (unsigned int*)image->getData();
tempData = new unsigned char[width * height * 3];
unsigned char *outPixel8 = tempData;

for(unsigned int i = 0; i < length; ++i, ++inPixel32)
{
*outPixel8++ = (*inPixel32 >> 0) & 0xFF; // R
*outPixel8++ = (*inPixel32 >> 8) & 0xFF; // G
*outPixel8++ = (*inPixel32 >> 16) & 0xFF; // B
}
}

initWithData(tempData, pixelFormat, width, height, imageSize);

if (tempData != image->getData())
{
delete [] tempData;
}

m_bHasPremultipliedAlpha = image->isPremultipliedAlpha();
return true;

}

//StringForFormat returns the pixel format.
func (t *Texture) StringForFormat() string {

	switch t.PixelFormat {
	case Texture2DPixelFormat_RGBA8888:
		return "RGBA8888"

	case Texture2DPixelFormat_RGB888:
		return "RGB888"

	case Texture2DPixelFormat_RGB565:
		return "RGB565"

	case Texture2DPixelFormat_RGBA4444:
		return "RGBA4444"

	case Texture2DPixelFormat_RGB5A1:
		return "RGB5A1"

	case Texture2DPixelFormat_AI88:
		return "AI88"

	case Texture2DPixelFormat_A8:
		return "A8"

	case Texture2DPixelFormat_I8:
		return "I8"

	case Texture2DPixelFormat_PVRTC4:
		return "PVRTC4"

	case Texture2DPixelFormat_PVRTC2:
		return "PVRTC2"

	default:
		panic("unrecognized pixel format")

	}

}

/** sets the default pixel format for UIImagescontains alpha channel.
If the UIImage contains alpha channel, then the options are:
- generate 32-bit textures: Texture2DPixelFormat_RGBA8888 (default one)
- generate 24-bit textures: Texture2DPixelFormat_RGB888
- generate 16-bit textures: Texture2DPixelFormat_RGBA4444
- generate 16-bit textures: Texture2DPixelFormat_RGB5A1
- generate 16-bit textures: Texture2DPixelFormat_RGB565
- generate 8-bit textures: Texture2DPixelFormat_A8 (only use it if you use just 1 color)
How does it work ?
- If the image is an RGBA (with Alpha) then the default pixel format will be used (it can be a 8-bit, 16-bit or 32-bit texture)
- If the image is an RGB (without Alpha) then: If the default pixel format is RGBA8888 then a RGBA8888 (32-bit) will be used. Otherwise a RGB565 (16-bit texture) will be used.
This parameter is not valid for PVR / PVR.CCZ images.

*/
//static void setDefaultAlphaPixelFormat(CCTexture2DPixelFormat format)

/** returns the alpha pixel format
@since v0.8
@js getDefaultAlphaPixelFormat
*/
//static CCTexture2DPixelFormat defaultAlphaPixelFormat()

/** treats (or not) PVR files as if they have alpha premultiplied.
Since it is impossible to know at runtime if the PVR images have the alpha channel premultiplied, it is
possible load them as if they have (or not) the alpha channel premultiplied.

By default it is disabled.

@since v0.99.5
*/
//static void PVRImagesHavePremultipliedAlpha(bool haveAlphaPremultiplied)

////void* keepData(void *data, unsigned int length)
///** Initializes a texture from a string with dimensions, alignment, font name and font size */
//bool initWithString(const char *text,  const char *fontName, float fontSize, const CCSize& dimensions, CCTextAlignment hAlignment, CCVerticalTextAlignment vAlignment)
///** Initializes a texture from a string with font name and font size */
//bool initWithString(const char *text, const char *fontName, float fontSize)
///** Initializes a texture from a string using a text definition*/
//bool initWithString(const char *text, ccFontDefinition *textDefinition)
