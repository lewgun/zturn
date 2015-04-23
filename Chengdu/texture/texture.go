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
	"github.com/golang/mobile/gl/glutil"
	"github.com/lewgun/mobile/geom"
	"github.com/lewgun/mobile/gl"
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
}

//initWithData Initializes with a texture2d with data
func (t *Texture) initWithData(
	data []byte,
	pixelFormat Texture2DPixelFormat,
	pixelsWide uint,
	pixelsHigh uint,
	contentSize geom.Size) bool {

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

func (t *Texture) HasPremultipliedAlpha() {

}

func (t *Texture) HasMipmaps() {}

func (t *Texture) initPremultipliedATextureWithImage(
	image *glutil.Image,
	pixelsWide uint,
	pixelsHigh uint) bool {

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
//static void setDefaultAlphaPixelFormat(CCTexture2DPixelFormat format);

/** returns the alpha pixel format
@since v0.8
@js getDefaultAlphaPixelFormat
*/
//static CCTexture2DPixelFormat defaultAlphaPixelFormat();

/** treats (or not) PVR files as if they have alpha premultiplied.
Since it is impossible to know at runtime if the PVR images have the alpha channel premultiplied, it is
possible load them as if they have (or not) the alpha channel premultiplied.

By default it is disabled.

@since v0.99.5
*/
//static void PVRImagesHavePremultipliedAlpha(bool haveAlphaPremultiplied);

//// returns the pixel format.
//const char* stringForFormat();

////void* keepData(void *data, unsigned int length);
///** Initializes a texture from a string with dimensions, alignment, font name and font size */
//bool initWithString(const char *text,  const char *fontName, float fontSize, const CCSize& dimensions, CCTextAlignment hAlignment, CCVerticalTextAlignment vAlignment);
///** Initializes a texture from a string with font name and font size */
//bool initWithString(const char *text, const char *fontName, float fontSize);
///** Initializes a texture from a string using a text definition*/
//bool initWithString(const char *text, ccFontDefinition *textDefinition);
