package qtqml

import (
	"reflect"
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

var Material = struct {
	// Theme
	Light  int
	Dark   int
	System int

	_ int // 分隔符号，重新计数
	// Color
	Red        int
	Pink       int
	Purple     int
	DeepPurple int
	Indigo     int
	Blue       int
	LightBlue  int
	Cyan       int
	Teal       int
	Green      int
	LightGreen int
	Lime       int
	Yellow     int
	Amber      int
	Orange     int
	DeepOrange int
	Brown      int
	Grey       int
	BlueGrey   int
}{
	// 0, 1, 2,
}

// 自动给常量赋值, 等于在结构体中的Index
func init() {
	InitEnums(&Material)
}

// must a struct
func InitEnums[T any](vx *T) {
	ve := reflect.ValueOf(vx).Elem()
	ty := ve.Type()
	var idx int
	for i := 0; i < ve.NumField(); i++ {
		if ty.Field(i).Name == "_" {
			idx = 0
			continue
		}

		f := ve.Field(i)
		f.Set(reflect.ValueOf(idx))
		idx++
	}
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
