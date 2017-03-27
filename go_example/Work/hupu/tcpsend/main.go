package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"github.com/codyguo/logs"
	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

type SendMsg struct {
	Flag    int16
	SeServe int16
	Len     int32
	Msg     string
}

type MyMainWindow struct {
	*walk.MainWindow
	flagLb  *walk.Label
	flagLe  *walk.LineEdit
	msgLb   *walk.Label
	msgText *walk.TextEdit

	msgBtn *walk.PushButton
	msgOut *walk.TextEdit
}

func (mw *MyMainWindow) EncodeMsg() {
	defer func() {
		mw.flagLe.SetEnabled(true)
		mw.msgText.SetEnabled(true)
	}()

	mw.flagLe.SetEnabled(false)
	mw.msgText.SetEnabled(false)
	mw.msgOut.SetText("")
	flagLe, err := strconv.Atoi(mw.flagLe.Text())
	flag := int16(flagLe)
	if 0 == flag || err != nil {
		walk.MsgBox(mw, "错误提示", "消息ID为数字，请重新输入消息ID！", walk.MsgBoxIconError)
		mw.flagLe.SetText("")
		return
	}
	msg := mw.msgText.Text()
	msg = strings.TrimSpace(msg)
	lenMsg := int32(len(msg))
	send := SendMsg{
		Flag: flag,
		Len:  lenMsg,
		Msg:  msg,
	}
	// logs.Notice(send)
	// logs.Noticef("Flag == %04x, ReServe == %04x, Len == %08x, Msg == %x", send.Flag, send.SeServe, send.Len, send.Msg)

	// result := encodeMsg(send)
	// logs.Notice(result)
	msgOut := encodeMsg(send)
	mw.msgOut.SetText(msgOut)
}

func main() {
	mw := new(MyMainWindow)
	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "iMan消息结构化",
		MinSize:  Size{435, 380},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{AssignTo: &mw.flagLb, Text: "消息ID:"},
					LineEdit{AssignTo: &mw.flagLe},
					Label{AssignTo: &mw.msgLb, Text: "消息体:"},
					TextEdit{AssignTo: &mw.msgText},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						AssignTo: &mw.msgBtn,
						Text:     "结构化",
						OnClicked: func() {
							go mw.EncodeMsg()
						},
					},
					PushButton{
						Text: "复制",
						OnClicked: func() {
							walk.Clipboard().Clear()
							if err := walk.Clipboard().SetText(mw.msgOut.Text()); err != nil {
								walk.MsgBox(mw, "错误提示", err.Error(), walk.MsgBoxIconError)
							}
						},
					},
				},
			},
			GroupBox{
				Title:  "结构化消息",
				Layout: VBox{},
				Children: []Widget{
					TextEdit{AssignTo: &mw.msgOut, ReadOnly: true},
				},
			},
		},
	}).Create(); err != nil {
		logs.Fatal(err)
	}

	mw.Run()
}

func encodeMsg(send SendMsg) string {
	flag := binaryMsg(send.Flag)
	seServe := binaryMsg(send.SeServe)
	lenMsg := binaryMsg(send.Len)
	msg := fmt.Sprintf("%x", send.Msg)

	var buf bytes.Buffer
	buf.Write([]byte(flag))
	buf.Write([]byte(seServe))
	buf.Write([]byte(lenMsg))
	buf.Write([]byte(msg))

	return buf.String()
}

func binaryMsg(data interface{}) string {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)
	checkErr(err)

	return fmt.Sprintf("%x", buf.Bytes())
}

func checkErr(err error) {
	if err != nil {
		logs.Fatal(err)
	}
}

// flag := fmt.Sprintf("%04x", send.Flag)
// seServe := fmt.Sprintf("%04x", send.SeServe)
// lenMsg := fmt.Sprintf("%08x", send.Len)
// msg := fmt.Sprintf("%x", send.Msg)

// var buf bytes.Buffer
// buf.Write([]byte(flag[2:] + flag[0:2]))
// buf.Write([]byte(seServe[2:] + seServe[0:2]))
// buf.Write([]byte(lenMsg[6:] + lenMsg[4:6] + lenMsg[2:4] + lenMsg[0:2]))
// buf.Write([]byte(msg))
