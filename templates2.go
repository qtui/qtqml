package qtqml

import (
	"github.com/kitech/gopp"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtrt"
)

type QQuickControl struct {
	*QQuickItem
}

func QQuickControlFromptr(ptr voidptr) *QQuickControl {
	return &QQuickControl{QQuickItemFromptr(ptr)}
}
func (me *QQuickControl) Dtor() { qtrt.Callany0(me) }

func NewQQuickControl(parent ...QQuickItemITF) *QQuickControl {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickControlFromptr(rv)
}

type QQuickAbstractButton struct {
	*QQuickControl
}

func QQuickAbstractButtonFromptr(ptr voidptr) *QQuickAbstractButton {
	return &QQuickAbstractButton{QQuickControlFromptr(ptr)}
}

type QQuickButton struct {
	*QQuickAbstractButton
}

func QQuickButtonFromptr(ptr voidptr) *QQuickButton {
	return &QQuickButton{QQuickAbstractButtonFromptr(ptr)}
}
func (me *QQuickButton) Dtor() { qtrt.Callany0(me) }

func NewQQuickButton(parent ...QQuickItemITF) *QQuickButton {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickButtonFromptr(rv)
}

type QQuickCheckBox struct {
	*QQuickAbstractButton
}

func QQuickCheckBoxFromptr(ptr voidptr) *QQuickCheckBox {
	return &QQuickCheckBox{QQuickAbstractButtonFromptr(ptr)}
}
func (me *QQuickCheckBox) Dtor() { qtrt.Callany0(me) }

func NewQQuickCheckBox(parent ...QQuickItemITF) *QQuickCheckBox {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickCheckBoxFromptr(rv)
}

type QQuickComboBox struct {
	*QQuickControl
}

func QQuickComboBoxFromptr(ptr voidptr) *QQuickComboBox {
	return &QQuickComboBox{QQuickControlFromptr(ptr)}
}
func (me *QQuickComboBox) Dtor() { qtrt.Callany0(me) }

func NewQQuickComboBox(parent ...QQuickItemITF) *QQuickComboBox {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickComboBoxFromptr(rv)
}

type QQuickSpinBox struct {
	*QQuickControl
}

func QQuickSpinBoxFromptr(ptr voidptr) *QQuickSpinBox {
	return &QQuickSpinBox{QQuickControlFromptr(ptr)}
}
func (me *QQuickSpinBox) Dtor() { qtrt.Callany0(me) }

func NewQQuickSpinBox(parent ...QQuickItemITF) *QQuickSpinBox {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickSpinBoxFromptr(rv)
}

type QQuickPopup struct {
	*qtcore.QObject
}

func QQuickPopupFromptr(ptr voidptr) *QQuickPopup {
	return &QQuickPopup{qtcore.QObjectFromptr(ptr)}
}

type QQuickDialog struct {
	*QQuickPopup
}

type QQuickLabel struct {
	*QQuickText
}

func QQuickLabelFromptr(ptr voidptr) *QQuickLabel {
	return &QQuickLabel{QQuickTextFromptr(ptr)}
}
func (me *QQuickLabel) Dtor() { qtrt.Callany0(me) }

func NewQQuickLabel(parent ...QQuickItemITF) *QQuickLabel {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickLabelFromptr(rv)
}

type QQuickIcon struct {
	*qtrt.CObject
}

func QQuickIconFromptr(ptr voidptr) *QQuickIcon {
	return &QQuickIcon{qtrt.CObjectFromptr(ptr)}
}
func (me *QQuickIcon) Dtor() { qtrt.Callany0(me) }

func NewQQuickIcon() *QQuickIcon {
	rv := qtrt.Callany[voidptr](nil)
	return QQuickIconFromptr(rv)
}

func (me *QQuickIcon) SetName(name string) { qtrt.Callany0(me, name) }
func (me *QQuickIcon) SetSource(uo *qtcore.QUrl) {
	qtrt.Callany0(me, uo)
}

func (me *QQuickIcon) SetWidth(w int)  { qtrt.Callany0(me, w) }
func (me *QQuickIcon) SetHeight(h int) { qtrt.Callany0(me, h) }
func (me *QQuickIcon) SetCache(b bool) { qtrt.Callany0(me, b) }

type QQuickAction struct {
	*qtcore.QObject
}

func QQuickActionFromptr(ptr voidptr) *QQuickAction {
	return &QQuickAction{qtcore.QObjectFromptr(ptr)}
}
func (me *QQuickAction) Dtor() { qtrt.Callany0(me) }

func NewQQuickAction(parent ...qtcore.QObjectITF) *QQuickAction {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickActionFromptr(rv)
}

func (me *QQuickAction) SetText(text string) {
	qtrt.Callany0(me, text)
}
func (me *QQuickAction) SetIcon(icon *QQuickIcon) {
	qtrt.Callany0(me, icon)
}
func (me *QQuickAction) SetEnabled(b bool)   { qtrt.Callany0(me, b) }
func (me *QQuickAction) SetChecked(b bool)   { qtrt.Callany0(me, b) }
func (me *QQuickAction) SetCheckable(b bool) { qtrt.Callany0(me, b) }

type QQuickPane struct {
	*QQuickControl
}

func QQuickPaneFromptr(ptr voidptr) *QQuickPane {
	return &QQuickPane{QQuickControlFromptr(ptr)}
}
func (me *QQuickPane) Dtor() { qtrt.Callany0(me) }

func NewQQuickPane(parent ...qtcore.QObjectITF) *QQuickPane {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickPaneFromptr(rv)
}

type QQuickPage struct {
	*QQuickPane
}

func QQuickPageFromptr(ptr voidptr) *QQuickPage {
	return &QQuickPage{QQuickPaneFromptr(ptr)}
}
func (me *QQuickPage) Dtor() { qtrt.Callany0(me) }

func NewQQuickPage(parent ...qtcore.QObjectITF) *QQuickPage {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickPageFromptr(rv)
}

type QQuickScrollView struct {
	*QQuickPane
}

func QQuickScrollViewFromptr(ptr voidptr) *QQuickScrollView {
	return &QQuickScrollView{QQuickPaneFromptr(ptr)}
}
func (me *QQuickScrollView) Dtor() { qtrt.Callany0(me) }

func NewQQuickScrollView(parent ...qtcore.QObjectITF) *QQuickScrollView {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickScrollViewFromptr(rv)
}

type QQuickContainer struct {
	*QQuickControl
}

func QQuickContainerFromptr(ptr voidptr) *QQuickContainer {
	return &QQuickContainer{QQuickControlFromptr(ptr)}
}
func (me *QQuickContainer) Dtor() { qtrt.Callany0(me) }

func NewQQuickContainer(parent ...qtcore.QObjectITF) *QQuickContainer {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickContainerFromptr(rv)
}

func (me *QQuickContainer) Count() int { return qtrt.Callany[int](me) }

func (me *QQuickContainer) AddItem(item *QQuickItemITF) {
	qtrt.Callany0(me, item)
}
func (me *QQuickContainer) RemoveItem(item *QQuickItemITF) {
	qtrt.Callany0(me, item)
}
func (me *QQuickContainer) InsetItem(idx int, item QQuickItemITF) {
	qtrt.Callany0(me, idx, item)
}
func (me *QQuickContainer) ItemAt(idx int) (item QQuickItemITF) {
	rv := qtrt.Callany[voidptr](me, idx)
	return QQuickItemFromptr(rv)
}
func (me *QQuickContainer) CurrentIndex() int { return qtrt.Callany[int](me) }
func (me *QQuickContainer) CurrentItem() *QQuickItem {
	rv := qtrt.Callany[voidptr](me)
	return QQuickItemFromptr(rv)
}
func (me *QQuickContainer) SetCurrentIndex(idx int) { qtrt.Callany0(me, idx) }
func (me *QQuickContainer) IncrementCurrentIndex()  { qtrt.Callany0(me) }
func (me *QQuickContainer) DecrementCurrentIndex()  { qtrt.Callany0(me) }

type QQuickSplitView struct {
	*QQuickContainer
}

func QQuickSplitViewFromptr(ptr voidptr) *QQuickSplitView {
	return &QQuickSplitView{QQuickContainerFromptr(ptr)}
}
func (me *QQuickSplitView) Dtor() { qtrt.Callany0(me) }

func NewQQuickSplitView(parent ...qtcore.QObjectITF) *QQuickSplitView {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickSplitViewFromptr(rv)
}

type QQuickContentItem struct {
	*QQuickItem
}

func QQuickContentItemFromptr(ptr voidptr) *QQuickContentItem {
	return &QQuickContentItem{QQuickItemFromptr(ptr)}
}
func (me *QQuickContentItem) Dtor() { qtrt.Callany0(me) }

func NewQQuickContentItem(parent ...qtcore.QObjectITF) *QQuickContentItem {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickContentItemFromptr(rv)
}

type QQuickProgressBar struct {
	*QQuickControl
}

func QQuickProgressBarFromptr(ptr voidptr) *QQuickProgressBar {
	return &QQuickProgressBar{QQuickControlFromptr(ptr)}
}
func (me *QQuickProgressBar) Dtor() { qtrt.Callany0(me) }

func NewQQuickProgressBar(parent ...qtcore.QObjectITF) *QQuickProgressBar {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickProgressBarFromptr(rv)
}

type QQuickStackView struct {
	*QQuickControl
}

func QQuickStackViewFromptr(ptr voidptr) *QQuickStackView {
	return &QQuickStackView{QQuickControlFromptr(ptr)}
}
func (me *QQuickStackView) Dtor() { qtrt.Callany0(me) }

func NewQQuickStackView(parent ...qtcore.QObjectITF) *QQuickStackView {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickStackViewFromptr(rv)
}

type QQuickTextArea struct {
	*QQuickTextEdit
}

func QQuickTextAreaFromptr(ptr voidptr) *QQuickTextArea {
	return &QQuickTextArea{QQuickTextEditFromptr(ptr)}
}

func NewQQuickTextArea(parent ...QQuickItemITF) *QQuickTextArea {
	// v := gopp.FirstofGv(parent)
	// log.Println(v, v.QQuickItemPTR().Dbgstr())
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickTextAreaFromptr(rv)
}

type QQuickSwipe struct {
	*qtcore.QObject // todo?
}

func QQuickSwipeFromptr(ptr voidptr) *QQuickSwipe {
	return &QQuickSwipe{qtcore.QObjectFromptr(ptr)}
}
func (me *QQuickSwipe) Dtor() { qtrt.Callany0(me) }

func NewQQuickSwipe(parent ...qtcore.QObjectITF) *QQuickSwipe {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickSwipeFromptr(rv)
}

type QQuickMenu struct {
	*QQuickPopup
}

func QQuickMenuFromptr(ptr voidptr) *QQuickMenu {
	return &QQuickMenu{QQuickPopupFromptr(ptr)}
}

func (me *QQuickMenu) Dtor() { qtrt.Callany0(me) }

func NewQQuickMenu(parent ...qtcore.QObjectITF) *QQuickMenu {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickMenuFromptr(rv)
}

func (me *QQuickMenu) AddItem(item *QQuickItemITF) {

}
func (me *QQuickMenu) InsertItem(idx int, item *QQuickItemITF) {

}
func (me *QQuickMenu) ItemAt(idx int) (item *QQuickItem) {
	return
}
func (me *QQuickMenu) RemoteItem(item *QQuickItem) {

}
func (me *QQuickMenu) SetTitle(title string) {

}
func (me *QQuickMenu) SetIcon(icon any) {

}

func (me *QQuickMenu) AddMenu(item *QQuickMenu) {
	qtrt.Callany0(me, item)
}
func (me *QQuickMenu) InsertMenu(idx int, item *QQuickMenu) {
	qtrt.Callany0(me, idx, item)
}
func (me *QQuickMenu) MenuAt(idx int) (item *QQuickMenu) {
	rv := qtrt.Callany[voidptr](me, idx)
	return QQuickMenuFromptr(rv)
}
func (me *QQuickMenu) RemoveMenu(item *QQuickMenu) {
	qtrt.Callany0(me, item)
}

func (me *QQuickMenu) AddAction(item *QQuickAction) {
	qtrt.Callany0(me, item)
}
func (me *QQuickMenu) InsertAction(idx int, item *QQuickAction) {
	qtrt.Callany0(me, idx, item)
}
func (me *QQuickMenu) ActionAt(idx int) (item *QQuickAction) {
	rv := qtrt.Callany[voidptr](me, idx)
	return QQuickActionFromptr(rv)
}
func (me *QQuickMenu) RemoveAction(item *QQuickAction) {
	qtrt.Callany0(me, item)
}
func (me *QQuickMenu) Popup(menuItem ...QQuickItemITF) {
	qtrt.Callany0(me, gopp.FirstofGv(menuItem))
}
func (me *QQuickMenu) Popupz2(pos qtcore.QPointF, menuItem QQuickItemITF) {
	qtrt.Callany0(me, pos, menuItem)
}

func (me *QQuickMenu) Dismiss() { qtrt.Callany0(me) }
