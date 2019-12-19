package user

import (
	"github.com/golang/mock/gomock"
	"github.com/qiuhoude/go-web/prepare/24_gomock/mock"
	"testing"
)

func TestUser_GetUserInfo(t *testing.T) {
	// 返回 gomock.Controller，它代表 mock 生态系统中的顶级控件.定义了 mock 对象的范围、生命周期和期待值.
	// 另外它在多个 goroutine 中是安全的
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id int64 = 1
	// 创建一个新的 mock 实例
	mockMale := mock.NewMockMale(ctl)

	// 声明给定的调用应按顺序进行（是对 gomock.After 的二次封装）

	gomock.InOrder(
		// 这里有三个步骤，EXPECT()返回一个允许调用者设置期望和返回值的对象.Get(id) 是设置入参并调用 mock 实例中的方法.Return(nil)
		// 是设置先前调用的方法出参.简单来说，就是设置入参并调用，最后设置返回值
		mockMale.EXPECT().Get(id).Return(nil).MaxTimes(10),
	)

	user := NewUser(mockMale)
	err := user.GetUserInfo(id)
	if err != nil {
		t.Errorf("user.GetUserInfo err: %v", err)
	}
}
