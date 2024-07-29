package qtqml

import (
	"runtime"

	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
	"github.com/qtui/qtclzsz"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtrt"
)

// 想到一种减少go包大小的方式，把方法写成成员变量函数。减少go符号表

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

func (me *QJSEngine) GlobalObject() *QJSValue {
	rovp := cgopp.Mallocpg(qtclzsz.Get("QJSValue"))
	qtrt.CallanyRov[int](rovp, me)
	return QJSValueFromptr(rovp)
}
func (me *QJSEngine) NewQObject(obj qtcore.QObjectITF) *QJSValue {
	rovp := cgopp.Mallocpg(qtclzsz.Get("QJSValue"))
	qtrt.CallanyRov[int](rovp, me, obj)
	return QJSValueFromptr(rovp)
}

type QJSValue struct {
	*qtrt.CObject
}

func QJSValueFromptr(ptr voidptr) *QJSValue {
	return &QJSValue{qtrt.CObjectFromptr(ptr)}
}

type QQmlEngine struct {
	*QJSEngine
}

func QQmlEngineFromptr(ptr voidptr) *QQmlEngine {
	return &QQmlEngine{QJSEngineFromptr(ptr)}
}

func (me *QQmlEngine) ContextForObject(qobj qtcore.QObjectITF) *QQmlContext {
	return QQmlEngine_ContextForObject(qobj)
}
func QQmlEngine_ContextForObject(qobj qtcore.QObjectITF) *QQmlContext {
	rv := qtrt.Callany[voidptr](nil, qobj)
	return QQmlContextFromptr(rv)
}

type QQmlApplicationEngine struct {
	*QQmlEngine
}

func QQmlApplicationEngineFromptr(ptr voidptr) *QQmlApplicationEngine {
	return &QQmlApplicationEngine{QQmlEngineFromptr(ptr)}
}

func (me *QQmlApplicationEngine) Dtor() { qtrt.Callany0(me) }

func NewQQmlApplicationEngine(parent ...qtcore.QObjectITF) *QQmlApplicationEngine {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQmlApplicationEngineFromptr(rv)
	// name := "QQmlApplicationEngineNew"
	// fnsym := qtrt.GetQtSymAddr(name)
	// log.Println(fnsym)
	// rv := cgopp.Litfficall(fnsym)
	// log.Println(rv)
	// return QQmlApplicationEngineFromptr(rv)
}

// todo 这个经常出现crash？？？+++
// 可能是要尽快初始化QQmlApplicationEngine，在QApplication之后
func (me *QQmlApplicationEngine) Load(file string) {
	qtrt.Callany0(me, file)
	// name := "QQmlApplicationEngineLoad1"
	// fnsym := qtrt.GetQtSymAddr(name)
	// file4c := cgopp.CStringaf(file)
	// // cgopp.FfiCall[int](fnsym, me.GetCthis(), file4c)
	// cgopp.Litfficall(fnsym, me.GetCthis(), file4c)

}

// todo 这个方法不是Qt带的
func (me *QQmlApplicationEngine) RootObject() *qtcore.QObject {
	symname := "QQmlApplicationEngineRootObject1"
	fnsym := qtrt.GetQtSymAddr(symname)
	rv := cgopp.FfiCall[voidptr](fnsym, me.GetCthis())
	return qtcore.QObjectFromptr(rv)
}

func (me *QQmlApplicationEngine) RootObjects() *qtcore.QList {
	rovp := cgopp.Mallocpg(qtclzsz.Get("QList"))
	qtrt.CallanyRov[voidptr](rovp, me)
	return qtcore.QListFromptr(rovp)
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

func NewQQmlProperty(qobj qtcore.QObjectITF, name string, ctx *QQmlContext) *QQmlProperty {
	// hotfix overload name resolusion conflict with
	// QQmlProperty::QQmlProperty(QObject*, QString const&, QQmlEngine*)
	if true {
		cthis := cgopp.Mallocpg(qtclzsz.Get("QQmlProperty"))
		fnadr := qtrt.GetQtSymAddr("__ZN12QQmlPropertyC1EP7QObjectRK7QStringP11QQmlContext")
		qstr := qtcore.QString_FromUtf8(name)
		cgopp.FfiCall[int](fnadr, cthis, qobj.QObjectPTR().GetCthis(), qstr.GetCthis(), ctx.GetCthis())
		me := QQmlPropertyFromptr(cthis)
		runtime.SetFinalizer(me, QQmlPropertyDtor)
		return me
	}

	rv := qtrt.Callany[voidptr](nil, qobj, name, ctx)
	me := QQmlPropertyFromptr(rv)
	runtime.SetFinalizer(me, QQmlPropertyDtor)
	return me
}

// __ZN12QQmlPropertyC2EP7QObjectRK7QStringP10QQmlEngine not work, dont use
// QQmlProperty::QQmlProperty(QObject*, QString const&, QQmlEngine*)
// func NewQQmlProperty(qobj qtcore.QObjectITF, name string, ctx *QQmlEngine) *QQmlProperty {
// 	rv := qtrt.Callany[voidptr](nil, qobj, name, ctx)
// 	me := QQmlPropertyFromptr(rv)
// 	runtime.SetFinalizer(me, QQmlPropertyDtor)
// 	return me
// }

func (me *QQmlProperty) IsValid() bool { return qtrt.Callany[bool](me) }
func (me *QQmlProperty) PropertyTypeName() string {
	rv := qtrt.Callany[voidptr](me)
	return cgopp.GoString(rv)
}
func (me *QQmlProperty) Name() string {
	rovp := cgopp.Mallocpg(qtclzsz.Get("QString"))
	qtrt.CallanyRov[voidptr](rovp, me)
	qstr := qtcore.QStringFromptr(rovp)
	defer qstr.Dtor()

	return qstr.ToUtf8().ConstData()
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

func (me *QQmlProperty) Object() *qtcore.QObject {
	rv := qtrt.Callany[voidptr](me)
	return qtcore.QObjectFromptr(rv)
}
