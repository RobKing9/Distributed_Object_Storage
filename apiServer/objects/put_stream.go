package objects

import (
	"Distributed_Storage_Deploy/apiServer/heartbeat"
	"Distributed_Storage_Deploy/src/lib/rs"
	"fmt"
)

func putStream(hash string, size int64) (*rs.RSPutStream, error) {
	servers := heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS, nil)
	if len(servers) != rs.ALL_SHARDS {
		return nil, fmt.Errorf("cannot find enough dataServer")
	}

	return rs.NewRSPutStream(servers, hash, size)
}
