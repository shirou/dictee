

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractListModel>
#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QLayout>
#include <QList>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QSize>
#include <QString>
#include <QStringList>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>


typedef QHash<qint32, QByteArray> typed01680;
typedef QHash<qint32, QByteArray> typed01680;
class Title: public QObject
{
Q_OBJECT
Q_PROPERTY(QString titleName READ titleName WRITE setTitleName NOTIFY titleNameChanged)
Q_PROPERTY(QString dictName READ dictName WRITE setDictName NOTIFY dictNameChanged)
public:
	Title(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Title_Title_QRegisterMetaType();Title_Title_QRegisterMetaTypes();callbackTitle_Constructor(this);};
	QString titleName() { return ({ Moc_PackedString tempVal = callbackTitle_TitleName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTitleName(QString titleName) { QByteArray t5aded7 = titleName.toUtf8(); Moc_PackedString titleNamePacked = { const_cast<char*>(t5aded7.prepend("WHITESPACE").constData()+10), t5aded7.size()-10 };callbackTitle_SetTitleName(this, titleNamePacked); };
	void Signal_TitleNameChanged(QString titleName) { QByteArray t5aded7 = titleName.toUtf8(); Moc_PackedString titleNamePacked = { const_cast<char*>(t5aded7.prepend("WHITESPACE").constData()+10), t5aded7.size()-10 };callbackTitle_TitleNameChanged(this, titleNamePacked); };
	QString dictName() { return ({ Moc_PackedString tempVal = callbackTitle_DictName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setDictName(QString dictName) { QByteArray t06e45a = dictName.toUtf8(); Moc_PackedString dictNamePacked = { const_cast<char*>(t06e45a.prepend("WHITESPACE").constData()+10), t06e45a.size()-10 };callbackTitle_SetDictName(this, dictNamePacked); };
	void Signal_DictNameChanged(QString dictName) { QByteArray t06e45a = dictName.toUtf8(); Moc_PackedString dictNamePacked = { const_cast<char*>(t06e45a.prepend("WHITESPACE").constData()+10), t06e45a.size()-10 };callbackTitle_DictNameChanged(this, dictNamePacked); };
	 ~Title() { callbackTitle_DestroyTitle(this); };
	bool event(QEvent * e) { return callbackTitle_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTitle_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackTitle_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTitle_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTitle_CustomEvent(this, event); };
	void deleteLater() { callbackTitle_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTitle_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTitle_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTitle_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTitle_TimerEvent(this, event); };
	
	QString titleNameDefault() { return _titleName; };
	void setTitleNameDefault(QString p) { if (p != _titleName) { _titleName = p; titleNameChanged(_titleName); } };
	QString dictNameDefault() { return _dictName; };
	void setDictNameDefault(QString p) { if (p != _dictName) { _dictName = p; dictNameChanged(_dictName); } };
signals:
	void titleNameChanged(QString titleName);
	void dictNameChanged(QString dictName);
public slots:
private:
	QString _titleName;
	QString _dictName;
};

Q_DECLARE_METATYPE(Title*)


void Title_Title_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class TitleModel: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(typed01680 roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<Title*> title READ title WRITE setTitle NOTIFY titleChanged)
public:
	TitleModel(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");TitleModel_TitleModel_QRegisterMetaType();TitleModel_TitleModel_QRegisterMetaTypes();callbackTitleModel_Constructor(this);};
	typed01680 roles() { return *static_cast<QHash<qint32, QByteArray>*>(callbackTitleModel_Roles(this)); };
	void setRoles(typed01680 roles) { callbackTitleModel_SetRoles(this, ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(typed01680 roles) { callbackTitleModel_RolesChanged(this, ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<Title*> title() { return *static_cast<QList<Title*>*>(callbackTitleModel_Title(this)); };
	void setTitle(QList<Title*> title) { callbackTitleModel_SetTitle(this, ({ QList<Title*>* tmpValue = new QList<Title*>(title); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_TitleChanged(QList<Title*> title) { callbackTitleModel_TitleChanged(this, ({ QList<Title*>* tmpValue = new QList<Title*>(title); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackTitleModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackTitleModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackTitleModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackTitleModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackTitleModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackTitleModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackTitleModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackTitleModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackTitleModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackTitleModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackTitleModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackTitleModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackTitleModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackTitleModel_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackTitleModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackTitleModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackTitleModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackTitleModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackTitleModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackTitleModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackTitleModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackTitleModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackTitleModel_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackTitleModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackTitleModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackTitleModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackTitleModel_ModelReset(this); };
	void resetInternalData() { callbackTitleModel_ResetInternalData(this); };
	void revert() { callbackTitleModel_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackTitleModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackTitleModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackTitleModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackTitleModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackTitleModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackTitleModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackTitleModel_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return *static_cast<QHash<int, QByteArray>*>(callbackTitleModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return *static_cast<QMap<int, QVariant>*>(callbackTitleModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackTitleModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackTitleModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackTitleModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return *static_cast<QList<QModelIndex>*>(callbackTitleModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackTitleModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackTitleModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackTitleModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackTitleModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackTitleModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackTitleModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackTitleModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackTitleModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackTitleModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackTitleModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackTitleModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackTitleModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTitleModel_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackTitleModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTitleModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTitleModel_CustomEvent(this, event); };
	void deleteLater() { callbackTitleModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTitleModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTitleModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTitleModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTitleModel_TimerEvent(this, event); };
	
	typed01680 rolesDefault() { return _roles; };
	void setRolesDefault(typed01680 p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<Title*> titleDefault() { return _title; };
	void setTitleDefault(QList<Title*> p) { if (p != _title) { _title = p; titleChanged(_title); } };
signals:
	void rolesChanged(typed01680 roles);
	void titleChanged(QList<Title*> title);
public slots:
	void addTitle(Title* v0) { callbackTitleModel_AddTitle(this, v0); };
	void removeTitle(qint32 row) { callbackTitleModel_RemoveTitle(this, row); };
private:
	typed01680 _roles;
	QList<Title*> _title;
};

Q_DECLARE_METATYPE(TitleModel*)


void TitleModel_TitleModel_QRegisterMetaTypes() {
	qRegisterMetaType<typed01680>("typed01680");
}

class DictTab: public QObject
{
Q_OBJECT
Q_PROPERTY(QString dictName READ dictName WRITE setDictName NOTIFY dictNameChanged)
public:
	DictTab(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");DictTab_DictTab_QRegisterMetaType();DictTab_DictTab_QRegisterMetaTypes();callbackDictTab_Constructor(this);};
	QString dictName() { return ({ Moc_PackedString tempVal = callbackDictTab_DictName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setDictName(QString dictName) { QByteArray t06e45a = dictName.toUtf8(); Moc_PackedString dictNamePacked = { const_cast<char*>(t06e45a.prepend("WHITESPACE").constData()+10), t06e45a.size()-10 };callbackDictTab_SetDictName(this, dictNamePacked); };
	void Signal_DictNameChanged(QString dictName) { QByteArray t06e45a = dictName.toUtf8(); Moc_PackedString dictNamePacked = { const_cast<char*>(t06e45a.prepend("WHITESPACE").constData()+10), t06e45a.size()-10 };callbackDictTab_DictNameChanged(this, dictNamePacked); };
	 ~DictTab() { callbackDictTab_DestroyDictTab(this); };
	bool event(QEvent * e) { return callbackDictTab_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackDictTab_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackDictTab_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackDictTab_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackDictTab_CustomEvent(this, event); };
	void deleteLater() { callbackDictTab_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackDictTab_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackDictTab_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackDictTab_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackDictTab_TimerEvent(this, event); };
	
	QString dictNameDefault() { return _dictName; };
	void setDictNameDefault(QString p) { if (p != _dictName) { _dictName = p; dictNameChanged(_dictName); } };
signals:
	void dictNameChanged(QString dictName);
public slots:
private:
	QString _dictName;
};

Q_DECLARE_METATYPE(DictTab*)


void DictTab_DictTab_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class DictTabModel: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(typed01680 roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<DictTab*> dictTab READ dictTab WRITE setDictTab NOTIFY dictTabChanged)
public:
	DictTabModel(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");DictTabModel_DictTabModel_QRegisterMetaType();DictTabModel_DictTabModel_QRegisterMetaTypes();callbackDictTabModel_Constructor(this);};
	typed01680 roles() { return *static_cast<QHash<qint32, QByteArray>*>(callbackDictTabModel_Roles(this)); };
	void setRoles(typed01680 roles) { callbackDictTabModel_SetRoles(this, ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(typed01680 roles) { callbackDictTabModel_RolesChanged(this, ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<DictTab*> dictTab() { return *static_cast<QList<DictTab*>*>(callbackDictTabModel_DictTab(this)); };
	void setDictTab(QList<DictTab*> dictTab) { callbackDictTabModel_SetDictTab(this, ({ QList<DictTab*>* tmpValue = new QList<DictTab*>(dictTab); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_DictTabChanged(QList<DictTab*> dictTab) { callbackDictTabModel_DictTabChanged(this, ({ QList<DictTab*>* tmpValue = new QList<DictTab*>(dictTab); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackDictTabModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackDictTabModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackDictTabModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackDictTabModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackDictTabModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackDictTabModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackDictTabModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackDictTabModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackDictTabModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackDictTabModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackDictTabModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackDictTabModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackDictTabModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackDictTabModel_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackDictTabModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackDictTabModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackDictTabModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackDictTabModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackDictTabModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackDictTabModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackDictTabModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackDictTabModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackDictTabModel_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackDictTabModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackDictTabModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackDictTabModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackDictTabModel_ModelReset(this); };
	void resetInternalData() { callbackDictTabModel_ResetInternalData(this); };
	void revert() { callbackDictTabModel_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackDictTabModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackDictTabModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackDictTabModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackDictTabModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackDictTabModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackDictTabModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackDictTabModel_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return *static_cast<QHash<int, QByteArray>*>(callbackDictTabModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return *static_cast<QMap<int, QVariant>*>(callbackDictTabModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackDictTabModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackDictTabModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackDictTabModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return *static_cast<QList<QModelIndex>*>(callbackDictTabModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackDictTabModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackDictTabModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackDictTabModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackDictTabModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackDictTabModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackDictTabModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackDictTabModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackDictTabModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackDictTabModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackDictTabModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackDictTabModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackDictTabModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackDictTabModel_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackDictTabModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackDictTabModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackDictTabModel_CustomEvent(this, event); };
	void deleteLater() { callbackDictTabModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackDictTabModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackDictTabModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackDictTabModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackDictTabModel_TimerEvent(this, event); };
	
	typed01680 rolesDefault() { return _roles; };
	void setRolesDefault(typed01680 p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<DictTab*> dictTabDefault() { return _dictTab; };
	void setDictTabDefault(QList<DictTab*> p) { if (p != _dictTab) { _dictTab = p; dictTabChanged(_dictTab); } };
signals:
	void rolesChanged(typed01680 roles);
	void dictTabChanged(QList<DictTab*> dictTab);
public slots:
	void addDictTab(DictTab* v0) { callbackDictTabModel_AddDictTab(this, v0); };
	void removeDictTab(qint32 row) { callbackDictTabModel_RemoveDictTab(this, row); };
private:
	typed01680 _roles;
	QList<DictTab*> _dictTab;
};

Q_DECLARE_METATYPE(DictTabModel*)


void DictTabModel_DictTabModel_QRegisterMetaTypes() {
	qRegisterMetaType<typed01680>("typed01680");
}

class QmlBridge: public QObject
{
Q_OBJECT
public:
	QmlBridge(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlBridge_QmlBridge_QRegisterMetaType();QmlBridge_QmlBridge_QRegisterMetaTypes();callbackQmlBridge_Constructor(this);};
	void Signal_OnSearchInputFinished(QString text) { QByteArray t372ea0 = text.toUtf8(); Moc_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQmlBridge_OnSearchInputFinished(this, textPacked); };
	void Signal_OnBackButtonClicked() { callbackQmlBridge_OnBackButtonClicked(this); };
	void Signal_OnForwardButtonClicked() { callbackQmlBridge_OnForwardButtonClicked(this); };
	void Signal_OnTitleClicked(qint32 index) { callbackQmlBridge_OnTitleClicked(this, index); };
	void Signal_SetViewText(QString text) { QByteArray t372ea0 = text.toUtf8(); Moc_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQmlBridge_SetViewText(this, textPacked); };
	 ~QmlBridge() { callbackQmlBridge_DestroyQmlBridge(this); };
	bool event(QEvent * e) { return callbackQmlBridge_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlBridge_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQmlBridge_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlBridge_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlBridge_CustomEvent(this, event); };
	void deleteLater() { callbackQmlBridge_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlBridge_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlBridge_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQmlBridge_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlBridge_TimerEvent(this, event); };
	
signals:
	void onSearchInputFinished(QString text);
	void onBackButtonClicked();
	void onForwardButtonClicked();
	void onTitleClicked(qint32 index);
	void setViewText(QString text);
public slots:
private:
};

Q_DECLARE_METATYPE(QmlBridge*)


void QmlBridge_QmlBridge_QRegisterMetaTypes() {
}

void DictTabModel_AddDictTab(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<DictTabModel*>(ptr), "addDictTab", Q_ARG(DictTab*, static_cast<DictTab*>(v0)));
}

void DictTabModel_RemoveDictTab(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<DictTabModel*>(ptr), "removeDictTab", Q_ARG(qint32, row));
}

struct Moc_PackedList DictTabModel_Roles(void* ptr)
{
	return ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(static_cast<DictTabModel*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList DictTabModel_RolesDefault(void* ptr)
{
	return ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(static_cast<DictTabModel*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void DictTabModel_SetRoles(void* ptr, void* roles)
{
	static_cast<DictTabModel*>(ptr)->setRoles(*static_cast<QHash<qint32, QByteArray>*>(roles));
}

void DictTabModel_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<DictTabModel*>(ptr)->setRolesDefault(*static_cast<QHash<qint32, QByteArray>*>(roles));
}

void DictTabModel_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<DictTabModel*>(ptr), static_cast<void (DictTabModel::*)(QHash<qint32, QByteArray>)>(&DictTabModel::rolesChanged), static_cast<DictTabModel*>(ptr), static_cast<void (DictTabModel::*)(QHash<qint32, QByteArray>)>(&DictTabModel::Signal_RolesChanged));
}

void DictTabModel_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<DictTabModel*>(ptr), static_cast<void (DictTabModel::*)(QHash<qint32, QByteArray>)>(&DictTabModel::rolesChanged), static_cast<DictTabModel*>(ptr), static_cast<void (DictTabModel::*)(QHash<qint32, QByteArray>)>(&DictTabModel::Signal_RolesChanged));
}

void DictTabModel_RolesChanged(void* ptr, void* roles)
{
	static_cast<DictTabModel*>(ptr)->rolesChanged(*static_cast<QHash<qint32, QByteArray>*>(roles));
}

struct Moc_PackedList DictTabModel_DictTab(void* ptr)
{
	return ({ QList<DictTab*>* tmpValue = new QList<DictTab*>(static_cast<DictTabModel*>(ptr)->dictTab()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList DictTabModel_DictTabDefault(void* ptr)
{
	return ({ QList<DictTab*>* tmpValue = new QList<DictTab*>(static_cast<DictTabModel*>(ptr)->dictTabDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void DictTabModel_SetDictTab(void* ptr, void* dictTab)
{
	static_cast<DictTabModel*>(ptr)->setDictTab(*static_cast<QList<DictTab*>*>(dictTab));
}

void DictTabModel_SetDictTabDefault(void* ptr, void* dictTab)
{
	static_cast<DictTabModel*>(ptr)->setDictTabDefault(*static_cast<QList<DictTab*>*>(dictTab));
}

void DictTabModel_ConnectDictTabChanged(void* ptr)
{
	QObject::connect(static_cast<DictTabModel*>(ptr), static_cast<void (DictTabModel::*)(QList<DictTab*>)>(&DictTabModel::dictTabChanged), static_cast<DictTabModel*>(ptr), static_cast<void (DictTabModel::*)(QList<DictTab*>)>(&DictTabModel::Signal_DictTabChanged));
}

void DictTabModel_DisconnectDictTabChanged(void* ptr)
{
	QObject::disconnect(static_cast<DictTabModel*>(ptr), static_cast<void (DictTabModel::*)(QList<DictTab*>)>(&DictTabModel::dictTabChanged), static_cast<DictTabModel*>(ptr), static_cast<void (DictTabModel::*)(QList<DictTab*>)>(&DictTabModel::Signal_DictTabChanged));
}

void DictTabModel_DictTabChanged(void* ptr, void* dictTab)
{
	static_cast<DictTabModel*>(ptr)->dictTabChanged(*static_cast<QList<DictTab*>*>(dictTab));
}

int DictTabModel_DictTabModel_QRegisterMetaType()
{
	return qRegisterMetaType<DictTabModel*>();
}

int DictTabModel_DictTabModel_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<DictTabModel*>(const_cast<const char*>(typeName));
}

int DictTabModel_DictTabModel_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<DictTabModel>();
#else
	return 0;
#endif
}

int DictTabModel_DictTabModel_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<DictTabModel>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int DictTabModel_____setItemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void DictTabModel_____setItemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* DictTabModel_____setItemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int DictTabModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void DictTabModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* DictTabModel_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int DictTabModel_____itemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void DictTabModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* DictTabModel_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

void* DictTabModel___setItemData_roles_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void DictTabModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* DictTabModel___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>;
}

struct Moc_PackedList DictTabModel___setItemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* DictTabModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void DictTabModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* DictTabModel___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* DictTabModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void DictTabModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* DictTabModel___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

int DictTabModel___dataChanged_roles_atList(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void DictTabModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* DictTabModel___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>;
}

void* DictTabModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void DictTabModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* DictTabModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>;
}

void* DictTabModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void DictTabModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* DictTabModel___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>;
}

void* DictTabModel___roleNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<int, QByteArray>*>(ptr)->value(i));
}

void DictTabModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* DictTabModel___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>;
}

struct Moc_PackedList DictTabModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* DictTabModel___itemData_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void DictTabModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* DictTabModel___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>;
}

struct Moc_PackedList DictTabModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* DictTabModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void DictTabModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* DictTabModel___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* DictTabModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void DictTabModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* DictTabModel___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* DictTabModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void DictTabModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* DictTabModel___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

int DictTabModel_____doSetRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void DictTabModel_____doSetRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* DictTabModel_____doSetRoleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int DictTabModel_____setRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void DictTabModel_____setRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* DictTabModel_____setRoleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

void* DictTabModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void DictTabModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* DictTabModel___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* DictTabModel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void DictTabModel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DictTabModel___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* DictTabModel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void DictTabModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DictTabModel___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* DictTabModel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void DictTabModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DictTabModel___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* DictTabModel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void DictTabModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DictTabModel___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* DictTabModel___roles_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<qint32, QByteArray>*>(ptr)->value(i));
}

void DictTabModel___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* DictTabModel___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<qint32, QByteArray>;
}

struct Moc_PackedList DictTabModel___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QHash<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* DictTabModel___setRoles_roles_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<qint32, QByteArray>*>(ptr)->value(i));
}

void DictTabModel___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* DictTabModel___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<qint32, QByteArray>;
}

struct Moc_PackedList DictTabModel___setRoles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QHash<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* DictTabModel___rolesChanged_roles_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<qint32, QByteArray>*>(ptr)->value(i));
}

void DictTabModel___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* DictTabModel___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<qint32, QByteArray>;
}

struct Moc_PackedList DictTabModel___rolesChanged_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QHash<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* DictTabModel___dictTab_atList(void* ptr, int i)
{
	return const_cast<DictTab*>(static_cast<QList<DictTab*>*>(ptr)->at(i));
}

void DictTabModel___dictTab_setList(void* ptr, void* i)
{
	static_cast<QList<DictTab*>*>(ptr)->append(static_cast<DictTab*>(i));
}

void* DictTabModel___dictTab_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<DictTab*>;
}

void* DictTabModel___setDictTab_dictTab_atList(void* ptr, int i)
{
	return const_cast<DictTab*>(static_cast<QList<DictTab*>*>(ptr)->at(i));
}

void DictTabModel___setDictTab_dictTab_setList(void* ptr, void* i)
{
	static_cast<QList<DictTab*>*>(ptr)->append(static_cast<DictTab*>(i));
}

void* DictTabModel___setDictTab_dictTab_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<DictTab*>;
}

void* DictTabModel___dictTabChanged_dictTab_atList(void* ptr, int i)
{
	return const_cast<DictTab*>(static_cast<QList<DictTab*>*>(ptr)->at(i));
}

void DictTabModel___dictTabChanged_dictTab_setList(void* ptr, void* i)
{
	static_cast<QList<DictTab*>*>(ptr)->append(static_cast<DictTab*>(i));
}

void* DictTabModel___dictTabChanged_dictTab_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<DictTab*>;
}

int DictTabModel_____roles_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<qint32>*>(ptr)->at(i);
}

void DictTabModel_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* DictTabModel_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>;
}

int DictTabModel_____setRoles_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<qint32>*>(ptr)->at(i);
}

void DictTabModel_____setRoles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* DictTabModel_____setRoles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>;
}

int DictTabModel_____rolesChanged_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<qint32>*>(ptr)->at(i);
}

void DictTabModel_____rolesChanged_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* DictTabModel_____rolesChanged_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>;
}

void* DictTabModel_NewDictTabModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new DictTabModel(static_cast<QWindow*>(parent));
	} else {
		return new DictTabModel(static_cast<QObject*>(parent));
	}
}

void DictTabModel_DestroyDictTabModel(void* ptr)
{
	static_cast<DictTabModel*>(ptr)->~DictTabModel();
}

char DictTabModel_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* DictTabModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<DictTabModel*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* DictTabModel_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<DictTabModel*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long DictTabModel_FlagsDefault(void* ptr, void* index)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

char DictTabModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char DictTabModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char DictTabModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char DictTabModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char DictTabModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char DictTabModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char DictTabModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char DictTabModel_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char DictTabModel_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char DictTabModel_SubmitDefault(void* ptr)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::submit();
}

void DictTabModel_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<DictTabModel*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void DictTabModel_ResetInternalDataDefault(void* ptr)
{
	static_cast<DictTabModel*>(ptr)->QAbstractListModel::resetInternalData();
}

void DictTabModel_RevertDefault(void* ptr)
{
	static_cast<DictTabModel*>(ptr)->QAbstractListModel::revert();
}

void DictTabModel_SortDefault(void* ptr, int column, long long order)
{
	static_cast<DictTabModel*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList DictTabModel_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<DictTabModel*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList DictTabModel_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<DictTabModel*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* DictTabModel_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::mimeData(*static_cast<QList<QModelIndex>*>(indexes));
}

void* DictTabModel_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<DictTabModel*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* DictTabModel_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<DictTabModel*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList DictTabModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<DictTabModel*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* DictTabModel_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<DictTabModel*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString DictTabModel_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t672968 = static_cast<DictTabModel*>(ptr)->QAbstractListModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t672968.prepend("WHITESPACE").constData()+10), t672968.size()-10 }; });
}

void* DictTabModel_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* DictTabModel_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<DictTabModel*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long DictTabModel_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long DictTabModel_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::supportedDropActions();
}

char DictTabModel_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char DictTabModel_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char DictTabModel_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int DictTabModel_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int DictTabModel_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char DictTabModel_EventDefault(void* ptr, void* e)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char DictTabModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<DictTabModel*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void DictTabModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<DictTabModel*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void DictTabModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<DictTabModel*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void DictTabModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<DictTabModel*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void DictTabModel_DeleteLaterDefault(void* ptr)
{
	static_cast<DictTabModel*>(ptr)->QAbstractListModel::deleteLater();
}

void DictTabModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<DictTabModel*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void DictTabModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<DictTabModel*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}



void QmlBridge_ConnectOnSearchInputFinished(void* ptr)
{
	QObject::connect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::onSearchInputFinished), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::Signal_OnSearchInputFinished));
}

void QmlBridge_DisconnectOnSearchInputFinished(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::onSearchInputFinished), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::Signal_OnSearchInputFinished));
}

void QmlBridge_OnSearchInputFinished(void* ptr, struct Moc_PackedString text)
{
	static_cast<QmlBridge*>(ptr)->onSearchInputFinished(QString::fromUtf8(text.data, text.len));
}

void QmlBridge_ConnectOnBackButtonClicked(void* ptr)
{
	QObject::connect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)()>(&QmlBridge::onBackButtonClicked), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)()>(&QmlBridge::Signal_OnBackButtonClicked));
}

void QmlBridge_DisconnectOnBackButtonClicked(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)()>(&QmlBridge::onBackButtonClicked), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)()>(&QmlBridge::Signal_OnBackButtonClicked));
}

void QmlBridge_OnBackButtonClicked(void* ptr)
{
	static_cast<QmlBridge*>(ptr)->onBackButtonClicked();
}

void QmlBridge_ConnectOnForwardButtonClicked(void* ptr)
{
	QObject::connect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)()>(&QmlBridge::onForwardButtonClicked), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)()>(&QmlBridge::Signal_OnForwardButtonClicked));
}

void QmlBridge_DisconnectOnForwardButtonClicked(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)()>(&QmlBridge::onForwardButtonClicked), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)()>(&QmlBridge::Signal_OnForwardButtonClicked));
}

void QmlBridge_OnForwardButtonClicked(void* ptr)
{
	static_cast<QmlBridge*>(ptr)->onForwardButtonClicked();
}

void QmlBridge_ConnectOnTitleClicked(void* ptr)
{
	QObject::connect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(qint32)>(&QmlBridge::onTitleClicked), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(qint32)>(&QmlBridge::Signal_OnTitleClicked));
}

void QmlBridge_DisconnectOnTitleClicked(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(qint32)>(&QmlBridge::onTitleClicked), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(qint32)>(&QmlBridge::Signal_OnTitleClicked));
}

void QmlBridge_OnTitleClicked(void* ptr, int index)
{
	static_cast<QmlBridge*>(ptr)->onTitleClicked(index);
}

void QmlBridge_ConnectSetViewText(void* ptr)
{
	QObject::connect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::setViewText), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::Signal_SetViewText));
}

void QmlBridge_DisconnectSetViewText(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::setViewText), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::Signal_SetViewText));
}

void QmlBridge_SetViewText(void* ptr, struct Moc_PackedString text)
{
	static_cast<QmlBridge*>(ptr)->setViewText(QString::fromUtf8(text.data, text.len));
}

int QmlBridge_QmlBridge_QRegisterMetaType()
{
	return qRegisterMetaType<QmlBridge*>();
}

int QmlBridge_QmlBridge_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlBridge*>(const_cast<const char*>(typeName));
}

int QmlBridge_QmlBridge_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge>();
#else
	return 0;
#endif
}

int QmlBridge_QmlBridge_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QmlBridge___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QmlBridge___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlBridge___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QmlBridge___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlBridge___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlBridge___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlBridge___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlBridge___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlBridge___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlBridge___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QmlBridge___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QmlBridge_NewQmlBridge(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QWindow*>(parent));
	} else {
		return new QmlBridge(static_cast<QObject*>(parent));
	}
}

void QmlBridge_DestroyQmlBridge(void* ptr)
{
	static_cast<QmlBridge*>(ptr)->~QmlBridge();
}

void QmlBridge_DestroyQmlBridgeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QmlBridge_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlBridge*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlBridge_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlBridge*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QmlBridge_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlBridge_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlBridge_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlBridge*>(ptr)->QObject::deleteLater();
}

void QmlBridge_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



struct Moc_PackedString Title_TitleName(void* ptr)
{
	return ({ QByteArray tdca8d4 = static_cast<Title*>(ptr)->titleName().toUtf8(); Moc_PackedString { const_cast<char*>(tdca8d4.prepend("WHITESPACE").constData()+10), tdca8d4.size()-10 }; });
}

struct Moc_PackedString Title_TitleNameDefault(void* ptr)
{
	return ({ QByteArray tc24b90 = static_cast<Title*>(ptr)->titleNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tc24b90.prepend("WHITESPACE").constData()+10), tc24b90.size()-10 }; });
}

void Title_SetTitleName(void* ptr, struct Moc_PackedString titleName)
{
	static_cast<Title*>(ptr)->setTitleName(QString::fromUtf8(titleName.data, titleName.len));
}

void Title_SetTitleNameDefault(void* ptr, struct Moc_PackedString titleName)
{
	static_cast<Title*>(ptr)->setTitleNameDefault(QString::fromUtf8(titleName.data, titleName.len));
}

void Title_ConnectTitleNameChanged(void* ptr)
{
	QObject::connect(static_cast<Title*>(ptr), static_cast<void (Title::*)(QString)>(&Title::titleNameChanged), static_cast<Title*>(ptr), static_cast<void (Title::*)(QString)>(&Title::Signal_TitleNameChanged));
}

void Title_DisconnectTitleNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Title*>(ptr), static_cast<void (Title::*)(QString)>(&Title::titleNameChanged), static_cast<Title*>(ptr), static_cast<void (Title::*)(QString)>(&Title::Signal_TitleNameChanged));
}

void Title_TitleNameChanged(void* ptr, struct Moc_PackedString titleName)
{
	static_cast<Title*>(ptr)->titleNameChanged(QString::fromUtf8(titleName.data, titleName.len));
}

struct Moc_PackedString Title_DictName(void* ptr)
{
	return ({ QByteArray td78b09 = static_cast<Title*>(ptr)->dictName().toUtf8(); Moc_PackedString { const_cast<char*>(td78b09.prepend("WHITESPACE").constData()+10), td78b09.size()-10 }; });
}

struct Moc_PackedString Title_DictNameDefault(void* ptr)
{
	return ({ QByteArray taf48dd = static_cast<Title*>(ptr)->dictNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(taf48dd.prepend("WHITESPACE").constData()+10), taf48dd.size()-10 }; });
}

void Title_SetDictName(void* ptr, struct Moc_PackedString dictName)
{
	static_cast<Title*>(ptr)->setDictName(QString::fromUtf8(dictName.data, dictName.len));
}

void Title_SetDictNameDefault(void* ptr, struct Moc_PackedString dictName)
{
	static_cast<Title*>(ptr)->setDictNameDefault(QString::fromUtf8(dictName.data, dictName.len));
}

void Title_ConnectDictNameChanged(void* ptr)
{
	QObject::connect(static_cast<Title*>(ptr), static_cast<void (Title::*)(QString)>(&Title::dictNameChanged), static_cast<Title*>(ptr), static_cast<void (Title::*)(QString)>(&Title::Signal_DictNameChanged));
}

void Title_DisconnectDictNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Title*>(ptr), static_cast<void (Title::*)(QString)>(&Title::dictNameChanged), static_cast<Title*>(ptr), static_cast<void (Title::*)(QString)>(&Title::Signal_DictNameChanged));
}

void Title_DictNameChanged(void* ptr, struct Moc_PackedString dictName)
{
	static_cast<Title*>(ptr)->dictNameChanged(QString::fromUtf8(dictName.data, dictName.len));
}

int Title_Title_QRegisterMetaType()
{
	return qRegisterMetaType<Title*>();
}

int Title_Title_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Title*>(const_cast<const char*>(typeName));
}

int Title_Title_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Title>();
#else
	return 0;
#endif
}

int Title_Title_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Title>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Title___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void Title___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Title___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* Title___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Title___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Title___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Title___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Title___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Title___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Title___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Title___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Title___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Title___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void Title___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Title___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* Title_NewTitle(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Title(static_cast<QWindow*>(parent));
	} else {
		return new Title(static_cast<QObject*>(parent));
	}
}

void Title_DestroyTitle(void* ptr)
{
	static_cast<Title*>(ptr)->~Title();
}

void Title_DestroyTitleDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Title_EventDefault(void* ptr, void* e)
{
	return static_cast<Title*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Title_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Title*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Title_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Title*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Title_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Title*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Title_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Title*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Title_DeleteLaterDefault(void* ptr)
{
	static_cast<Title*>(ptr)->QObject::deleteLater();
}

void Title_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Title*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Title_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Title*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void TitleModel_AddTitle(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<TitleModel*>(ptr), "addTitle", Q_ARG(Title*, static_cast<Title*>(v0)));
}

void TitleModel_RemoveTitle(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<TitleModel*>(ptr), "removeTitle", Q_ARG(qint32, row));
}

struct Moc_PackedList TitleModel_Roles(void* ptr)
{
	return ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(static_cast<TitleModel*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList TitleModel_RolesDefault(void* ptr)
{
	return ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(static_cast<TitleModel*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void TitleModel_SetRoles(void* ptr, void* roles)
{
	static_cast<TitleModel*>(ptr)->setRoles(*static_cast<QHash<qint32, QByteArray>*>(roles));
}

void TitleModel_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<TitleModel*>(ptr)->setRolesDefault(*static_cast<QHash<qint32, QByteArray>*>(roles));
}

void TitleModel_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<TitleModel*>(ptr), static_cast<void (TitleModel::*)(QHash<qint32, QByteArray>)>(&TitleModel::rolesChanged), static_cast<TitleModel*>(ptr), static_cast<void (TitleModel::*)(QHash<qint32, QByteArray>)>(&TitleModel::Signal_RolesChanged));
}

void TitleModel_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<TitleModel*>(ptr), static_cast<void (TitleModel::*)(QHash<qint32, QByteArray>)>(&TitleModel::rolesChanged), static_cast<TitleModel*>(ptr), static_cast<void (TitleModel::*)(QHash<qint32, QByteArray>)>(&TitleModel::Signal_RolesChanged));
}

void TitleModel_RolesChanged(void* ptr, void* roles)
{
	static_cast<TitleModel*>(ptr)->rolesChanged(*static_cast<QHash<qint32, QByteArray>*>(roles));
}

struct Moc_PackedList TitleModel_Title(void* ptr)
{
	return ({ QList<Title*>* tmpValue = new QList<Title*>(static_cast<TitleModel*>(ptr)->title()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList TitleModel_TitleDefault(void* ptr)
{
	return ({ QList<Title*>* tmpValue = new QList<Title*>(static_cast<TitleModel*>(ptr)->titleDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void TitleModel_SetTitle(void* ptr, void* title)
{
	static_cast<TitleModel*>(ptr)->setTitle(*static_cast<QList<Title*>*>(title));
}

void TitleModel_SetTitleDefault(void* ptr, void* title)
{
	static_cast<TitleModel*>(ptr)->setTitleDefault(*static_cast<QList<Title*>*>(title));
}

void TitleModel_ConnectTitleChanged(void* ptr)
{
	QObject::connect(static_cast<TitleModel*>(ptr), static_cast<void (TitleModel::*)(QList<Title*>)>(&TitleModel::titleChanged), static_cast<TitleModel*>(ptr), static_cast<void (TitleModel::*)(QList<Title*>)>(&TitleModel::Signal_TitleChanged));
}

void TitleModel_DisconnectTitleChanged(void* ptr)
{
	QObject::disconnect(static_cast<TitleModel*>(ptr), static_cast<void (TitleModel::*)(QList<Title*>)>(&TitleModel::titleChanged), static_cast<TitleModel*>(ptr), static_cast<void (TitleModel::*)(QList<Title*>)>(&TitleModel::Signal_TitleChanged));
}

void TitleModel_TitleChanged(void* ptr, void* title)
{
	static_cast<TitleModel*>(ptr)->titleChanged(*static_cast<QList<Title*>*>(title));
}

int TitleModel_TitleModel_QRegisterMetaType()
{
	return qRegisterMetaType<TitleModel*>();
}

int TitleModel_TitleModel_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TitleModel*>(const_cast<const char*>(typeName));
}

int TitleModel_TitleModel_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TitleModel>();
#else
	return 0;
#endif
}

int TitleModel_TitleModel_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TitleModel>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int TitleModel_____setItemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void TitleModel_____setItemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TitleModel_____setItemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int TitleModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void TitleModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TitleModel_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int TitleModel_____itemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void TitleModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TitleModel_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

void* TitleModel___setItemData_roles_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void TitleModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* TitleModel___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>;
}

struct Moc_PackedList TitleModel___setItemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TitleModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void TitleModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TitleModel___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* TitleModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void TitleModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TitleModel___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

int TitleModel___dataChanged_roles_atList(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void TitleModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* TitleModel___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>;
}

void* TitleModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void TitleModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* TitleModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>;
}

void* TitleModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void TitleModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* TitleModel___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>;
}

void* TitleModel___roleNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<int, QByteArray>*>(ptr)->value(i));
}

void TitleModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* TitleModel___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>;
}

struct Moc_PackedList TitleModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TitleModel___itemData_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void TitleModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* TitleModel___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>;
}

struct Moc_PackedList TitleModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TitleModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void TitleModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TitleModel___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* TitleModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void TitleModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TitleModel___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* TitleModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void TitleModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TitleModel___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

int TitleModel_____doSetRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void TitleModel_____doSetRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TitleModel_____doSetRoleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int TitleModel_____setRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void TitleModel_____setRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TitleModel_____setRoleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

void* TitleModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void TitleModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TitleModel___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* TitleModel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void TitleModel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TitleModel___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* TitleModel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void TitleModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TitleModel___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* TitleModel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void TitleModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TitleModel___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* TitleModel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void TitleModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TitleModel___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* TitleModel___roles_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<qint32, QByteArray>*>(ptr)->value(i));
}

void TitleModel___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* TitleModel___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<qint32, QByteArray>;
}

struct Moc_PackedList TitleModel___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QHash<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TitleModel___setRoles_roles_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<qint32, QByteArray>*>(ptr)->value(i));
}

void TitleModel___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* TitleModel___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<qint32, QByteArray>;
}

struct Moc_PackedList TitleModel___setRoles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QHash<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TitleModel___rolesChanged_roles_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<qint32, QByteArray>*>(ptr)->value(i));
}

void TitleModel___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* TitleModel___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<qint32, QByteArray>;
}

struct Moc_PackedList TitleModel___rolesChanged_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QHash<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TitleModel___title_atList(void* ptr, int i)
{
	return const_cast<Title*>(static_cast<QList<Title*>*>(ptr)->at(i));
}

void TitleModel___title_setList(void* ptr, void* i)
{
	static_cast<QList<Title*>*>(ptr)->append(static_cast<Title*>(i));
}

void* TitleModel___title_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Title*>;
}

void* TitleModel___setTitle_title_atList(void* ptr, int i)
{
	return const_cast<Title*>(static_cast<QList<Title*>*>(ptr)->at(i));
}

void TitleModel___setTitle_title_setList(void* ptr, void* i)
{
	static_cast<QList<Title*>*>(ptr)->append(static_cast<Title*>(i));
}

void* TitleModel___setTitle_title_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Title*>;
}

void* TitleModel___titleChanged_title_atList(void* ptr, int i)
{
	return const_cast<Title*>(static_cast<QList<Title*>*>(ptr)->at(i));
}

void TitleModel___titleChanged_title_setList(void* ptr, void* i)
{
	static_cast<QList<Title*>*>(ptr)->append(static_cast<Title*>(i));
}

void* TitleModel___titleChanged_title_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Title*>;
}

int TitleModel_____roles_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<qint32>*>(ptr)->at(i);
}

void TitleModel_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* TitleModel_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>;
}

int TitleModel_____setRoles_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<qint32>*>(ptr)->at(i);
}

void TitleModel_____setRoles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* TitleModel_____setRoles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>;
}

int TitleModel_____rolesChanged_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<qint32>*>(ptr)->at(i);
}

void TitleModel_____rolesChanged_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* TitleModel_____rolesChanged_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>;
}

void* TitleModel_NewTitleModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TitleModel(static_cast<QWindow*>(parent));
	} else {
		return new TitleModel(static_cast<QObject*>(parent));
	}
}

void TitleModel_DestroyTitleModel(void* ptr)
{
	static_cast<TitleModel*>(ptr)->~TitleModel();
}

char TitleModel_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* TitleModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<TitleModel*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* TitleModel_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<TitleModel*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long TitleModel_FlagsDefault(void* ptr, void* index)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

char TitleModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char TitleModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char TitleModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char TitleModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char TitleModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char TitleModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char TitleModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char TitleModel_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char TitleModel_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char TitleModel_SubmitDefault(void* ptr)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::submit();
}

void TitleModel_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<TitleModel*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void TitleModel_ResetInternalDataDefault(void* ptr)
{
	static_cast<TitleModel*>(ptr)->QAbstractListModel::resetInternalData();
}

void TitleModel_RevertDefault(void* ptr)
{
	static_cast<TitleModel*>(ptr)->QAbstractListModel::revert();
}

void TitleModel_SortDefault(void* ptr, int column, long long order)
{
	static_cast<TitleModel*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList TitleModel_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<TitleModel*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList TitleModel_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<TitleModel*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TitleModel_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::mimeData(*static_cast<QList<QModelIndex>*>(indexes));
}

void* TitleModel_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<TitleModel*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* TitleModel_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<TitleModel*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList TitleModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<TitleModel*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TitleModel_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<TitleModel*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString TitleModel_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t8d30e6 = static_cast<TitleModel*>(ptr)->QAbstractListModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t8d30e6.prepend("WHITESPACE").constData()+10), t8d30e6.size()-10 }; });
}

void* TitleModel_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* TitleModel_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<TitleModel*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long TitleModel_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long TitleModel_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::supportedDropActions();
}

char TitleModel_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char TitleModel_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char TitleModel_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int TitleModel_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int TitleModel_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char TitleModel_EventDefault(void* ptr, void* e)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char TitleModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TitleModel*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TitleModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TitleModel*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void TitleModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TitleModel*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TitleModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TitleModel*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void TitleModel_DeleteLaterDefault(void* ptr)
{
	static_cast<TitleModel*>(ptr)->QAbstractListModel::deleteLater();
}

void TitleModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TitleModel*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void TitleModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TitleModel*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}



struct Moc_PackedString DictTab_DictName(void* ptr)
{
	return ({ QByteArray tbe4738 = static_cast<DictTab*>(ptr)->dictName().toUtf8(); Moc_PackedString { const_cast<char*>(tbe4738.prepend("WHITESPACE").constData()+10), tbe4738.size()-10 }; });
}

struct Moc_PackedString DictTab_DictNameDefault(void* ptr)
{
	return ({ QByteArray t96edaa = static_cast<DictTab*>(ptr)->dictNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t96edaa.prepend("WHITESPACE").constData()+10), t96edaa.size()-10 }; });
}

void DictTab_SetDictName(void* ptr, struct Moc_PackedString dictName)
{
	static_cast<DictTab*>(ptr)->setDictName(QString::fromUtf8(dictName.data, dictName.len));
}

void DictTab_SetDictNameDefault(void* ptr, struct Moc_PackedString dictName)
{
	static_cast<DictTab*>(ptr)->setDictNameDefault(QString::fromUtf8(dictName.data, dictName.len));
}

void DictTab_ConnectDictNameChanged(void* ptr)
{
	QObject::connect(static_cast<DictTab*>(ptr), static_cast<void (DictTab::*)(QString)>(&DictTab::dictNameChanged), static_cast<DictTab*>(ptr), static_cast<void (DictTab::*)(QString)>(&DictTab::Signal_DictNameChanged));
}

void DictTab_DisconnectDictNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<DictTab*>(ptr), static_cast<void (DictTab::*)(QString)>(&DictTab::dictNameChanged), static_cast<DictTab*>(ptr), static_cast<void (DictTab::*)(QString)>(&DictTab::Signal_DictNameChanged));
}

void DictTab_DictNameChanged(void* ptr, struct Moc_PackedString dictName)
{
	static_cast<DictTab*>(ptr)->dictNameChanged(QString::fromUtf8(dictName.data, dictName.len));
}

int DictTab_DictTab_QRegisterMetaType()
{
	return qRegisterMetaType<DictTab*>();
}

int DictTab_DictTab_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<DictTab*>(const_cast<const char*>(typeName));
}

int DictTab_DictTab_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<DictTab>();
#else
	return 0;
#endif
}

int DictTab_DictTab_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<DictTab>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* DictTab___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void DictTab___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* DictTab___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* DictTab___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void DictTab___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DictTab___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* DictTab___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void DictTab___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DictTab___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* DictTab___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void DictTab___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DictTab___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* DictTab___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void DictTab___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DictTab___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* DictTab_NewDictTab(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new DictTab(static_cast<QWindow*>(parent));
	} else {
		return new DictTab(static_cast<QObject*>(parent));
	}
}

void DictTab_DestroyDictTab(void* ptr)
{
	static_cast<DictTab*>(ptr)->~DictTab();
}

void DictTab_DestroyDictTabDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char DictTab_EventDefault(void* ptr, void* e)
{
	return static_cast<DictTab*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char DictTab_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<DictTab*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void DictTab_ChildEventDefault(void* ptr, void* event)
{
	static_cast<DictTab*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void DictTab_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<DictTab*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void DictTab_CustomEventDefault(void* ptr, void* event)
{
	static_cast<DictTab*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void DictTab_DeleteLaterDefault(void* ptr)
{
	static_cast<DictTab*>(ptr)->QObject::deleteLater();
}

void DictTab_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<DictTab*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void DictTab_TimerEventDefault(void* ptr, void* event)
{
	static_cast<DictTab*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
