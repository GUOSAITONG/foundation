package response

// 实例
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 获取用户信息（业务方法）
func GetUser(id string) Result {
	if id == "" {
		return Error(400, "用户ID不能为空")
	}

	// 模拟数据库查询
	if id == "123" {
		return Success(User{
			ID:   "123",
			Name: "张三",
		})
	}

	return Error(404, "用户不存在")
}
