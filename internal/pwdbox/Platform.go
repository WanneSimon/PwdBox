package pwdbox

// 平台实体
type Platform struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`   // 平台名称
	Site   string `json:"site"`   // 官网
	Remark string `json:"remark"` // 备注
	Num    int    `json:"num"`    // 序号
	// CreateTime time.Time `json:"create_time"` // 创建时间
	CreateTime string `json:"create_time"` // 创建时间
	Img        string `json:"img"`         // 图片
}

type PlatformMapper struct {
	Get    func(id int64) (Platform, error) `args:"id"`
	Update func(p Platform) error
	Save   func(p Platform) error
	Delete func(id int64) error                                         `args:"id"`
	MaxNum func() (int, error)                                          // maybe nil
	Count  func(name string) (int, error)                               `args:"name"`
	List   func(name string, start int, offset int) ([]Platform, error) `args:"name,start,offset"`
}
