package texture

import (
	"fmt"
	"unsafe"
)

/****************************************************************************
Copyright (c) 2010-2012 cocos2d-x.org
Copyright (c) 2008-2010 Ricardo Quesada
Copyright (c) 2011      Zynga Inc.
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

/* A class that implements a Texture Atlas.
Supported features:
* The atlas file can be a PVRTC, PNG or any other format supported by Texture2D
* Quads can be updated in runtime
* Quads can be added in runtime
* Quads can be removed in runtime
* Quads can be re-ordered in runtime
* The TextureAtlas capacity can be increased or decreased in runtime
* OpenGL component: V3F, C4B, T2F.
The quads are rendered using an OpenGL ES VBO.
To render the quads using an interleaved vertex array list, you should modify the ccConfig.h file
*/
type Atlas struct {
	Indices *uint16

	VAOName uint

	BuffersVBO [2]uint //0: vertex  1: indices
	IsDirty    bool    //indicates whether or not the array buffer of the VBO needs to be updated

	/** quantity of quads that are going to be drawn */
	TotalQuads uint

	/** quantity of quads that can be stored with the current texture atlas size */
	Capacity uint

	/** Texture of the texture atlas */
	Texture *Texture

	/** Quads that are going to be rendered */
	Quads *V3F_C4B_T2F_Quad
}

func (a *Atlas) String() string {
	return fmt.Sprintf("<CCTextureAtlas | totalQuads = %u>", p.TotalQuads)
}

/** creates a TextureAtlas with an filename and with an initial capacity for Quads.
* The TextureAtlas capacity can be increased in runtime.
 */
func (a *Atlas) NewAtlas(file string, cap uint) *Atlas {

	a := &Atlas{}

	if !a.InitWithFile(file, cap) {
		return nil
	}

	return a
}

/** initializes a TextureAtlas with a filename and with a certain capacity for Quads.
* The TextureAtlas capacity can be increased in runtime.
*
* WARNING: Do not reinitialize the TextureAtlas because it will leak memory (issue #706)
 */
func (a *Atlas) InitWithFile(file string, cap uint) bool {
	tex := sharedCache().AddImage(file)

	if tex == nil {
		return false
	}

	return a.initWithTexture(tex, cap)
}

/** creates a TextureAtlas with a previously initialized Texture2D object, and
* with an initial capacity for n Quads.
* The TextureAtlas capacity can be increased in runtime.
 */
func (a *Atlas) NewAtlasWithTexture(tex *Texture, cap uint) *Atlas {
	a := &Atlas{}

	if !a.InitWithFile(tex, cap) {
		return nil
	}

	return a
}

/** initializes a TextureAtlas with a previously initialized Texture2D object, and
* with an initial capacity for Quads.
* The TextureAtlas capacity can be increased in runtime.
*
* WARNING: Do not reinitialize the TextureAtlas because it will leak memory (issue #706)
 */

func (a *Atlas) initWithTexture(tex *Texture, cap uint) {

	if a.Quads != nil || a.Indices != nil {
		return
	}

	a.Capacity = cap

	a.TotalQuads = 0
	a.Texture = tex

	a.Quads = make([]byte, cap*unsafe.Sizeof(V3F_C4B_T2F_Quad))
	a.Indices = make([]byte, cap*6*unsafe.Sizeof(uint16))

	if !(a.Quads && a.Indices) && cap > 0 {
		return false
	}
	a.setupIndices()

	a.setupVBOandVAO()
	a.setupVBO()

	a.IsDirty = true
	return true

}

/** updates a Quad (texture, vertex and color) at a certain index
* index must be between 0 and the atlas capacity - 1
@since v0.8
*/
func (a *Atlas) updateQuad(quad *V3F_C4B_T2F_Quad, index uint) {
	if index < 0 || index >= a.Capacity {
		return
	}

	a.TotalQuads = Max(index+1, a.TotalQuads)
	a.Quads[index] = *quad
	a.IsDirty = true

}

/** Inserts a Quad (texture, vertex and color) at a certain index
index must be between 0 and the atlas capacity - 1
@since v0.8
*/
func (a *Atlas) insertQuad(quad *V3F_C4B_T2F_Quad, index uint) {

	if index < 0 || index >= a.Capacity {
		return
	}

	a.TotalQuads++
	if index > a.Capacity {
		return
	}

	remaining := (a.TotalQuads - 1) - index

	if remaining > 0 {
		// texture coordinates
		//memmove( &m_pQuads[index+1],&m_pQuads[index], sizeof(m_pQuads[0]) * remaining );
	}

	a.Quads[index] = *quad
	a.IsDirty = true

}

/** Inserts a c array of quads at a given index
index must be between 0 and the atlas capacity - 1
this method doesn't enlarge the array when amount + index > totalQuads
@since v1.1
*/
func (a *Atlas) insertQuads(quads *V3F_C4B_T2F_Quad,
	index uint,
	amount uint) {
	if index+amount > a.Capacity {
		return
	}
	a.TotalQuads += amount

	if a.TotalQuads > a.Capacity {
		return
	}

	remaining := (a.TotalQuads - 1) - index - amount

	if remaining > 0 {
		// tex coordinates
		//memmove( &m_pQuads[index+amount],&m_pQuads[index], sizeof(m_pQuads[0]) * remaining );
	}

	max := index + amount
	j := 0

	for i := index; i < max; i++ {
		a.Quads[index] = quads[j]
		index++
		j++
	}

	a.IsDirty = true
}

/** Removes the quad that is located at a certain index and inserts it at a new index
This operation is faster than removing and inserting in a quad in 2 different steps
@since v0.7.2
*/
func (a *Atlas) insertQuadFromIndex(old, new uint) {
	if old == new {
		return
	}

	if new < 0 || new >= a.TotalQuads {
		return
	}
	if old < 0 || old >= a.TotalQuads {
		return
	}

	howMany := old - new
	if howMany <= 0 {
		howMany = new - old
	}

	dst := old
	src := old + 1

	if old > new {
		dst = new + 1
		src = new
	}

	quadsBackup := a.Quads[old]
	// memmove( &m_pQuads[dst],&m_pQuads[src], sizeof(m_pQuads[0]) * howMany );
	a.Quads[new] = quadsBackup
	a.IsDirty = true

}

/** removes a quad at a given index number.
The capacity remains the same, but the total number of quads to be drawn is reduced in 1
@since v0.7.2
*/
func (a *Atlas) removeQuadAtIndex(index uint) {
	if index >= a.TotalQuads {
		return
	}

	remaining := (a.TotalQuads - 1) - index
	if remaining != 0 {
		// texture coordinates
		//memmove( &m_pQuads[index],&m_pQuads[index+1], sizeof(m_pQuads[0]) * remaining );
	}

	a.TotalQuads--
	a.IsDirty = true
}

/** removes a amount of quads starting from index
@since 1.1
*/
func (a *Atlas) removeQuadsAtIndex(index, amount uint) {
	if index+amount > a.TotalQuads {
		return
	}

	remaining := a.TotalQuads - (index + amount)
	a.TotalQuads -= amount

	if remaining != 0 {
		memmove(&m_pQuads[index], &m_pQuads[index+amount], sizeof(m_pQuads[0])*remaining)
	}

	a.IsDirty = true
}

/** removes all Quads.
The TextureAtlas capacity remains untouched. No memory is freed.
The total number of quads to be drawn will be 0
@since v0.7.2
*/
func (a *Atlas) removeAllQuads() {
	a.TotalQuads = 0
}

/** resize the capacity of the CCTextureAtlas.
* The new capacity can be lower or higher than the current one
* It returns YES if the resize was successful.
* If it fails to resize the capacity it will return NO with a new capacity of 0.
 */
func (a *Atlas) resizeCapacity(new uint) {
	if a.Capacity == new {
		return true
	}

	old := a.Capacity
	a.TotalQuads = min(a.TotalQuads, new)
	a.Capacity = new

	var (
		tempQuads   *V3F_C4B_T2F_Quad
		tempIndices *uint16
	)

	if a.Quads == nil {
		tempQuads = make([]byte, a.Capacity*unsafe.Sizeof(a.Quads[0]))

	} else {

		//		tmpQuads = (ccV3F_C4B_T2F_Quad*)realloc( m_pQuads, sizeof(m_pQuads[0]) * m_uCapacity );
		//		if (tmpQuads != NULL && m_uCapacity > uOldCapactiy)
		//		{
		//			memset(tmpQuads+uOldCapactiy, 0, (m_uCapacity - uOldCapactiy)*sizeof(m_pQuads[0]) );
		//		}
	}

	if a.Indices == nil {
		tempQuads = make([]byte, a.Capacity*unsafe.Sizeof(a.Quads[0]))

	} else {

		//tmpIndices = (GLushort*)realloc( m_pIndices, sizeof(m_pIndices[0]) * m_uCapacity * 6 );
		//if (tmpIndices != NULL && m_uCapacity > uOldCapactiy)
		//{
		//memset( tmpIndices+uOldCapactiy, 0, (m_uCapacity-uOldCapactiy) * 6 * sizeof(m_pIndices[0]) );
		//}
	}

	if tempQuads == nil || tempIndices == nil {
		a.Quads = nil
		a.Indices = nil
		a.Capacity = 0
		a.TotalQuads = 0
		return false
	}
	a.Quads = tempQuads
	a.Indices = tempIndices

	a.setupIndices()
	a.mapBuffers()
	a.IsDirty = true
	return true

}

/**
Used internally by CCParticleBatchNode
don't use this unless you know what you're doing
@since 1.1
*/
func (a *Atlas) increaseTotalQuadsWith(amount uint) {
	a.TotalQuads += amount
}

/** Moves an amount of quads from oldIndex at newIndex
@since v1.1
*/
func (a *Atlas) moveQuadsFromIndex(oldIndex, amount, newIndex uint) {
	if newIndex+amount > a.TotalQuads {
		return
	}
	if oldIndex >= a.TotalQuads {
		return
	}

	if oldIndex == newIndex {
		return
	}

	quadSize := unsafe.Sizeof(V3F_C4B_T2F_Quad)
	tempQuads := make([]byte, quadSize*amount)

	memcpy(tempQuads, &m_pQuads[oldIndex], quadSize*amount)

	//	if (newIndex < oldIndex)
	//	{
	//		// move quads from newIndex to newIndex + amount to make room for buffer
	//		memmove( &m_pQuads[newIndex], &m_pQuads[newIndex+amount], (oldIndex-newIndex)*quadSize);
	//	}
	//else
	//{
	//// move quads above back
	//memmove( &m_pQuads[oldIndex], &m_pQuads[oldIndex+amount], (newIndex-oldIndex)*quadSize);
	//}
	//memcpy( &m_pQuads[newIndex], tempQuads, amount*quadSize);

	a.IsDirty = true
}

/**
Moves quads from index till totalQuads to the newIndex
Used internally by CCParticleBatchNode
This method doesn't enlarge the array if newIndex + quads to be moved > capacity
@since 1.1
*/
func (a *Atlas) moveQuadsFromIndex(index, new uint) {
	if new+(a.TotalQuads-index) > a.Capacity {
		return
	}
	//memmove(m_pQuads + newIndex,m_pQuads + index, (m_uTotalQuads - index) * sizeof(m_pQuads[0]));

}

/**
Ensures that after a realloc quads are still empty
Used internally by CCParticleBatchNode
@since 1.1
*/
func (a *Atlas) fillWithEmptyQuadsFromIndex(index, amount uint) {
	//	ccV3F_C4B_T2F_Quad quad;
	//	memset(&quad, 0, sizeof(quad));
	//
	//	unsigned int to = index + amount;
	//	for (unsigned int i = index ; i < to ; i++)
	//	{
	//	m_pQuads[i] = quad;
	//	}
}

/** draws n quads
* n can't be greater than the capacity of the Atlas
 */
func (a *Atlas) drawNumberOfQuads(n int) {
	a.drawNumberOfQuads(n, 0)
}

/** draws n quads from an index (offset).
n + start can't be greater than the capacity of the atlas
@since v1.0
*/
func (a *Atlas) drawNumberOfQuads(n, start uint) {
	if n == 0 {
		return
	}
	/*
		 ccGLBindTexture2D(m_pTexture->getName());

	#if CC_TEXTURE_ATLAS_USE_VAO

	    //
	    // Using VBO and VAO
	    //

	    // XXX: update is done in draw... perhaps it should be done in a timer
	    if (m_bDirty)
	    {
	        glBindBuffer(GL_ARRAY_BUFFER, m_pBuffersVBO[0]);
	        // option 1: subdata
	        //glBufferSubData(GL_ARRAY_BUFFER, sizeof(m_pQuads[0])*start, sizeof(m_pQuads[0]) * n , &m_pQuads[start] );

			// option 2: data
	        //		glBufferData(GL_ARRAY_BUFFER, sizeof(quads_[0]) * (n-start), &quads_[start], GL_DYNAMIC_DRAW);

			// option 3: orphaning + glMapBuffer
			glBufferData(GL_ARRAY_BUFFER, sizeof(m_pQuads[0]) * (n-start), NULL, GL_DYNAMIC_DRAW);
			void *buf = glMapBuffer(GL_ARRAY_BUFFER, GL_WRITE_ONLY);
			memcpy(buf, m_pQuads, sizeof(m_pQuads[0])* (n-start));
			glUnmapBuffer(GL_ARRAY_BUFFER);

			glBindBuffer(GL_ARRAY_BUFFER, 0);

	        m_bDirty = false;
	    }

	    ccGLBindVAO(m_uVAOname);

	#if CC_REBIND_INDICES_BUFFER
	    glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, m_pBuffersVBO[1]);
	#endif

	#if CC_TEXTURE_ATLAS_USE_TRIANGLE_STRIP
	    glDrawElements(GL_TRIANGLE_STRIP, (GLsizei) n*6, GL_UNSIGNED_SHORT, (GLvoid*) (start*6*sizeof(m_pIndices[0])) );
	#else
	    glDrawElements(GL_TRIANGLES, (GLsizei) n*6, GL_UNSIGNED_SHORT, (GLvoid*) (start*6*sizeof(m_pIndices[0])) );
	#endif // CC_TEXTURE_ATLAS_USE_TRIANGLE_STRIP

	#if CC_REBIND_INDICES_BUFFER
	    glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, 0);
	#endif

	//    glBindVertexArray(0);

	#else // ! CC_TEXTURE_ATLAS_USE_VAO

	    //
	    // Using VBO without VAO
	    //

	#define kQuadSize sizeof(m_pQuads[0].bl)
	    glBindBuffer(GL_ARRAY_BUFFER, m_pBuffersVBO[0]);

	    // XXX: update is done in draw... perhaps it should be done in a timer
	    if (m_bDirty)
	    {
	        glBufferSubData(GL_ARRAY_BUFFER, sizeof(m_pQuads[0])*start, sizeof(m_pQuads[0]) * n , &m_pQuads[start] );
	        m_bDirty = false;
	    }

	    ccGLEnableVertexAttribs(kCCVertexAttribFlag_PosColorTex);

	    // vertices
	    glVertexAttribPointer(kCCVertexAttrib_Position, 3, GL_FLOAT, GL_FALSE, kQuadSize, (GLvoid*) offsetof(ccV3F_C4B_T2F, vertices));

	    // colors
	    glVertexAttribPointer(kCCVertexAttrib_Color, 4, GL_UNSIGNED_BYTE, GL_TRUE, kQuadSize, (GLvoid*) offsetof(ccV3F_C4B_T2F, colors));

	    // tex coords
	    glVertexAttribPointer(kCCVertexAttrib_TexCoords, 2, GL_FLOAT, GL_FALSE, kQuadSize, (GLvoid*) offsetof(ccV3F_C4B_T2F, texCoords));

	    glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, m_pBuffersVBO[1]);

	#if CC_TEXTURE_ATLAS_USE_TRIANGLE_STRIP
	    glDrawElements(GL_TRIANGLE_STRIP, (GLsizei)n*6, GL_UNSIGNED_SHORT, (GLvoid*) (start*6*sizeof(m_pIndices[0])));
	#else
	    glDrawElements(GL_TRIANGLES, (GLsizei)n*6, GL_UNSIGNED_SHORT, (GLvoid*) (start*6*sizeof(m_pIndices[0])));
	#endif // CC_TEXTURE_ATLAS_USE_TRIANGLE_STRIP

	    glBindBuffer(GL_ARRAY_BUFFER, 0);
	    glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, 0);

	#endif // CC_TEXTURE_ATLAS_USE_VAO

	    CC_INCREMENT_GL_DRAWS(1);
	    CHECK_GL_ERROR_DEBUG();
	*/
}

/** draws all the Atlas's Quads
 */
func (a *Atlas) drawQuads() {
	a.drawNumberOfQuads(a.TotalQuads, 0)
}

/** listen the event that coming to foreground on Android
 */
func (a *Atlas) listenBackToForeground( /*obj *CCObject*/ ) {
	//	#if CC_TEXTURE_ATLAS_USE_VAO
	//setupVBOandVAO();
	//#else
	//setupVBO();
	//#endif
	//
	//// set m_bDirty to true to force it rebinding buffer
	//m_bDirty = true;
}

/** whether or not the array buffer of the VBO needs to be updated*/
/** specify if the array buffer of the VBO needs to be updated */

func (a *Atlas) setupIndices() {

	if a.Capacity == 0 {
		return
	}

	for i := 0; i < a.Capacity; i++ {

		//	#if CC_TEXTURE_ATLAS_USE_TRIANGLE_STRIP
		a.Indices[i*6+0] = i*4 + 0
		a.Indices[i*6+1] = i*4 + 0
		a.Indices[i*6+2] = i*4 + 2
		a.Indices[i*6+3] = i*4 + 1
		a.Indices[i*6+4] = i*4 + 3
		a.Indices[i*6+5] = i*4 + 3
		//	#else
		//	p.Indices[i*6+0] = i*4+0;
		//	p.Indices[i*6+1] = i*4+1;
		//	p.Indices[i*6+2] = i*4+2;
		//
		//	// inverted index. issue #179
		//	p.Indices[i*6+3] = i*4+3;
		//	p.Indices[i*6+4] = i*4+2;
		//	p.Indices[i*6+5] = i*4+1;
		//	#endif
	}
}

func (a *Atlas) mapBuffers() {
	//	// Avoid changing the element buffer for whatever VAO might be bound.
	//	ccGLBindVAO(0);
	//
	//	glBindBuffer(GL_ARRAY_BUFFER, m_pBuffersVBO[0]);
	//	glBufferData(GL_ARRAY_BUFFER, sizeof(m_pQuads[0]) * m_uCapacity, m_pQuads, GL_DYNAMIC_DRAW);
	//	glBindBuffer(GL_ARRAY_BUFFER, 0);
	//
	//	glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, m_pBuffersVBO[1]);
	//	glBufferData(GL_ELEMENT_ARRAY_BUFFER, sizeof(m_pIndices[0]) * m_uCapacity * 6, m_pIndices, GL_STATIC_DRAW);
	//	glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, 0);
	//
	//	CHECK_GL_ERROR_DEBUG();
}

func (a *Atlas) setupVBOandVAO() {
	//	glGenVertexArrays(1, &m_uVAOname);
	//	ccGLBindVAO(m_uVAOname);
	//
	//	#define kQuadSize sizeof(m_pQuads[0].bl)
	//
	//	glGenBuffers(2, &m_pBuffersVBO[0]);
	//
	//	glBindBuffer(GL_ARRAY_BUFFER, m_pBuffersVBO[0]);
	//	glBufferData(GL_ARRAY_BUFFER, sizeof(m_pQuads[0]) * m_uCapacity, m_pQuads, GL_DYNAMIC_DRAW);
	//
	//	// vertices
	//	glEnableVertexAttribArray(kCCVertexAttrib_Position);
	//	glVertexAttribPointer(kCCVertexAttrib_Position, 3, GL_FLOAT, GL_FALSE, kQuadSize, (GLvoid*) offsetof( ccV3F_C4B_T2F, vertices));
	//
	//// colors
	//glEnableVertexAttribArray(kCCVertexAttrib_Color);
	//glVertexAttribPointer(kCCVertexAttrib_Color, 4, GL_UNSIGNED_BYTE, GL_TRUE, kQuadSize, (GLvoid*) offsetof( ccV3F_C4B_T2F, colors));
	//
	//// tex coords
	//glEnableVertexAttribArray(kCCVertexAttrib_TexCoords);
	//glVertexAttribPointer(kCCVertexAttrib_TexCoords, 2, GL_FLOAT, GL_FALSE, kQuadSize, (GLvoid*) offsetof( ccV3F_C4B_T2F, texCoords));
	//
	//glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, m_pBuffersVBO[1]);
	//glBufferData(GL_ELEMENT_ARRAY_BUFFER, sizeof(m_pIndices[0]) * m_uCapacity * 6, m_pIndices, GL_STATIC_DRAW);
	//
	//// Must unbind the VAO before changing the element buffer.
	//ccGLBindVAO(0);
	//glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, 0);
	//glBindBuffer(GL_ARRAY_BUFFER, 0);
	//
	//CHECK_GL_ERROR_DEBUG();
}

func (a *Atlas) setupVBO() {
	//	glGenBuffers(2, &m_pBuffersVBO[0]);
	//
	//	mapBuffers();
}

func (a *Atlas) Release() {
//	glDeleteBuffers(2, a.BuffersVBO)
//
//	glDeleteVertexArrays(1, &a.VAOName)
//	ccGLBindVAO(0)

}
