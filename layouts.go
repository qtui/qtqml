package qtqml

import (
	"github.com/kitech/gopp"
	"github.com/qtui/qtrt"
)

type QQuickLayout struct {
	*QQuickItem
}

func QQuickLayoutFromptr(ptr voidptr) *QQuickLayout {
	return &QQuickLayout{QQuickItemFromptr(ptr)}
}

type QQuickGridLayoutBase struct {
	*QQuickLayout
}

func QQuickGridLayoutBaseFromptr(ptr voidptr) *QQuickGridLayoutBase {
	return &QQuickGridLayoutBase{QQuickLayoutFromptr(ptr)}
}

type QQuickGridLayout struct {
	*QQuickGridLayoutBase
}

func QQuickGridLayoutFromptr(ptr voidptr) *QQuickGridLayout {
	return &QQuickGridLayout{QQuickGridLayoutBaseFromptr(ptr)}
}

func (me *QQuickGridLayout) Dtor() { qtrt.Callany0(me) }
func NewQQuickGridLayout(parent ...QQuickItemITF) *QQuickGridLayout {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickGridLayoutFromptr(rv)
}

type QQuickLinearLayout struct {
	*QQuickGridLayoutBase
}

func QQuickLinearLayoutFromptr(ptr voidptr) *QQuickLinearLayout {
	return &QQuickLinearLayout{QQuickGridLayoutBaseFromptr(ptr)}
}

type QQuickRowLayout struct {
	*QQuickLinearLayout
}

func QQuickRowLayoutFromptr(ptr voidptr) *QQuickRowLayout {
	return &QQuickRowLayout{QQuickLinearLayoutFromptr(ptr)}
}
func (me *QQuickRowLayout) Dtor() { qtrt.Callany0(me) }
func NewQQuickRowLayout(parent ...QQuickItemITF) *QQuickColumnLayout {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickColumnLayoutFromptr(rv)
}

type QQuickColumnLayout struct {
	*QQuickLinearLayout
}

func QQuickColumnLayoutFromptr(ptr voidptr) *QQuickColumnLayout {
	return &QQuickColumnLayout{QQuickLinearLayoutFromptr(ptr)}
}
func (me *QQuickColumnLayout) Dtor() { qtrt.Callany0(me) }
func NewQQuickColumnLayout(parent ...QQuickItemITF) *QQuickColumnLayout {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickColumnLayoutFromptr(rv)
}
