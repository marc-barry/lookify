package types

import (
	"encoding/json"
)

const (
	NotificationMessageType = "notification"
	StateMessageType        = "state"
	UpdateMessageType       = "update"
)

type ExaBGPReaderMessage struct {
	ExaBGP       string          `json:"exabgp"`
	Time         uint64          `json:"time"`
	Host         string          `json:"host"`
	PID          string          `json:"pid"`
	PPID         string          `json:"ppid"`
	Type         string          `json:"type"`
	Neighbor     json.RawMessage `json:"neighbor"`
	Notification json.RawMessage `json:"notification"`
}

type UpdateType struct {
	IP      string `json:"ip"`
	Address struct {
		Local string `json:"local"`
		Peer  string `json:"peer"`
	} `json:"address"`
	ASN struct {
		Local string `json:"local"`
		Peer  string `json:"peer"`
	} `json:"asn"`
	Message struct {
		Update struct {
			Attribute *struct {
				Origin            string    `json:"origin,omitempty"`
				ASPath            []uint32  `json:"as-path,omitempty"`
				ConfederationPath []uint32  `json:"confederation-path"`
				AtomicAggegate    bool      `json:"atomic-aggregate"`
				Community         [][]int32 `json:"community,omitempty"`
			} `json:"attribute,omitempty"`
			Announce *struct {
				IPv4Unicast *map[string]interface{} `json:"ipv4 unicast,omitempty"`
			} `json:"announce,omitempty"`
			Withdraw *struct {
				IPv4Unicast *map[string]interface{} `json:"ipv4 unicast,omitempty"`
			} `json:"withdraw,omitempty"`
		} `json:"update"`
	} `json:"message"`
}

type NotificationType string

type StateType struct {
	IP      string `json:"ip"`
	Address struct {
		Local string `json:"local"`
		Peer  string `json:"peer"`
	} `json:"address"`
	ASN struct {
		Local string `json:"local"`
		Peer  string `json:"peer"`
	} `json:"asn"`
	State  string `json:"state"`
	Reason string `json:"reason,omitempty"`
}

func UnmarshalExaBGPReaderMessage(b []byte) (m *ExaBGPReaderMessage, err error) {
	m = &ExaBGPReaderMessage{}
	err = json.Unmarshal(b, m)
	if err != nil {
		m = nil
	}
	return
}

func UnmarshalUpdateType(r json.RawMessage) (t *UpdateType, err error) {
	t = &UpdateType{}
	err = json.Unmarshal(r, t)
	if err != nil {
		t = nil
	}
	return
}

func UnmarshalNotificationType(r json.RawMessage) (t *NotificationType, err error) {
	var n NotificationType
	t = &n
	err = json.Unmarshal(r, t)
	if err != nil {
		t = nil
	}
	return
}

func UnmarshalStateType(r json.RawMessage) (t *StateType, err error) {
	t = &StateType{}
	err = json.Unmarshal(r, t)
	if err != nil {
		t = nil
	}
	return
}
