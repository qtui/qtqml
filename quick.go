package qtqml

import (
	"github.com/kitech/gopp"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtrt"
)

type QQuickItem struct {
	*qtcore.QObject
}

func QQuickItemFromptr(ptr voidptr) *QQuickItem {
	return &QQuickItem{qtcore.QObjectFromptr(ptr)}
}

func (me *QQuickItem) Dtor() {
	qtrt.Callany0(me)
}

func NewQQuickItem(parent ...*QQuickItem) *QQuickItem {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickItemFromptr(rv)
}

func (me *QQuickItem) SetZ(v float64)              { qtrt.Callany0(me, v) }
func (me *QQuickItem) SetOpacity(v float64)        { qtrt.Callany0(me, v) }
func (me *QQuickItem) SetVisible(v bool)           { qtrt.Callany0(me, v) }
func (me *QQuickItem) SetHeight(v float64)         { qtrt.Callany0(me, v) }
func (me *QQuickItem) SetWidth(v float64)          { qtrt.Callany0(me, v) }
func (me *QQuickItem) SetImplicitHeight(v float64) { qtrt.Callany0(me, v) }
func (me *QQuickItem) SetImplicitWidth(v float64)  { qtrt.Callany0(me, v) }
func (me *QQuickItem) SetX(v float64)              { qtrt.Callany0(me, v) }
func (me *QQuickItem) SetY(v float64)              { qtrt.Callany0(me, v) }
func (me *QQuickItem) IsVisible() (v bool)         { return qtrt.Callany[bool](me) }
func (me *QQuickItem) IsEnabled() (v bool)         { return qtrt.Callany[bool](me) }
func (me *QQuickItem) X() (v float64)              { return qtrt.Callany[float64](me) }
func (me *QQuickItem) Y() (v float64)              { return qtrt.Callany[float64](me) }
func (me *QQuickItem) Width() (v float64)          { return qtrt.Callany[float64](me) }
func (me *QQuickItem) Height() (v float64)         { return qtrt.Callany[float64](me) }
func (me *QQuickItem) ImplicitWidth() (v float64)  { return qtrt.Callany[float64](me) }
func (me *QQuickItem) ImplicitHeight() (v float64) { return qtrt.Callany[float64](me) }

func (me *QQuickItem) ViewportItem() *QQuickItem {
	rv := qtrt.Callany[voidptr](me)
	return QQuickItemFromptr(rv)
}

func (me *QQuickItem) Polish() { qtrt.Callany0(me) }

// /////////
type QQuickWindow struct {
	*qtcore.QObject
}

func QQuickWindowFromptr(ptr voidptr) *QQuickWindow {
	return &QQuickWindow{qtcore.QObjectFromptr(ptr)}
}

type QQuickView struct {
	*QQuickWindow
}

func QQuickViewFromptr(ptr voidptr) *QQuickView {
	return &QQuickView{QQuickWindowFromptr(ptr)}
}

type QQuickFlickable struct {
	*QQuickItem
}

func QQuickFlickableFromptr(ptr voidptr) *QQuickFlickable {
	return &QQuickFlickable{QQuickItemFromptr(ptr)}
}

type QQuickItemView struct {
	*QQuickFlickable
}

func QQuickItemViewFromptr(ptr voidptr) *QQuickItemView {
	return &QQuickItemView{QQuickFlickableFromptr(ptr)}
}

type QQuickListView struct {
	*QQuickItemView
}

func QQuickListViewFromptr(ptr voidptr) *QQuickListView {
	return &QQuickListView{QQuickItemViewFromptr(ptr)}
}

type QQuickImplicitSizeItem struct {
	*QQuickItem
}

func QQuickImplicitSizeItemFromptr(ptr voidptr) *QQuickImplicitSizeItem {
	return &QQuickImplicitSizeItem{QQuickItemFromptr(ptr)}
}

type QQuickLoader struct {
	*QQuickImplicitSizeItem
}

func QQuickLoaderFromptr(ptr voidptr) *QQuickLoader {
	return &QQuickLoader{QQuickImplicitSizeItemFromptr(ptr)}
}

type QQuickText struct {
}

type QQuickTextEdit struct {
}
