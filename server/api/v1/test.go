package v1

import (
    "github.com/gin-gonic/gin"

    "gin-vue-admin/model/response"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /test/testT [post]
func TestT(c *gin.Context) {
    response.Ok(c)
}

