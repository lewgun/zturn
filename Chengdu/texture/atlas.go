package texture

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
}

/** creates a TextureAtlas with an filename and with an initial capacity for Quads.
* The TextureAtlas capacity can be increased in runtime.
 */
func (a *Atlas) NewAtlas(file string, cap uint) *Atlas {
}

/** initializes a TextureAtlas with a filename and with a certain capacity for Quads.
* The TextureAtlas capacity can be increased in runtime.
*
* WARNING: Do not reinitialize the TextureAtlas because it will leak memory (issue #706)
 */
func (a *Atlas) InitWithFile(file string, cap uint) bool

/** creates a TextureAtlas with a previously initialized Texture2D object, and
* with an initial capacity for n Quads.
* The TextureAtlas capacity can be increased in runtime.
 */
func (a *Atlas) NewAtlasWithTexture(tex *Texture, cap uint) *Atlas {
}

/** initializes a TextureAtlas with a previously initialized Texture2D object, and
* with an initial capacity for Quads.
* The TextureAtlas capacity can be increased in runtime.
*
* WARNING: Do not reinitialize the TextureAtlas because it will leak memory (issue #706)
 */

func (a *Atlas) initWithTexture(tex *Texture, cap uint) {
}

/** updates a Quad (texture, vertex and color) at a certain index
* index must be between 0 and the atlas capacity - 1
@since v0.8
*/
func (a *Atlas) updateQuad(quad *ccV3F_C4B_T2F_Quad, index uint) {

}

/** Inserts a Quad (texture, vertex and color) at a certain index
index must be between 0 and the atlas capacity - 1
@since v0.8
*/
func (a *Atlas) insertQuad(quad *V3F_C4B_T2F_Quad, index uint) {

}

/** Inserts a c array of quads at a given index
index must be between 0 and the atlas capacity - 1
this method doesn't enlarge the array when amount + index > totalQuads
@since v1.1
*/
func (a *Atlas) insertQuads(quads *V3F_C4B_T2F_Quad,
	index uint,
	amount uint)

/** Removes the quad that is located at a certain index and inserts it at a new index
This operation is faster than removing and inserting in a quad in 2 different steps
@since v0.7.2
*/
func (a *Atlas) insertQuadFromIndex(from, new uint)

/** removes a quad at a given index number.
The capacity remains the same, but the total number of quads to be drawn is reduced in 1
@since v0.7.2
*/
func (a *Atlas) removeQuadAtIndex(index uint)

/** removes a amount of quads starting from index
@since 1.1
*/
func (a *Atlas) removeQuadsAtIndex(index, amount uint)

/** removes all Quads.
The TextureAtlas capacity remains untouched. No memory is freed.
The total number of quads to be drawn will be 0
@since v0.7.2
*/
func (a *Atlas) removeAllQuads()

/** resize the capacity of the CCTextureAtlas.
* The new capacity can be lower or higher than the current one
* It returns YES if the resize was successful.
* If it fails to resize the capacity it will return NO with a new capacity of 0.
 */
func (a *Atlas) resizeCapacity(n uint)

/**
Used internally by CCParticleBatchNode
don't use this unless you know what you're doing
@since 1.1
*/
func (a *Atlas) increaseTotalQuadsWith(amount uint)

/** Moves an amount of quads from oldIndex at newIndex
@since v1.1
*/
func (a *Atlas) moveQuadsFromIndex(oldIndex, amount, newIndex uint)

/**
Moves quads from index till totalQuads to the newIndex
Used internally by CCParticleBatchNode
This method doesn't enlarge the array if newIndex + quads to be moved > capacity
@since 1.1
*/
func (a *Atlas) moveQuadsFromIndex(current, new uint)

/**
Ensures that after a realloc quads are still empty
Used internally by CCParticleBatchNode
@since 1.1
*/
func (a *Atlas) fillWithEmptyQuadsFromIndex(index, amount uint)

/** draws n quads
* n can't be greater than the capacity of the Atlas
 */
func (a *Atlas) drawNumberOfQuads(n int) {
}

/** draws n quads from an index (offset).
n + start can't be greater than the capacity of the atlas
@since v1.0
*/
func (a *Atlas) drawNumberOfQuads(n, start uint)

/** draws all the Atlas's Quads
 */
func (a *Atlas) drawQuads()

/** listen the event that coming to foreground on Android
 */
func (a *Atlas) listenBackToForeground( /*obj *CCObject*/ ) {
}

/** whether or not the array buffer of the VBO needs to be updated*/
/** specify if the array buffer of the VBO needs to be updated */

func (a *Atlas) setupIndices(void) {}

func (a *Atlas) mapBuffers() {
}

func (a *Atlas) setupVBOandVAO(void) {}

func (a *Atlas) setupVBO() {
}
