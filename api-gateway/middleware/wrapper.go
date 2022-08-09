package middleware

import (
	"github.com/afex/hystrix-go/hystrix"
)

func NewUserWrapper() {
	commandName := "my-endpoint"
	hystrix.ConfigureCommand(commandName, hystrix.CommandConfig{
		Timeout:                1000 * 30,
		ErrorPercentThreshold:  1,
		SleepWindow:            10000,
		MaxConcurrentRequests:  1000,
		RequestVolumeThreshold: 5,
	})
}

//func Hystrix(commandName string,handlerFunc gin.HandlerFunc) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var resp interface{}
//		err := hystrix.Do(commandName, func() error {
//			resp, err :=
//		}, func(e error) error {
//			return errors.New("hystrix fallback")
//		})
//		c.Next()
//	}
//}
