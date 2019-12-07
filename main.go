package main

import (
	"fmt"
	"github.com/var-rain/Good/jni"
)

type Jni struct{}

func init() {
	jni.RegisterMethod(new(Jni))
}

func (j Jni) JniOnLoad(vm jni.Javavm, reserved uintptr) jni.JInt {
	env, _ := vm.GetEnv(jni.JNI_VERSION_1_6)
	class := env.FindClass("net.lingin.good.Main")
	methods := []jni.JNINativeMethod{
		{
			Name:      "sayHello",
			Signature: "()V",
			FnPtr:     SayHello,
		},
	}
	env.RegisterNatives(class, methods, jni.JInt(len(methods)))
	return jni.JNI_VERSION_1_6
}

func (j Jni) JniOnUnload(vm jni.Javavm, reserved uintptr) {
}

//export SayHello
func SayHello() {
	println("Hello Golang for java with jni")
}

func main() {
	var p = SayHello
	fmt.Println(p)
}
