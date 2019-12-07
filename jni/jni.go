package jni

/*

#cgo CFLAGS: -ID:/Oracle/Java/include -ID:/Oracle/Java/include/win32

#include <jni.h>
#include <stdlib.h>

// ------------------------------- JavaVM --------------------------------

static inline jint DestroyJavaVM(JavaVM *vm) {
    return (*vm)->DestroyJavaVM(vm);
}

static inline jint AttachCurrentThread(JavaVM *vm, void **penv, void *args) {
    return (*vm)->AttachCurrentThread(vm, penv, args);
}

static inline jint DetachCurrentThread(JavaVM *vm) {
    return (*vm)->DetachCurrentThread(vm);
}

static inline jint GetEnv(JavaVM *vm, void **penv, jint version) {
    return (*vm)->GetEnv(vm, penv, version);
}

static inline jint AttachCurrentThreadAsDaemon(JavaVM *vm, void **penv, void *args) {
    return (*vm)->AttachCurrentThreadAsDaemon(vm, penv, args);
}

// ------------------------------- JNIEnv --------------------------------

static inline jint GetVersion(JNIEnv *env) {
    return (*env)->GetVersion(env);
}

static inline jclass DefineClass(JNIEnv *env, const char *name, jobject loader, const jbyte *buf, jsize len) {
    return (*env)->DefineClass(env, name, loader, buf, len);
}

static inline jclass FindClass(JNIEnv *env, const char *name) {
    return (*env)->FindClass(env, name);
}

static inline jmethodID FromReflectedMethod(JNIEnv *env, jobject method) {
    return (*env)->FromReflectedMethod(env, method);
}

static inline jfieldID FromReflectedField(JNIEnv *env, jobject field) {
    return (*env)->FromReflectedField(env, field);
}

static inline jobject ToReflectedMethod(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic) {
    return (*env)->ToReflectedMethod(env, cls, methodID, isStatic);
}

static inline jclass GetSuperclass(JNIEnv *env, jclass sub) {
    return (*env)->GetSuperclass(env, sub);
}

static inline jboolean IsAssignableFrom(JNIEnv *env, jclass sub, jclass sup) {
    return (*env)->IsAssignableFrom(env, sub, sup);
}

static inline jobject ToReflectedField(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic) {
    return (*env)->ToReflectedField(env, cls, fieldID, isStatic);
}

static inline jint Throw(JNIEnv *env, jthrowable obj) {
    return (*env)->Throw(env, obj);
}

static inline jint ThrowNew(JNIEnv *env, jclass clazz, const char *msg) {
    return (*env)->ThrowNew(env, clazz, msg);
}

static inline jthrowable ExceptionOccurred(JNIEnv *env) {
    return (*env)->ExceptionOccurred(env);
}

static inline void ExceptionDescribe(JNIEnv *env) {
    (*env)->ExceptionDescribe(env);
}

static inline void ExceptionClear(JNIEnv *env) {
    (*env)->ExceptionClear(env);
}

static inline void FatalError(JNIEnv *env, const char *msg) {
    (*env)->FatalError(env, msg);
}

static inline jint PushLocalFrame(JNIEnv *env, jint capacity) {
    return (*env)->PushLocalFrame(env, capacity);
}

static inline jobject PopLocalFrame(JNIEnv *env, jobject result) {
    return (*env)->PopLocalFrame(env, result);
}

static inline jobject NewGlobalRef(JNIEnv *env, jobject lobj) {
    return (*env)->NewGlobalRef(env, lobj);
}

static inline void DeleteGlobalRef(JNIEnv *env, jobject gref) {
    (*env)->DeleteGlobalRef(env, gref);
}

static inline void DeleteLocalRef(JNIEnv *env, jobject obj) {
    (*env)->DeleteLocalRef(env, obj);
}

static inline jboolean IsSameObject(JNIEnv *env, jobject obj1, jobject obj2) {
    return (*env)->IsSameObject(env, obj1, obj2);
}

static inline jobject NewLocalRef(JNIEnv *env, jobject ref) {
    return (*env)->NewLocalRef(env, ref);
}

static inline jint EnsureLocalCapacity(JNIEnv *env, jint capacity) {
    return (*env)->EnsureLocalCapacity(env, capacity);
}

static inline jobject AllocObject(JNIEnv *env, jclass clazz) {
    return (*env)->AllocObject(env, clazz);
}

//static inline jobject NewObject(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jobject result;
//    va_start(args, methodID);
//    result = (*env)->NewObjectV(env, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jobject NewObjectV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
//    return (*env)->NewObjectV(env, clazz, methodID, args);
//}

static inline jobject NewObjectA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    return (*env)->NewObjectA(env, clazz, methodID, args);
}

static inline jclass GetObjectClass(JNIEnv *env, jobject obj) {
    return (*env)->GetObjectClass(env, obj);
}

static inline jboolean IsInstanceOf(JNIEnv *env, jobject obj, jclass clazz) {
    return (*env)->IsInstanceOf(env, obj, clazz);
}

static inline jmethodID GetMethodID(JNIEnv *env, jclass clazz, const char *name, const char *sig) {
    return (*env)->GetMethodID(env, clazz, name, sig);
}

//static inline jobject CallObjectMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
//    va_list args;
//    jobject result;
//    va_start(args, methodID);
//    result = (*env)->CallObjectMethodV(env, obj, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jobject CallObjectMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
//    return (*env)->CallObjectMethodV(env, obj, methodID, args);
//}

static inline jobject CallObjectMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    return (*env)->CallObjectMethodA(env, obj, methodID, args);
}

//static inline jboolean CallBooleanMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
//    va_list args;
//    jboolean result;
//    va_start(args, methodID);
//    result = (*env)->CallBooleanMethodV(env, obj, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jboolean CallBooleanMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
//    return (*env)->CallBooleanMethodV(env, obj, methodID, args);
//}

static inline jboolean CallBooleanMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    return (*env)->CallBooleanMethodA(env, obj, methodID, args);
}

//static inline jbyte CallByteMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
//    va_list args;
//    jbyte result;
//    va_start(args, methodID);
//    result = (*env)->CallByteMethodV(env, obj, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jbyte CallByteMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
//    return (*env)->CallByteMethodV(env, obj, methodID, args);
//}

static inline jbyte CallByteMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    return (*env)->CallByteMethodA(env, obj, methodID, args);
}

//static inline jchar CallCharMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
//    va_list args;
//    jchar result;
//    va_start(args, methodID);
//    result = (*env)->CallCharMethodV(env, obj, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jchar CallCharMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
//    return (*env)->CallCharMethodV(env, obj, methodID, args);
//}

static inline jchar CallCharMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    return (*env)->CallCharMethodA(env, obj, methodID, args);
}

//static inline jshort CallShortMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
//    va_list args;
//    jshort result;
//    va_start(args, methodID);
//    result = (*env)->CallShortMethodV(env, obj, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jshort CallShortMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
//    return (*env)->CallShortMethodV(env, obj, methodID, args);
//}

static inline jshort CallShortMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    return (*env)->CallShortMethodA(env, obj, methodID, args);
}

//static inline jint CallIntMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
//    va_list args;
//    jint result;
//    va_start(args, methodID);
//    result = (*env)->CallIntMethodV(env, obj, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jint CallIntMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
//    return (*env)->CallIntMethodV(env, obj, methodID, args);
//}

static inline jint CallIntMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    return (*env)->CallIntMethodA(env, obj, methodID, args);
}

//static inline jlong CallLongMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
//    va_list args;
//    jlong result;
//    va_start(args, methodID);
//    result = (*env)->CallLongMethodV(env, obj, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jlong CallLongMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
//    return (*env)->CallLongMethodV(env, obj, methodID, args);
//}

static inline jlong CallLongMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    return (*env)->CallLongMethodA(env, obj, methodID, args);
}

//static inline jfloat CallFloatMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
//    va_list args;
//    jfloat result;
//    va_start(args, methodID);
//    result = (*env)->CallFloatMethodV(env, obj, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jfloat CallFloatMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
//    return (*env)->CallFloatMethodV(env, obj, methodID, args);
//}

static inline jfloat CallFloatMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    return (*env)->CallFloatMethodA(env, obj, methodID, args);
}

//static inline jdouble CallDoubleMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
//    va_list args;
//    jdouble result;
//    va_start(args, methodID);
//    result = (*env)->CallDoubleMethodV(env, obj, methodID, args);
//    va_end(args);
//    return result;
//}
//
//static inline jdouble CallDoubleMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
//    return (*env)->CallDoubleMethodV(env, obj, methodID, args);
//}

static inline jdouble CallDoubleMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    return (*env)->CallDoubleMethodA(env, obj, methodID, args);
}

//static inline void CallVoidMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
//    va_list args;
//    va_start(args, methodID);
//    (*env)->CallVoidMethodV(env, obj, methodID, args);
//    va_end(args);
//}

//static inline void CallVoidMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
//    return (*env)->CallVoidMethodV(env, obj, methodID, args);
//}

static inline void CallVoidMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    return (*env)->CallVoidMethodA(env, obj, methodID, args);
}

//static inline jobject CallNonvirtualObjectMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jobject result;
//    va_start(args, methodID);
//    result = (*env)->CallNonvirtualObjectMethodV(env, obj, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jobject CallNonvirtualObjectMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
//                                                  va_list args) {
//    return (*env)->CallNonvirtualObjectMethodV(env, obj, clazz, methodID, args);
//}

static inline jobject CallNonvirtualObjectMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
                                                  const jvalue *args) {
    return (*env)->CallNonvirtualObjectMethodA(env, obj, clazz, methodID, args);
}

//static inline jboolean CallNonvirtualBooleanMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jboolean result;
//    va_start(args, methodID);
//    result = (*env)->CallNonvirtualBooleanMethodV(env, obj, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jboolean CallNonvirtualBooleanMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
//                                                    va_list args) {
//    return (*env)->CallNonvirtualBooleanMethodV(env, obj, clazz, methodID, args);
//}

static inline jboolean CallNonvirtualBooleanMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
                                                    const jvalue *args) {
    return (*env)->CallNonvirtualBooleanMethodA(env, obj, clazz, methodID, args);
}

//static inline jbyte CallNonvirtualByteMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jbyte result;
//    va_start(args, methodID);
//    result = (*env)->CallNonvirtualByteMethodV(env, obj, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jbyte CallNonvirtualByteMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
//                                              va_list args) {
//    return (*env)->CallNonvirtualByteMethodV(env, obj, clazz, methodID, args);
//}

static inline jbyte CallNonvirtualByteMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
                                              const jvalue *args) {
    return (*env)->CallNonvirtualByteMethodA(env, obj, clazz, methodID, args);
}

//static inline jchar CallNonvirtualCharMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jchar result;
//    va_start(args, methodID);
//    result = (*env)->CallNonvirtualCharMethodV(env, obj, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jchar CallNonvirtualCharMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
//                                              va_list args) {
//    return (*env)->CallNonvirtualCharMethodV(env, obj, clazz, methodID, args);
//}

static inline jchar CallNonvirtualCharMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
                                              const jvalue *args) {
    return (*env)->CallNonvirtualCharMethodA(env, obj, clazz, methodID, args);
}

//static inline jshort CallNonvirtualShortMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jshort result;
//    va_start(args, methodID);
//    result = (*env)->CallNonvirtualShortMethodV(env, obj, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jshort CallNonvirtualShortMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
//                                                va_list args) {
//    return (*env)->CallNonvirtualShortMethodV(env, obj, clazz, methodID, args);
//}

static inline jshort CallNonvirtualShortMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
                                                const jvalue *args) {
    return (*env)->CallNonvirtualShortMethodA(env, obj, clazz, methodID, args);
}

//static inline jint CallNonvirtualIntMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jint result;
//    va_start(args, methodID);
//    result = (*env)->CallNonvirtualIntMethodV(env, obj, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jint CallNonvirtualIntMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
//                                            va_list args) {
//    return (*env)->CallNonvirtualIntMethodV(env, obj, clazz, methodID, args);
//}

static inline jint CallNonvirtualIntMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
                                            const jvalue *args) {
    return (*env)->CallNonvirtualIntMethodA(env, obj, clazz, methodID, args);
}

//static inline jlong CallNonvirtualLongMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jlong result;
//    va_start(args, methodID);
//    result = (*env)->CallNonvirtualLongMethodV(env, obj, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jlong CallNonvirtualLongMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
//                                              va_list args) {
//    return (*env)->CallNonvirtualLongMethodV(env, obj, clazz, methodID, args);
//}

static inline jlong CallNonvirtualLongMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
                                              const jvalue *args) {
    return (*env)->CallNonvirtualLongMethodA(env, obj, clazz, methodID, args);
}

//static inline jfloat CallNonvirtualFloatMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jfloat result;
//    va_start(args, methodID);
//    result = (*env)->CallNonvirtualFloatMethodV(env, obj, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jfloat CallNonvirtualFloatMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
//                                                va_list args) {
//    return (*env)->CallNonvirtualFloatMethodV(env, obj, clazz, methodID, args);
//}

static inline jfloat CallNonvirtualFloatMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
                                                const jvalue *args) {
    return (*env)->CallNonvirtualFloatMethodA(env, obj, clazz, methodID, args);
}

//static inline jdouble CallNonvirtualDoubleMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jdouble result;
//    va_start(args, methodID);
//    result = (*env)->CallNonvirtualDoubleMethodV(env, obj, clazz, methodID, args);
//    va_end(args);
//    return result;
//}
//
//static inline jdouble CallNonvirtualDoubleMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
//                                                  va_list args) {
//    return (*env)->CallNonvirtualDoubleMethodV(env, obj, clazz, methodID, args);
//}

static inline jdouble CallNonvirtualDoubleMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
                                                  const jvalue *args) {
    return (*env)->CallNonvirtualDoubleMethodA(env, obj, clazz, methodID, args);
}

//static inline void CallNonvirtualVoidMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    va_start(args, methodID);
//    (*env)->CallNonvirtualVoidMethodV(env, obj, clazz, methodID, args);
//    va_end(args);
//}

//static inline void CallNonvirtualVoidMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
//                                             va_list args) {
//    (*env)->CallNonvirtualVoidMethodV(env, obj, clazz, methodID, args);
//}

static inline void CallNonvirtualVoidMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,
                                             const jvalue *args) {
    (*env)->CallNonvirtualVoidMethodA(env, obj, clazz, methodID, args);
}

static inline jfieldID GetFieldID(JNIEnv *env, jclass clazz, const char *name, const char *sig) {
    return (*env)->GetFieldID(env, clazz, name, sig);
}

static inline jobject GetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    return (*env)->GetObjectField(env, obj, fieldID);
}

static inline jboolean GetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    return (*env)->GetBooleanField(env, obj, fieldID);
}

static inline jbyte GetByteField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    return (*env)->GetByteField(env, obj, fieldID);
}

static inline jchar GetCharField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    return (*env)->GetCharField(env, obj, fieldID);
}

static inline jshort GetShortField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    return (*env)->GetShortField(env, obj, fieldID);
}

static inline jint GetIntField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    return (*env)->GetIntField(env, obj, fieldID);
}

static inline jlong GetLongField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    return (*env)->GetLongField(env, obj, fieldID);
}

static inline jfloat GetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    return (*env)->GetFloatField(env, obj, fieldID);
}

static inline jdouble GetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    return (*env)->GetDoubleField(env, obj, fieldID);
}

static inline void SetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID, jobject val) {
    (*env)->SetObjectField(env, obj, fieldID, val);
}

static inline void SetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID, jboolean val) {
    (*env)->SetBooleanField(env, obj, fieldID, val);
}

static inline void SetByteField(JNIEnv *env, jobject obj, jfieldID fieldID, jbyte val) {
    (*env)->SetByteField(env, obj, fieldID, val);
}

static inline void SetCharField(JNIEnv *env, jobject obj, jfieldID fieldID, jchar val) {
    (*env)->SetCharField(env, obj, fieldID, val);
}

static inline void SetShortField(JNIEnv *env, jobject obj, jfieldID fieldID, jshort val) {
    (*env)->SetShortField(env, obj, fieldID, val);
}

static inline void SetIntField(JNIEnv *env, jobject obj, jfieldID fieldID, jint val) {
    (*env)->SetIntField(env, obj, fieldID, val);
}

static inline void SetLongField(JNIEnv *env, jobject obj, jfieldID fieldID, jlong val) {
    (*env)->SetLongField(env, obj, fieldID, val);
}

static inline void SetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID, jfloat val) {
    (*env)->SetFloatField(env, obj, fieldID, val);
}

static inline void SetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID, jdouble val) {
    (*env)->SetDoubleField(env, obj, fieldID, val);
}

static inline jmethodID GetStaticMethodID(JNIEnv *env, jclass clazz, const char *name, const char *sig) {
    return (*env)->GetStaticMethodID(env, clazz, name, sig);
}

//static inline jobject CallStaticObjectMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jobject result;
//    va_start(args, methodID);
//    result = (*env)->CallStaticObjectMethodV(env, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jobject CallStaticObjectMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
//    return (*env)->CallStaticObjectMethodV(env, clazz, methodID, args);
//}

static inline jobject CallStaticObjectMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    return (*env)->CallStaticObjectMethodA(env, clazz, methodID, args);
}

//static inline jboolean CallStaticBooleanMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jboolean result;
//    va_start(args, methodID);
//    result = (*env)->CallStaticBooleanMethodV(env, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jboolean CallStaticBooleanMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
//    return (*env)->CallStaticBooleanMethodV(env, clazz, methodID, args);
//}

static inline jboolean CallStaticBooleanMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    return (*env)->CallStaticBooleanMethodA(env, clazz, methodID, args);
}

//static inline jbyte CallStaticByteMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jbyte result;
//    va_start(args, methodID);
//    result = (*env)->CallStaticByteMethodV(env, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jbyte CallStaticByteMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
//    return (*env)->CallStaticByteMethodV(env, clazz, methodID, args);
//}

static inline jbyte CallStaticByteMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    return (*env)->CallStaticByteMethodA(env, clazz, methodID, args);
}

//static inline jchar CallStaticCharMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jchar result;
//    va_start(args, methodID);
//    result = (*env)->CallStaticCharMethodV(env, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jchar CallStaticCharMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
//    return (*env)->CallStaticCharMethodV(env, clazz, methodID, args);
//}

static inline jchar CallStaticCharMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    return (*env)->CallStaticCharMethodA(env, clazz, methodID, args);
}

//static inline jshort CallStaticShortMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jshort result;
//    va_start(args, methodID);
//    result = (*env)->CallStaticShortMethodV(env, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jshort CallStaticShortMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
//    return (*env)->CallStaticShortMethodV(env, clazz, methodID, args);
//}

static inline jshort CallStaticShortMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    return (*env)->CallStaticShortMethodA(env, clazz, methodID, args);
}

//static inline jint CallStaticIntMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jint result;
//    va_start(args, methodID);
//    result = (*env)->CallStaticIntMethodV(env, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jint CallStaticIntMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
//    return (*env)->CallStaticIntMethodV(env, clazz, methodID, args);
//}

static inline jint CallStaticIntMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    return (*env)->CallStaticIntMethodA(env, clazz, methodID, args);
}

//static inline jlong CallStaticLongMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jlong result;
//    va_start(args, methodID);
//    result = (*env)->CallStaticLongMethodV(env, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jlong CallStaticLongMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
//    return (*env)->CallStaticLongMethodV(env, clazz, methodID, args);
//}

static inline jlong CallStaticLongMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    return (*env)->CallStaticLongMethodA(env, clazz, methodID, args);
}

//static inline jfloat CallStaticFloatMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jfloat result;
//    va_start(args, methodID);
//    result = (*env)->CallStaticFloatMethodV(env, clazz, methodID, args);
//    va_end(args);
//    return result;
//}

//static inline jfloat CallStaticFloatMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
//    return (*env)->CallStaticFloatMethodV(env, clazz, methodID, args);
//}

static inline jfloat CallStaticFloatMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    return (*env)->CallStaticFloatMethodA(env, clazz, methodID, args);
}

//static inline jdouble CallStaticDoubleMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
//    va_list args;
//    jdouble result;
//    va_start(args, methodID);
//    result = (*env)->CallStaticDoubleMethodV(env, clazz, methodID, args);
//    va_end(args);
//    return result;
//}
//
//static inline jdouble CallStaticDoubleMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
//    return (*env)->CallStaticDoubleMethodV(env, clazz, methodID, args);
//}

static inline jdouble CallStaticDoubleMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    return (*env)->CallStaticDoubleMethodA(env, clazz, methodID, args);
}

//static inline void CallStaticVoidMethod(JNIEnv *env, jclass cls, jmethodID methodID, ...) {
//    va_list args;
//    va_start(args, methodID);
//    (*env)->CallStaticVoidMethodV(env, cls, methodID, args);
//    va_end(args);
//}
//
//static inline void CallStaticVoidMethodV(JNIEnv *env, jclass cls, jmethodID methodID, va_list args) {
//    (*env)->CallStaticVoidMethodV(env, cls, methodID, args);
//}

static inline void CallStaticVoidMethodA(JNIEnv *env, jclass cls, jmethodID methodID, const jvalue *args) {
    (*env)->CallStaticVoidMethodA(env, cls, methodID, args);
}

static inline jfieldID GetStaticFieldID(JNIEnv *env, jclass clazz, const char *name, const char *sig) {
    return (*env)->GetStaticFieldID(env, clazz, name, sig);
}

static inline jobject GetStaticObjectField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    return (*env)->GetStaticObjectField(env, clazz, fieldID);
}

static inline jboolean GetStaticBooleanField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    return (*env)->GetStaticBooleanField(env, clazz, fieldID);
}

static inline jbyte GetStaticByteField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    return (*env)->GetStaticByteField(env, clazz, fieldID);
}

static inline jchar GetStaticCharField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    return (*env)->GetStaticCharField(env, clazz, fieldID);
}

static inline jshort GetStaticShortField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    return (*env)->GetStaticShortField(env, clazz, fieldID);
}

static inline jint GetStaticIntField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    return (*env)->GetStaticIntField(env, clazz, fieldID);
}

static inline jlong GetStaticLongField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    return (*env)->GetStaticLongField(env, clazz, fieldID);
}

static inline jfloat GetStaticFloatField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    return (*env)->GetStaticFloatField(env, clazz, fieldID);
}

static inline jdouble GetStaticDoubleField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    return (*env)->GetStaticDoubleField(env, clazz, fieldID);
}

static inline void SetStaticObjectField(JNIEnv *env, jclass clazz, jfieldID fieldID, jobject value) {
    (*env)->SetStaticObjectField(env, clazz, fieldID, value);
}

static inline void SetStaticBooleanField(JNIEnv *env, jclass clazz, jfieldID fieldID, jboolean value) {
    (*env)->SetStaticBooleanField(env, clazz, fieldID, value);
}

static inline void SetStaticByteField(JNIEnv *env, jclass clazz, jfieldID fieldID, jbyte value) {
    (*env)->SetStaticByteField(env, clazz, fieldID, value);
}

static inline void SetStaticCharField(JNIEnv *env, jclass clazz, jfieldID fieldID, jchar value) {
    (*env)->SetStaticCharField(env, clazz, fieldID, value);
}

static inline void SetStaticShortField(JNIEnv *env, jclass clazz, jfieldID fieldID, jshort value) {
    (*env)->SetStaticShortField(env, clazz, fieldID, value);
}

static inline void SetStaticIntField(JNIEnv *env, jclass clazz, jfieldID fieldID, jint value) {
    (*env)->SetStaticIntField(env, clazz, fieldID, value);
}

static inline void SetStaticLongField(JNIEnv *env, jclass clazz, jfieldID fieldID, jlong value) {
    (*env)->SetStaticLongField(env, clazz, fieldID, value);
}

static inline void SetStaticFloatField(JNIEnv *env, jclass clazz, jfieldID fieldID, jfloat value) {
    (*env)->SetStaticFloatField(env, clazz, fieldID, value);
}

static inline void SetStaticDoubleField(JNIEnv *env, jclass clazz, jfieldID fieldID, jdouble value) {
    (*env)->SetStaticDoubleField(env, clazz, fieldID, value);
}

static inline jstring NewString(JNIEnv *env, const jchar *unicode, jsize len) {
    return (*env)->NewString(env, unicode, len);
}

static inline jsize GetStringLength(JNIEnv *env, jstring str) {
    return (*env)->GetStringLength(env, str);
}

static inline const jchar *GetStringChars(JNIEnv *env, jstring str, jboolean *isCopy) {
    return (*env)->GetStringChars(env, str, isCopy);
}

static inline void ReleaseStringChars(JNIEnv *env, jstring str, const jchar *chars) {
    (*env)->ReleaseStringChars(env, str, chars);
}

static inline jstring NewStringUTF(JNIEnv *env, const char *utf) {
    return (*env)->NewStringUTF(env, utf);
}

static inline jsize GetStringUTFLength(JNIEnv *env, jstring str) {
    return (*env)->GetStringUTFLength(env, str);
}

static inline const char *GetStringUTFChars(JNIEnv *env, jstring str, jboolean *isCopy) {
    return (*env)->GetStringUTFChars(env, str, isCopy);
}

static inline void ReleaseStringUTFChars(JNIEnv *env, jstring str, const char *chars) {
    (*env)->ReleaseStringUTFChars(env, str, chars);
}

static inline jsize GetArrayLength(JNIEnv *env, jarray array) {
    return (*env)->GetArrayLength(env, array);
}

static inline jobjectArray NewObjectArray(JNIEnv *env, jsize len, jclass clazz, jobject init) {
    return (*env)->NewObjectArray(env, len, clazz, init);
}

static inline jobject GetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index) {
    return (*env)->GetObjectArrayElement(env, array, index);
}

static inline void SetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index, jobject val) {
    (*env)->SetObjectArrayElement(env, array, index, val);
}

static inline jbooleanArray NewBooleanArray(JNIEnv *env, jsize len) {
    return (*env)->NewBooleanArray(env, len);
}

static inline jbyteArray NewByteArray(JNIEnv *env, jsize len) {
    return (*env)->NewByteArray(env, len);
}

static inline jcharArray NewCharArray(JNIEnv *env, jsize len) {
    return (*env)->NewCharArray(env, len);
}

static inline jshortArray NewShortArray(JNIEnv *env, jsize len) {
    return (*env)->NewShortArray(env, len);
}

static inline jintArray NewIntArray(JNIEnv *env, jsize len) {
    return (*env)->NewIntArray(env, len);
}

static inline jlongArray NewLongArray(JNIEnv *env, jsize len) {
    return (*env)->NewLongArray(env, len);
}

static inline jfloatArray NewFloatArray(JNIEnv *env, jsize len) {
    return (*env)->NewFloatArray(env, len);
}

static inline jdoubleArray NewDoubleArray(JNIEnv *env, jsize len) {
    return (*env)->NewDoubleArray(env, len);
}

static inline jboolean *GetBooleanArrayElements(JNIEnv *env, jbooleanArray array, jboolean *isCopy) {
    return (*env)->GetBooleanArrayElements(env, array, isCopy);
}

static inline jbyte *GetByteArrayElements(JNIEnv *env, jbyteArray array, jboolean *isCopy) {
    return (*env)->GetByteArrayElements(env, array, isCopy);
}

static inline jchar *GetCharArrayElements(JNIEnv *env, jcharArray array, jboolean *isCopy) {
    return (*env)->GetCharArrayElements(env, array, isCopy);
}

static inline jshort *GetShortArrayElements(JNIEnv *env, jshortArray array, jboolean *isCopy) {
    return (*env)->GetShortArrayElements(env, array, isCopy);
}

static inline jint *GetIntArrayElements(JNIEnv *env, jintArray array, jboolean *isCopy) {
    return (*env)->GetIntArrayElements(env, array, isCopy);
}

static inline jlong *GetLongArrayElements(JNIEnv *env, jlongArray array, jboolean *isCopy) {
    return (*env)->GetLongArrayElements(env, array, isCopy);
}

static inline jfloat *GetFloatArrayElements(JNIEnv *env, jfloatArray array, jboolean *isCopy) {
    return (*env)->GetFloatArrayElements(env, array, isCopy);
}

static inline jdouble *GetDoubleArrayElements(JNIEnv *env, jdoubleArray array, jboolean *isCopy) {
    return (*env)->GetDoubleArrayElements(env, array, isCopy);
}

static inline void ReleaseBooleanArrayElements(JNIEnv *env, jbooleanArray array, jboolean *elems, jint mode) {
    (*env)->ReleaseBooleanArrayElements(env, array, elems, mode);
}

static inline void ReleaseByteArrayElements(JNIEnv *env, jbyteArray array, jbyte *elems, jint mode) {
    (*env)->ReleaseByteArrayElements(env, array, elems, mode);
}

static inline void ReleaseCharArrayElements(JNIEnv *env, jcharArray array, jchar *elems, jint mode) {
    (*env)->ReleaseCharArrayElements(env, array, elems, mode);
}

static inline void ReleaseShortArrayElements(JNIEnv *env, jshortArray array, jshort *elems, jint mode) {
    (*env)->ReleaseShortArrayElements(env, array, elems, mode);
}

static inline void ReleaseIntArrayElements(JNIEnv *env, jintArray array, jint *elems, jint mode) {
    (*env)->ReleaseIntArrayElements(env, array, elems, mode);
}

static inline void ReleaseLongArrayElements(JNIEnv *env, jlongArray array, jlong *elems, jint mode) {
    (*env)->ReleaseLongArrayElements(env, array, elems, mode);
}

static inline void ReleaseFloatArrayElements(JNIEnv *env, jfloatArray array, jfloat *elems, jint mode) {
    (*env)->ReleaseFloatArrayElements(env, array, elems, mode);
}

static inline void ReleaseDoubleArrayElements(JNIEnv *env, jdoubleArray array, jdouble *elems, jint mode) {
    (*env)->ReleaseDoubleArrayElements(env, array, elems, mode);
}

static inline void GetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf) {
    (*env)->GetBooleanArrayRegion(env, array, start, l, buf);
}

static inline void GetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf) {
    (*env)->GetByteArrayRegion(env, array, start, len, buf);
}

static inline void GetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf) {
    (*env)->GetCharArrayRegion(env, array, start, len, buf);
}

static inline void GetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf) {
    (*env)->GetShortArrayRegion(env, array, start, len, buf);
}

static inline void GetIntArrayRegion(JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf) {
    (*env)->GetIntArrayRegion(env, array, start, len, buf);
}

static inline void GetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf) {
    (*env)->GetLongArrayRegion(env, array, start, len, buf);
}

static inline void GetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf) {
    (*env)->GetFloatArrayRegion(env, array, start, len, buf);
}

static inline void GetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf) {
    (*env)->GetDoubleArrayRegion(env, array, start, len, buf);
}

static inline void SetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start, jsize l, const jboolean *buf) {
    (*env)->SetBooleanArrayRegion(env, array, start, l, buf);
}

static inline void SetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start, jsize len, const jbyte *buf) {
    (*env)->SetByteArrayRegion(env, array, start, len, buf);
}

static inline void SetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start, jsize len, const jchar *buf) {
    (*env)->SetCharArrayRegion(env, array, start, len, buf);
}

static inline void SetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start, jsize len, const jshort *buf) {
    (*env)->SetShortArrayRegion(env, array, start, len, buf);
}

static inline void SetIntArrayRegion(JNIEnv *env, jintArray array, jsize start, jsize len, const jint *buf) {
    (*env)->SetIntArrayRegion(env, array, start, len, buf);
}

static inline void SetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start, jsize len, const jlong *buf) {
    (*env)->SetLongArrayRegion(env, array, start, len, buf);
}

static inline void SetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start, jsize len, const jfloat *buf) {
    (*env)->SetFloatArrayRegion(env, array, start, len, buf);
}

static inline void SetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start, jsize len, const jdouble *buf) {
    (*env)->SetDoubleArrayRegion(env, array, start, len, buf);
}

static inline jint RegisterNatives(JNIEnv *env, jclass clazz, const JNINativeMethod *methods, jint nMethods) {
   return (*env)->RegisterNatives(env, clazz, methods, nMethods);
}

static inline jint UnregisterNatives(JNIEnv *env, jclass clazz) {
    return (*env)->UnregisterNatives(env, clazz);
}

static inline jint MonitorEnter(JNIEnv *env, jobject obj) {
    return (*env)->MonitorEnter(env, obj);
}

static inline jint MonitorExit(JNIEnv *env, jobject obj) {
    return (*env)->MonitorExit(env, obj);
}

static inline jint GetJavaVM(JNIEnv *env, JavaVM **vm) {
    return (*env)->GetJavaVM(env, vm);
}

static inline void GetStringRegion(JNIEnv *env, jstring str, jsize start, jsize len, jchar *buf) {
    (*env)->GetStringRegion(env, str, start, len, buf);
}

static inline void GetStringUTFRegion(JNIEnv *env, jstring str, jsize start, jsize len, char *buf) {
    (*env)->GetStringUTFRegion(env, str, start, len, buf);
}

static inline void *GetPrimitiveArrayCritical(JNIEnv *env, jarray array, jboolean *isCopy) {
    (*env)->GetPrimitiveArrayCritical(env, array, isCopy);
}

static inline void ReleasePrimitiveArrayCritical(JNIEnv *env, jarray array, void *carray, jint mode) {
    (*env)->ReleasePrimitiveArrayCritical(env, array, carray, mode);
}

static inline const jchar *GetStringCritical(JNIEnv *env, jstring string, jboolean *isCopy) {
    return (*env)->GetStringCritical(env, string, isCopy);
}

static inline void ReleaseStringCritical(JNIEnv *env, jstring string, const jchar *cstring) {
    (*env)->ReleaseStringCritical(env, string, cstring);
}

static inline jweak NewWeakGlobalRef(JNIEnv *env, jobject obj) {
    return (*env)->NewWeakGlobalRef(env, obj);
}

static inline void DeleteWeakGlobalRef(JNIEnv *env, jweak ref) {
    (*env)->DeleteWeakGlobalRef(env, ref);
}

static inline jboolean ExceptionCheck(JNIEnv *env) {
    return (*env)->ExceptionCheck(env);
}

static inline jobject NewDirectByteBuffer(JNIEnv *env, void *address, jlong capacity) {
    return (*env)->NewDirectByteBuffer(env, address, capacity);
}

static inline void *GetDirectBufferAddress(JNIEnv *env, jobject buf) {
    return (*env)->GetDirectBufferAddress(env, buf);
}

static inline jlong GetDirectBufferCapacity(JNIEnv *env, jobject buf) {
    return (*env)->GetDirectBufferCapacity(env, buf);
}

// New JNI 1.6 Features
static inline jobjectRefType GetObjectRefType(JNIEnv *env, jobject obj) {
	return (*env)->GetObjectRefType(env, obj);
}
*/
import "C"
import (
	"unsafe"
)

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
type JSize JInt

// Reference types, in Go.
type JObject uintptr
type JClass JObject
type JString JObject
type JArray JObject
type JObjectArray JArray
type JBooleanArray JArray
type JByteArray JArray
type JCharArray JArray
type JShortArray JArray
type JIntArray JArray
type JLongArray JArray
type JFloatArray JArray
type JDoubleArray JArray
type JThrowable JObject
type JWeak JObject

type JFieldID uintptr  // field IDs
type JMethodID uintptr // method IDs

type JNIEnv uintptr
type Javavm uintptr

type JValue struct {
	z JBoolean
	b JByte
	c JChar
	s JShort
	i JInt
	j JLong
	f JFloat
	d JDouble
	l JObject
}

const (
	JNIInvalidRefType    = 0
	JNILocalRefType      = 1
	JNIGlobalRefType     = 2
	JNIWeakGlobalRefType = 3
)

type JObjectRefType int

type JNINativeMethod struct {
	Name      string
	Signature string
	FnPtr     interface{}
}

type JavaVMOption struct {
	OptionString string
	ExtraInfo    uintptr
}

type JavaVMInitArgs struct {
	Version            JInt
	NOptions           JInt
	Options            []JavaVMOption
	IgnoreUnrecognized JBoolean
}

type JavaVMAttachArgs struct {
	Version JInt
	Name    string
	Group   JObject
}

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

type Method interface {
	JniOnLoad(vm Javavm, reserved uintptr) JInt
	JniOnUnload(vm Javavm, reserved uintptr)
}

func jvmPtr(vm Javavm) *C.JavaVM {
	return (*C.JavaVM)(unsafe.Pointer(vm))
}

func envPtr(env JNIEnv) *C.JNIEnv {
	return (*C.JNIEnv)(unsafe.Pointer(env))
}

func cInt(value JInt) C.jint {
	return C.jint(value)
}

func cSize(value JSize) C.jsize {
	return C.jsize(value)
}

func cBoolean(value JBoolean) C.jboolean {
	return C.jboolean(value)
}

func cByte(value JByte) C.jbyte {
	return C.jbyte(value)
}

func cLong(value JLong) C.jlong {
	return C.jlong(value)
}

func cChar(value JChar) C.jchar {
	return C.jchar(value)
}

func cShort(value JShort) C.jshort {
	return C.jshort(value)
}

func cFloat(value JFloat) C.jfloat {
	return C.jfloat(value)
}

func cDouble(value JDouble) C.jdouble {
	return C.jdouble(value)
}

func cObject(value JObject) C.jobject {
	return C.jobject(value)
}

func cClass(value JClass) C.jclass {
	return C.jclass(value)
}

func cMethodID(value JMethodID) C.jmethodID {
	return C.jmethodID(unsafe.Pointer(value))
}

func cFieldID(value JFieldID) C.jfieldID {
	return C.jfieldID(unsafe.Pointer(value))
}

func cThrowable(value JThrowable) C.jthrowable {
	return C.jthrowable(value)
}

func cValues(value []JValue) *C.jvalue {
	if len(value) == 0 {
		return nil
	}
	return (*C.jvalue)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&value))))
}

func free(ptr unsafe.Pointer) {
	C.free(ptr)
}

// ------------------------------- JavaVM --------------------------------

func (vm Javavm) DestroyJavaVM() JInt {
	return JInt(C.DestroyJavaVM(jvmPtr(vm)))
}

func (vm Javavm) AttachCurrentThread(args unsafe.Pointer) (JNIEnv, JInt) {
	var env unsafe.Pointer
	result := JInt(C.AttachCurrentThread(jvmPtr(vm), &env, args))
	return JNIEnv(env), result
}

func (vm Javavm) DetachCurrentThread() JInt {
	return JInt(C.DetachCurrentThread(jvmPtr(vm)))
}

func (vm Javavm) GetEnv(version JInt) (JNIEnv, JInt) {
	var env unsafe.Pointer
	result := JInt(C.GetEnv(jvmPtr(vm), &env, cInt(version)))
	return JNIEnv(env), result
}

func (vm Javavm) AttachCurrentThreadAsDaemon(args unsafe.Pointer) (JNIEnv, JInt) {
	var env unsafe.Pointer
	result := JInt(C.AttachCurrentThreadAsDaemon(jvmPtr(vm), &env, args))
	return JNIEnv(env), result
}

// ------------------------------- JNIEnv --------------------------------

func (env JNIEnv) GetVersion() JInt {
	return JInt(C.GetVersion(envPtr(env)))
}

func (env JNIEnv) DefineClass(name string, loader JObject, buf []JByte, len JSize) JClass {
	cName := C.CString(name)
	defer free(unsafe.Pointer(cName))
	return JClass(C.DefineClass(envPtr(env), cName, cObject(loader), (*C.jbyte)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&buf)))), cSize(len)))
}

func (env JNIEnv) FindClass(name string) JClass {
	cName := C.CString(name)
	defer free(unsafe.Pointer(cName))
	return JClass(C.FindClass(envPtr(env), cName))
}

func (env JNIEnv) FromReflectedMethod(method JObject) JMethodID {
	return JMethodID(unsafe.Pointer(C.FromReflectedMethod(envPtr(env), cObject(method))))
}

func (env JNIEnv) FromReflectedField(field JObject) JFieldID {
	return JFieldID(unsafe.Pointer(C.FromReflectedField(envPtr(env), cObject(field))))
}

func (env JNIEnv) ToReflectedMethod(cls JClass, methodID JMethodID, isStatic JBoolean) JObject {
	return JObject(C.ToReflectedMethod(envPtr(env), cClass(cls), cMethodID(methodID), cBoolean(isStatic)))
}

func (env JNIEnv) GetSuperclass(sub JClass) JClass {
	return JClass(C.GetSuperclass(envPtr(env), cClass(sub)))
}

func (env JNIEnv) IsAssignableFrom(sub JClass, sup JClass) JBoolean {
	return JBoolean(C.IsAssignableFrom(envPtr(env), cClass(sub), cClass(sup)))
}

func (env JNIEnv) ToReflectedField(cls JClass, fieldID JFieldID, isStatic JBoolean) JObject {
	return JObject(C.ToReflectedField(envPtr(env), cClass(cls), cFieldID(fieldID), cBoolean(isStatic)))
}

func (env JNIEnv) Throw(obj JThrowable) JInt {
	return JInt(C.Throw(envPtr(env), cThrowable(obj)))
}

func (env JNIEnv) ThrowNew(clazz JClass, msg string) JInt {
	cMsg := C.CString(msg)
	defer free(unsafe.Pointer(cMsg))
	return JInt(C.ThrowNew(envPtr(env), cClass(clazz), cMsg))
}

func (env JNIEnv) ExceptionOccurred() JThrowable {
	return JThrowable(C.ExceptionOccurred(envPtr(env)))
}

func (env JNIEnv) ExceptionDescribe() {
	C.ExceptionDescribe(envPtr(env))
}

func (env JNIEnv) ExceptionClear() {
	C.ExceptionClear(envPtr(env))
}

func (env JNIEnv) FatalError(msg string) {
	cMsg := C.CString(msg)
	defer free(unsafe.Pointer(cMsg))
	C.FatalError(envPtr(env), cMsg)
}

func (env JNIEnv) PushLocalFrame(capacity JInt) JInt {
	return JInt(C.PushLocalFrame(envPtr(env), cInt(capacity)))
}

func (env JNIEnv) PopLocalFrame(result JObject) JObject {
	return JObject(C.PopLocalFrame(envPtr(env), cObject(result)))
}

func (env JNIEnv) NewGlobalRef(lobj JObject) JObject {
	return JObject(C.NewGlobalRef(envPtr(env), cObject(lobj)))
}

func (env JNIEnv) DeleteGlobalRef(gref JObject) {
	C.DeleteGlobalRef(envPtr(env), cObject(gref))
}

func (env JNIEnv) DeleteLocalRef(obj JObject) {
	C.DeleteLocalRef(envPtr(env), cObject(obj))
}

func (env JNIEnv) IsSameObject(obj1 JObject, obj2 JObject) JBoolean {
	return JBoolean(C.IsSameObject(envPtr(env), cObject(obj1), cObject(obj2)))
}

func (env JNIEnv) NewLocalRef(ref JObject) JObject {
	return JObject(C.NewLocalRef(envPtr(env), cObject(ref)))
}

func (env JNIEnv) EnsureLocalCapacity(capacity JInt) JInt {
	return JInt(C.EnsureLocalCapacity(envPtr(env), cInt(capacity)))
}

func (env JNIEnv) AllocObject(clazz JClass) JObject {
	return JObject(C.AllocObject(envPtr(env), cClass(clazz)))
}

func (env JNIEnv) NewObjectA(clazz JClass, methodID JMethodID, args ...JValue) JObject {
	return JObject(C.NewObjectA(envPtr(env), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) GetObjectClass(obj JObject) JClass {
	return JClass(C.GetObjectClass(envPtr(env), cObject(obj)))
}

func (env JNIEnv) IsInstanceOf(obj JObject, clazz JClass) JBoolean {
	return JBoolean(C.IsInstanceOf(envPtr(env), cObject(obj), cClass(clazz)))
}

func (env JNIEnv) GetMethodID(clazz JClass, name string, sig string) JMethodID {
	cName := C.CString(name)
	cSig := C.CString(sig)
	defer free(unsafe.Pointer(cName))
	defer free(unsafe.Pointer(cSig))
	return JMethodID(unsafe.Pointer(C.GetMethodID(envPtr(env), cClass(clazz), cName, cSig)))
}

func (env JNIEnv) CallObjectMethodA(obj JObject, methodID JMethodID, args ...JValue) JObject {
	return JObject(C.CallObjectMethodA(envPtr(env), cObject(obj), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallBooleanMethodA(obj JObject, methodID JMethodID, args ...JValue) JBoolean {
	return JBoolean(C.CallBooleanMethodA(envPtr(env), cObject(obj), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallByteMethodA(obj JObject, methodID JMethodID, args ...JValue) JByte {
	return JByte(C.CallByteMethodA(envPtr(env), cObject(obj), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallCharMethodA(obj JObject, methodID JMethodID, args ...JValue) JChar {
	return JChar(C.CallCharMethodA(envPtr(env), cObject(obj), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallShortMethodA(obj JObject, methodID JMethodID, args ...JValue) JShort {
	return JShort(C.CallShortMethodA(envPtr(env), cObject(obj), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallIntMethodA(obj JObject, methodID JMethodID, args ...JValue) JInt {
	return JInt(C.CallIntMethodA(envPtr(env), cObject(obj), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallLongMethodA(obj JObject, methodID JMethodID, args ...JValue) JLong {
	return JLong(C.CallLongMethodA(envPtr(env), cObject(obj), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallFloatMethodA(obj JObject, methodID JMethodID, args ...JValue) JFloat {
	return JFloat(C.CallFloatMethodA(envPtr(env), cObject(obj), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallDoubleMethodA(obj JObject, methodID JMethodID, args ...JValue) JDouble {
	return JDouble(C.CallDoubleMethodA(envPtr(env), cObject(obj), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallVoidMethodA(obj JObject, methodID JMethodID, args ...JValue) {
	C.CallVoidMethodA(envPtr(env), cObject(obj), cMethodID(methodID), cValues(args))
}

func (env JNIEnv) CallNonvirtualObjectMethodA(obj JObject, clazz JClass, methodID JMethodID, args ...JValue) JObject {
	return JObject(C.CallNonvirtualObjectMethodA(envPtr(env), cObject(obj), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallNonvirtualBooleanMethodA(obj JObject, clazz JClass, methodID JMethodID, args ...JValue) JBoolean {
	return JBoolean(C.CallNonvirtualBooleanMethodA(envPtr(env), cObject(obj), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallNonvirtualByteMethodA(obj JObject, clazz JClass, methodID JMethodID, args ...JValue) JByte {
	return JByte(C.CallNonvirtualByteMethodA(envPtr(env), cObject(obj), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallNonvirtualCharMethodA(obj JObject, clazz JClass, methodID JMethodID, args ...JValue) JChar {
	return JChar(C.CallNonvirtualCharMethodA(envPtr(env), cObject(obj), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallNonvirtualShortMethodA(obj JObject, clazz JClass, methodID JMethodID, args ...JValue) JShort {
	return JShort(C.CallNonvirtualShortMethodA(envPtr(env), cObject(obj), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallNonvirtualIntMethodA(obj JObject, clazz JClass, methodID JMethodID, args ...JValue) JInt {
	return JInt(C.CallNonvirtualIntMethodA(envPtr(env), cObject(obj), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallNonvirtualLongMethodA(obj JObject, clazz JClass, methodID JMethodID, args ...JValue) JLong {
	return JLong(C.CallNonvirtualLongMethodA(envPtr(env), cObject(obj), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallNonvirtualFloatMethodA(obj JObject, clazz JClass, methodID JMethodID, args ...JValue) JFloat {
	return JFloat(C.CallNonvirtualFloatMethodA(envPtr(env), cObject(obj), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallNonvirtualDoubleMethodA(obj JObject, clazz JClass, methodID JMethodID, args ...JValue) JDouble {
	return JDouble(C.CallNonvirtualDoubleMethodA(envPtr(env), cObject(obj), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallNonvirtualVoidMethodA(obj JObject, clazz JClass, methodID JMethodID, args ...JValue) {
	C.CallNonvirtualVoidMethodA(envPtr(env), cObject(obj), cClass(clazz), cMethodID(methodID), cValues(args))
}

func (env JNIEnv) GetFieldID(clazz JClass, name string, sig string) JFieldID {
	cName := C.CString(name)
	cSig := C.CString(sig)
	defer free(unsafe.Pointer(cName))
	defer free(unsafe.Pointer(cSig))
	return JFieldID(unsafe.Pointer(C.GetFieldID(envPtr(env), cClass(clazz), cName, cSig)))
}

func (env JNIEnv) GetObjectField(obj JObject, fieldID JFieldID) JObject {
	return JObject(C.GetObjectField(envPtr(env), cObject(obj), cFieldID(fieldID)))
}

func (env JNIEnv) GetBooleanField(obj JObject, fieldID JFieldID) JBoolean {
	return JBoolean(C.GetBooleanField(envPtr(env), cObject(obj), cFieldID(fieldID)))
}

func (env JNIEnv) GetByteField(obj JObject, fieldID JFieldID) JByte {
	return JByte(C.GetByteField(envPtr(env), cObject(obj), cFieldID(fieldID)))
}

func (env JNIEnv) GetCharField(obj JObject, fieldID JFieldID) JChar {
	return JChar(C.GetCharField(envPtr(env), cObject(obj), cFieldID(fieldID)))
}

func (env JNIEnv) GetShortField(obj JObject, fieldID JFieldID) JShort {
	return JShort(C.GetShortField(envPtr(env), cObject(obj), cFieldID(fieldID)))
}

func (env JNIEnv) GetIntField(obj JObject, fieldID JFieldID) JInt {
	return JInt(C.GetIntField(envPtr(env), cObject(obj), cFieldID(fieldID)))
}

func (env JNIEnv) GetLongField(obj JObject, fieldID JFieldID) JLong {
	return JLong(C.GetLongField(envPtr(env), cObject(obj), cFieldID(fieldID)))
}

func (env JNIEnv) GetFloatField(obj JObject, fieldID JFieldID) JFloat {
	return JFloat(C.GetFloatField(envPtr(env), cObject(obj), cFieldID(fieldID)))
}

func (env JNIEnv) GetDoubleField(obj JObject, fieldID JFieldID) JDouble {
	return JDouble(C.GetDoubleField(envPtr(env), cObject(obj), cFieldID(fieldID)))
}

func (env JNIEnv) SetObjectField(obj JObject, fieldID JFieldID, val JObject) {
	C.SetObjectField(envPtr(env), cObject(obj), cFieldID(fieldID), cObject(val))
}

func (env JNIEnv) SetBooleanField(obj JObject, fieldID JFieldID, val JBoolean) {
	C.SetBooleanField(envPtr(env), cObject(obj), cFieldID(fieldID), cBoolean(val))
}

func (env JNIEnv) SetByteField(obj JObject, fieldID JFieldID, val JByte) {
	C.SetByteField(envPtr(env), cObject(obj), cFieldID(fieldID), cByte(val))
}

func (env JNIEnv) SetCharField(obj JObject, fieldID JFieldID, val JChar) {
	C.SetCharField(envPtr(env), cObject(obj), cFieldID(fieldID), cChar(val))
}

func (env JNIEnv) SetShortField(obj JObject, fieldID JFieldID, val JShort) {
	C.SetShortField(envPtr(env), cObject(obj), cFieldID(fieldID), cShort(val))
}

func (env JNIEnv) SetIntField(obj JObject, fieldID JFieldID, val JInt) {
	C.SetIntField(envPtr(env), cObject(obj), cFieldID(fieldID), cInt(val))
}

func (env JNIEnv) SetLongField(obj JObject, fieldID JFieldID, val JLong) {
	C.SetLongField(envPtr(env), cObject(obj), cFieldID(fieldID), cLong(val))
}

func (env JNIEnv) SetFloatField(obj JObject, fieldID JFieldID, val JFloat) {
	C.SetFloatField(envPtr(env), cObject(obj), cFieldID(fieldID), cFloat(val))
}

func (env JNIEnv) SetDoubleField(obj JObject, fieldID JFieldID, val JDouble) {
	C.SetDoubleField(envPtr(env), cObject(obj), cFieldID(fieldID), cDouble(val))
}

func (env JNIEnv) GetStaticMethodID(clazz JClass, name string, sig string) JMethodID {
	cName := C.CString(name)
	cSig := C.CString(sig)
	defer free(unsafe.Pointer(cName))
	defer free(unsafe.Pointer(cSig))
	return JMethodID(unsafe.Pointer(C.GetStaticMethodID(envPtr(env), cClass(clazz), cName, cSig)))
}

func (env JNIEnv) CallStaticObjectMethodA(clazz JClass, methodID JMethodID, args ...JValue) JObject {
	return JObject(C.CallStaticObjectMethodA(envPtr(env), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallStaticBooleanMethodA(clazz JClass, methodID JMethodID, args ...JValue) JBoolean {
	return JBoolean(C.CallStaticBooleanMethodA(envPtr(env), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallStaticByteMethodA(clazz JClass, methodID JMethodID, args ...JValue) JByte {
	return JByte(C.CallStaticByteMethodA(envPtr(env), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallStaticCharMethodA(clazz JClass, methodID JMethodID, args ...JValue) JChar {
	return JChar(C.CallStaticCharMethodA(envPtr(env), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallStaticShortMethodA(clazz JClass, methodID JMethodID, args ...JValue) JShort {
	return JShort(C.CallStaticShortMethodA(envPtr(env), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallStaticIntMethodA(clazz JClass, methodID JMethodID, args ...JValue) JInt {
	return JInt(C.CallStaticIntMethodA(envPtr(env), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallStaticLongMethodA(clazz JClass, methodID JMethodID, args ...JValue) JLong {
	return JLong(C.CallStaticLongMethodA(envPtr(env), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallStaticFloatMethodA(clazz JClass, methodID JMethodID, args ...JValue) JFloat {
	return JFloat(C.CallStaticFloatMethodA(envPtr(env), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallStaticDoubleMethodA(clazz JClass, methodID JMethodID, args ...JValue) JDouble {
	return JDouble(C.CallStaticDoubleMethodA(envPtr(env), cClass(clazz), cMethodID(methodID), cValues(args)))
}

func (env JNIEnv) CallStaticVoidMethodA(cls JClass, methodID JMethodID, args ...JValue) {
	C.CallStaticVoidMethodA(envPtr(env), cClass(cls), cMethodID(methodID), cValues(args))
}

func (env JNIEnv) GetStaticFieldID(clazz JClass, name string, sig string) JFieldID {
	cName := C.CString(name)
	cSig := C.CString(sig)
	defer free(unsafe.Pointer(cName))
	defer free(unsafe.Pointer(cSig))
	return JFieldID(unsafe.Pointer(C.GetStaticFieldID(envPtr(env), cClass(clazz), cName, cSig)))
}

func (env JNIEnv) GetStaticObjectField(clazz JClass, fieldID JFieldID) JObject {
	return JObject(C.GetStaticObjectField(envPtr(env), cClass(clazz), cFieldID(fieldID)))
}

func (env JNIEnv) GetStaticBooleanField(clazz JClass, fieldID JFieldID) JBoolean {
	return JBoolean(C.GetStaticBooleanField(envPtr(env), cClass(clazz), cFieldID(fieldID)))
}

func (env JNIEnv) GetStaticByteField(clazz JClass, fieldID JFieldID) JByte {
	return JByte(C.GetStaticByteField(envPtr(env), cClass(clazz), cFieldID(fieldID)))
}

func (env JNIEnv) GetStaticCharField(clazz JClass, fieldID JFieldID) JChar {
	return JChar(C.GetStaticCharField(envPtr(env), cClass(clazz), cFieldID(fieldID)))
}

func (env JNIEnv) GetStaticShortField(clazz JClass, fieldID JFieldID) JShort {
	return JShort(C.GetStaticShortField(envPtr(env), cClass(clazz), cFieldID(fieldID)))
}

func (env JNIEnv) GetStaticIntField(clazz JClass, fieldID JFieldID) JInt {
	return JInt(C.GetStaticIntField(envPtr(env), cClass(clazz), cFieldID(fieldID)))
}

func (env JNIEnv) GetStaticLongField(clazz JClass, fieldID JFieldID) JLong {
	return JLong(C.GetStaticLongField(envPtr(env), cClass(clazz), cFieldID(fieldID)))
}

func (env JNIEnv) GetStaticFloatField(clazz JClass, fieldID JFieldID) JFloat {
	return JFloat(C.GetStaticFloatField(envPtr(env), cClass(clazz), cFieldID(fieldID)))
}

func (env JNIEnv) GetStaticDoubleField(clazz JClass, fieldID JFieldID) JDouble {
	return JDouble(C.GetStaticDoubleField(envPtr(env), cClass(clazz), cFieldID(fieldID)))
}

func (env JNIEnv) SetStaticObjectField(clazz JClass, fieldID JFieldID, value JObject) {
	C.SetStaticObjectField(envPtr(env), cClass(clazz), cFieldID(fieldID), cObject(value))
}

func (env JNIEnv) SetStaticBooleanField(clazz JClass, fieldID JFieldID, value JBoolean) {
	C.SetStaticBooleanField(envPtr(env), cClass(clazz), cFieldID(fieldID), cBoolean(value))
}

func (env JNIEnv) SetStaticByteField(clazz JClass, fieldID JFieldID, value JByte) {
	C.SetStaticByteField(envPtr(env), cClass(clazz), cFieldID(fieldID), cByte(value))
}

func (env JNIEnv) SetStaticCharField(clazz JClass, fieldID JFieldID, value JChar) {
	C.SetStaticCharField(envPtr(env), cClass(clazz), cFieldID(fieldID), cChar(value))
}

func (env JNIEnv) SetStaticShortField(clazz JClass, fieldID JFieldID, value JShort) {
	C.SetStaticShortField(envPtr(env), cClass(clazz), cFieldID(fieldID), cShort(value))
}

func (env JNIEnv) SetStaticIntField(clazz JClass, fieldID JFieldID, value JInt) {
	C.SetStaticIntField(envPtr(env), cClass(clazz), cFieldID(fieldID), cInt(value))
}

func (env JNIEnv) SetStaticLongField(clazz JClass, fieldID JFieldID, value JLong) {
	C.SetStaticLongField(envPtr(env), cClass(clazz), cFieldID(fieldID), cLong(value))
}

func (env JNIEnv) SetStaticFloatField(clazz JClass, fieldID JFieldID, value JFloat) {
	C.SetStaticFloatField(envPtr(env), cClass(clazz), cFieldID(fieldID), cFloat(value))
}

func (env JNIEnv) SetStaticDoubleField(clazz JClass, fieldID JFieldID, value JDouble) {
	C.SetStaticDoubleField(envPtr(env), cClass(clazz), cFieldID(fieldID), cDouble(value))
}

func (env JNIEnv) NewString(unicode string, len JSize) JString {
	return JString(C.NewString(envPtr(env), (*C.jchar)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&unicode)))), cSize(len)))
}

func (env JNIEnv) GetStringLength(str JString) JSize {
	return JSize(C.GetStringLength(envPtr(env), C.jstring(str)))
}

func (env JNIEnv) GetStringChars(str JString, isCopy JBoolean) string {
	chars := C.GetStringChars(envPtr(env), C.jstring(str), (*C.jboolean)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&isCopy)))))
	return string(*chars)
}

func (env JNIEnv) ReleaseStringChars(str JString, chars string) {
	C.ReleaseStringChars(envPtr(env), C.jstring(str), (*C.jchar)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&chars)))))
}

func (env JNIEnv) NewStringUTF(utf string) JString {
	cUtf := C.CString(utf)
	defer free(unsafe.Pointer(cUtf))
	return JString(C.NewStringUTF(envPtr(env), cUtf))
}

func (env JNIEnv) GetStringUTFLength(str JString) JSize {
	return JSize(C.GetStringUTFLength(envPtr(env), C.jstring(str)))
}

func (env JNIEnv) GetStringUTFChars(str JString, isCopy JBoolean) string {
	return string(*(C.GetStringUTFChars(envPtr(env), C.jstring(str), (*C.jboolean)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&isCopy)))))))
}

func (env JNIEnv) ReleaseStringUTFChars(str JString, chars string) {
	C.ReleaseStringUTFChars(envPtr(env), C.jstring(str), (*C.char)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&chars)))))
}

func (env JNIEnv) GetArrayLength(array JArray) JSize {
	return JSize(C.GetArrayLength(envPtr(env), C.jarray(array)))
}

func (env JNIEnv) NewObjectArray(len JSize, clazz JClass, init JObject) JObjectArray {
	return JObjectArray(C.NewObjectArray(envPtr(env), cSize(len), cClass(clazz), cObject(init)))
}

func (env JNIEnv) GetObjectArrayElement(array JObjectArray, index JSize) JObject {
	return JObject(C.GetObjectArrayElement(envPtr(env), C.jobjectArray(array), cSize(index)))
}

func (env JNIEnv) SetObjectArrayElement(array JObjectArray, index JSize, val JObject) {
	C.SetObjectArrayElement(envPtr(env), C.jobjectArray(array), cSize(index), cObject(val))
}

func (env JNIEnv) NewBooleanArray(len JSize) JBooleanArray {
	return JBooleanArray(C.NewBooleanArray(envPtr(env), cSize(len)))
}

func (env JNIEnv) NewByteArray(len JSize) JByteArray {
	return JByteArray(C.NewByteArray(envPtr(env), cSize(len)))
}

func (env JNIEnv) NewCharArray(len JSize) JCharArray {
	return JCharArray(C.NewCharArray(envPtr(env), cSize(len)))
}

func (env JNIEnv) NewShortArray(len JSize) JShortArray {
	return JShortArray(C.NewShortArray(envPtr(env), cSize(len)))
}

func (env JNIEnv) NewIntArray(len JSize) JIntArray {
	return JIntArray(C.NewIntArray(envPtr(env), cSize(len)))
}

func (env JNIEnv) NewLongArray(len JSize) JLongArray {
	return JLongArray(C.NewLongArray(envPtr(env), cSize(len)))
}

func (env JNIEnv) NewFloatArray(len JSize) JFloatArray {
	return JFloatArray(C.NewFloatArray(envPtr(env), cSize(len)))
}

func (env JNIEnv) NewDoubleArray(len JSize) JDoubleArray {
	return JDoubleArray(C.NewDoubleArray(envPtr(env), cSize(len)))
}

//func (env JNIEnv) GetBooleanArrayElements(array JBooleanArray, isCopy JBoolean) []JBoolean {
//	return []JBoolean(C.GetBooleanArrayElements(envPtr(env), C.jbooleanArray(array), (*C.jboolean)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&isCopy))))))
//}

//func (env JNIEnv) GetByteArrayElements(array JByteArray, isCopy JBoolean) []JByte {
//	return []JByte(C.GetByteArrayElements(envPtr(env), C.jbyteArray(array), (*C.jboolean)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&isCopy))))))
//}

//func (env JNIEnv) GetCharArrayElements(array JCharArray, isCopy JBoolean) []JChar {
//	return []JChar(C.GetCharArrayElements(envPtr(env), C.jcharArray(array), (*C.jboolean)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&isCopy))))))
//}

//func (env JNIEnv) GetShortArrayElements(array JShortArray, isCopy JBoolean) []JShort {
//	return []JShort(C.GetShortArrayElements(envPtr(env), C.jshortArray(array), (*C.jboolean)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&isCopy))))))
//}

//func (env JNIEnv) GetIntArrayElements(array JIntArray, isCopy JBoolean) []JInt {
//	return []JInt(C.GetIntArrayElements(envPtr(env), C.jintArray(array), (*C.jboolean)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&isCopy))))))
//}

//func (env JNIEnv) GetLongArrayElements(array JLongArray, isCopy JBoolean) []JLong {
//	return []JLong(C.GetLongArrayElements(envPtr(env), C.jlongArray(array), (*C.jboolean)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&isCopy))))))
//}

//func (env JNIEnv) GetFloatArrayElements(array JFloatArray, isCopy JBoolean) []JFloat {
//	return []JFloat(C.GetFloatArrayElements(envPtr(env), C.jfloatArray(array), (*C.jboolean)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&isCopy))))))
//}

//func (env JNIEnv) GetDoubleArrayElements(array JDoubleArray, isCopy JBoolean) []JDouble {
//	return []JDouble(C.GetDoubleArrayElements(envPtr(env), C.jdoubleArray(array), (*C.jboolean)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&isCopy))))))
//}

//func (env JNIEnv) ReleaseBooleanArrayElements(array JBooleanArray, elems []JBoolean, mode JInt) {
//	C.ReleaseBooleanArrayElements(envPtr(env), C.jbooleanArray(array), elems, cInt(mode))
//}

//func (env JNIEnv) void ReleaseByteArrayElements(jbyteArray array, jbyte *elems, jint mode) {
//C.ReleaseByteArrayElements(envPtr(env), array, elems, mode);
//}
//
//func (env JNIEnv) void ReleaseCharArrayElements(jcharArray array, jchar *elems, jint mode) {
//C.ReleaseCharArrayElements(envPtr(env), array, elems, mode);
//}
//
//func (env JNIEnv) void ReleaseShortArrayElements(jshortArray array, jshort *elems, jint mode) {
//C.ReleaseShortArrayElements(envPtr(env), array, elems, mode);
//}
//
//func (env JNIEnv) void ReleaseIntArrayElements(jintArray array, jint *elems, jint mode) {
//C.ReleaseIntArrayElements(envPtr(env), array, elems, mode);
//}
//
//func (env JNIEnv) void ReleaseLongArrayElements(jlongArray array, jlong *elems, jint mode) {
//C.ReleaseLongArrayElements(envPtr(env), array, elems, mode);
//}
//
//func (env JNIEnv) void ReleaseFloatArrayElements(jfloatArray array, jfloat *elems, jint mode) {
//C.ReleaseFloatArrayElements(envPtr(env), array, elems, mode);
//}
//
//func (env JNIEnv) void ReleaseDoubleArrayElements(jdoubleArray array, jdouble *elems, jint mode) {
//C.ReleaseDoubleArrayElements(envPtr(env), array, elems, mode);
//}

func (env JNIEnv) GetBooleanArrayRegion(array JBooleanArray, start JSize, l JSize, buf []JBoolean) {
	C.GetBooleanArrayRegion(envPtr(env), C.jbooleanArray(array), cSize(start), cSize(l), (*C.jboolean)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&buf)))))
}

//func (env JNIEnv) void GetByteArrayRegion(jbyteArray array, jsize start, len JSize, jbyte *buf) {
//C.GetByteArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void GetCharArrayRegion(jcharArray array, jsize start, len JSize, jchar *buf) {
//C.GetCharArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void GetShortArrayRegion(jshortArray array, jsize start, len JSize, jshort *buf) {
//C.GetShortArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void GetIntArrayRegion(jintArray array, jsize start, len JSize, jint *buf) {
//C.GetIntArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void GetLongArrayRegion(jlongArray array, jsize start, len JSize, jlong *buf) {
//C.GetLongArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void GetFloatArrayRegion(jfloatArray array, jsize start, len JSize, jfloat *buf) {
//C.GetFloatArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void GetDoubleArrayRegion(jdoubleArray array, jsize start, len JSize, jdouble *buf) {
//C.GetDoubleArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void SetBooleanArrayRegion(jbooleanArray array, jsize start, jsize l, const jboolean *buf) {
//C.SetBooleanArrayRegion(envPtr(env), array, start, l, buf);
//}
//
//func (env JNIEnv) void SetByteArrayRegion(jbyteArray array, jsize start, len JSize, const jbyte *buf) {
//C.SetByteArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void SetCharArrayRegion(jcharArray array, jsize start, len JSize, const jchar *buf) {
//C.SetCharArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void SetShortArrayRegion(jshortArray array, jsize start, len JSize, const jshort *buf) {
//C.SetShortArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void SetIntArrayRegion(jintArray array, jsize start, len JSize, const jint *buf) {
//C.SetIntArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void SetLongArrayRegion(jlongArray array, jsize start, len JSize, const jlong *buf) {
//C.SetLongArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void SetFloatArrayRegion(jfloatArray array, jsize start, len JSize, const jfloat *buf) {
//C.SetFloatArrayRegion(envPtr(env), array, start, len, buf);
//}
//
//func (env JNIEnv) void SetDoubleArrayRegion(jdoubleArray array, jsize start, len JSize, const jdouble *buf) {
//C.SetDoubleArrayRegion(envPtr(env), array, start, len, buf);
//}

func (env JNIEnv) RegisterNatives(clazz JClass, methods []JNINativeMethod, nMethods JInt) JInt {
	return JInt(C.RegisterNatives(envPtr(env), cClass(clazz), (*C.JNINativeMethod)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&methods)))), cInt(nMethods)))
}

func (env JNIEnv) UnregisterNatives(clazz JClass) JInt {
	return JInt(C.UnregisterNatives(envPtr(env), cClass(clazz)))
}

//func (env JNIEnv) jint MonitorEnter(obj JObject) {
//return C.MonitorEnter(envPtr(env), cObject(obj));
//}
//
//func (env JNIEnv) jint MonitorExit(obj JObject) {
//return C.MonitorExit(envPtr(env), cObject(obj));
//}
//
//func (env JNIEnv) jint GetJavaVM(JavaVM **vm) {
//return C.GetJavaVM(envPtr(env), vm);
//}
//
//func (env JNIEnv) void GetStringRegion(jstring str, jsize start, len JSize, jchar *buf) {
//C.GetStringRegion(envPtr(env), str, start, len, buf);
//}
//
//func (env JNIEnv) void GetStringUTFRegion(jstring str, jsize start, len JSize, char *buf) {
//C.GetStringUTFRegion(envPtr(env), str, start, len, buf);
//}
//
//func (env JNIEnv) void *GetPrimitiveArrayCritical(jarray array, jboolean *isCopy) {
//C.GetPrimitiveArrayCritical(envPtr(env), array, isCopy);
//}
//
//func (env JNIEnv) void ReleasePrimitiveArrayCritical(jarray array, void *carray, jint mode) {
//C.ReleasePrimitiveArrayCritical(envPtr(env), array, carray, mode);
//}
//
//func (env JNIEnv) const jchar *GetStringCritical(jstring string, jboolean *isCopy) {
//return C.GetStringCritical(envPtr(env), string, isCopy);
//}
//
//func (env JNIEnv) void ReleaseStringCritical(jstring string, const jchar *cstring) {
//C.ReleaseStringCritical(envPtr(env), string, cstring);
//}
//
//func (env JNIEnv) jweak NewWeakGlobalRef(obj JObject) {
//return C.NewWeakGlobalRef(envPtr(env), cObject(obj));
//}
//
//func (env JNIEnv) void DeleteWeakGlobalRef(jweak ref) {
//C.DeleteWeakGlobalRef(envPtr(env), ref);
//}
//
//func (env JNIEnv) jboolean ExceptionCheck(JNIEnv *env) {
//return C.ExceptionCheck(envPtr(env));
//}
//
//func (env JNIEnv) JObject NewDirectByteBuffer(void *address, jlong capacity) {
//return C.NewDirectByteBuffer(envPtr(env), address, capacity);
//}
//
//func (env JNIEnv) void *GetDirectBufferAddress(JObject buf) {
//return C.GetDirectBufferAddress(envPtr(env), buf);
//}

func (env JNIEnv) GetDirectBufferCapacity(buf JObject) JLong {
	return JLong(C.GetDirectBufferCapacity(envPtr(env), cObject(buf)))
}

// New JNI 1.6 Features
func (env JNIEnv) GetObjectRefType(obj JObject) JObjectRefType {
	return JObjectRefType(C.GetObjectRefType(envPtr(env), cObject(obj)))
}
