package qtqml

import (
	"runtime"

	"github.com/kitech/gopp/cgopp"
	"github.com/qtui/qtclzsz"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtrt"
)

type QQuickAttachedPropertyPropagator struct {
	*qtcore.QObject
}

func QQuickAttachedPropertyPropagatorFromptr(ptr voidptr) *QQuickAttachedPropertyPropagator {
	return &QQuickAttachedPropertyPropagator{qtcore.QObjectFromptr(ptr)}
}

type QQuickMaterialStyle struct {
	*QQuickAttachedPropertyPropagator
}

func QQuickMaterialStyleFromptr(ptr voidptr) *QQuickMaterialStyle {
	return &QQuickMaterialStyle{QQuickAttachedPropertyPropagatorFromptr(ptr)}
}

func (me *QQuickMaterialStyle) Foreground() *qtcore.QVariant {
	rovp := cgopp.Mallocpg(qtclzsz.Get("QVariant"))
	qtrt.CallanyRov[int](rovp, me)

	rv := qtcore.QVariantFromptr(rovp)
	runtime.SetFinalizer(rv, qtcore.QVariantDtor)
	return rv
}
func (me *QQuickMaterialStyle) Background() *qtcore.QVariant {
	rovp := cgopp.Mallocpg(qtclzsz.Get("QVariant"))
	qtrt.CallanyRov[int](rovp, me)

	rv := qtcore.QVariantFromptr(rovp)
	runtime.SetFinalizer(rv, qtcore.QVariantDtor)
	return rv
}

func (me *QQuickMaterialStyle) SetForeground(v *qtcore.QVariant) {
	qtrt.Callany0(me, v)
}
func (me *QQuickMaterialStyle) SetBackground(v *qtcore.QVariant) {
	qtrt.Callany0(me, v)
}

// todo colors
func (me *QQuickMaterialStyle) primaryColor() {

}
