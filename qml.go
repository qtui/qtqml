package qtqml

import (
	"runtime"

	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
	"github.com/qtui/qtclzsz"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtrt"
)

// 想到一种减少go包大小的方式，把方法写成成员变量函数。减少go的符号表

type QQmlContext struct {
	*qtcore.QObject
}

func QQmlContextFromptr(ptr voidptr) *QQmlContext {
	return &QQmlContext{qtcore.QObjectFromptr(ptr)}
}

type QJSEngine struct {
	*qtcore.QObject
}

func QJSEngineFromptr(ptr voidptr) *QJSEngine {
	return &QJSEngine{qtcore.QObjectFromptr(ptr)}
}

func (me *QJSEngine) CollectGarbage() { qtrt.Callany0(me) }
func (me *QJSEngine) ObjectOwnership(obj qtcore.QObjectITF) int {
	return QJSEngine_objectOwnership(obj)
}
func QJSEngine_objectOwnership(obj qtcore.QObjectITF) int {
	return qtrt.Callany[int](nil, obj)
}

func (me *QJSEngine) SetUiLanguage(lang string) {
	qtrt.Callany0(me, lang)
}

func (me *QJSEngine) HasError() bool { return qtrt.Callany[bool](me) }

type QQmlEngine struct {
	*QJSEngine
}

func QQmlEngineFromptr(ptr voidptr) *QQmlEngine {
	return &QQmlEngine{QJSEngineFromptr(ptr)}
}

func (me *QQmlEngine) ContextForObject(obj qtcore.QObjectITF) *QQmlContext {
	rv := qtrt.Callany[voidptr](me, obj)
	return QQmlContextFromptr(rv)
}

func (me *QQmlEngine) contextForObject(qobj qtcore.QObjectITF) *QQmlContext {
	return QQmlEngine_contextForObject(qobj)
}
func QQmlEngine_contextForObject(qobj qtcore.QObjectITF) *QQmlContext {
	rv := qtrt.Callany[voidptr](nil, qobj)
	return QQmlContextFromptr(rv)
}

type QQmlApplicationEngine struct {
	*QQmlEngine
}

func QQmlApplicationEngineFromptr(ptr voidptr) *QQmlApplicationEngine {
	return &QQmlApplicationEngine{QQmlEngineFromptr(ptr)}
}

func NewQQmlApplicationEngine(parent ...qtcore.QObjectITF) *QQmlApplicationEngine {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQmlApplicationEngineFromptr(rv)
}

func (me *QQmlApplicationEngine) Load(file string) {
	qtrt.Callany0(me, file)
}

// todo
func (me *QQmlApplicationEngine) RootObject() *qtcore.QObject {
	symname := "QQmlApplicationEngineRootObject1"
	fnsym := qtrt.GetQtSymAddr(symname)
	rv := cgopp.FfiCall[voidptr](fnsym, me.GetCthis())
	return qtcore.QObjectFromptr(rv)
}

// //////////
type QQmlProperty struct {
	*qtrt.CObject
}

func QQmlPropertyFromptr(ptr voidptr) *QQmlProperty {
	return &QQmlProperty{qtrt.CObjectFromptr(ptr)}
}

func (me *QQmlProperty) Dtor()          { qtrt.Callany0(me) }
func QQmlPropertyDtor(me *QQmlProperty) { me.Dtor() }

// QQmlProperty(obj, QString(name), ctx);
func NewQQmlProperty(qobj qtcore.QObjectITF, name string, ctx any) *QQmlProperty {
	rv := qtrt.Callany[voidptr](nil, qobj, name, ctx)
	return QQmlPropertyFromptr(rv)
}

func (me *QQmlProperty) PropertyTypeName() string {
	rv := qtrt.Callany[voidptr](me)
	return cgopp.GoString(rv)
}

func (me *QQmlProperty) Write(valx any) bool {
	qvar := qtcore.NewQVariant(valx)
	defer qvar.Dtor()

	rv := qtrt.Callany[bool](me, qvar)
	return rv
}
func (me *QQmlProperty) Read() *qtcore.QVariant {
	rovp := cgopp.Mallocpg(qtclzsz.Get("QVariant"))
	qtrt.CallanyRov[int](rovp, me)

	qvar := qtcore.QVariantFromptr(rovp)
	runtime.SetFinalizer(qvar, qtcore.QVariantDtor)
	return qvar
}
