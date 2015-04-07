package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/Shopify/exabgp-util/types"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadMessage(t *testing.T) {
	Convey("Testing ability to read an ExaBGP message", t, func() {
		b, err := ioutil.ReadFile("../files/json/state_connected.json")
		So(err, ShouldBeNil)
		So(b, ShouldNotBeEmpty)

		compacted := bytes.Buffer{}
		err = json.Compact(&compacted, b)
		So(err, ShouldBeNil)

		r := bytes.NewReader(b)

		m, err := ReadMessage(r)
		So(err, ShouldBeNil)
		So(m, ShouldNotBeNil)

		So(m.Type, ShouldEqual, types.StateMessageType)

		t, err := types.UnmarshalStateType(m.Neighbor)
		So(err, ShouldBeNil)
		So(t, ShouldNotBeNil)

		u, err := json.Marshal(t)
		So(err, ShouldBeNil)
		So(u, ShouldNotBeEmpty)

		compacted = bytes.Buffer{}
		err = json.Compact(&compacted, m.Neighbor)
		So(err, ShouldBeNil)

		So(compacted.String(), ShouldEqual, string(u))
	})
}
func TestScanMessage(t *testing.T) {
	Convey("Testing ability to scan a ExaBGP messages", t, func() {
		buf := bytes.Buffer{}

		b, err := ioutil.ReadFile("../files/json/state_connected.json")
		So(err, ShouldBeNil)
		So(b, ShouldNotBeEmpty)
		compacted := bytes.Buffer{}
		err = json.Compact(&compacted, b)
		_, err = buf.Write(compacted.Bytes())
		So(err, ShouldBeNil)
		_, _ = buf.WriteString("\n")

		b, err = ioutil.ReadFile("../files/json/notification_shutdown.json")
		So(err, ShouldBeNil)
		So(b, ShouldNotBeEmpty)
		compacted = bytes.Buffer{}
		err = json.Compact(&compacted, b)
		_, err = buf.Write(compacted.Bytes())
		So(err, ShouldBeNil)
		_, _ = buf.WriteString("\n")

		b, err = ioutil.ReadFile("../files/json/update_announce.json")
		So(err, ShouldBeNil)
		So(b, ShouldNotBeEmpty)
		compacted = bytes.Buffer{}
		err = json.Compact(&compacted, b)
		_, err = buf.Write(compacted.Bytes())
		So(err, ShouldBeNil)
		_, _ = buf.WriteString("\n")

		mChan := ScanMessage(&buf)

		m := <-mChan
		So(m, ShouldNotBeNil)
		So(m.Type, ShouldEqual, types.StateMessageType)

		m = <-mChan
		So(m, ShouldNotBeNil)
		So(m.Type, ShouldEqual, types.NotificationMessageType)

		m = <-mChan
		So(m, ShouldNotBeNil)
		So(m.Type, ShouldEqual, types.UpdateMessageType)

		So(len(mChan), ShouldEqual, 0)

		select {
		case x, more := <-mChan:
			So(x, ShouldBeNil)
			So(more, ShouldBeFalse)
		}
	})
}
