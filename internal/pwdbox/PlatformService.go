package pwdbox

import (
	_ "embed"
	"log"
	"sync"
	"time"

	"github.com/zhuxiujia/GoMybatis"
)

//go:embed  PlatformMapper.xml
var PlatformMapperXml []byte
var PlatformMapperInstance PlatformMapper

// var globalString string = "this is init-string"

type PlatformService struct {
	IsCreated bool // 是否初始化完成(链接数据库)
}

// 用单例吧
var PlatformServiceInstance PlatformService
var platformServiceLock sync.Mutex

// 获取 PlatformService，需要先初始化数据库链接
func GetPlatformService() PlatformService {
	if PlatformServiceInstance.IsCreated {
		return PlatformServiceInstance
	}
	platformServiceLock.Lock()
	if !PlatformServiceInstance.IsCreated {
		// SqlEngine.WriteMapperPtr(&(platformServiceInstance.mapper), PlatformMapperXml)
		PlatformServiceInstance.Init(SqlEngine)
		PlatformServiceInstance.IsCreated = true
	}
	platformServiceLock.Unlock()
	return PlatformServiceInstance
}

func (ps *PlatformService) Init(sqlEngine GoMybatis.GoMybatisEngine) {
	sqlEngine.WriteMapperPtr(&PlatformMapperInstance, PlatformMapperXml)
}

func (ps *PlatformService) Save(entity Platform) Platform {
	if entity.Name == "" {
		panic("require username")
	}
	maxNum, err := PlatformMapperInstance.MaxNum()
	if err != nil {
		log.Println(err)
	} else {
		entity.Num = maxNum + 1
	}

	entity.Id = GenID()
	entity.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	PlatformMapperInstance.Save(entity)
	return entity
}

func (ps *PlatformService) Update(entity Platform) Platform {
	// if entity.Id == 0 {
	// 	panic("unknown id")
	// }
	if entity.Name == "" {
		panic("require username")
	}
	PlatformMapperInstance.Update(entity)
	return entity
}

func (ps *PlatformService) Get(id int64) Platform {
	en, _ := PlatformMapperInstance.Get(id)
	return en
}

func (ps *PlatformService) Delete(id int64) {
	PlatformMapperInstance.Delete(id)
}

func (ps *PlatformService) List(name string, page int, size int) []Platform {
	var start = (page - 1) * size
	list, err := PlatformMapperInstance.List(name, start, size)
	if err != nil {
		log.Println(err)
		return []Platform{}
	}
	return list
}

func (ps *PlatformService) PageList(name string, page int, size int) PageData[Platform] {
	total, err := PlatformMapperInstance.Count(name)
	if err != nil {
		log.Println(err)
		total = 0
	}
	if total == 0 {
		return EmptyPageData[Platform]()
	}

	var start = (page - 1) * size
	list, err := PlatformMapperInstance.List(name, start, size)
	if err != nil {
		log.Println(err)
		list = []Platform{}
	}
	return NewPageData[Platform](page, size, total, list)
}
