package qtqml

import "github.com/qtui/qtcore"

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

type QQuickSwipe struct {
	*qtcore.QObject
}
