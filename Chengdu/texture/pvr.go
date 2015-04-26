package texture

import "strings"

const (
	PVRTextureFlagTypeMask = 0xff
)

//
// XXX DO NO ALTER THE ORDER IN THIS LIST XXX
//
var PVRTableFormats = [...]PVRTexturePixelFormatInfo{

	// 0: BGRA_8888
	{GL_RGBA, GL_BGRA, GL_UNSIGNED_BYTE, 32, false, true, Texture2DPixelFormat_RGBA8888},
	// 1: RGBA_8888
	{GL_RGBA, GL_RGBA, GL_UNSIGNED_BYTE, 32, false, true, Texture2DPixelFormat_RGBA8888},
	// 2: RGBA_4444
	{GL_RGBA, GL_RGBA, GL_UNSIGNED_SHORT_4_4_4_4, 16, false, true, Texture2DPixelFormat_RGBA4444},
	// 3: RGBA_5551
	{GL_RGBA, GL_RGBA, GL_UNSIGNED_SHORT_5_5_5_1, 16, false, true, Texture2DPixelFormat_RGB5A1},
	// 4: RGB_565
	{GL_RGB, GL_RGB, GL_UNSIGNED_SHORT_5_6_5, 16, false, false, Texture2DPixelFormat_RGB565},
	// 5: RGB_888
	{GL_RGB, GL_RGB, GL_UNSIGNED_BYTE, 24, false, false, Texture2DPixelFormat_RGB888},
	// 6: A_8
	{GL_ALPHA, GL_ALPHA, GL_UNSIGNED_BYTE, 8, false, false, Texture2DPixelFormat_A8},
	// 7: L_8
	{GL_LUMINANCE, GL_LUMINANCE, GL_UNSIGNED_BYTE, 8, false, false, Texture2DPixelFormat_I8},
	// 8: LA_88
	{GL_LUMINANCE_ALPHA, GL_LUMINANCE_ALPHA, GL_UNSIGNED_BYTE, 16, false, true, Texture2DPixelFormat_AI88},

	//// Not all platforms include GLES/gl2ext.h so these PVRTC enums are not always
	//// available.
	//#ifdef GL_COMPRESSED_RGB_PVRTC_2BPPV1_IMG
	//// 9: PVRTC 2BPP RGB
	//{GL_COMPRESSED_RGB_PVRTC_2BPPV1_IMG, 0xFFFFFFFF, 0xFFFFFFFF, 2, true, false, Texture2DPixelFormat_PVRTC2},
	//// 10: PVRTC 2BPP RGBA
	//{GL_COMPRESSED_RGBA_PVRTC_2BPPV1_IMG, 0xFFFFFFFF, 0xFFFFFFFF, 2, true, true, Texture2DPixelFormat_PVRTC2},
	//// 11: PVRTC 4BPP RGB
	//{GL_COMPRESSED_RGB_PVRTC_4BPPV1_IMG, 0xFFFFFFFF, 0xFFFFFFFF, 4, true, false, Texture2DPixelFormat_PVRTC4},
	//// 12: PVRTC 4BPP RGBA
	//{GL_COMPRESSED_RGBA_PVRTC_4BPPV1_IMG, 0xFFFFFFFF, 0xFFFFFFFF, 4, true, true, Texture2DPixelFormat_PVRTC4},
	//#endif
}

/**
@brief Structure which can tell where mipmap begins and how long is it
*/
type PVRMipmap struct {
	addr *byte
	len  uint
}

type PVRTexturePixelFormatInfo struct {
	internalFormat int
	format         int
	typ            int
	bpp            uint32
	compressed     bool
	alpha          bool
	PixelFormat    Texture2DPixelFormat
}

type pixelFormatHash struct {
	pixelFormat     uint64
	pixelFormatInfo *PVRTexturePixelFormatInfo
}

// Values taken from PVRTexture.h from http://www.imgtec.com
const (
	PVR2TextureFlagMipmap       = (1 << 8)  // has mip map levels
	PVR2TextureFlagTwiddle      = (1 << 9)  // is twiddled
	PVR2TextureFlagBumpmap      = (1 << 10) // has normals encoded for a bump map
	PVR2TextureFlagTiling       = (1 << 11) // is bordered for tiled pvr
	PVR2TextureFlagCubemap      = (1 << 12) // is a cubemap/skybox
	PVR2TextureFlagFalseMipCol  = (1 << 13) // are there false colored MIP levels
	PVR2TextureFlagVolume       = (1 << 14) // is this a volume texture
	PVR2TextureFlagAlpha        = (1 << 15) // v2.1 is there transparency info in the texture
	PVR2TextureFlagVerticalFlip = (1 << 16) // v2.1 is the texture vertically flipped
)

/**
@brief Determine how many mipmaps can we have.
Its same as define but it respects namespaces
*/
const (
	PVRMipMapMax                      = 16
	PVR3TextureFlagPremultipliedAlpha = (1 << 1) // has premultiplied alpha
)

const (
	gPVRTexIdentifier = "PVR!"
)

// v2
type PVR2TexturePixelFormat uint8

const (
	PVR2TexturePixelFormat_RGBA_4444 PVR2TexturePixelFormat = 0x10 + iota
	PVR2TexturePixelFormat_RGBA_5551
	PVR2TexturePixelFormat_RGBA_8888
	PVR2TexturePixelFormat_RGB_565
	PVR2TexturePixelFormat_RGB_555 // unsupported
	PVR2TexturePixelFormat_RGB_888
	PVR2TexturePixelFormat_I_8
	PVR2TexturePixelFormat_AI_88
	PVR2TexturePixelFormat_PVRTC_2BPP_RGBA
	PVR2TexturePixelFormat_PVRTC_4BPP_RGBA
	PVR2TexturePixelFormat_BGRA_8888
	PVR2TexturePixelFormat_A_8
)

// v3
/* supported predefined formats */
const (
	PVR3TexturePixelFormat_PVRTC_2BPP_RGB = iota
	PVR3TexturePixelFormat_PVRTC_2BPP_RGBA
	PVR3TexturePixelFormat_PVRTC_4BPP_RGB
	PVR3TexturePixelFormat_PVRTC_4BPP_RGBA
)

/* supported channel type formats */
const (
	PVR3TexturePixelFormat_BGRA_8888 = uint64(0x0808080861726762)
	PVR3TexturePixelFormat_RGBA_8888 = uint64(0x0808080861626772)
	PVR3TexturePixelFormat_RGBA_4444 = uint64(0x0404040461626772)
	PVR3TexturePixelFormat_RGBA_5551 = uint64(0x0105050561626772)
	PVR3TexturePixelFormat_RGB_565   = uint64(0x0005060500626772)
	PVR3TexturePixelFormat_RGB_888   = uint64(0x0008080800626772)
	PVR3TexturePixelFormat_A_8       = uint64(0x0000000800000061)
	PVR3TexturePixelFormat_L_8       = uint64(0x000000080000006c)
	PVR3TexturePixelFormat_LA_88     = uint64(0x000008080000616c)
)

// v2
var v2PixelFormatHash = [...]pixelFormatHash{

	{PVR2TexturePixelFormat_BGRA_8888, &PVRTableFormats[0]},
	{PVR2TexturePixelFormat_RGBA_8888, &PVRTableFormats[1]},
	{PVR2TexturePixelFormat_RGBA_4444, &PVRTableFormats[2]},
	{PVR2TexturePixelFormat_RGBA_5551, &PVRTableFormats[3]},
	{PVR2TexturePixelFormat_RGB_565, &PVRTableFormats[4]},
	{PVR2TexturePixelFormat_RGB_888, &PVRTableFormats[5]},
	{PVR2TexturePixelFormat_A_8, &PVRTableFormats[6]},
	{PVR2TexturePixelFormat_I_8, &PVRTableFormats[7]},
	{PVR2TexturePixelFormat_AI_88, &PVRTableFormats[8]},

	//#ifdef GL_COMPRESSED_RGB_PVRTC_2BPPV1_IMG
	//{ kPVR2TexturePixelFormat_PVRTC_2BPP_RGBA,	&PVRTableFormats[10] },
	//{ kPVR2TexturePixelFormat_PVRTC_4BPP_RGBA,	&PVRTableFormats[12] },
	//#endif
}

var (
	PVR2MaxTableElements = len(v2PixelFormatHash)
)

// v3
var v3PixelFormatHash = [...]pixelFormatHash{

	{PVR3TexturePixelFormat_BGRA_8888, &PVRTableFormats[0]},
	{PVR3TexturePixelFormat_RGBA_8888, &PVRTableFormats[1]},
	{PVR3TexturePixelFormat_RGBA_4444, &PVRTableFormats[2]},
	{PVR3TexturePixelFormat_RGBA_5551, &PVRTableFormats[3]},
	{PVR3TexturePixelFormat_RGB_565, &PVRTableFormats[4]},
	{PVR3TexturePixelFormat_RGB_888, &PVRTableFormats[5]},
	{PVR3TexturePixelFormat_A_8, &PVRTableFormats[6]},
	{PVR3TexturePixelFormat_L_8, &PVRTableFormats[7]},
	{PVR3TexturePixelFormat_LA_88, &PVRTableFormats[8]},

	//#ifdef GL_COMPRESSED_RGB_PVRTC_2BPPV1_IMG
	//{kPVR3TexturePixelFormat_PVRTC_2BPP_RGB,	&PVRTableFormats[9] },
	//{kPVR3TexturePixelFormat_PVRTC_2BPP_RGBA,	&PVRTableFormats[10] },
	//{kPVR3TexturePixelFormat_PVRTC_4BPP_RGB,	&PVRTableFormats[11] },
	//{kPVR3TexturePixelFormat_PVRTC_4BPP_RGBA,	&PVRTableFormats[12] },
	//#endif
}

//Tells How large is tableFormats
var (
	PVR3MaxTableElements = len(v3PixelFormatHash)
)

type PVRv2TexHeader struct {
	headerLen       uint
	height          uint
	width           uint
	numMipmaps      uint
	flags           uint
	dataLen         uint
	bpp             uint
	bitmaskRed      uint
	bitmaskGreen    uint
	bitmaskBlueuint uint
	bitmaskAlpha    uint
	pvrTag          uint
	numSurfs        uint
}

//#ifdef _MSC_VER
//#pragma pack(push,1)
//#endif
type PVRv3TexHeader struct {
	version       uint32
	flags         uint32
	pixelFormat   uint64
	colorSpace    uint32
	channelType   uint32
	height        uint32
	width         uint32
	depth         uint32
	numOfSurfaces uint32
	numOfFaces    uint32
	numOfMipmaps  uint32
	metadataLen   uint32
}

/** TexturePVR

Object that loads PVR images.
Supported PVR formats:
   - RGBA8888
   - BGRA8888
   - RGBA4444
   - RGBA5551
   - RGB565
   - A8
   - I8
   - AI88
   - PVRTC 4BPP
   - PVRTC 2BPP

Limitations:
   Pre-generated mipmaps, such as PVR textures with mipmap levels embedded in file,
   are only supported if all individual sprites are of _square_ size.
   To use mipmaps with non-square textures, instead call Texture2D#generateMipmap on the sheet texture itself
   (and to save space, save the PVR sprite sheet without mip maps included).

*/

type PVR struct {

	// pointer to mipmap images
	Mipmaps [PVRMipMapMax]PVRMipmap
	/** how many mipmaps the texture has.
	  1 means one level (level 0 number of mipmap used)
	  NumberOfMipmaps uint
	  /** texture width */
	Width uint
	/** texture height */
	Height uint
	/** texture id name */
	Name uint
	/** whether or not the texture has alpha */
	HasAlpha bool
	/** whether or not the texture has premultiplied alpha */
	HasPremultipliedAlpha bool
	/** whether or not the texture should use hasPremultipliedAlpha instead of global default */
	ForcePremultipliedAlpha bool

	// cocos2d integration
	RetainName bool
	Format     Texture2DPixelFormat

	PixelFormatInfo *PVRTexturePixelFormatInfo
}

func (p *PVR) unpackPVRv2Data(data []byte) bool {

	var (
		success     bool
		header      PVRv2TexHeader
		flags       uint
		pvrTag      uint
		dataLen     uint
		dataOffset  uint
		dataSize    uint
		blockSize   uint
		widthBlocks uint
		heightBlocks
		width       uint
		height      uint
		bpp         uint
		formatFlags uint
	)

	dataLen := len(data)

	//unsigned char *bytes = NULL

	//Cast first sizeof(PVRTexHeader) bytes of data stream as PVRTexHeader
	header = (*PVRv2TexHeader)(data)

	//Make sure that tag is in correct formatting
	//pvrTag = CC_SWAP_INT32_LITTLE_TO_HOST(header.pvrTag)

	if gPVRTexIdentifier[0] != (char)(((pvrTag>>0)&0xff)) ||
		gPVRTexIdentifier[1] != (char)(((pvrTag>>8)&0xff)) ||
		gPVRTexIdentifier[2] != (char)(((pvrTag>>16)&0xff)) ||
		gPVRTexIdentifier[3] != (char)(((pvrTag>>24)&0xff)) {
		return false
	}

	//CCConfiguration *configuration = CCConfiguration::sharedConfiguration()

	//flags = CC_SWAP_INT32_LITTLE_TO_HOST(header->flags)
	formatFlags = flags & PVRTextureFlagTypeMask

	var flipped bool
	if flags&kPVR2TextureFlagVerticalFlip != 0 {
		flipped = true
	}

	if flipped {
		//CCLOG("cocos2d: WARNING: Image is flipped. Regenerate it using PVRTexTool")
	}
	//
	//if (! configuration->supportsNPOT() &&
	//(header.width != ccNextPOT(header.width) || header.height != ccNextPOT(header.height))){
	////CCLOG("cocos2d: ERROR: Loading an NPOT texture (%dx%d) but is not supported on this device", header->width, header->height)
	//return false
	//}

	pvr2TableElements := PVR2MaxTableElements
	//if (! CCConfiguration::sharedConfiguration()->supportsPVRTC())
	//{
	//pvr2TableElements = 9
	//}

	for i := 0; i < pvr2TableElements; i++ {
		//Does image format in table fits to the one parsed from header?
		if v2PixelFormatHash[i].pixelFormat == formatFlags {
			p.PixelFormatInfo = v2PixelFormatHash[i].pixelFormatInfo

			//Reset num of mipmaps
			p.NumberOfMipmaps = 0

			//Get size of mipmap
			width = CC_SWAP_INT32_LITTLE_TO_HOST(header.width)
			p.Width = width

			height = CC_SWAP_INT32_LITTLE_TO_HOST(header.height)
			p.Height = height

			//Do we use alpha ?
			if CC_SWAP_INT32_LITTLE_TO_HOST(header.bitmaskAlpha) {
				t.HasAlpha = true
			} else {
				t.HasAlpha = false
			}

			//Get ptr to where data starts..
			dataLen = CC_SWAP_INT32_LITTLE_TO_HOST(header.dataLen)

			//Move by size of header
			bytes = (*byte)(data) + sizeof(PVRv2TexHeader)
			t.Format = t.PixelFormatInfo.PixelFormat
			bpp = t.PixelFormatInfo.bpp

			// Calculate the data size for each texture level and respect the minimum number of blocks
			for dataOffset < dataLength {
				switch formatFlags {
				case PVR2TexturePixelFormat_PVRTC_2BPP_RGBA:
					blockSize = 8 * 4 // Pixel by pixel block size for 2bpp
					widthBlocks = width / 8
					heightBlocks = height / 4

				case PVR2TexturePixelFormat_PVRTC_4BPP_RGBA:
					blockSize = 4 * 4 // Pixel by pixel block size for 4bpp
					widthBlocks = width / 4
					heightBlocks = height / 4

				case PVR2TexturePixelFormat_BGRA_8888:
				//if (CCConfiguration::sharedConfiguration()->supportsBGRA8888() == false)
				//{
				//CCLOG("cocos2d: TexturePVR. BGRA8888 not supported on this device")
				//return false
				//}
				default:
					blockSize = 1
					widthBlocks = width
					heightBlocks = height

				}

				// Clamp to minimum number of blocks
				if widthBlocks < 2 {
					widthBlocks = 2
				}

				if heightBlocks < 2 {
					heightBlocks = 2
				}

				dataSize = widthBlocks * heightBlocks * ((blockSize * bpp) / 8)
				packetLen := (dataLen - dataOffset)
				if packetLen > dataSize {
					packetLen = dataSize
				}

				//Make record to the mipmaps array and increment counter
				t.Mipmaps[t.NumberOfMipmaps].addr = bytes + dataOffset
				t.Mipmaps[t.NumberOfMipmaps].len = packetLen
				t.NumberOfMipmaps++

				////Check that we didn't overflow
				//CCAssert(m_uNumberOfMipmaps < CC_PVRMIPMAP_MAX,
				//"TexturePVR: Maximum number of mipmaps reached. Increase the CC_PVRMIPMAP_MAX value")

				dataOffset += packetLen

				//Update width and height to the next lower power of two
				width = MAX(width>>1, 1)
				height = MAX(height>>1, 1)
			}

			//Mark pass as success
			success = true
			break
		}
	}

	if !success {
		//CCLOG("cocos2d: WARNING: Unsupported PVR Pixel Format: 0x%2x. Re-encode it with a OpenGL pixel format variant", formatFlags)
	}

	return success

}
func (p *PVR) unpackPVRv3Data(data []byte) bool {

	dataLen = len(data)

	if dataLen < sizeof(PVRv3TexHeader) {
		return false
	}

	header := (*PVRv3TexHeader)(data)

	// validate version
	if CC_SWAP_INT32_BIG_TO_HOST(header.version) != 0x50565203 {
		//  CCLOG("cocos2d: WARNING: pvr file version mismatch")
		return false
	}

	// parse pixel format
	pixelFormat := header.pixelFormat

	infoValid := false

	pvr3TableElements := PVR3MaxTableElements
	//    if (! CCConfiguration::sharedConfiguration()->supportsPVRTC())
	//{
	//pvr3TableElements = 9
	//}

	for i := 0; i < pvr3TableElements; i++ {
		if v3PixelFormatHash[i].pixelFormat == pixelFormat {
			p.PixelFormatInfo = v3PixelFormatHash[i].pixelFormatInfo
			p.HasAlpha = p.PixelFormatInfo.alpha
			infoValid = true
			break
		}
	}

	// unsupported / bad pixel format
	if !infoValid {
		//CCLOG("cocos2d: WARNING: unsupported pvr pixelformat: %lx", (unsigned long)pixelFormat )
		return false
	}

	// flags
	flags := CC_SWAP_INT32_LITTLE_TO_HOST(header.flags)

	// PVRv3 specifies premultiply alpha in a flag -- should always respect this in PVRv3 files
	t.ForcePremultipliedAlpha = true
	if flags & PVR3TextureFlagPremultipliedAlpha {
		t.HasPremultipliedAlpha = true
	}

	// sizing
	width := CC_SWAP_INT32_LITTLE_TO_HOST(header.width)
	height := CC_SWAP_INT32_LITTLE_TO_HOST(header.height)
	p.Width = width
	p.Height = height

	var (
		dataOffset   uint32
		dataSize     uint32
		blockSize    uint32
		widthBlocks  uint32
		heightBlocks uint32

		bytes *uint8
	)

	dataOffset = (sizeof(PVRv3TexHeader) + header.metadataLen)
	bytes = data

	p.NumberOfMipmaps = p.numOfMipmaps
	//CCAssert(m_uNumberOfMipmaps < CC_PVRMIPMAP_MAX, "TexturePVR: Maximum number of mimpaps reached. Increate the CC_PVRMIPMAP_MAX value")

	for i := 0; i < p.numOfMipmaps; i++ {

		switch pixelFormat {

		case PVR3TexturePixelFormat_PVRTC_2BPP_RGB:
			fallthrough

		case PVR3TexturePixelFormat_PVRTC_2BPP_RGBA:
			blockSize = 8 * 4 // Pixel by pixel block size for 2bpp
			widthBlocks = width / 8
			heightBlocks = height / 4

		case PVR3TexturePixelFormat_PVRTC_4BPP_RGB:
			fallthrough
		case PVR3TexturePixelFormat_PVRTC_4BPP_RGBA:
			blockSize = 4 * 4 // Pixel by pixel block size for 4bpp
			widthBlocks = width / 4
			heightBlocks = height / 4

		case PVR3TexturePixelFormat_BGRA_8888:
		//if( ! CCConfiguration::sharedConfiguration()->supportsBGRA8888())
		//{
		//CCLOG("cocos2d: TexturePVR. BGRA8888 not supported on this device")
		//return false
		//}
		default:
			blockSize = 1
			widthBlocks = width
			heightBlocks = height

		}

		// Clamp to minimum number of blocks
		if widthBlocks < 2 {
			widthBlocks = 2
		}
		if heightBlocks < 2 {
			heightBlocks = 2
		}

		dataSize = widthBlocks * heightBlocks * ((blockSize * p.PixelFormatInfo.bpp) / 8)
		//packetLen := ((uint)dataLen-dataOffset)
		packetLen := 0
		if packetLen > dataSize {
			packetLen = dataSize
		}

		p.Mipmaps[i].addr = bytes + dataOffset
		p.Mipmaps[i].len = packetLen

		dataOffset += packetLen
		//CCAssert(dataOffset <= dataLength, "CCTexurePVR: Invalid lenght")

		width = MAX(width>>1, 1)
		height = MAX(height>>1, 1)
	}

	return true

}
func (p *PVR) createGLTexture() bool {
	width := p.Width
	height = p.Height

	var err GLenum

	if p.NumOfMipmaps > 0 {
		if p.Name != 0 {
			ccGLDeleteTexture(p.Name)
		}

		// From PVR sources: "PVR files are never row aligned."
		glPixelStorei(GL_UNPACK_ALIGNMENT, 1)

		glGenTextures(1, &m_uName)
		ccGLBindTexture2D(m_uName)

		// Default: Anti alias.
		if m_uNumberOfMipmaps == 1 {
			glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR)
		}
	} else {
		glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR_MIPMAP_NEAREST)
	}
	glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR)

	glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_CLAMP_TO_EDGE)
	glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_CLAMP_TO_EDGE)

	//CHECK_GL_ERROR_DEBUG() // clean possible GL error

	internalFormat := p.PixelFormatInfo.internalFormat

	format := p.PixelFormatInfo.format
	typ = p.PixelFormatInfo.typ
	compressed := p.PixelFormatInfo.compressed

	// Generate textures with mipmaps
	for i := 0; i < p.NumOfMipmaps; i++ {
		//if (compressed && ! CCConfiguration::sharedConfiguration()->supportsPVRTC()){
		////CCLOG("cocos2d: WARNING: PVRTC images are not supported")
		//return false
		//}

		data := p.Mipmaps[i].addr
		dataLen = p.Mipmaps[i].len

		if compressed {
			glCompressedTexImage2D(GL_TEXTURE_2D, i, internalFormat, width, height, 0, datalen, data)
		} else {
			glTexImage2D(GL_TEXTURE_2D, i, internalFormat, width, height, 0, format, typ, data)
		}

		if i > 0 && (width != height || ccNextPOT(width) != width) {
			//CCLOG("cocos2d: TexturePVR. WARNING. Mipmap level %u is not squared. Texture won't render correctly. width=%u != height=%u", i, width, height)
		}

		err = glGetError()
		if err != GL_NO_ERROR {
			//CCLOG("cocos2d: TexturePVR: Error uploading compressed texture level: %u . glError: 0x%04X", i, err)
			return false
		}

		width = MAX(width>>1, 1)
		height = MAX(height>>1, 1)
	}

	return true

}

/*initWithContentsOfFile initializes a TexturePVR with a path */
func (p *PVR) initWithContentsOfFile(path string) bool {

	var (
		pvrData *byte
		pvrLen  = 0
	)

	path = strings.ToLower(path)

	if strings.Contains(path, ".ccz") {

		pvrLen = InflateCCZFile(path, &pvrData)
	} else if strings.Contains(path, ".gz") {
		pvrLen = InflateGZipFile(path, &pvrData)
	} else {

		pvrData = sharedFileUtils().getFileData(path, "rb", (*uint64)&pvrLen)
	}

	if pvrLen < 0 {
		return false
	}

	p.NumberOfMipmaps = 0

	p.Name = 0
	p.Width = 0
	p.Height = 0
	p.PixelFormatInfo = NULL
	p.HasAlpha = false
	p.ForcePremultipliedAlpha = false
	p.HasPremultipliedAlpha = false

	p.RetainName = false // cocos2d integration

	if !((p.unpackPVRv2Data(pvrData, pvrLen) || p.unpackPVRv3Data(pvrData, pvrLen)) && p.createGLTexture()) {

		return false
	}

	return true

}

func (p *PVR) Release() {
	if p.Name != 0 && !p.RetainName {
		GLDeleteTexture(p.Name)
	}
}

/** NewPVR creates and initializes a TexturePVR with a path */
func NewPVR(path string) *PVR {

	p := &PVR{}
	if !p.initWithContentsOfFile(path) {
		return nil
	}
	return &p

}
