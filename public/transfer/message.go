package transfer

import "time"

type Message struct {
	DeviceId int64
	Messages []MessageItem
}

// 单条消息投递
type MessageItem struct {
	MessageId      int64     `json:"message_id"`       // 消息id
	SenderType     int       `json:"sender_type"`      // 发送者类型
	SenderId       int64     `json:"sender_id"`        // 发送者id
	SenderDeviceId int64     `json:"sender_device_id"` // 发送者设备id
	ReceiverType   int       `json:"receiver_type"`    // 接收者类型
	ReceiverId     int64     `json:"receiver_id"`      // 接收者id
	Type           int       `json:"type"`             // 消息类型
	Content        string    `json:"content"`          // 消息内容
	Sequence       int64     `json:"sequence"`         // 消息序列
	SendTime       time.Time `json:"send_time"`        // 消息发送时间戳，精确到毫秒
}

type MessageACK struct {
	MessageId    int64     `json:"message_id"`    // 消息id
	DeviceId     int64     `json:"device_id"`     // 设备id
	UserId       int64     `json:"user_id"`       // 用户id
	SyncSequence int64     `json:"sync_sequence"` // 消息序列
	ReceiveTime  time.Time `json:"receive_time"`  // 消息接收时间戳，精确到毫秒
}
