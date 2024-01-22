package model

type UserFrom struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	ChainAccount string `json:"chainAccount"`
	Name         string `json:"name"`
	CardId       string `json:"cardId"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
}

type LoginFrom struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUserFrom struct {
	Username     string `json:"username"`
	ChainAccount string `json:"chainAccount"`
	Name         string `json:"name"`
	CardId       string `json:"cardId"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
}

type MessageResponse struct {
	Message string      // 消息字段
	Data    interface{} // 数据字段，可以是任意类型
}
