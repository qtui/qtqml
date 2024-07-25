package qtqml

import (
	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtrt"
)

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
