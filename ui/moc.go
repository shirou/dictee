package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QmlBridge_ITF interface {
	std_core.QObject_ITF
	QmlBridge_PTR() *QmlBridge
}

func (ptr *QmlBridge) QmlBridge_PTR() *QmlBridge {
	return ptr
}

func (ptr *QmlBridge) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlBridge) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlBridge(ptr QmlBridge_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlBridge_PTR().Pointer()
	}
	return nil
}

func NewQmlBridgeFromPointer(ptr unsafe.Pointer) *QmlBridge {
	var n *QmlBridge
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QmlBridge)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QmlBridge:
			n = deduced

		case *std_core.QObject:
			n = &QmlBridge{QObject: *deduced}

		default:
			n = new(QmlBridge)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackQmlBridge_Constructor
func callbackQmlBridge_Constructor(ptr unsafe.Pointer) {
	gPtr := NewQmlBridgeFromPointer(ptr)
	qt.Register(ptr, gPtr)
	gPtr.init()
}

//export callbackQmlBridge_OnSearchInputFinished
func callbackQmlBridge_OnSearchInputFinished(ptr unsafe.Pointer, text C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "onSearchInputFinished"); signal != nil {
		signal.(func(string))(cGoUnpackString(text))
	}

}

func (ptr *QmlBridge) ConnectOnSearchInputFinished(f func(text string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "onSearchInputFinished") {
			C.QmlBridge_ConnectOnSearchInputFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "onSearchInputFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "onSearchInputFinished", func(text string) {
				signal.(func(string))(text)
				f(text)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onSearchInputFinished", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectOnSearchInputFinished() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectOnSearchInputFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "onSearchInputFinished")
	}
}

func (ptr *QmlBridge) OnSearchInputFinished(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QmlBridge_OnSearchInputFinished(ptr.Pointer(), C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackQmlBridge_OnBackButtonClicked
func callbackQmlBridge_OnBackButtonClicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "onBackButtonClicked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectOnBackButtonClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "onBackButtonClicked") {
			C.QmlBridge_ConnectOnBackButtonClicked(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "onBackButtonClicked"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "onBackButtonClicked", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onBackButtonClicked", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectOnBackButtonClicked() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectOnBackButtonClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "onBackButtonClicked")
	}
}

func (ptr *QmlBridge) OnBackButtonClicked() {
	if ptr.Pointer() != nil {
		C.QmlBridge_OnBackButtonClicked(ptr.Pointer())
	}
}

//export callbackQmlBridge_OnForwardButtonClicked
func callbackQmlBridge_OnForwardButtonClicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "onForwardButtonClicked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectOnForwardButtonClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "onForwardButtonClicked") {
			C.QmlBridge_ConnectOnForwardButtonClicked(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "onForwardButtonClicked"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "onForwardButtonClicked", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onForwardButtonClicked", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectOnForwardButtonClicked() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectOnForwardButtonClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "onForwardButtonClicked")
	}
}

func (ptr *QmlBridge) OnForwardButtonClicked() {
	if ptr.Pointer() != nil {
		C.QmlBridge_OnForwardButtonClicked(ptr.Pointer())
	}
}

//export callbackQmlBridge_OnTitleClicked
func callbackQmlBridge_OnTitleClicked(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "onTitleClicked"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QmlBridge) ConnectOnTitleClicked(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "onTitleClicked") {
			C.QmlBridge_ConnectOnTitleClicked(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "onTitleClicked"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "onTitleClicked", func(index int) {
				signal.(func(int))(index)
				f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onTitleClicked", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectOnTitleClicked() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectOnTitleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "onTitleClicked")
	}
}

func (ptr *QmlBridge) OnTitleClicked(index int) {
	if ptr.Pointer() != nil {
		C.QmlBridge_OnTitleClicked(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQmlBridge_SetViewText
func callbackQmlBridge_SetViewText(ptr unsafe.Pointer, text C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setViewText"); signal != nil {
		signal.(func(string))(cGoUnpackString(text))
	}

}

func (ptr *QmlBridge) ConnectSetViewText(f func(text string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "setViewText") {
			C.QmlBridge_ConnectSetViewText(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "setViewText"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setViewText", func(text string) {
				signal.(func(string))(text)
				f(text)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setViewText", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectSetViewText() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectSetViewText(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "setViewText")
	}
}

func (ptr *QmlBridge) SetViewText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QmlBridge_SetViewText(ptr.Pointer(), C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func QmlBridge_QRegisterMetaType() int {
	return int(int32(C.QmlBridge_QmlBridge_QRegisterMetaType()))
}

func (ptr *QmlBridge) QRegisterMetaType() int {
	return int(int32(C.QmlBridge_QmlBridge_QRegisterMetaType()))
}

func QmlBridge_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge_QmlBridge_QRegisterMetaType2(typeNameC)))
}

func (ptr *QmlBridge) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge_QmlBridge_QRegisterMetaType2(typeNameC)))
}

func QmlBridge_QmlRegisterType() int {
	return int(int32(C.QmlBridge_QmlBridge_QmlRegisterType()))
}

func (ptr *QmlBridge) QmlRegisterType() int {
	return int(int32(C.QmlBridge_QmlBridge_QmlRegisterType()))
}

func QmlBridge_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridge_QmlBridge_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridge_QmlBridge_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.QmlBridge___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QmlBridge) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlBridge___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QmlBridge) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlBridge___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QmlBridge___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QmlBridge) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlBridge___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QmlBridge___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QmlBridge) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlBridge___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlBridge___findChildren_newList(ptr.Pointer()))
}

func (ptr *QmlBridge) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlBridge___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlBridge___children_newList(ptr.Pointer()))
}

func NewQmlBridge(parent std_core.QObject_ITF) *QmlBridge {
	var tmpValue = NewQmlBridgeFromPointer(C.QmlBridge_NewQmlBridge(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQmlBridge_DestroyQmlBridge
func callbackQmlBridge_DestroyQmlBridge(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QmlBridge"); signal != nil {
		signal.(func())()
	} else {
		NewQmlBridgeFromPointer(ptr).DestroyQmlBridgeDefault()
	}
}

func (ptr *QmlBridge) ConnectDestroyQmlBridge(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QmlBridge"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QmlBridge")
	}
}

func (ptr *QmlBridge) DestroyQmlBridge() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DestroyQmlBridge(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QmlBridge) DestroyQmlBridgeDefault() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DestroyQmlBridgeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlBridge_Event
func callbackQmlBridge_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QmlBridge) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlBridge_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQmlBridge_EventFilter
func callbackQmlBridge_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QmlBridge) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlBridge_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQmlBridge_ChildEvent
func callbackQmlBridge_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlBridge) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlBridge_ConnectNotify
func callbackQmlBridge_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge_CustomEvent
func callbackQmlBridge_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlBridge) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQmlBridge_DeleteLater
func callbackQmlBridge_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQmlBridgeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlBridge) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlBridge_Destroyed
func callbackQmlBridge_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQmlBridge_DisconnectNotify
func callbackQmlBridge_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge_ObjectNameChanged
func callbackQmlBridge_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQmlBridge_TimerEvent
func callbackQmlBridge_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlBridge) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type Title_ITF interface {
	std_core.QObject_ITF
	Title_PTR() *Title
}

func (ptr *Title) Title_PTR() *Title {
	return ptr
}

func (ptr *Title) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Title) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromTitle(ptr Title_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Title_PTR().Pointer()
	}
	return nil
}

func NewTitleFromPointer(ptr unsafe.Pointer) *Title {
	var n *Title
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Title)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Title:
			n = deduced

		case *std_core.QObject:
			n = &Title{QObject: *deduced}

		default:
			n = new(Title)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackTitle_Constructor
func callbackTitle_Constructor(ptr unsafe.Pointer) {
	gPtr := NewTitleFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackTitle_TitleName
func callbackTitle_TitleName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "titleName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTitleFromPointer(ptr).TitleNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Title) ConnectTitleName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "titleName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "titleName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleName", f)
		}
	}
}

func (ptr *Title) DisconnectTitleName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "titleName")
	}
}

func (ptr *Title) TitleName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Title_TitleName(ptr.Pointer()))
	}
	return ""
}

func (ptr *Title) TitleNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Title_TitleNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTitle_SetTitleName
func callbackTitle_SetTitleName(ptr unsafe.Pointer, titleName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTitleName"); signal != nil {
		signal.(func(string))(cGoUnpackString(titleName))
	} else {
		NewTitleFromPointer(ptr).SetTitleNameDefault(cGoUnpackString(titleName))
	}
}

func (ptr *Title) ConnectSetTitleName(f func(titleName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTitleName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTitleName", func(titleName string) {
				signal.(func(string))(titleName)
				f(titleName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTitleName", f)
		}
	}
}

func (ptr *Title) DisconnectSetTitleName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTitleName")
	}
}

func (ptr *Title) SetTitleName(titleName string) {
	if ptr.Pointer() != nil {
		var titleNameC *C.char
		if titleName != "" {
			titleNameC = C.CString(titleName)
			defer C.free(unsafe.Pointer(titleNameC))
		}
		C.Title_SetTitleName(ptr.Pointer(), C.struct_Moc_PackedString{data: titleNameC, len: C.longlong(len(titleName))})
	}
}

func (ptr *Title) SetTitleNameDefault(titleName string) {
	if ptr.Pointer() != nil {
		var titleNameC *C.char
		if titleName != "" {
			titleNameC = C.CString(titleName)
			defer C.free(unsafe.Pointer(titleNameC))
		}
		C.Title_SetTitleNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: titleNameC, len: C.longlong(len(titleName))})
	}
}

//export callbackTitle_TitleNameChanged
func callbackTitle_TitleNameChanged(ptr unsafe.Pointer, titleName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "titleNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(titleName))
	}

}

func (ptr *Title) ConnectTitleNameChanged(f func(titleName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleNameChanged") {
			C.Title_ConnectTitleNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "titleNameChanged", func(titleName string) {
				signal.(func(string))(titleName)
				f(titleName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleNameChanged", f)
		}
	}
}

func (ptr *Title) DisconnectTitleNameChanged() {
	if ptr.Pointer() != nil {
		C.Title_DisconnectTitleNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleNameChanged")
	}
}

func (ptr *Title) TitleNameChanged(titleName string) {
	if ptr.Pointer() != nil {
		var titleNameC *C.char
		if titleName != "" {
			titleNameC = C.CString(titleName)
			defer C.free(unsafe.Pointer(titleNameC))
		}
		C.Title_TitleNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: titleNameC, len: C.longlong(len(titleName))})
	}
}

//export callbackTitle_DictName
func callbackTitle_DictName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "dictName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTitleFromPointer(ptr).DictNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Title) ConnectDictName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "dictName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dictName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dictName", f)
		}
	}
}

func (ptr *Title) DisconnectDictName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "dictName")
	}
}

func (ptr *Title) DictName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Title_DictName(ptr.Pointer()))
	}
	return ""
}

func (ptr *Title) DictNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Title_DictNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTitle_SetDictName
func callbackTitle_SetDictName(ptr unsafe.Pointer, dictName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setDictName"); signal != nil {
		signal.(func(string))(cGoUnpackString(dictName))
	} else {
		NewTitleFromPointer(ptr).SetDictNameDefault(cGoUnpackString(dictName))
	}
}

func (ptr *Title) ConnectSetDictName(f func(dictName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDictName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDictName", func(dictName string) {
				signal.(func(string))(dictName)
				f(dictName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDictName", f)
		}
	}
}

func (ptr *Title) DisconnectSetDictName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDictName")
	}
}

func (ptr *Title) SetDictName(dictName string) {
	if ptr.Pointer() != nil {
		var dictNameC *C.char
		if dictName != "" {
			dictNameC = C.CString(dictName)
			defer C.free(unsafe.Pointer(dictNameC))
		}
		C.Title_SetDictName(ptr.Pointer(), C.struct_Moc_PackedString{data: dictNameC, len: C.longlong(len(dictName))})
	}
}

func (ptr *Title) SetDictNameDefault(dictName string) {
	if ptr.Pointer() != nil {
		var dictNameC *C.char
		if dictName != "" {
			dictNameC = C.CString(dictName)
			defer C.free(unsafe.Pointer(dictNameC))
		}
		C.Title_SetDictNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: dictNameC, len: C.longlong(len(dictName))})
	}
}

//export callbackTitle_DictNameChanged
func callbackTitle_DictNameChanged(ptr unsafe.Pointer, dictName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "dictNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(dictName))
	}

}

func (ptr *Title) ConnectDictNameChanged(f func(dictName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dictNameChanged") {
			C.Title_ConnectDictNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dictNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dictNameChanged", func(dictName string) {
				signal.(func(string))(dictName)
				f(dictName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dictNameChanged", f)
		}
	}
}

func (ptr *Title) DisconnectDictNameChanged() {
	if ptr.Pointer() != nil {
		C.Title_DisconnectDictNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dictNameChanged")
	}
}

func (ptr *Title) DictNameChanged(dictName string) {
	if ptr.Pointer() != nil {
		var dictNameC *C.char
		if dictName != "" {
			dictNameC = C.CString(dictName)
			defer C.free(unsafe.Pointer(dictNameC))
		}
		C.Title_DictNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: dictNameC, len: C.longlong(len(dictName))})
	}
}

func Title_QRegisterMetaType() int {
	return int(int32(C.Title_Title_QRegisterMetaType()))
}

func (ptr *Title) QRegisterMetaType() int {
	return int(int32(C.Title_Title_QRegisterMetaType()))
}

func Title_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Title_Title_QRegisterMetaType2(typeNameC)))
}

func (ptr *Title) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Title_Title_QRegisterMetaType2(typeNameC)))
}

func Title_QmlRegisterType() int {
	return int(int32(C.Title_Title_QmlRegisterType()))
}

func (ptr *Title) QmlRegisterType() int {
	return int(int32(C.Title_Title_QmlRegisterType()))
}

func Title_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Title_Title_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Title) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Title_Title_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Title) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.Title___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Title) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Title___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Title) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Title___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *Title) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Title___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Title) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Title___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Title) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.Title___findChildren_newList2(ptr.Pointer()))
}

func (ptr *Title) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Title___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Title) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Title___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Title) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.Title___findChildren_newList3(ptr.Pointer()))
}

func (ptr *Title) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Title___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Title) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Title___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Title) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Title___findChildren_newList(ptr.Pointer()))
}

func (ptr *Title) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Title___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Title) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Title___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Title) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Title___children_newList(ptr.Pointer()))
}

func NewTitle(parent std_core.QObject_ITF) *Title {
	var tmpValue = NewTitleFromPointer(C.Title_NewTitle(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTitle_DestroyTitle
func callbackTitle_DestroyTitle(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Title"); signal != nil {
		signal.(func())()
	} else {
		NewTitleFromPointer(ptr).DestroyTitleDefault()
	}
}

func (ptr *Title) ConnectDestroyTitle(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Title"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Title", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Title", f)
		}
	}
}

func (ptr *Title) DisconnectDestroyTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Title")
	}
}

func (ptr *Title) DestroyTitle() {
	if ptr.Pointer() != nil {
		C.Title_DestroyTitle(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Title) DestroyTitleDefault() {
	if ptr.Pointer() != nil {
		C.Title_DestroyTitleDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTitle_Event
func callbackTitle_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Title) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Title_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackTitle_EventFilter
func callbackTitle_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Title) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Title_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackTitle_ChildEvent
func callbackTitle_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTitleFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Title) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Title_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTitle_ConnectNotify
func callbackTitle_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTitleFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Title) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Title_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTitle_CustomEvent
func callbackTitle_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewTitleFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Title) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Title_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTitle_DeleteLater
func callbackTitle_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewTitleFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Title) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Title_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTitle_Destroyed
func callbackTitle_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackTitle_DisconnectNotify
func callbackTitle_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTitleFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Title) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Title_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTitle_ObjectNameChanged
func callbackTitle_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackTitle_TimerEvent
func callbackTitle_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTitleFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Title) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Title_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type TitleModel_ITF interface {
	std_core.QAbstractListModel_ITF
	TitleModel_PTR() *TitleModel
}

func (ptr *TitleModel) TitleModel_PTR() *TitleModel {
	return ptr
}

func (ptr *TitleModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *TitleModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromTitleModel(ptr TitleModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.TitleModel_PTR().Pointer()
	}
	return nil
}

func NewTitleModelFromPointer(ptr unsafe.Pointer) *TitleModel {
	var n *TitleModel
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(TitleModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *TitleModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &TitleModel{QAbstractListModel: *deduced}

		default:
			n = new(TitleModel)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackTitleModel_Constructor
func callbackTitleModel_Constructor(ptr unsafe.Pointer) {
	gPtr := NewTitleModelFromPointer(ptr)
	qt.Register(ptr, gPtr)
	gPtr.init()
}

//export callbackTitleModel_AddTitle
func callbackTitleModel_AddTitle(ptr unsafe.Pointer, v0 unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addTitle"); signal != nil {
		signal.(func(*Title))(NewTitleFromPointer(v0))
	}

}

func (ptr *TitleModel) ConnectAddTitle(f func(v0 *Title)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addTitle"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addTitle", func(v0 *Title) {
				signal.(func(*Title))(v0)
				f(v0)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addTitle", f)
		}
	}
}

func (ptr *TitleModel) DisconnectAddTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addTitle")
	}
}

func (ptr *TitleModel) AddTitle(v0 Title_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel_AddTitle(ptr.Pointer(), PointerFromTitle(v0))
	}
}

//export callbackTitleModel_RemoveTitle
func callbackTitleModel_RemoveTitle(ptr unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "removeTitle"); signal != nil {
		signal.(func(int))(int(int32(row)))
	}

}

func (ptr *TitleModel) ConnectRemoveTitle(f func(row int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeTitle"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "removeTitle", func(row int) {
				signal.(func(int))(row)
				f(row)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeTitle", f)
		}
	}
}

func (ptr *TitleModel) DisconnectRemoveTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeTitle")
	}
}

func (ptr *TitleModel) RemoveTitle(row int) {
	if ptr.Pointer() != nil {
		C.TitleModel_RemoveTitle(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackTitleModel_Roles
func callbackTitleModel_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__roles_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__roles_newList())
		for k, v := range NewTitleModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TitleModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "roles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "roles", func() map[int]*std_core.QByteArray {
				signal.(func() map[int]*std_core.QByteArray)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "roles", f)
		}
	}
}

func (ptr *TitleModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *TitleModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewTitleModelFromPointer(l.data).__roles_keyList() {
				out[i] = NewTitleModelFromPointer(l.data).__roles_atList(i)
			}
			return out
		}(C.TitleModel_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *TitleModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewTitleModelFromPointer(l.data).__roles_keyList() {
				out[i] = NewTitleModelFromPointer(l.data).__roles_atList(i)
			}
			return out
		}(C.TitleModel_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackTitleModel_SetRoles
func callbackTitleModel_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewTitleModelFromPointer(l.data).__setRoles_keyList() {
				out[i] = NewTitleModelFromPointer(l.data).__setRoles_roles_atList(i)
			}
			return out
		}(roles))
	} else {
		NewTitleModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewTitleModelFromPointer(l.data).__setRoles_keyList() {
				out[i] = NewTitleModelFromPointer(l.data).__setRoles_roles_atList(i)
			}
			return out
		}(roles))
	}
}

func (ptr *TitleModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRoles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", f)
		}
	}
}

func (ptr *TitleModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *TitleModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.TitleModel_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *TitleModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.TitleModel_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackTitleModel_RolesChanged
func callbackTitleModel_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewTitleModelFromPointer(l.data).__rolesChanged_keyList() {
				out[i] = NewTitleModelFromPointer(l.data).__rolesChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

func (ptr *TitleModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.TitleModel_ConnectRolesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rolesChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", f)
		}
	}
}

func (ptr *TitleModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.TitleModel_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *TitleModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.TitleModel_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackTitleModel_Title
func callbackTitleModel_Title(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "title"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__title_newList())
			for _, v := range signal.(func() []*Title)() {
				tmpList.__title_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__title_newList())
		for _, v := range NewTitleModelFromPointer(ptr).TitleDefault() {
			tmpList.__title_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TitleModel) ConnectTitle(f func() []*Title) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "title"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "title", func() []*Title {
				signal.(func() []*Title)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "title", f)
		}
	}
}

func (ptr *TitleModel) DisconnectTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "title")
	}
}

func (ptr *TitleModel) Title() []*Title {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Title {
			var out = make([]*Title, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).__title_atList(i)
			}
			return out
		}(C.TitleModel_Title(ptr.Pointer()))
	}
	return make([]*Title, 0)
}

func (ptr *TitleModel) TitleDefault() []*Title {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Title {
			var out = make([]*Title, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).__title_atList(i)
			}
			return out
		}(C.TitleModel_TitleDefault(ptr.Pointer()))
	}
	return make([]*Title, 0)
}

//export callbackTitleModel_SetTitle
func callbackTitleModel_SetTitle(ptr unsafe.Pointer, title C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setTitle"); signal != nil {
		signal.(func([]*Title))(func(l C.struct_Moc_PackedList) []*Title {
			var out = make([]*Title, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).__setTitle_title_atList(i)
			}
			return out
		}(title))
	} else {
		NewTitleModelFromPointer(ptr).SetTitleDefault(func(l C.struct_Moc_PackedList) []*Title {
			var out = make([]*Title, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).__setTitle_title_atList(i)
			}
			return out
		}(title))
	}
}

func (ptr *TitleModel) ConnectSetTitle(f func(title []*Title)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTitle"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTitle", func(title []*Title) {
				signal.(func([]*Title))(title)
				f(title)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTitle", f)
		}
	}
}

func (ptr *TitleModel) DisconnectSetTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTitle")
	}
}

func (ptr *TitleModel) SetTitle(title []*Title) {
	if ptr.Pointer() != nil {
		C.TitleModel_SetTitle(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__setTitle_title_newList())
			for _, v := range title {
				tmpList.__setTitle_title_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *TitleModel) SetTitleDefault(title []*Title) {
	if ptr.Pointer() != nil {
		C.TitleModel_SetTitleDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__setTitle_title_newList())
			for _, v := range title {
				tmpList.__setTitle_title_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackTitleModel_TitleChanged
func callbackTitleModel_TitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "titleChanged"); signal != nil {
		signal.(func([]*Title))(func(l C.struct_Moc_PackedList) []*Title {
			var out = make([]*Title, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).__titleChanged_title_atList(i)
			}
			return out
		}(title))
	}

}

func (ptr *TitleModel) ConnectTitleChanged(f func(title []*Title)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleChanged") {
			C.TitleModel_ConnectTitleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", func(title []*Title) {
				signal.(func([]*Title))(title)
				f(title)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", f)
		}
	}
}

func (ptr *TitleModel) DisconnectTitleChanged() {
	if ptr.Pointer() != nil {
		C.TitleModel_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleChanged")
	}
}

func (ptr *TitleModel) TitleChanged(title []*Title) {
	if ptr.Pointer() != nil {
		C.TitleModel_TitleChanged(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__titleChanged_title_newList())
			for _, v := range title {
				tmpList.__titleChanged_title_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func TitleModel_QRegisterMetaType() int {
	return int(int32(C.TitleModel_TitleModel_QRegisterMetaType()))
}

func (ptr *TitleModel) QRegisterMetaType() int {
	return int(int32(C.TitleModel_TitleModel_QRegisterMetaType()))
}

func TitleModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TitleModel_TitleModel_QRegisterMetaType2(typeNameC)))
}

func (ptr *TitleModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TitleModel_TitleModel_QRegisterMetaType2(typeNameC)))
}

func TitleModel_QmlRegisterType() int {
	return int(int32(C.TitleModel_TitleModel_QmlRegisterType()))
}

func (ptr *TitleModel) QmlRegisterType() int {
	return int(int32(C.TitleModel_TitleModel_QmlRegisterType()))
}

func TitleModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TitleModel_TitleModel_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TitleModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TitleModel_TitleModel_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TitleModel) ____setItemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TitleModel_____setItemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TitleModel) ____setItemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TitleModel_____setItemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TitleModel) ____setItemData_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel_____setItemData_keyList_newList(ptr.Pointer()))
}

func (ptr *TitleModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TitleModel_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TitleModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TitleModel_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TitleModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel_____roleNames_keyList_newList(ptr.Pointer()))
}

func (ptr *TitleModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TitleModel_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TitleModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TitleModel_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TitleModel) ____itemData_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel_____itemData_keyList_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __setItemData_roles_atList(i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.TitleModel___setItemData_roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *TitleModel) __setItemData_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___setItemData_roles_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __setItemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).____setItemData_keyList_atList(i)
			}
			return out
		}(C.TitleModel___setItemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TitleModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.TitleModel___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TitleModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___changePersistentIndexList_from_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.TitleModel___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TitleModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___changePersistentIndexList_to_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TitleModel___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TitleModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.TitleModel___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TitleModel) __dataChanged_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___dataChanged_roles_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQPersistentModelIndexFromPointer(C.TitleModel___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *TitleModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___layoutAboutToBeChanged_parents_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQPersistentModelIndexFromPointer(C.TitleModel___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *TitleModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___layoutChanged_parents_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __roleNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.TitleModel___roleNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TitleModel) __roleNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___roleNames_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).____roleNames_keyList_atList(i)
			}
			return out
		}(C.TitleModel___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TitleModel) __itemData_atList(i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.TitleModel___itemData_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *TitleModel) __itemData_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___itemData_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).____itemData_keyList_atList(i)
			}
			return out
		}(C.TitleModel___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TitleModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.TitleModel___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TitleModel) __mimeData_indexes_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___mimeData_indexes_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.TitleModel___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TitleModel) __match_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___match_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.TitleModel___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TitleModel) __persistentIndexList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___persistentIndexList_newList(ptr.Pointer()))
}

func (ptr *TitleModel) ____doSetRoleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TitleModel_____doSetRoleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TitleModel) ____doSetRoleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TitleModel_____doSetRoleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TitleModel) ____doSetRoleNames_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel_____doSetRoleNames_keyList_newList(ptr.Pointer()))
}

func (ptr *TitleModel) ____setRoleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TitleModel_____setRoleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TitleModel) ____setRoleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TitleModel_____setRoleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TitleModel) ____setRoleNames_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel_____setRoleNames_keyList_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.TitleModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TitleModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.TitleModel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TitleModel) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___findChildren_newList2(ptr.Pointer()))
}

func (ptr *TitleModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.TitleModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TitleModel) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___findChildren_newList3(ptr.Pointer()))
}

func (ptr *TitleModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.TitleModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TitleModel) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___findChildren_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.TitleModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TitleModel) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___children_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __roles_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.TitleModel___roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TitleModel) __roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___roles_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).____roles_keyList_atList(i)
			}
			return out
		}(C.TitleModel___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TitleModel) __setRoles_roles_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.TitleModel___setRoles_roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TitleModel) __setRoles_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___setRoles_roles_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __setRoles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).____setRoles_keyList_atList(i)
			}
			return out
		}(C.TitleModel___setRoles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TitleModel) __rolesChanged_roles_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.TitleModel___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TitleModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___rolesChanged_roles_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __rolesChanged_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).____rolesChanged_keyList_atList(i)
			}
			return out
		}(C.TitleModel___rolesChanged_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TitleModel) __title_atList(i int) *Title {
	if ptr.Pointer() != nil {
		var tmpValue = NewTitleFromPointer(C.TitleModel___title_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __title_setList(i Title_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___title_setList(ptr.Pointer(), PointerFromTitle(i))
	}
}

func (ptr *TitleModel) __title_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___title_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __setTitle_title_atList(i int) *Title {
	if ptr.Pointer() != nil {
		var tmpValue = NewTitleFromPointer(C.TitleModel___setTitle_title_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __setTitle_title_setList(i Title_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___setTitle_title_setList(ptr.Pointer(), PointerFromTitle(i))
	}
}

func (ptr *TitleModel) __setTitle_title_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___setTitle_title_newList(ptr.Pointer()))
}

func (ptr *TitleModel) __titleChanged_title_atList(i int) *Title {
	if ptr.Pointer() != nil {
		var tmpValue = NewTitleFromPointer(C.TitleModel___titleChanged_title_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TitleModel) __titleChanged_title_setList(i Title_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel___titleChanged_title_setList(ptr.Pointer(), PointerFromTitle(i))
	}
}

func (ptr *TitleModel) __titleChanged_title_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel___titleChanged_title_newList(ptr.Pointer()))
}

func (ptr *TitleModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TitleModel_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TitleModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TitleModel_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TitleModel) ____roles_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel_____roles_keyList_newList(ptr.Pointer()))
}

func (ptr *TitleModel) ____setRoles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TitleModel_____setRoles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TitleModel) ____setRoles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TitleModel_____setRoles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TitleModel) ____setRoles_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel_____setRoles_keyList_newList(ptr.Pointer()))
}

func (ptr *TitleModel) ____rolesChanged_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TitleModel_____rolesChanged_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TitleModel) ____rolesChanged_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TitleModel_____rolesChanged_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TitleModel) ____rolesChanged_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.TitleModel_____rolesChanged_keyList_newList(ptr.Pointer()))
}

func NewTitleModel(parent std_core.QObject_ITF) *TitleModel {
	var tmpValue = NewTitleModelFromPointer(C.TitleModel_NewTitleModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *TitleModel) DestroyTitleModel() {
	if ptr.Pointer() != nil {
		C.TitleModel_DestroyTitleModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTitleModel_DropMimeData
func callbackTitleModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TitleModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackTitleModel_Index
func callbackTitleModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewTitleModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *TitleModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.TitleModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTitleModel_Sibling
func callbackTitleModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewTitleModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *TitleModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.TitleModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTitleModel_Flags
func callbackTitleModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewTitleModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TitleModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.TitleModel_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackTitleModel_InsertColumns
func callbackTitleModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TitleModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackTitleModel_InsertRows
func callbackTitleModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TitleModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackTitleModel_MoveColumns
func callbackTitleModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *TitleModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackTitleModel_MoveRows
func callbackTitleModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *TitleModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackTitleModel_RemoveColumns
func callbackTitleModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TitleModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackTitleModel_RemoveRows
func callbackTitleModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TitleModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackTitleModel_SetData
func callbackTitleModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *TitleModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackTitleModel_SetHeaderData
func callbackTitleModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *TitleModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackTitleModel_SetItemData
func callbackTitleModel_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			var out = make(map[int]*std_core.QVariant, int(l.len))
			for _, i := range NewTitleModelFromPointer(l.data).__setItemData_keyList() {
				out[i] = NewTitleModelFromPointer(l.data).__setItemData_roles_atList(i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		var out = make(map[int]*std_core.QVariant, int(l.len))
		for _, i := range NewTitleModelFromPointer(l.data).__setItemData_keyList() {
			out[i] = NewTitleModelFromPointer(l.data).__setItemData_roles_atList(i)
		}
		return out
	}(roles)))))
}

func (ptr *TitleModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

//export callbackTitleModel_Submit
func callbackTitleModel_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *TitleModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_SubmitDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackTitleModel_ColumnsAboutToBeInserted
func callbackTitleModel_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTitleModel_ColumnsAboutToBeMoved
func callbackTitleModel_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackTitleModel_ColumnsAboutToBeRemoved
func callbackTitleModel_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTitleModel_ColumnsInserted
func callbackTitleModel_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTitleModel_ColumnsMoved
func callbackTitleModel_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackTitleModel_ColumnsRemoved
func callbackTitleModel_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTitleModel_DataChanged
func callbackTitleModel_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackTitleModel_FetchMore
func callbackTitleModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewTitleModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *TitleModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackTitleModel_HeaderDataChanged
func callbackTitleModel_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackTitleModel_LayoutAboutToBeChanged
func callbackTitleModel_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			var out = make([]*std_core.QPersistentModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackTitleModel_LayoutChanged
func callbackTitleModel_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			var out = make([]*std_core.QPersistentModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackTitleModel_ModelAboutToBeReset
func callbackTitleModel_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackTitleModel_ModelReset
func callbackTitleModel_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackTitleModel_ResetInternalData
func callbackTitleModel_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewTitleModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *TitleModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.TitleModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackTitleModel_Revert
func callbackTitleModel_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewTitleModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *TitleModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.TitleModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackTitleModel_RowsAboutToBeInserted
func callbackTitleModel_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackTitleModel_RowsAboutToBeMoved
func callbackTitleModel_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackTitleModel_RowsAboutToBeRemoved
func callbackTitleModel_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTitleModel_RowsInserted
func callbackTitleModel_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTitleModel_RowsMoved
func callbackTitleModel_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackTitleModel_RowsRemoved
func callbackTitleModel_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTitleModel_Sort
func callbackTitleModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewTitleModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *TitleModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.TitleModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackTitleModel_RoleNames
func callbackTitleModel_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewTitleModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TitleModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewTitleModelFromPointer(l.data).__roleNames_keyList() {
				out[i] = NewTitleModelFromPointer(l.data).__roleNames_atList(i)
			}
			return out
		}(C.TitleModel_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackTitleModel_ItemData
func callbackTitleModel_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__itemData_newList())
		for k, v := range NewTitleModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TitleModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			var out = make(map[int]*std_core.QVariant, int(l.len))
			for _, i := range NewTitleModelFromPointer(l.data).__itemData_keyList() {
				out[i] = NewTitleModelFromPointer(l.data).__itemData_atList(i)
			}
			return out
		}(C.TitleModel_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackTitleModel_MimeData
func callbackTitleModel_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			var out = make([]*std_core.QModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewTitleModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		var out = make([]*std_core.QModelIndex, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewTitleModelFromPointer(l.data).__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *TitleModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQMimeDataFromPointer(C.TitleModel_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackTitleModel_Buddy
func callbackTitleModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewTitleModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TitleModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.TitleModel_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTitleModel_Parent
func callbackTitleModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewTitleModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TitleModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.TitleModel_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTitleModel_Match
func callbackTitleModel_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewTitleModelFromPointer(NewTitleModelFromPointer(nil).__match_newList())
		for _, v := range NewTitleModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TitleModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			var out = make([]*std_core.QModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewTitleModelFromPointer(l.data).__match_atList(i)
			}
			return out
		}(C.TitleModel_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackTitleModel_Span
func callbackTitleModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewTitleModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TitleModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.TitleModel_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackTitleModel_MimeTypes
func callbackTitleModel_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewTitleModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *TitleModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.TitleModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackTitleModel_Data
func callbackTitleModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewTitleModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *TitleModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.TitleModel_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackTitleModel_HeaderData
func callbackTitleModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewTitleModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *TitleModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.TitleModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackTitleModel_SupportedDragActions
func callbackTitleModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewTitleModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *TitleModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.TitleModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackTitleModel_SupportedDropActions
func callbackTitleModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewTitleModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *TitleModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.TitleModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackTitleModel_CanDropMimeData
func callbackTitleModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TitleModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackTitleModel_CanFetchMore
func callbackTitleModel_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TitleModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackTitleModel_HasChildren
func callbackTitleModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TitleModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackTitleModel_ColumnCount
func callbackTitleModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewTitleModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *TitleModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TitleModel_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackTitleModel_RowCount
func callbackTitleModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewTitleModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *TitleModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TitleModel_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackTitleModel_Event
func callbackTitleModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TitleModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackTitleModel_EventFilter
func callbackTitleModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTitleModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TitleModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.TitleModel_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackTitleModel_ChildEvent
func callbackTitleModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTitleModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TitleModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTitleModel_ConnectNotify
func callbackTitleModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTitleModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TitleModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTitleModel_CustomEvent
func callbackTitleModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewTitleModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TitleModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTitleModel_DeleteLater
func callbackTitleModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewTitleModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TitleModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TitleModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTitleModel_Destroyed
func callbackTitleModel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackTitleModel_DisconnectNotify
func callbackTitleModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTitleModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TitleModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTitleModel_ObjectNameChanged
func callbackTitleModel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackTitleModel_TimerEvent
func callbackTitleModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTitleModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TitleModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TitleModel_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type DictTab_ITF interface {
	std_core.QObject_ITF
	DictTab_PTR() *DictTab
}

func (ptr *DictTab) DictTab_PTR() *DictTab {
	return ptr
}

func (ptr *DictTab) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *DictTab) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromDictTab(ptr DictTab_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.DictTab_PTR().Pointer()
	}
	return nil
}

func NewDictTabFromPointer(ptr unsafe.Pointer) *DictTab {
	var n *DictTab
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(DictTab)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *DictTab:
			n = deduced

		case *std_core.QObject:
			n = &DictTab{QObject: *deduced}

		default:
			n = new(DictTab)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackDictTab_Constructor
func callbackDictTab_Constructor(ptr unsafe.Pointer) {
	gPtr := NewDictTabFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackDictTab_DictName
func callbackDictTab_DictName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "dictName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewDictTabFromPointer(ptr).DictNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *DictTab) ConnectDictName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "dictName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dictName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dictName", f)
		}
	}
}

func (ptr *DictTab) DisconnectDictName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "dictName")
	}
}

func (ptr *DictTab) DictName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.DictTab_DictName(ptr.Pointer()))
	}
	return ""
}

func (ptr *DictTab) DictNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.DictTab_DictNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackDictTab_SetDictName
func callbackDictTab_SetDictName(ptr unsafe.Pointer, dictName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setDictName"); signal != nil {
		signal.(func(string))(cGoUnpackString(dictName))
	} else {
		NewDictTabFromPointer(ptr).SetDictNameDefault(cGoUnpackString(dictName))
	}
}

func (ptr *DictTab) ConnectSetDictName(f func(dictName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDictName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDictName", func(dictName string) {
				signal.(func(string))(dictName)
				f(dictName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDictName", f)
		}
	}
}

func (ptr *DictTab) DisconnectSetDictName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDictName")
	}
}

func (ptr *DictTab) SetDictName(dictName string) {
	if ptr.Pointer() != nil {
		var dictNameC *C.char
		if dictName != "" {
			dictNameC = C.CString(dictName)
			defer C.free(unsafe.Pointer(dictNameC))
		}
		C.DictTab_SetDictName(ptr.Pointer(), C.struct_Moc_PackedString{data: dictNameC, len: C.longlong(len(dictName))})
	}
}

func (ptr *DictTab) SetDictNameDefault(dictName string) {
	if ptr.Pointer() != nil {
		var dictNameC *C.char
		if dictName != "" {
			dictNameC = C.CString(dictName)
			defer C.free(unsafe.Pointer(dictNameC))
		}
		C.DictTab_SetDictNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: dictNameC, len: C.longlong(len(dictName))})
	}
}

//export callbackDictTab_DictNameChanged
func callbackDictTab_DictNameChanged(ptr unsafe.Pointer, dictName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "dictNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(dictName))
	}

}

func (ptr *DictTab) ConnectDictNameChanged(f func(dictName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dictNameChanged") {
			C.DictTab_ConnectDictNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dictNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dictNameChanged", func(dictName string) {
				signal.(func(string))(dictName)
				f(dictName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dictNameChanged", f)
		}
	}
}

func (ptr *DictTab) DisconnectDictNameChanged() {
	if ptr.Pointer() != nil {
		C.DictTab_DisconnectDictNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dictNameChanged")
	}
}

func (ptr *DictTab) DictNameChanged(dictName string) {
	if ptr.Pointer() != nil {
		var dictNameC *C.char
		if dictName != "" {
			dictNameC = C.CString(dictName)
			defer C.free(unsafe.Pointer(dictNameC))
		}
		C.DictTab_DictNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: dictNameC, len: C.longlong(len(dictName))})
	}
}

func DictTab_QRegisterMetaType() int {
	return int(int32(C.DictTab_DictTab_QRegisterMetaType()))
}

func (ptr *DictTab) QRegisterMetaType() int {
	return int(int32(C.DictTab_DictTab_QRegisterMetaType()))
}

func DictTab_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.DictTab_DictTab_QRegisterMetaType2(typeNameC)))
}

func (ptr *DictTab) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.DictTab_DictTab_QRegisterMetaType2(typeNameC)))
}

func DictTab_QmlRegisterType() int {
	return int(int32(C.DictTab_DictTab_QmlRegisterType()))
}

func (ptr *DictTab) QmlRegisterType() int {
	return int(int32(C.DictTab_DictTab_QmlRegisterType()))
}

func DictTab_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.DictTab_DictTab_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *DictTab) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.DictTab_DictTab_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *DictTab) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.DictTab___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *DictTab) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.DictTab___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *DictTab) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTab___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *DictTab) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.DictTab___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DictTab) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DictTab___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DictTab) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.DictTab___findChildren_newList2(ptr.Pointer()))
}

func (ptr *DictTab) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.DictTab___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DictTab) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DictTab___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DictTab) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.DictTab___findChildren_newList3(ptr.Pointer()))
}

func (ptr *DictTab) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.DictTab___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DictTab) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DictTab___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DictTab) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTab___findChildren_newList(ptr.Pointer()))
}

func (ptr *DictTab) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.DictTab___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DictTab) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DictTab___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DictTab) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTab___children_newList(ptr.Pointer()))
}

func NewDictTab(parent std_core.QObject_ITF) *DictTab {
	var tmpValue = NewDictTabFromPointer(C.DictTab_NewDictTab(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackDictTab_DestroyDictTab
func callbackDictTab_DestroyDictTab(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~DictTab"); signal != nil {
		signal.(func())()
	} else {
		NewDictTabFromPointer(ptr).DestroyDictTabDefault()
	}
}

func (ptr *DictTab) ConnectDestroyDictTab(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~DictTab"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~DictTab", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~DictTab", f)
		}
	}
}

func (ptr *DictTab) DisconnectDestroyDictTab() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~DictTab")
	}
}

func (ptr *DictTab) DestroyDictTab() {
	if ptr.Pointer() != nil {
		C.DictTab_DestroyDictTab(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *DictTab) DestroyDictTabDefault() {
	if ptr.Pointer() != nil {
		C.DictTab_DestroyDictTabDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackDictTab_Event
func callbackDictTab_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *DictTab) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTab_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackDictTab_EventFilter
func callbackDictTab_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *DictTab) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTab_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackDictTab_ChildEvent
func callbackDictTab_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewDictTabFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *DictTab) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DictTab_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackDictTab_ConnectNotify
func callbackDictTab_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDictTabFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *DictTab) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.DictTab_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDictTab_CustomEvent
func callbackDictTab_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewDictTabFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *DictTab) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DictTab_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackDictTab_DeleteLater
func callbackDictTab_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewDictTabFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *DictTab) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.DictTab_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackDictTab_Destroyed
func callbackDictTab_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackDictTab_DisconnectNotify
func callbackDictTab_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDictTabFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *DictTab) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.DictTab_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDictTab_ObjectNameChanged
func callbackDictTab_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackDictTab_TimerEvent
func callbackDictTab_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewDictTabFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *DictTab) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DictTab_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type DictTabModel_ITF interface {
	std_core.QAbstractListModel_ITF
	DictTabModel_PTR() *DictTabModel
}

func (ptr *DictTabModel) DictTabModel_PTR() *DictTabModel {
	return ptr
}

func (ptr *DictTabModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *DictTabModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromDictTabModel(ptr DictTabModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.DictTabModel_PTR().Pointer()
	}
	return nil
}

func NewDictTabModelFromPointer(ptr unsafe.Pointer) *DictTabModel {
	var n *DictTabModel
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(DictTabModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *DictTabModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &DictTabModel{QAbstractListModel: *deduced}

		default:
			n = new(DictTabModel)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackDictTabModel_Constructor
func callbackDictTabModel_Constructor(ptr unsafe.Pointer) {
	gPtr := NewDictTabModelFromPointer(ptr)
	qt.Register(ptr, gPtr)
	gPtr.init()
}

//export callbackDictTabModel_AddDictTab
func callbackDictTabModel_AddDictTab(ptr unsafe.Pointer, v0 unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addDictTab"); signal != nil {
		signal.(func(*DictTab))(NewDictTabFromPointer(v0))
	}

}

func (ptr *DictTabModel) ConnectAddDictTab(f func(v0 *DictTab)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addDictTab"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addDictTab", func(v0 *DictTab) {
				signal.(func(*DictTab))(v0)
				f(v0)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addDictTab", f)
		}
	}
}

func (ptr *DictTabModel) DisconnectAddDictTab() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addDictTab")
	}
}

func (ptr *DictTabModel) AddDictTab(v0 DictTab_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel_AddDictTab(ptr.Pointer(), PointerFromDictTab(v0))
	}
}

//export callbackDictTabModel_RemoveDictTab
func callbackDictTabModel_RemoveDictTab(ptr unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "removeDictTab"); signal != nil {
		signal.(func(int))(int(int32(row)))
	}

}

func (ptr *DictTabModel) ConnectRemoveDictTab(f func(row int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeDictTab"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "removeDictTab", func(row int) {
				signal.(func(int))(row)
				f(row)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeDictTab", f)
		}
	}
}

func (ptr *DictTabModel) DisconnectRemoveDictTab() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeDictTab")
	}
}

func (ptr *DictTabModel) RemoveDictTab(row int) {
	if ptr.Pointer() != nil {
		C.DictTabModel_RemoveDictTab(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackDictTabModel_Roles
func callbackDictTabModel_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__roles_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__roles_newList())
		for k, v := range NewDictTabModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *DictTabModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "roles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "roles", func() map[int]*std_core.QByteArray {
				signal.(func() map[int]*std_core.QByteArray)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "roles", f)
		}
	}
}

func (ptr *DictTabModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *DictTabModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewDictTabModelFromPointer(l.data).__roles_keyList() {
				out[i] = NewDictTabModelFromPointer(l.data).__roles_atList(i)
			}
			return out
		}(C.DictTabModel_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *DictTabModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewDictTabModelFromPointer(l.data).__roles_keyList() {
				out[i] = NewDictTabModelFromPointer(l.data).__roles_atList(i)
			}
			return out
		}(C.DictTabModel_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackDictTabModel_SetRoles
func callbackDictTabModel_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewDictTabModelFromPointer(l.data).__setRoles_keyList() {
				out[i] = NewDictTabModelFromPointer(l.data).__setRoles_roles_atList(i)
			}
			return out
		}(roles))
	} else {
		NewDictTabModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewDictTabModelFromPointer(l.data).__setRoles_keyList() {
				out[i] = NewDictTabModelFromPointer(l.data).__setRoles_roles_atList(i)
			}
			return out
		}(roles))
	}
}

func (ptr *DictTabModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRoles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", f)
		}
	}
}

func (ptr *DictTabModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *DictTabModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.DictTabModel_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *DictTabModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.DictTabModel_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackDictTabModel_RolesChanged
func callbackDictTabModel_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewDictTabModelFromPointer(l.data).__rolesChanged_keyList() {
				out[i] = NewDictTabModelFromPointer(l.data).__rolesChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

func (ptr *DictTabModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.DictTabModel_ConnectRolesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rolesChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", f)
		}
	}
}

func (ptr *DictTabModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.DictTabModel_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *DictTabModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.DictTabModel_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackDictTabModel_DictTab
func callbackDictTabModel_DictTab(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "dictTab"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__dictTab_newList())
			for _, v := range signal.(func() []*DictTab)() {
				tmpList.__dictTab_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__dictTab_newList())
		for _, v := range NewDictTabModelFromPointer(ptr).DictTabDefault() {
			tmpList.__dictTab_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *DictTabModel) ConnectDictTab(f func() []*DictTab) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "dictTab"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dictTab", func() []*DictTab {
				signal.(func() []*DictTab)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dictTab", f)
		}
	}
}

func (ptr *DictTabModel) DisconnectDictTab() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "dictTab")
	}
}

func (ptr *DictTabModel) DictTab() []*DictTab {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*DictTab {
			var out = make([]*DictTab, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).__dictTab_atList(i)
			}
			return out
		}(C.DictTabModel_DictTab(ptr.Pointer()))
	}
	return make([]*DictTab, 0)
}

func (ptr *DictTabModel) DictTabDefault() []*DictTab {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*DictTab {
			var out = make([]*DictTab, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).__dictTab_atList(i)
			}
			return out
		}(C.DictTabModel_DictTabDefault(ptr.Pointer()))
	}
	return make([]*DictTab, 0)
}

//export callbackDictTabModel_SetDictTab
func callbackDictTabModel_SetDictTab(ptr unsafe.Pointer, dictTab C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setDictTab"); signal != nil {
		signal.(func([]*DictTab))(func(l C.struct_Moc_PackedList) []*DictTab {
			var out = make([]*DictTab, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).__setDictTab_dictTab_atList(i)
			}
			return out
		}(dictTab))
	} else {
		NewDictTabModelFromPointer(ptr).SetDictTabDefault(func(l C.struct_Moc_PackedList) []*DictTab {
			var out = make([]*DictTab, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).__setDictTab_dictTab_atList(i)
			}
			return out
		}(dictTab))
	}
}

func (ptr *DictTabModel) ConnectSetDictTab(f func(dictTab []*DictTab)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDictTab"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDictTab", func(dictTab []*DictTab) {
				signal.(func([]*DictTab))(dictTab)
				f(dictTab)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDictTab", f)
		}
	}
}

func (ptr *DictTabModel) DisconnectSetDictTab() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDictTab")
	}
}

func (ptr *DictTabModel) SetDictTab(dictTab []*DictTab) {
	if ptr.Pointer() != nil {
		C.DictTabModel_SetDictTab(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__setDictTab_dictTab_newList())
			for _, v := range dictTab {
				tmpList.__setDictTab_dictTab_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *DictTabModel) SetDictTabDefault(dictTab []*DictTab) {
	if ptr.Pointer() != nil {
		C.DictTabModel_SetDictTabDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__setDictTab_dictTab_newList())
			for _, v := range dictTab {
				tmpList.__setDictTab_dictTab_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackDictTabModel_DictTabChanged
func callbackDictTabModel_DictTabChanged(ptr unsafe.Pointer, dictTab C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dictTabChanged"); signal != nil {
		signal.(func([]*DictTab))(func(l C.struct_Moc_PackedList) []*DictTab {
			var out = make([]*DictTab, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).__dictTabChanged_dictTab_atList(i)
			}
			return out
		}(dictTab))
	}

}

func (ptr *DictTabModel) ConnectDictTabChanged(f func(dictTab []*DictTab)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dictTabChanged") {
			C.DictTabModel_ConnectDictTabChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dictTabChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dictTabChanged", func(dictTab []*DictTab) {
				signal.(func([]*DictTab))(dictTab)
				f(dictTab)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dictTabChanged", f)
		}
	}
}

func (ptr *DictTabModel) DisconnectDictTabChanged() {
	if ptr.Pointer() != nil {
		C.DictTabModel_DisconnectDictTabChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dictTabChanged")
	}
}

func (ptr *DictTabModel) DictTabChanged(dictTab []*DictTab) {
	if ptr.Pointer() != nil {
		C.DictTabModel_DictTabChanged(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__dictTabChanged_dictTab_newList())
			for _, v := range dictTab {
				tmpList.__dictTabChanged_dictTab_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func DictTabModel_QRegisterMetaType() int {
	return int(int32(C.DictTabModel_DictTabModel_QRegisterMetaType()))
}

func (ptr *DictTabModel) QRegisterMetaType() int {
	return int(int32(C.DictTabModel_DictTabModel_QRegisterMetaType()))
}

func DictTabModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.DictTabModel_DictTabModel_QRegisterMetaType2(typeNameC)))
}

func (ptr *DictTabModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.DictTabModel_DictTabModel_QRegisterMetaType2(typeNameC)))
}

func DictTabModel_QmlRegisterType() int {
	return int(int32(C.DictTabModel_DictTabModel_QmlRegisterType()))
}

func (ptr *DictTabModel) QmlRegisterType() int {
	return int(int32(C.DictTabModel_DictTabModel_QmlRegisterType()))
}

func DictTabModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.DictTabModel_DictTabModel_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *DictTabModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.DictTabModel_DictTabModel_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *DictTabModel) ____setItemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DictTabModel_____setItemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *DictTabModel) ____setItemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.DictTabModel_____setItemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *DictTabModel) ____setItemData_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel_____setItemData_keyList_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DictTabModel_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *DictTabModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.DictTabModel_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *DictTabModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel_____roleNames_keyList_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DictTabModel_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *DictTabModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.DictTabModel_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *DictTabModel) ____itemData_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel_____itemData_keyList_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __setItemData_roles_atList(i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.DictTabModel___setItemData_roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *DictTabModel) __setItemData_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___setItemData_roles_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __setItemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).____setItemData_keyList_atList(i)
			}
			return out
		}(C.DictTabModel___setItemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *DictTabModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.DictTabModel___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *DictTabModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___changePersistentIndexList_from_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.DictTabModel___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *DictTabModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___changePersistentIndexList_to_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DictTabModel___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *DictTabModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.DictTabModel___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *DictTabModel) __dataChanged_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___dataChanged_roles_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQPersistentModelIndexFromPointer(C.DictTabModel___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *DictTabModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___layoutAboutToBeChanged_parents_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQPersistentModelIndexFromPointer(C.DictTabModel___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *DictTabModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___layoutChanged_parents_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __roleNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.DictTabModel___roleNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *DictTabModel) __roleNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___roleNames_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).____roleNames_keyList_atList(i)
			}
			return out
		}(C.DictTabModel___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *DictTabModel) __itemData_atList(i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.DictTabModel___itemData_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *DictTabModel) __itemData_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___itemData_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).____itemData_keyList_atList(i)
			}
			return out
		}(C.DictTabModel___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *DictTabModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.DictTabModel___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *DictTabModel) __mimeData_indexes_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___mimeData_indexes_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.DictTabModel___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *DictTabModel) __match_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___match_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.DictTabModel___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *DictTabModel) __persistentIndexList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___persistentIndexList_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) ____doSetRoleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DictTabModel_____doSetRoleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *DictTabModel) ____doSetRoleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.DictTabModel_____doSetRoleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *DictTabModel) ____doSetRoleNames_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel_____doSetRoleNames_keyList_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) ____setRoleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DictTabModel_____setRoleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *DictTabModel) ____setRoleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.DictTabModel_____setRoleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *DictTabModel) ____setRoleNames_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel_____setRoleNames_keyList_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.DictTabModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *DictTabModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.DictTabModel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DictTabModel) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___findChildren_newList2(ptr.Pointer()))
}

func (ptr *DictTabModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.DictTabModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DictTabModel) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___findChildren_newList3(ptr.Pointer()))
}

func (ptr *DictTabModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.DictTabModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DictTabModel) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___findChildren_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.DictTabModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DictTabModel) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___children_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __roles_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.DictTabModel___roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *DictTabModel) __roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___roles_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).____roles_keyList_atList(i)
			}
			return out
		}(C.DictTabModel___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *DictTabModel) __setRoles_roles_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.DictTabModel___setRoles_roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *DictTabModel) __setRoles_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___setRoles_roles_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __setRoles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).____setRoles_keyList_atList(i)
			}
			return out
		}(C.DictTabModel___setRoles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *DictTabModel) __rolesChanged_roles_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.DictTabModel___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *DictTabModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___rolesChanged_roles_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __rolesChanged_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).____rolesChanged_keyList_atList(i)
			}
			return out
		}(C.DictTabModel___rolesChanged_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *DictTabModel) __dictTab_atList(i int) *DictTab {
	if ptr.Pointer() != nil {
		var tmpValue = NewDictTabFromPointer(C.DictTabModel___dictTab_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __dictTab_setList(i DictTab_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___dictTab_setList(ptr.Pointer(), PointerFromDictTab(i))
	}
}

func (ptr *DictTabModel) __dictTab_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___dictTab_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __setDictTab_dictTab_atList(i int) *DictTab {
	if ptr.Pointer() != nil {
		var tmpValue = NewDictTabFromPointer(C.DictTabModel___setDictTab_dictTab_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __setDictTab_dictTab_setList(i DictTab_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___setDictTab_dictTab_setList(ptr.Pointer(), PointerFromDictTab(i))
	}
}

func (ptr *DictTabModel) __setDictTab_dictTab_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___setDictTab_dictTab_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) __dictTabChanged_dictTab_atList(i int) *DictTab {
	if ptr.Pointer() != nil {
		var tmpValue = NewDictTabFromPointer(C.DictTabModel___dictTabChanged_dictTab_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DictTabModel) __dictTabChanged_dictTab_setList(i DictTab_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel___dictTabChanged_dictTab_setList(ptr.Pointer(), PointerFromDictTab(i))
	}
}

func (ptr *DictTabModel) __dictTabChanged_dictTab_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel___dictTabChanged_dictTab_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DictTabModel_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *DictTabModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.DictTabModel_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *DictTabModel) ____roles_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel_____roles_keyList_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) ____setRoles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DictTabModel_____setRoles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *DictTabModel) ____setRoles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.DictTabModel_____setRoles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *DictTabModel) ____setRoles_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel_____setRoles_keyList_newList(ptr.Pointer()))
}

func (ptr *DictTabModel) ____rolesChanged_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DictTabModel_____rolesChanged_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *DictTabModel) ____rolesChanged_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.DictTabModel_____rolesChanged_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *DictTabModel) ____rolesChanged_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.DictTabModel_____rolesChanged_keyList_newList(ptr.Pointer()))
}

func NewDictTabModel(parent std_core.QObject_ITF) *DictTabModel {
	var tmpValue = NewDictTabModelFromPointer(C.DictTabModel_NewDictTabModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *DictTabModel) DestroyDictTabModel() {
	if ptr.Pointer() != nil {
		C.DictTabModel_DestroyDictTabModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackDictTabModel_DropMimeData
func callbackDictTabModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *DictTabModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackDictTabModel_Index
func callbackDictTabModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewDictTabModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *DictTabModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.DictTabModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackDictTabModel_Sibling
func callbackDictTabModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewDictTabModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *DictTabModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.DictTabModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackDictTabModel_Flags
func callbackDictTabModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewDictTabModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *DictTabModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.DictTabModel_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackDictTabModel_InsertColumns
func callbackDictTabModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *DictTabModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackDictTabModel_InsertRows
func callbackDictTabModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *DictTabModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackDictTabModel_MoveColumns
func callbackDictTabModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *DictTabModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackDictTabModel_MoveRows
func callbackDictTabModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *DictTabModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackDictTabModel_RemoveColumns
func callbackDictTabModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *DictTabModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackDictTabModel_RemoveRows
func callbackDictTabModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *DictTabModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackDictTabModel_SetData
func callbackDictTabModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *DictTabModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackDictTabModel_SetHeaderData
func callbackDictTabModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *DictTabModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackDictTabModel_SetItemData
func callbackDictTabModel_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			var out = make(map[int]*std_core.QVariant, int(l.len))
			for _, i := range NewDictTabModelFromPointer(l.data).__setItemData_keyList() {
				out[i] = NewDictTabModelFromPointer(l.data).__setItemData_roles_atList(i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		var out = make(map[int]*std_core.QVariant, int(l.len))
		for _, i := range NewDictTabModelFromPointer(l.data).__setItemData_keyList() {
			out[i] = NewDictTabModelFromPointer(l.data).__setItemData_roles_atList(i)
		}
		return out
	}(roles)))))
}

func (ptr *DictTabModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

//export callbackDictTabModel_Submit
func callbackDictTabModel_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *DictTabModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_SubmitDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackDictTabModel_ColumnsAboutToBeInserted
func callbackDictTabModel_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackDictTabModel_ColumnsAboutToBeMoved
func callbackDictTabModel_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackDictTabModel_ColumnsAboutToBeRemoved
func callbackDictTabModel_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackDictTabModel_ColumnsInserted
func callbackDictTabModel_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackDictTabModel_ColumnsMoved
func callbackDictTabModel_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackDictTabModel_ColumnsRemoved
func callbackDictTabModel_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackDictTabModel_DataChanged
func callbackDictTabModel_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackDictTabModel_FetchMore
func callbackDictTabModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewDictTabModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *DictTabModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackDictTabModel_HeaderDataChanged
func callbackDictTabModel_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackDictTabModel_LayoutAboutToBeChanged
func callbackDictTabModel_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			var out = make([]*std_core.QPersistentModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackDictTabModel_LayoutChanged
func callbackDictTabModel_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			var out = make([]*std_core.QPersistentModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackDictTabModel_ModelAboutToBeReset
func callbackDictTabModel_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackDictTabModel_ModelReset
func callbackDictTabModel_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackDictTabModel_ResetInternalData
func callbackDictTabModel_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewDictTabModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *DictTabModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.DictTabModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackDictTabModel_Revert
func callbackDictTabModel_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewDictTabModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *DictTabModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.DictTabModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackDictTabModel_RowsAboutToBeInserted
func callbackDictTabModel_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackDictTabModel_RowsAboutToBeMoved
func callbackDictTabModel_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackDictTabModel_RowsAboutToBeRemoved
func callbackDictTabModel_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackDictTabModel_RowsInserted
func callbackDictTabModel_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackDictTabModel_RowsMoved
func callbackDictTabModel_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackDictTabModel_RowsRemoved
func callbackDictTabModel_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackDictTabModel_Sort
func callbackDictTabModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewDictTabModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *DictTabModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.DictTabModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackDictTabModel_RoleNames
func callbackDictTabModel_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewDictTabModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *DictTabModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewDictTabModelFromPointer(l.data).__roleNames_keyList() {
				out[i] = NewDictTabModelFromPointer(l.data).__roleNames_atList(i)
			}
			return out
		}(C.DictTabModel_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackDictTabModel_ItemData
func callbackDictTabModel_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__itemData_newList())
		for k, v := range NewDictTabModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *DictTabModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			var out = make(map[int]*std_core.QVariant, int(l.len))
			for _, i := range NewDictTabModelFromPointer(l.data).__itemData_keyList() {
				out[i] = NewDictTabModelFromPointer(l.data).__itemData_atList(i)
			}
			return out
		}(C.DictTabModel_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackDictTabModel_MimeData
func callbackDictTabModel_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			var out = make([]*std_core.QModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewDictTabModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		var out = make([]*std_core.QModelIndex, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewDictTabModelFromPointer(l.data).__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *DictTabModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQMimeDataFromPointer(C.DictTabModel_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackDictTabModel_Buddy
func callbackDictTabModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewDictTabModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *DictTabModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.DictTabModel_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackDictTabModel_Parent
func callbackDictTabModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewDictTabModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *DictTabModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.DictTabModel_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackDictTabModel_Match
func callbackDictTabModel_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewDictTabModelFromPointer(NewDictTabModelFromPointer(nil).__match_newList())
		for _, v := range NewDictTabModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *DictTabModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			var out = make([]*std_core.QModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewDictTabModelFromPointer(l.data).__match_atList(i)
			}
			return out
		}(C.DictTabModel_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackDictTabModel_Span
func callbackDictTabModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewDictTabModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *DictTabModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.DictTabModel_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackDictTabModel_MimeTypes
func callbackDictTabModel_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewDictTabModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *DictTabModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.DictTabModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackDictTabModel_Data
func callbackDictTabModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewDictTabModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *DictTabModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.DictTabModel_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackDictTabModel_HeaderData
func callbackDictTabModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewDictTabModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *DictTabModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.DictTabModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackDictTabModel_SupportedDragActions
func callbackDictTabModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewDictTabModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *DictTabModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.DictTabModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackDictTabModel_SupportedDropActions
func callbackDictTabModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewDictTabModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *DictTabModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.DictTabModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackDictTabModel_CanDropMimeData
func callbackDictTabModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *DictTabModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackDictTabModel_CanFetchMore
func callbackDictTabModel_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *DictTabModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackDictTabModel_HasChildren
func callbackDictTabModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *DictTabModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackDictTabModel_ColumnCount
func callbackDictTabModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewDictTabModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *DictTabModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DictTabModel_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackDictTabModel_RowCount
func callbackDictTabModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewDictTabModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *DictTabModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DictTabModel_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackDictTabModel_Event
func callbackDictTabModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *DictTabModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackDictTabModel_EventFilter
func callbackDictTabModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDictTabModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *DictTabModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DictTabModel_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackDictTabModel_ChildEvent
func callbackDictTabModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewDictTabModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *DictTabModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackDictTabModel_ConnectNotify
func callbackDictTabModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDictTabModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *DictTabModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDictTabModel_CustomEvent
func callbackDictTabModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewDictTabModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *DictTabModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackDictTabModel_DeleteLater
func callbackDictTabModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewDictTabModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *DictTabModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.DictTabModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackDictTabModel_Destroyed
func callbackDictTabModel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackDictTabModel_DisconnectNotify
func callbackDictTabModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDictTabModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *DictTabModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDictTabModel_ObjectNameChanged
func callbackDictTabModel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackDictTabModel_TimerEvent
func callbackDictTabModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewDictTabModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *DictTabModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DictTabModel_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
