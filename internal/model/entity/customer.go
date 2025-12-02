package entity

// Customer 外卖系统用户实体（对应实验一的user表）
type Customer struct {
    Uid      int    `json:"uid"      orm:"uid"`
    Name     string `json:"name"     orm:"name"`
    Password string `json:"password" orm:"password"`
    Phone    string `json:"phone"    orm:"phone"`
    Address  string `json:"address"  orm:"address"`
}