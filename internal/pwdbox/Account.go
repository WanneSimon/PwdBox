package pwdbox

// 账户实体
type Account struct {
	Id         int64  `json:"id"`
	PlatformId int64  `json:"platform_id"` // 平台
	Username   string `json:"username"`    // 用户名
	Password   string `json:"password"`    // 密码
	Phone      string `json:"phone"`       // 绑定电话（多个,隔开）
	Email      string `json:"email"`       // 绑定邮箱（多个,隔开）
	Remark     string `json:"remark"`      // 备注
	CreateTime string `json:"create_time"` // 创建时间
}

type AccountMapper struct {
	Get    func(id int64) (Account, error) `args:"id"`
	Update func(p Account) error
	Save   func(p Account) error
	Delete func(id int64) error `args:"id"`
	// MaxNum func() (int, error)                                                              // maybe nil
	Count func(platformId int64, username string, phone string, email string) (int, error) `args:"platform_id,username,phone,email"`
	List  func(platformId int64, username string, phone string, email string,
		start int, offset int) ([]Account, error) `args:"platform_id,username,phone,email,start,offset"`
}
