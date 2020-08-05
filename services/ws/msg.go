package ws

const (
	CATE_HEART       = 1 // 心跳包 保活
	CATE_PRIVATE_MSG = 2 // 私信
	CATE_COMMENT     = 3 // 发布的主题评论或恢复评论通知
	CATE_TOPIC       = 4 // 关注的人发布主题通知
)

type Message struct {
	Id       int    `json:"id,omitempty"`
	Category uint8  `json:"category" description:"消息类型"`
	Ack      uint8  `json:"ack" description:"2 交流信息 1 应答"`
	NonceStr string `json:"nonce_str"  description:"随机字符串"`
	Body     string `json:"body,omitempty"`
}
