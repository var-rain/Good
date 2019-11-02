package jni

/*
#include<jni.h>

JNIEXPORT jint JNI_OnLoad(JavaVM* vm, void* reserved){
	return JniOnLoad(vm,reserved);
}

JNIEXPORT void JNI_OnUnload(JavaVM* vm, void* reserved){
	JniOnUnload(vm,reserved);
}
*/
import "C"
import "unsafe"

// Primitive types that match up with Java equivalents.
type JBoolean uint8  // unsigned 8 bits
type JByte int8      // signed 8 bits
type JChar uint16    // unsigned 16 bits
type JShort int16    // signed 16 bits
type JInt int32      // signed 32 bits
type JLong int64     // signed 64 bits
type JFloat float32  // 32-bit IEEE 754
type JDouble float64 // 64-bit IEEE 754

// "cardinal indices and sizes"
type JSize int32

// Reference types, in Go.
type JObject unsafe.Pointer
type JClass unsafe.Pointer
type JString unsafe.Pointer
type JArray unsafe.Pointer
type JObjectArray unsafe.Pointer
type JBooleanArray unsafe.Pointer
type JByteArray unsafe.Pointer
type JCharArray unsafe.Pointer
type JShortArray unsafe.Pointer
type JIntArray unsafe.Pointer
type JLongArray unsafe.Pointer
type JFloatArray unsafe.Pointer
type JDoubleArray unsafe.Pointer
type JThrowable unsafe.Pointer
type JWeak unsafe.Pointer

type JFieldID unsafe.Pointer  // field IDs
type JMethodID unsafe.Pointer // method IDs

type JValue unsafe.Pointer

type JObjectRefType int

const (
	JNIInvalidRefType    = 0
	JNILocalRefType      = 1
	JNIGlobalRefType     = 2
	JNIWeakGlobalRefType = 3
)

type JNINativeMethod struct {
	Name      *C.char
	Signature *C.char
	FnPtr     unsafe.Pointer
}

type JNIEnv unsafe.Pointer
type Javavm unsafe.Pointer

// Manifest constants.
const (
	JNI_FALSE = 0
	JNI_TRUE  = 1

	JNI_VERSION_1_1 = 0x00010001
	JNI_VERSION_1_2 = 0x00010002
	JNI_VERSION_1_4 = 0x00010004
	JNI_VERSION_1_6 = 0x00010006

	JNI_OK        = 0  // no error
	JNI_ERR       = -1 // generic error
	JNI_EDETACHED = -2 // thread detached from the VM
	JNI_EVERSION  = -3 // JNI version error
	JNI_ENOMEM    = -4 // Out of memory
	JNI_EEXIST    = -5 // VM already created
	JNI_EINVAL    = -6 // Invalid argument

	JNI_COMMIT = 1 // copy content, do not free buffer
	JNI_ABORT  = 2 // free buffer w/o copying back
)

////export JNI_OnLoad
//func JNI_OnLoad(vm Javavm, reserved unsafe.Pointer) JInt {
//
//}

////export JNI_OnUnload
//func JNI_OnUnload(vm Javavm, reserved unsafe.Pointer) {
//
//}
