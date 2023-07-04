package view

type HelloReq struct {
	Name string `json:"name" validate:"required" reg_err_info:"名字必填"`
	Age  int    `json:"age" validate:"gte=18" reg_err_info:"年龄不应小于18岁"`
}
