package Model

type Information struct {
	Id            int    `json:"id" xorm:"pk autoincr"`
	Title         string `json:"title"`
	School		  string `json:"school`
	Createtime    string `json:"createtime"`
	Userid        int    `json:"userid"`
	Usernickname  string `json:"usernickname"`
	Usertelephone string `json:"usertelephone"`
	Content       string `json:"content"`
	Praise        string `json:"praise"`
}

type PostpraiseH struct {
	Id int `json:"id" binding:"required"`
	Praise string `json:"praise" `
}

type Suggest struct {
	Id         int    `json:"id" binding:"required" xorm:"pk autoincr"`
	Content    string `json:"content" binding:"required"`
	Createtime string `json:"createtime" binding:"required"`
	Userid     int    `json:"userid" binding:"required"`
	School string `json:"school" binding:"required"`
}


type Updatee struct {
	Id int	`json:"id" xorm:"pk autoincr"`
	Version float64`json:"version"`
	Url string `json:"url"`
}

type User struct {
	Id         int    `json:"id" xorm:"pk autoincr"`
	Telephone  string `json:"telephone" xorm:"telephone"`
	School     string `json:"school" xorm:"school"`
	Nickname   string `json:"nickname" xorm:"nickname"`
	Sex        string `json:"sex" xorm:"sex"`
	Signature  string `json:"signature" xorm:"signature"`
	Createtime string `json:"createtime" xorm:"createtime"`
	Password   string `json:"password" xorm:"password"`
	Authority  string `json:"authority" xorm:"authority"`
	Praise     string `json:"praise" xorm:"praise"`
}

type GetUser struct {
	Id int `json:"id", binding:"required" xorm:"pk autoincr"`
	Telephone string `json:"telephone" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type Putsign struct {
	Signature string `json:"signature"`
}