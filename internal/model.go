package internal

type UserReqRes struct {
	ID       int64  `json:"id,required" form:"id" query:"id"`
	UserName string `json:"username" form:"username" query:"username"`
	Email    string `json:"email" form:"email" query:"email"`
	Age      int32  `json:"age" form:"age" query:"age"`
}
