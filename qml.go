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

func NewQQmlApplicationEngine(parent ...qtcore.QObjectITF) *QQmlApplicationEngine {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQmlApplicationEngineFromptr(rv)
}

// todo 这个经常出现crash？？？
func (me *QQmlApplicationEngine) Load(file string) {
	qtrt.Callany0(me, file)
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
func NewQQmlProperty(qobj qtcore.QObjectITF, name string, ctx any) *QQmlProperty {
	rv := qtrt.Callany[voidptr](nil, qobj, name, ctx)
	me := QQmlPropertyFromptr(rv)
	runtime.SetFinalizer(me, QQmlPropertyDtor)
	return me
}

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
