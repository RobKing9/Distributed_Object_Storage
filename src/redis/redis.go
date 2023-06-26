package redis

// 通过 Redis 将定位信息表加载到内存 实现对数据的去重

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/go-redis/redis"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 定义文件信息结构体

type File struct {
	Hash    string `json:"hash"`
	Name    string `json:"name"`
	ShardID int    `json:"shard_id"`
}

/*
	遍历指定目录下的所有文件，并计算出每个文件的散列值。将文件的散列值、文件名和分片 ID 保存到文件信息结构体中
	并使用 JSON 编码将其转换成字符串。最后，使用 HSet 命令将文件信息存储到 Redis 的 Hash 数据结构中。
	对于已经保存到 Redis 中的文件信息，我们可以使用 HExists 和 HGet 命令进行查询操作。
	如果需要删除某个文件信息，可以使用 HDel 命令进行删除。
*/

func collectObject() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})

	// 遍历指定目录下的所有文件，并将文件信息存储到 Redis 中
	dir := "/path/to/files"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// 计算文件的散列值
			hash := calculateFileHash(path)

			// 初始化文件信息结构体
			file := File{
				Hash:    hash,
				Name:    info.Name(),
				ShardID: 1,
			}
			fileJson, _ := json.Marshal(file)

			// 将文件信息存储到 Redis 中
			err := client.HSet("file_info", hash, fileJson).Err()
			if err != nil {
				panic(err)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// 查询某个散列值的文件是否存在
	exists, err := client.HExists("file_info", "hash1").Result()
	if err != nil {
		panic(err)
	}
	if exists {
		// 获取散列值为"hash1"的文件的信息
		fileJson, err := client.HGet("file_info", "hash1").Bytes()
		if err != nil {
			panic(err)
		}
		var file File
		json.Unmarshal(fileJson, &file)

		// 对获取到的文件信息进行处理
		// ...
	} else {
		// 文件不存在
		panic("file not found")
	}
}

// 计算文件的散列值
func calculateFileHash(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	// 使用 SHA256 计算文件的散列值
	hash := sha256.Sum256(data)
	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}
