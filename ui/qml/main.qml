import QtQuick 2.7
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.0

ApplicationWindow {
    id: applicationWindow
    visible: true
    width: 640
    height: 480
    title: qsTr("dictee")

    ColumnLayout {
        id: columnLayout
        width: 400
        height: 400
        transformOrigin: Item.Center

        RowLayout {
            id: inputLayout
            y: 0
            width: 400
            height: 40

            Button {
                id: backButton
                width: 30
                height: 30
                text: qsTr("<")
                padding: 12
                rightPadding: 8
                Layout.maximumHeight: 32
                Layout.maximumWidth: 32
                Layout.minimumHeight: 32
                Layout.minimumWidth: 32
                Layout.preferredHeight: 32
                Layout.preferredWidth: 32
                onClicked: QmlBridge.onBackButtonClicked()
            }

            Button {
                id: forwardButton
                width: 30
                height: 30
                text: qsTr(">")
                Layout.preferredHeight: 32
                Layout.preferredWidth: 32
                Layout.maximumHeight: 32
                Layout.maximumWidth: 32
                Layout.minimumHeight: 32
                Layout.minimumWidth: 32
                leftPadding: 8
                onClicked: QmlBridge.onForwardButtonClicked()
            }

            TextField {
                id: searchInput
                width: 120
                height: 30
                Layout.fillHeight: false
                Layout.fillWidth: true
                font.family: "Arial"
                font.pixelSize: 12
                focus: true
                placeholderText: qsTr("Search")
                onEditingFinished: QmlBridge.onSearchInputFinished(this.text)
            }

            Button {
                id: searchClear
                width: 30
                height: 30
                text: qsTr("x")
                Layout.maximumHeight: 32
                Layout.maximumWidth: 32
                Layout.minimumHeight: 32
                Layout.minimumWidth: 32
                Layout.preferredHeight: 32
                Layout.preferredWidth: 32
            }
        }

        TabBar {
            id: tabBar
            width: 400
            height: 32
            Layout.maximumHeight: 32
            Layout.minimumHeight: 32
            Layout.minimumWidth: 200
            Layout.preferredHeight: 32
            Layout.preferredWidth: 400
            Layout.fillWidth: true
            onCurrentIndexChanged: {
                 console.log("Tab Bar current index changed: "+ currentIndex)
            }
        }

        RowLayout {
            id: viewLayout
            width: 100
            height: 100

            Component {
                id: titleDelegate
                Rectangle {
                    id: wrapper
                    width: 100; height: 40
                    color: ListView.isCurrentItem ? "#2A61C7" : "#F6F6F6"
                    TextField {
                        text: titleName
                        color: ListView.isCurrentItem ? "white" : "black"
                        readOnly: true
                        horizontalAlignment: TextInput.AlignHCenter
                    }

                    MouseArea {
                        anchors.fill: parent
                        onClicked: {
                            titleList.currentIndex = index
                            QmlBridge.onTitleClicked(index)
                        }
                    }
                }
            }

            ListView {
                id: titleList
                width: 180; height: 200
                Layout.minimumWidth: 40
                Layout.maximumWidth: 230
                Layout.fillWidth: true
                Layout.fillHeight: true

                model: TitleModel
                delegate: titleDelegate
                focus: true
            }
            Text {
                id: textArea
                text: "hoge"
                textFormat: Text.RichText
                Layout.minimumHeight: 120
                Layout.minimumWidth: 200
                Layout.fillHeight: true
                Layout.fillWidth: true
                wrapMode: TextArea.Wrap
                linkActivated: console.log(link + " link activated")
                Connections {
                        target: QmlBridge
                        onSetViewText: {
				            textArea.text = text
                        }
                }
            }
        }
    }
}
