/*
 * @Author: dsz
 * @Date: 2021-08-15 21:11:14
 * @LastEditTime: 2021-08-15 21:16:05
 * @Description:code
 * @FilePath: \BlogApi\pkg\e\code.go
 */
package e

const (
	SUCCESS       = 100
	ERROR         = 500
	INVALID_PARAM = 400

	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)
