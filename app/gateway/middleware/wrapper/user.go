package wrapper

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/CocaineCong/grpc-todolist/pkg/e"
)

func NewServiceWrapper(name string) {
	c := &Config{
		Namespace:              name,
		Timeout:                1 * time.Second, // TODO 建议加在配置文件里面
		MaxConcurrentRequests:  100,
		RequestVolumeThreshold: 10,
		SleepWindow:            5 * time.Second,
		ErrorPercentThreshold:  50,
	}

	g := NewGroup(c)
	if err := g.Do(name, func() error {
		return errors.New(e.GetMsg(e.ErrorServiceUnavailable))
	}); err != nil {
		fmt.Println("err", err)
	}
}
