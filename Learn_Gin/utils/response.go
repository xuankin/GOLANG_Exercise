// Learn_Gin/utils/response.go
package utils

import "github.com/gin-gonic/gin"

// Cấu trúc phản hồi chuẩn
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Hàm helper trả về phản hồi thành công
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: "Success",
		Data:    data,
	})
}

// Hàm helper trả về phản hồi lỗi
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
