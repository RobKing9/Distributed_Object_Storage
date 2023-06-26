package heartbeat

import (
	"Distributed_Object_Storage/src/lib/rabbitmq"
	"strconv"
	"sync"
	"time"
)

// 接收和处理来自数据服务节点的心跳消息

// 缓存所有的数据节点 key为数据节点的心跳消息即监听地址,value 为时间
var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

func ListenHeartbeat() {
	q := rabbitmq.New(rabbitmq.RABBITMQ_SERVER)
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume()
	// 清除10s没收到心跳消息的数据节点
	go removeExpiredDataServer()
	for msg := range c {
		// 每个节点的监听地址
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			// 检查节点的最后一次心跳消息时间是否在10秒钟之前
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

// 获取所有的数据节点监听地址

func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s, _ := range dataServers {
		ds = append(ds, s)
	}
	return ds
}
