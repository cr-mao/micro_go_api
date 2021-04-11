package initalize


import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"go.uber.org/zap"
)

func InitSentinel(){
	// We should initialize Sentinel first.
	conf := config.NewDefaultConfig()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		zap.S().Fatal(err)
	}
	//规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "goods-list",  //商品立碑
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, // flow.Throttling     匀速  要计算平均请求数 如 1秒 10个。 那么100ms 就能通过1个。
			Threshold:              1,        // 测试用 1秒最多1个
			StatIntervalInMs:       1000,
		},
	})
	if err != nil {
		zap.S().Fatal(err)
		return
	}
	//
	//ch := make(chan struct{})
	//for i := 0; i < 10; i++ {
	//	e, b := sentinel.Entry(resName, sentinel.WithTrafficType(base.Inbound))
	//	if b != nil {
	//		fmt.Println("not pass")
	//		// Blocked. We could get the block reason from the BlockError.
	//	} else {
	//
	//		fmt.Println("pass")
	//		// Passed, wrap the logic here.
	//		// Be sure the entry is exited finally.
	//		e.Exit()
	//	}
	//}

}
