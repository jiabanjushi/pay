package V3

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
)

func SandboxCallBack( c  *gin.Context)  {

	//寻找这个账号最早的充值订
	p1 := model.PrepaidPhoneOrders{Username: "niubi001", Successfully: 0, AccountPractical: 1000, RechargeType: "USDT", RechargeAddress: "TW2HWaLWy9pwiRN4yLju6YKW3aQ6Fw8888", CollectionAddress:"TW2HWaLWy9pwiRN4yLju6YKW3aQ6Fw8889"}
	p1.UpdateMaxCreatedOfStatusToTwo(mysql.DB, viper.GetInt64("eth.OrderEffectivityTime"))


	tools.ReturnError200Data(c, nil, "Ok")

}