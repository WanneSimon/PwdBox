package pwdbox

import (
	_ "embed"
	"log"
	"sync"
	"time"

	"github.com/zhuxiujia/GoMybatis"
)

//go:embed  AccountMapper.xml
var accountMapperXml []byte
var accountMapper AccountMapper

type AccountService struct {
	IsCreated bool // 是否初始化完成(链接数据库)
}

// 用单例吧
var AccountServiceInstance AccountService
var accountServiceLock sync.Mutex

// 获取 AccountService，需要先初始化数据库链接
func GetAccountService() AccountService {
	if AccountServiceInstance.IsCreated {
		return AccountServiceInstance
	}

	accountServiceLock.Lock()
	if !AccountServiceInstance.IsCreated {
		// SqlEngine.WriteMapperPtr(&(accountServiceInstance.mapper), accountMapperXml)
		AccountServiceInstance.Init(SqlEngine)
		AccountServiceInstance.IsCreated = true
	}
	accountServiceLock.Unlock()
	return AccountServiceInstance
}

func (as *AccountService) Init(sqlEngine GoMybatis.GoMybatisEngine) {
	sqlEngine.WriteMapperPtr(&accountMapper, accountMapperXml)
}

func (as *AccountService) Save(entity Account) *Account {
	if entity.Username == "" {
		panic("require username")
	}
	entity.Id = GenID()
	entity.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	// TODO 密码加密
	if entity.Password != "" {
		data, err := EncryptToString(entity.Password, aesHolder.Key, aesHolder.IV)
		if err != nil {
			log.Println("加密失败")
			log.Println(err)
			return nil
		}
		entity.Password = data
	}

	accountMapper.Save(entity)
	return &entity
}

func (as *AccountService) Update(entity Account) Account {
	// if entity.Id == 0 {
	// 	panic("unknown id")
	// }
	if entity.Username == "" {
		panic("require username")
	}
	accountMapper.Update(entity)
	return entity
}

func (as *AccountService) Get(id int64) Account {
	en, _ := accountMapper.Get(id)
	en.Password = ""
	return en
}

func (as *AccountService) Delete(id int64) {
	accountMapper.Delete(id)
}

func (as *AccountService) List(platformId int64, username string, phone string, email string,
	page int, size int) []Account {
	var start = (page - 1) * size
	list, err := accountMapper.List(platformId, username, phone, email, start, size)
	if err != nil {
		log.Println(err)
		return []Account{}
	}
	return list
}

func (as *AccountService) PageList(platformId int64, username string, phone string, email string,
	page int, size int) PageData[Account] {
	total, err := accountMapper.Count(platformId, username, phone, email)
	if err != nil {
		log.Println(err)
		total = 0
	}
	if total == 0 {
		return EmptyPageData[Account]()
	}

	start := (page - 1) * size
	list, err := accountMapper.List(platformId, username, phone, email, start, size)
	if err != nil {
		log.Println(err)
		// return NewPageData[Account](page, size, total, []Account{})
		list = []Account{}
	}
	return NewPageData[Account](page, size, total, list)
}

func (as *AccountService) UpdatePwd(id int64, newPassword string) bool {
	po, err := accountMapper.Get(id)
	if err != nil {
		log.Printf("不存在账户: %d\n", id)
		log.Println(err)
		return false
	}

	data, err := EncryptToString(newPassword, aesHolder.Key, aesHolder.IV)
	if err != nil {
		log.Println("加密失败")
		log.Println(err)
		return false
	}

	po.Password = data
	accountMapper.Update(po)
	return true
}
