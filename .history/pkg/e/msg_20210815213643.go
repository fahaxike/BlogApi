/*
 * @Author: dsz
 * @Date: 2021-08-15 21:16:22
 * @LastEditTime: 2021-08-15 21:36:43
 * @Description:msg
 * @FilePath: \BlogApi\pkg\e\msg.go
 */
package e
import "log"


var MsgFlags=map[int]string{
	SUCCESS:"ok",
	ERROR:"fail",
	INVALID_PARAM:"请求参数错误",
	ERROR_EXIST_TAG:"已存在该标签签名",
	ERROR_NOT_EXIST_ARTICLE:"该文章不存在",
	ERROR_NOT_EXIST_TAG:"该标签不存在",

	ERROR_AUTH:"Token错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:"Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:"Token已超时",
	ERROR_AUTH_TOKEN:"Token生成失败",
	}
	func GetMsg(code int) (msg string){
		var ok bool
		msg,ok =MsgFlags[code]
		if !ok{
			msg=MsgFlags[ERROR]
		}
		return msg
	}


}