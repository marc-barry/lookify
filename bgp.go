package main

import (
	"fmt"
	"io"
	"sync"

	"container/ring"

	"github.com/Shopify/exabgp-util/types"
	"github.com/Shopify/exabgp-util/util"
	"github.com/Sirupsen/logrus"
)

const (
	NotificationsBufSize = 10
	StatesBufSize        = 10
)

type BGP struct {
	reader        io.Reader
	notifications *ring.Ring
	states        *ring.Ring
	updates       []*types.UpdateType
	nMu           sync.Mutex
	sMu           sync.Mutex
	uMu           sync.Mutex
}

func NewBGP(reader io.Reader) (*BGP, error) {
	return &BGP{reader: reader, notifications: ring.New(NotificationsBufSize), states: ring.New(StatesBufSize), updates: []*types.UpdateType{}}, nil
}

func (bgp *BGP) handleNotificationMessage(t *types.NotificationType) {
	Log.WithField("notification", fmt.Sprintf("%s", *t)).Debugf("Notification message")
	if *t != types.NotificationType("shutdown") {
		Log.Warnf("Unknown BGP notification message: %s", t)
	}
	mutexed(&bgp.nMu, func() {
		bgp.notifications.Value = t
		bgp.notifications = bgp.notifications.Next()
	})
}

func (bgp *BGP) handleStateMessage(t *types.StateType) {
	Log.WithField("state", fmt.Sprintf("%+v", *t)).Debugf("State message")
	switch t.State {
	case "connected":
		Log.Info("BGP session connected")
	case "up":
		Log.Info("BGP session up")
	case "down":
		Log.Info("BGP session down")
	default:
		Log.Warnf("Unknown BGP session state: %s", t.State)
	}

	mutexed(&bgp.sMu, func() {
		bgp.states.Value = t
		bgp.states = bgp.states.Next()
	})
}

func (bgp *BGP) handleUpdateMessage(t *types.UpdateType) {
	Log.WithField("message", fmt.Sprintf("%+v", *t)).Debugf("Update message")
	mutexed(&bgp.uMu, func() {
		bgp.updates = append(bgp.updates, t)
	})
}

func (bgp *BGP) ReadMessages() {
	Log.Info("Starting BGP message scanner")
	messageChan := util.ScanMessage(bgp.reader)

	for m := range messageChan {
		Log.WithFields(logrus.Fields{
			"type":    m.Type,
			"message": fmt.Sprintf("%+v", *m),
		}).Debugf("Received BGP message")
		switch m.Type {
		case types.NotificationMessageType:
			t, err := types.UnmarshalNotificationType(m.Notification)
			if err != nil {
				Log.WithField("error", err).Error("Error unmarshaling BGP notification message")
				continue
			}
			bgp.handleNotificationMessage(t)
		case types.StateMessageType:
			t, err := types.UnmarshalStateType(m.Neighbor)
			if err != nil {
				Log.WithField("error", err).Error("Error unmarshaling BGP state message")
				continue
			}
			bgp.handleStateMessage(t)
		case types.UpdateMessageType:
			t, err := types.UnmarshalUpdateType(m.Neighbor)
			if err != nil {
				Log.WithField("error", err).Error("Error unmarshaling BGP update message")
				continue
			}
			bgp.handleUpdateMessage(t)
		default:
			Log.WithField("type", m.Type).Warn("Unknown BGP message type")
		}
	}

	Log.Info("Stopped BGP message scanner")
}

func (bgp *BGP) Notifications() []*types.NotificationType {
	notifications := make([]*types.NotificationType, 0, NotificationsBufSize)
	mutexed(&bgp.nMu, func() {
		bgp.notifications.Do(
			func(n interface{}) {
				if n != nil {
					notifications = append(notifications, n.(*types.NotificationType))
				}
			})
	})
	return notifications
}

func (bgp *BGP) States() []*types.StateType {
	states := make([]*types.StateType, 0, StatesBufSize)
	mutexed(&bgp.sMu, func() {
		bgp.states.Do(
			func(n interface{}) {
				if n != nil {
					states = append(states, n.(*types.StateType))
				}
			})
	})
	return states
}

func (bgp *BGP) Updates() []*types.UpdateType {
	updates := make([]*types.UpdateType, 0, len(bgp.updates))
	mutexed(&bgp.uMu, func() {
		for _, u := range bgp.updates {
			updates = append(updates, u)
		}
	})
	return updates
}
