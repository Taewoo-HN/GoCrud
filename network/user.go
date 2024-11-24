package network

import (
	"awesomeProject/service"
	"awesomeProject/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router      *Network
	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}
		router.RegisterGET("/", userRouterInstance.get)
		router.RegisterPOST("/", userRouterInstance.create)
		router.RegisterDELETE("/", userRouterInstance.delete)
		router.RegisterUPDATE("/", userRouterInstance.update)
	})
	return userRouterInstance
}

func (u *userRouter) create(g *gin.Context) {
	var req types.CreateUserRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(g, types.CreateUserResponse{
			ApiResponse: types.NewAPIResponse(-1, "바인딩 에러", err.Error()),
		})
	} else if err = u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(g, types.CreateUserResponse{
			ApiResponse: types.NewAPIResponse(-1, "Create 에러", err.Error()),
		})
	} else {
		u.router.failedResponse(g, types.CreateUserResponse{
			ApiResponse: types.NewAPIResponse(1, "성공", nil),
		})
	}

}
func (u *userRouter) get(g *gin.Context) {
	fmt.Println("User Router Get")

	u.router.okResponse(g, &types.GetUserResponse{
		ApiResponse: types.NewAPIResponse(1, "성공입니다.", nil),
		Users:       u.userService.Get(),
	})
}
func (u *userRouter) update(g *gin.Context) {
	//var req types.UpdateUserRequest
	fmt.Println("User Router Update")
	u.userService.Update(nil, nil)
	u.router.okResponse(g, &types.UpdateUserResponse{
		ApiResponse: types.NewAPIResponse(1, "성공입니다.", nil),
	})
}
func (u *userRouter) delete(g *gin.Context) {
	var req types.DeleteUserRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(g, types.DeleteUserResponse{
			ApiResponse: types.NewAPIResponse(-1, "바인딩 에러", err.Error()),
		})
	} else if err = u.userService.Delete(req.ToUser()); err != nil {
		u.router.failedResponse(g, types.DeleteUserResponse{
			ApiResponse: types.NewAPIResponse(-1, "삭제 에러", err.Error()),
		})
	} else {
		u.router.okResponse(g, &types.DeleteUserResponse{
			ApiResponse: types.NewAPIResponse(1, "성공", nil),
		})
	}
	u.router.okResponse(g, &types.DeleteUserResponse{
		ApiResponse: types.NewAPIResponse(1, "성공입니다.", nil),
	})
}
