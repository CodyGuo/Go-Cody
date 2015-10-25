import QtQuick 2.0

Rectangle {
    width: 360
    height: 360
    Text {
        text: qsTr("Hello World")
        font.pixelSize: 18;
        font.bold: true;
        anchors.centerIn: parent
    }
    MouseArea {
        enabled: true
        anchors.fill: parent
        onClicked: {
            Qt.red;

//            Qt.quit();
        }
    }
}
