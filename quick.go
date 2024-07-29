package qtqml

import (
	"runtime"

	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
	"github.com/qtui/qtclzsz"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtrt"
)

type QQuickApplication struct {
	*qtcore.QObject
}

// todo Materail.theme =
type QQuickApplicationWindow struct {
	*QQuickItem
}

func QQuickApplicationWindowFromptr(ptr voidptr) *QQuickApplicationWindow {
	return &QQuickApplicationWindow{QQuickItemFromptr(ptr)}
}

// func NewQQuickApplicationWindow(parent ...qtcore.QWindow)
func NewQQuickApplicationWindow() *QQuickApplicationWindow {
	rv := qtrt.Callany[voidptr](nil, nil)
	return QQuickApplicationWindowFromptr(rv)
}

func (me *QQuickApplicationWindow) ContentItem() *QQuickItem {
	rv := qtrt.Callany[voidptr](me)
	return QQuickItemFromptr(rv)
}
func (me *QQuickApplicationWindow) Header() *QQuickItem {
	rv := qtrt.Callany[voidptr](me)
	return QQuickItemFromptr(rv)
}
func (me *QQuickApplicationWindow) Footer() *QQuickItem {
	rv := qtrt.Callany[voidptr](me)
	return QQuickItemFromptr(rv)
}
func (me *QQuickApplicationWindow) SetHeader(item *QQuickItem) {
	qtrt.Callany0(me, item)
}
func (me *QQuickApplicationWindow) SetFooter(item *QQuickItem) {
	qtrt.Callany0(me, item)
}

type QQuickItem struct {
	*qtcore.QObject
}
type QQuickItemITF interface {
	QQuickItemPTR() *QQuickItem
}

func (me *QQuickItem) QQuickItemPTR() *QQuickItem { return me }

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
func (me *QQuickItem) ChildItems() *qtcore.QList {
	rovp := cgopp.Mallocpg(qtclzsz.Get("QList"))
	qtrt.CallanyRov[int](rovp, me)
	lst := qtcore.QListFromptr(rovp)
	runtime.SetFinalizer(lst, qtcore.QListDtor)
	return lst
}

// /////////
type QQuickWindow struct {
	*qtcore.QWindow
}

func QQuickWindowFromptr(ptr voidptr) *QQuickWindow {
	return &QQuickWindow{qtcore.QWindowFromptr(ptr)}
}

func (me *QQuickWindow) ContentItem() *QQuickItem {
	rv := qtrt.Callany[voidptr](me)
	return QQuickItemFromptr(rv)
}

type QQuickView struct {
	*QQuickWindow
}

func QQuickViewFromptr(ptr voidptr) *QQuickView {
	return &QQuickView{QQuickWindowFromptr(ptr)}
}

func NewQQuickView(parent ...qtcore.QWindowITF) *QQuickView {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickViewFromptr(rv)
}

func (me *QQuickView) SetSource(file string) {
	uo := qtcore.NewQUrl(file)
	// log.Println(uo.Url())
	qtrt.Callany0(me, uo)
}

func (me *QQuickView) RootObject() *QQuickItem {
	rv := qtrt.Callany[voidptr](me)
	return QQuickItemFromptr(rv)
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
	*QQuickImplicitSizeItem
}

func QQuickTextFromptr(ptr voidptr) *QQuickText {
	return &QQuickText{QQuickImplicitSizeItemFromptr(ptr)}
}

func (me *QQuickText) SetText(text string) { qtrt.Callany0(me, text) }

type QQuickTextEdit struct {
	*QQuickItem
}

func QQuickTextEditFromptr(ptr voidptr) *QQuickTextEdit {
	return &QQuickTextEdit{QQuickItemFromptr(ptr)}
}
