package texture

type ETC struct {
	Name   uint
	width  uint
	height uint
}

func NewETC(path string) *ETC {

	e := &ETC{}
	if !e.init(path) {
		return nil
	}
	return e
}

func (e *ETC) init(path string) bool {
	/*
	   	// Only Android supports ETC file format
	   	#if (CC_TARGET_PLATFORM == CC_PLATFORM_ANDROID)
	   bool ret = loadTexture(CCFileUtils::sharedFileUtils()->fullPathForFilename(file).c_str());
	   return ret;
	   #else
	   return false;
	   #endif
	*/
}

//// Call back function for java
//#if (CC_TARGET_PLATFORM == CC_PLATFORM_ANDROID)
//#define  LOG_TAG    "CCTextureETC.cpp"
//#define  LOGD(...)  __android_log_print(ANDROID_LOG_DEBUG,LOG_TAG,__VA_ARGS__)
//
//static unsigned int sWidth = 0;
//static unsigned int sHeight = 0;
//static unsigned char *sData = NULL;
//static unsigned int sLength = 0;
//
//extern "C"
//{
//JNIEXPORT void JNICALL Java_org_cocos2dx_lib_Cocos2dxETCLoader_nativeSetTextureInfo(JNIEnv* env, jobject thiz, jint width, jint height, jbyteArray data, jint dataLength)
//{
//sWidth = (unsigned int)width;
//sHeight = (unsigned int)height;
//sLength = dataLength;
//sData = new unsigned char[sLength];
//env->GetByteArrayRegion(data, 0, sLength, (jbyte*)sData);
//}
//}
//#endif

func (e *ETC) LoadTexture(file string) bool {
	//	#if (CC_TARGET_PLATFORM == CC_PLATFORM_ANDROID)
	//JniMethodInfo t;
	//if (JniHelper::getStaticMethodInfo(t, "org/cocos2dx/lib/Cocos2dxETCLoader", "loadTexture", "(Ljava/lang/String;)Z"))
	//{
	//jstring stringArg1 = t.env->NewStringUTF(file);
	//jboolean ret = t.env->CallStaticBooleanMethod(t.classID, t.methodID, stringArg1);
	//
	//t.env->DeleteLocalRef(stringArg1);
	//t.env->DeleteLocalRef(t.classID);
	//
	//if (ret)
	//{
	//_width = sWidth;
	//_height = sHeight;
	//
	//
	//glGenTextures(1, &_name);
	//glBindTexture(GL_TEXTURE_2D, _name);
	//
	//glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR);
	//glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_CLAMP_TO_EDGE);
	//glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_CLAMP_TO_EDGE);
	//
	//glCompressedTexImage2D(GL_TEXTURE_2D, 0, GL_ETC1_RGB8_OES, _width, _height, 0, sLength, sData);
	//
	//glBindTexture(GL_TEXTURE_2D, 0);
	//
	//delete [] sData;
	//sData = NULL;
	//
	//GLenum err = glGetError();
	//if (err != GL_NO_ERROR)
	//{
	//LOGD("width %d, height %d, lenght %d", _width, _height, sLength);
	//LOGD("cocos2d: TextureETC: Error uploading compressed texture %s glError: 0x%04X", file, err);
	//return false;
	//}
	//
	//return true;
	//}
	//else
	//{
	//return false;
	//}
	//}
	//#else
	//return false;
	//#endif
}
