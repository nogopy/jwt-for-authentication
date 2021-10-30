package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserIdFromContext(ctx *gin.Context) int {
	if customerId, ok := ctx.Get("user_id"); ok {
		iAreaId := 0
		switch customerId.(type) {
		case float64:
			iAreaId = int(customerId.(float64))
			break
		case string:
			iAreaId, _ = strconv.Atoi(customerId.(string))
			break
		case int:
			iAreaId = customerId.(int)
			break
		case int64:
			iAreaId = int(customerId.(int64))
			break
		case uint32:
			iAreaId = int(customerId.(uint32))
			break
		}

		return iAreaId
	}

	return 0
}

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
