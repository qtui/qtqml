package qtqml

import (
	"github.com/kitech/gopp"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtrt"
)

type QQuickControl struct {
	*QQuickItem
}

type QQuickAbstractButton struct {
	*QQuickControl
}

type QQuickButton struct {
	*QQuickAbstractButton
}

type QQuickCheckBox struct {
	*QQuickAbstractButton
}

type QQuickComboBox struct {
	*QQuickControl
}

type QQuickSpinBox struct {
	*QQuickControl
}

type QQuickPopup struct {
	*qtcore.QObject
}

type QQuickDialog struct {
	*QQuickPopup
}

type QQuickLabel struct {
	*QQuickText
}

type QQuickPane struct {
	*QQuickControl
}
type QQuickPage struct {
	*QQuickPane
}
type QQuickScrollView struct {
	*QQuickPane
}

type QQuickContainer struct {
	*QQuickControl
}

type QQuickSplitView struct {
	*QQuickContainer
}

type QQuickContentItem struct {
	*QQuickItem
}

type QQuickProgressBar struct {
	*QQuickControl
}

type QQuickStackView struct {
	*QQuickControl
}

type QQuickTextArea struct {
	*QQuickTextEdit
}

func QQuickTextAreaFromptr(ptr voidptr) *QQuickTextArea {
	return &QQuickTextArea{QQuickTextEditFromptr(ptr)}
}

func NewQQuickTextArea(parent ...QQuickItemITF) *QQuickTextArea {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QQuickTextAreaFromptr(rv)
}

type QQuickSwipe struct {
	*qtcore.QObject
}
