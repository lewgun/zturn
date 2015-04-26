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
func (t *Texture) DrawAtPoint(point geom.Point) {
	coordinates := []float32{
		0.0, t.MaxT,
		t.MaxS, t.MaxT,
		0.0, 0.0,
		t.MaxS, 0.0}

	width := float32(t.PixelsWide * t.MaxS)
	height := t.PixelsHigh * t.MaxT

	vertices := float32{
		point.X, point.y,
		width + point.x, point.y,
		point.x, height + point.y,
		width + point.x, height + point.y}

	//	ccGLEnableVertexAttribs( VertexAttribFlag_Position | VertexAttribFlag_TexCoords )
	//	m_pShaderProgram->use()
	//	m_pShaderProgram->setUniformsForBuiltins()
	//
	//	ccGLBindTexture2D( m_uName )
	//
	//
	//	#ifdef EMSCRIPTEN
	//	setGLBufferData(vertices, 8 * sizeof(GLfloat), 0)
	//	glVertexAttribPointer(VertexAttrib_Position, 2, GL_FLOAT, GL_FALSE, 0, 0)
	//
	//	setGLBufferData(coordinates, 8 * sizeof(GLfloat), 1)
	//	glVertexAttribPointer(VertexAttrib_TexCoords, 2, GL_FLOAT, GL_FALSE, 0, 0)
	//	#else
	//glVertexAttribPointer(VertexAttrib_Position, 2, GL_FLOAT, GL_FALSE, 0, vertices)
	//glVertexAttribPointer(VertexAttrib_TexCoords, 2, GL_FLOAT, GL_FALSE, 0, coordinates)
	//#endif // EMSCRIPTEN
	//
	//glDrawArrays(GL_TRIANGLE_STRIP, 0, 4)

}

//DrawInRect draws a texture inside a rect.
func (t *Texture) DrawInRect(rect geom.Rect) {

	coordinates := []float32{
		0.0, t.MaxT,
		t.MaxS, t.MaxT,
		0.0, 0.0,
		t.MaxS, 0.0}

	//	GLfloat    vertices[] = {    rect.origin.x,        rect.origin.y,                            /*0.0f,*/
	//	rect.origin.x + rect.size.width,        rect.origin.y,                            /*0.0f,*/
	//	rect.origin.x,                            rect.origin.y + rect.size.height,        /*0.0f,*/
	//	rect.origin.x + rect.size.width,        rect.origin.y + rect.size.height,        /*0.0f*/ }

	//	ccGLEnableVertexAttribs( VertexAttribFlag_Position | VertexAttribFlag_TexCoords )
	//	m_pShaderProgram->use()
	//	m_pShaderProgram->setUniformsForBuiltins()
	//
	//	ccGLBindTexture2D( m_uName )
	//
	//	#ifdef EMSCRIPTEN
	//	setGLBufferData(vertices, 8 * sizeof(GLfloat), 0)
	//	glVertexAttribPointer(VertexAttrib_Position, 2, GL_FLOAT, GL_FALSE, 0, 0)
	//
	//	setGLBufferData(coordinates, 8 * sizeof(GLfloat), 1)
	//	glVertexAttribPointer(VertexAttrib_TexCoords, 2, GL_FLOAT, GL_FALSE, 0, 0)
	//	#else
	//glVertexAttribPointer(VertexAttrib_Position, 2, GL_FLOAT, GL_FALSE, 0, vertices)
	//glVertexAttribPointer(VertexAttrib_TexCoords, 2, GL_FLOAT, GL_FALSE, 0, coordinates)
	//#endif // EMSCRIPTEN
	//glDrawArrays(GL_TRIANGLE_STRIP, 0, 4)

}

/**
Extensions to make it easy to create a CCTexture2D object from an image file.
Note that RGBA type textures will have their alpha premultiplied - use the blending mode (GL_ONE, GL_ONE_MINUS_SRC_ALPHA).
*/
/** Initializes a texture from a glutil.Image object */
func (t *Texture) initWithImage(img *glutil.Image) bool {
	if img == nil {
		//CCLOG("cocos2d: CCTexture2D. Can't create Texture. UIImage is nil")
		return false
	}

	//imageWidth := img->getWidth()
	//imageHeight := img->getHeight()

	//	CCConfiguration *conf = CCConfiguration::sharedConfiguration()
	//
	// maxTextureSize := conf.getMaxTextureSize()
	//if imageWidth > maxTextureSize || imageHeight > maxTextureSize {
	//
	////CCLOG("cocos2d: WARNING: Image (%u x %u) is bigger than the supported %u x %u", imageWidth, imageHeight, maxTextureSize, maxTextureSize)
	//return false
	//}

	// always load premultiplied images
	return initPremultipliedATextureWithImage(uiImage, imageWidth, imageHeight)

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

	pvr := NewPVR(file)
	bRet := pvr.initWithContentsOfFile(file)

	if bRet {
		//		pvr->setRetainName(true) // don't dealloc texture on release
		//
		//		t.Name = pvr.Name
		//		t.MaxS = 1.0
		//		t.MaxT = 1.0
		//		t.PixelsWide = pvr->getWidth()
		//		t.PixelsHigh = pvr->getHeight()
		//		t.ContentSize = CCSizeMake((float)m_uPixelsWide, (float)m_uPixelsHigh)
		//	t.HasPremultipliedAlpha = PVRHaveAlphaPremultiplied_
		//	t.PixelFormat = pvr->getFormat()
		//	t.HasMipmaps = pvr->getNumberOfMipmaps() > 1

	} else {
		//CCLOG("cocos2d: Couldn't load PVR image %s", file)
	}

	return bRet

}

/** Initializes a texture from a ETC file */
func (t *Texture) initWithETCFile(file string) bool {

	etc := NewETC()
	bRet := etc.init(file)

	if bRet {
		//t.Name = etc->getName()
		//t.MaxS = 1.0
		//t.MaxT = 1.0
		//t.PixelsWide = etc->getWidth()
		//t.PixelsHigh = etc->getHeight()
		//t.ContentSize = CCSizeMake((float)m_uPixelsWide, (float)m_uPixelsHigh)
		//t.HasPremultipliedAlpha = true

	} else {
		//CCLOG("cocos2d: Couldn't load ETC image %s", file)
	}

	return bRet
}

//SetTexParameters sets the min filter, mag filter, wrap s and wrap t texture parameters.
//If the texture size is NPOT (non power of 2), then in can only use GL_CLAMP_TO_EDGE in GL_TEXTURE_WRAP_{S,T}.
func (t *Texture) SetTexParameters(texParams *TexParams) {
	//CCAssert( (m_uPixelsWide == ccNextPOT(m_uPixelsWide) || texParams->wrapS == GL_CLAMP_TO_EDGE) &&
	//(m_uPixelsHigh == ccNextPOT(m_uPixelsHigh) || texParams->wrapT == GL_CLAMP_TO_EDGE),
	//"GL_CLAMP_TO_EDGE should be used in NPOT dimensions")
	//
	//ccGLBindTexture2D( m_uName )
	//glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, texParams->minFilter )
	//glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, texParams->magFilter )
	//glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, texParams->wrapS )
	//glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, texParams->wrapT )
	//
	//#if CC_ENABLE_CACHE_TEXTURE_DATA
	//VolatileTexture::setTexParameters(this, texParams)
	//#endif
}

// SetAntiAliasTexParameters sets antialias texture parameters:
// - GL_TEXTURE_MIN_FILTER = GL_LINEAR
// - GL_TEXTURE_MAG_FILTER = GL_LINEAR
func (t *Texture) SetAntiAliasTexParameters() {
	//ccGLBindTexture2D( m_uName )
	//
	//if( ! m_bHasMipmaps )
	//{
	//glTexParameteri( GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST )
	//}
	//else
	//{
	//glTexParameteri( GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST_MIPMAP_NEAREST )
	//}
	//
	//glTexParameteri( GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST )
	//#if CC_ENABLE_CACHE_TEXTURE_DATA
	//ccTexParams texParams = {m_bHasMipmaps?GL_NEAREST_MIPMAP_NEAREST:GL_NEAREST,GL_NEAREST,GL_NONE,GL_NONE}
	//VolatileTexture::setTexParameters(this, &texParams)
	//#endif
}

//SetAliasTexParameters sets alias texture parameters:
//- GL_TEXTURE_MIN_FILTER = GL_NEAREST
//- GL_TEXTURE_MAG_FILTER = GL_NEAREST
func (t *Texture) SetAliasTexParameters() {
	//ccGLBindTexture2D( m_uName )
	//
	//if( ! m_bHasMipmaps )
	//{
	//glTexParameteri( GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR )
	//}
	//else
	//{
	//glTexParameteri( GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR_MIPMAP_NEAREST )
	//}
	//
	//glTexParameteri( GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR )
	//#if CC_ENABLE_CACHE_TEXTURE_DATA
	//ccTexParams texParams = {m_bHasMipmaps?GL_LINEAR_MIPMAP_NEAREST:GL_LINEAR,GL_LINEAR,GL_NONE,GL_NONE}
	//VolatileTexture::setTexParameters(this, &texParams)
	//#endif
}

//GenerateMipmap Generates mipmap images for the texture.
//It only works if the texture size is POT (power of 2).
func (t *Texture) GenerateMipmap() {

	//CCAssert( m_uPixelsWide == ccNextPOT(m_uPixelsWide) && m_uPixelsHigh == ccNextPOT(m_uPixelsHigh), "Mipmap texture only works in POT textures")
	//ccGLBindTexture2D( m_uName )
	//glGenerateMipmap(GL_TEXTURE_2D)
	t.HasMipmaps = true
}

//BitsPerPixelForFormat returns the bits-per-pixel of the in-memory OpenGL texture
func (t *Texture) BitsPerPixelForCurrentFormat() uint {
	return t.BitsPerPixelForFormat(t.PixelFormat)
}

//BitsPerPixelForFormat Helper functions that returns bits per pixels for a given format.
func (t *Texture) BitsPerPixelForFormat(format Texture2DPixelFormat) uint {
	var ret uint

	switch format {
	case Texture2DPixelFormat_RGBA8888:
		ret = 32

	case Texture2DPixelFormat_RGB888:
		// It is 32 and not 24, since its internal representation uses 32 bits.
		ret = 32

	case Texture2DPixelFormat_RGB565:
		ret = 16

	case Texture2DPixelFormat_RGBA4444:
		ret = 16

	case Texture2DPixelFormat_RGB5A1:
		ret = 16

	case Texture2DPixelFormat_AI88:
		ret = 16

	case Texture2DPixelFormat_A8:
		ret = 8

	case Texture2DPixelFormat_I8:
		ret = 8

	case Texture2DPixelFormat_PVRTC4:
		ret = 4

	case Texture2DPixelFormat_PVRTC2:
		ret = 2

	default:
		ret = -1
		//CCAssert(false , "unrecognized pixel format")
		//CCLOG("bitsPerPixelForFormat: %ld, cannot give useful result", (long)format)

	}
	return ret
}

/** content size */
func (t *Texture) ContentSizeInPixels() geom.Size {

}

func (t *Texture) initPremultipliedATextureWithImage(
	image *glutil.Image,
	pixelsWide uint,
	pixelsHigh uint) bool {

	//          tempData := image->getData()

	var (
		inPixel32   *uint32
		inPixel8    *uint8
		outPixel16  *uint16
		pixelFormat CCTexture2DPixelFormat
	)

	var hasAlpha bool
	// hasAlpha := image->hasAlpha()
	//   imageSize := CCSizeMake((float)(image->getWidth()), (float)(image->getHeight()))

	//   bpp := image->getBitsPerComponent()

	// compute pixel format
	if hasAlpha {
		pixelFormat = g_defaultAlphaPixelFormat

	} else {
		if bpp >= 8 {
			pixelFormat = Texture2DPixelFormat_RGB888
		} else {
			pixelFormat = Texture2DPixelFormat_RGB565
		}

	}

	// Repack the pixel data into the right format

	length := width * height

	if pixelFormat == Texture2DPixelFormat_RGB565 {
		if hasAlpha {
			// Convert "RRRRRRRRRGGGGGGGGBBBBBBBBAAAAAAAA" to "RRRRRGGGGGGBBBBB"

			tempData := make([]byte, width*height*2)
			//outPixel16 = (*uint16)tempData
			//inPixel32 = (*uint32)image->getData()

			var outPixel16 *uint16
			var inPixel32 *uint32

			for i := 0; i < length; i++ {
				*outPixel16 =
					((((*inPixel32 >> 0) & 0xFF) >> 3) << 11) | // R
						((((*inPixel32 >> 8) & 0xFF) >> 2) << 5) | // G
						((((*inPixel32 >> 16) & 0xFF) >> 3) << 0) // B
				inPixel32++
				outPixel16++
			}
		} else {
			// Convert "RRRRRRRRRGGGGGGGGBBBBBBBB" to "RRRRRGGGGGGBBBBB"

			//tempData = make([]byte, width * height * 2)
			//outPixel16 = (*uint16)tempData
			//inPixel8 = (unsigned char*)image->getData()

			tempData := make([]byte, width*height*2)

			var outPixel16 *uint16
			var inPixel8 *byte

			for i := 0; i < length; i++ {
				//*outPixel16 =
				//(((*inPixel8++ & 0xFF) >> 3) << 11) |  // R
				//(((*inPixel8++ & 0xFF) >> 2) << 5)  |  // G
				//(((*inPixel8++ & 0xFF) >> 3) << 0)    // B
				outPixel16++
			}
		}
	} else if pixelFormat == Texture2DPixelFormat_RGBA4444 {
		// Convert "RRRRRRRRRGGGGGGGGBBBBBBBBAAAAAAAA" to "RRRRGGGGBBBBAAAA"

		//inPixel32 = (uint32*)image->getData()
		//tempData = make([]byte, width * height * 2)
		//outPixel16 = (uint16*)tempData

		tempData := make([]byte, width*height*2)

		var outPixel16 *uint16
		var inPixel32 *byte

		for i := 0; i < length; i++ {
			*outPixel16 =
				((((*inPixel32 >> 0) & 0xFF) >> 4) << 12) | // R
					((((*inPixel32 >> 8) & 0xFF) >> 4) << 8) | // G
					((((*inPixel32 >> 16) & 0xFF) >> 4) << 4) | // B
					((((*inPixel32 >> 24) & 0xFF) >> 4) << 0) // A
			inPixel32++
			outPixel16++
		}
	} else if pixelFormat == Texture2DPixelFormat_RGB5A1 {
		// Convert "RRRRRRRRRGGGGGGGGBBBBBBBBAAAAAAAA" to "RRRRRGGGGGBBBBBA"
		//inPixel32 = ( *uint32)image->getData()
		//tempData = make([]byte, width * height * 2)
		//outPixel16 = (*uint16)tempData

		tempData := make([]byte, width*height*2)
		var outPixel16 *uint16
		var inPixel32 *byte

		for i := 0; i < length; i++ {
			*outPixel16 =
				((((*inPixel32 >> 0) & 0xFF) >> 3) << 11) | // R
					((((*inPixel32 >> 8) & 0xFF) >> 3) << 6) | // G
					((((*inPixel32 >> 16) & 0xFF) >> 3) << 1) | // B
					((((*inPixel32 >> 24) & 0xFF) >> 7) << 0) // A
			outPixel16++
			inPixel32++
		}
	} else if pixelFormat == Texture2DPixelFormat_A8 {
		// Convert "RRRRRRRRRGGGGGGGGBBBBBBBBAAAAAAAA" to "AAAAAAAA"
		//inPixel32 = (uint32*)image->getData()
		//tempData = make([]byte, width * height)
		//outPixel8 := (*byte)tempData

		tempData := make([]byte, width*height)
		var outPixel8 *uint16
		var inPixel32 *byte

		for i := 0; i < length; i++ {
			*outPixel8 = (*inPixel32 >> 24) & 0xFF // A
			inPixel32++
			outPixel8++
		}
	}

	if hasAlpha && pixelFormat == Texture2DPixelFormat_RGB888 {
		// Convert "RRRRRRRRRGGGGGGGGBBBBBBBBAAAAAAAA" to "RRRRRRRRGGGGGGGGBBBBBBBB"
		//inPixel32 = (*uint32)image->getData()
		//tempData = make([]byte, width * height * 3)
		//outPixel8 := (*byte)tempData

		tempData := make([]byte, width*height)
		var outPixel8 *uint8
		var inPixel32 *uint32

		for i := 0; i < length; i++ {
			*outPixel8 = (*inPixel32 >> 0) & 0xFF // R
			outPixel8++
			*outPixel8 = (*inPixel32 >> 8) & 0xFF // G
			outPixel8++
			*outPixel8 = (*inPixel32 >> 16) & 0xFF // B
			outPixel8++
			inPixel32++
		}
	}

	t.initWithData(tempData, pixelFormat, width, height, imageSize)

	//if (tempData != image->getData())
	//{
	//delete [] tempData
	//}

	//t.HasPremultipliedAlpha = image->isPremultipliedAlpha()
	return true

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
