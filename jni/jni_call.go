package jni

import "C"

var reg Method

func RegisterMethod(method Method) {
	reg = method
}

//export JNI_OnLoad
func JNI_OnLoad(vm Javavm, reserved uintptr) JInt {
	if reg != nil {
		return reg.JniOnLoad(vm, reserved)
	} else {
		return JNI_VERSION_1_6
	}
}

//export JNI_OnUnload
func JNI_OnUnload(vm Javavm, reserved uintptr) {
	if reg != nil {
		reg.JniOnUnload(vm, reserved)
	}
}
