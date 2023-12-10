package internal

import (
	"context"
	"entry/kitex_gen/user_service"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req UserReqRes
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest,
			map[string]interface{}{
				"err": err.Error(),
			},
		)
		return
	}
	if req.Email == "" || req.UserName == "" {
		c.JSON(consts.StatusBadRequest,
			map[string]interface{}{
				"msg": "InvalidParameters",
			},
		)
		return
	}
	err = GetRpcClient().SaveUser(ctx,
		&user_service.User{
			Id:       req.ID,
			Username: req.UserName,
			Age:      req.Age,
			Email:    req.Email,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(consts.StatusCreated, map[string]interface{}{
		"msg": "created",
	})
}
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req UserReqRes
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest,
			map[string]interface{}{
				"err": err.Error(),
			},
		)
		return
	}
	resp, err := GetRpcClient().GetUser(ctx, req.ID)
	if err != nil {
		c.JSON(consts.StatusBadRequest,
			map[string]interface{}{
				"err": err.Error(),
			},
		)
		return
	}
	log.Println(resp)
	c.JSON(consts.StatusOK, map[string]interface{}{
		"user": resp,
	})
}
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req UserReqRes
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest,
			map[string]interface{}{
				"err": err.Error(),
			},
		)
		return
	}
	err = GetRpcClient().DeleteUser(ctx, req.ID)
	if err != nil {
		c.JSON(consts.StatusBadRequest,
			map[string]interface{}{
				"err": err.Error(),
			},
		)
		return
	}
	c.JSON(consts.StatusOK, map[string]interface{}{
		"msg": "deleted",
	})
}
