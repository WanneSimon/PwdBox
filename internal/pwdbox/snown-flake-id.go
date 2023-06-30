package pwdbox

import (
	"log"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node
var inited bool = false
var initLock sync.Mutex

func InitSnowFlakeId(startTime string, machineID int64) (err error) {
	var st time.Time
	// 格式化 1月2号下午3时4分5秒  2006年
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		log.Println(err)
		return
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// 生成 64 位的 雪花 ID
func GenID() int64 {
	if !inited {
		initLock.Lock()
		// err := Init("2022-01-01", 1)
		if !inited {
			err := InitSnowFlakeId(time.Now().Format("2006-01-02"), 1)
			if err != nil {
				panic(err)
			}
			inited = true
		}
		initLock.Unlock()
	}

	return node.Generate().Int64()
}
