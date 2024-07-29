package qtqml

import (
	"log"
	"testing"
)

func TestQQmlProperty1(t *testing.T) {
	var qmlape *QQmlApplicationEngine // todo test cannot get???
	if true {
		return
	}

	var robj = qmlape.RootObject()
	var qmlctx = qmlape.ContextForObject(robj)
	log.Println(qmlctx)
	log.Println(qmlctx.GetCthis(), qmlctx.Dbgstr())
	// log.Println(qmlctx.Dbgstr())
	// log.Println(qvar.ToPtr())
	xprop := NewQQmlProperty(robj, "Material.theme", qmlctx)
	log.Println(xprop)
	log.Println(xprop.IsValid(), xprop.PropertyTypeName()) // QQuickMaterialStyle::Theme
	qvar := xprop.Read()
	log.Println(qvar)
	log.Println(qvar.GetCthis())
	log.Println(qvar.ToPtr(), qvar.TypeName()) // QQuickMaterialStyle::Theme
	log.Println(qvar.ToString())               // works

	themeo := xprop.Object()
	log.Println(themeo)
	log.Println(themeo.Dbgstr())

	xprop = NewQQmlProperty(robj, "Material.background", qmlctx)
	log.Println(xprop, xprop.IsValid(), xprop.GetCthis(), xprop.Object(), xprop.Object().Dbgstr())
	qvar = xprop.Read()
	log.Println(xprop.PropertyTypeName(), qvar.ToString(), qvar.TypeName())

}
