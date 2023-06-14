package snowflake

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	once sync.Once
	instance *snowflake.Node
)

func newSnowflake() *snowflake.Node {
	once.Do(func() {
		instance, _ = snowflake.NewNode(1)
	})
	return instance
}

func GenerateID() snowflake.ID {
	return newSnowflake().Generate()
}
