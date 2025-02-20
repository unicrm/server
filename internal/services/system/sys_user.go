package system

import (
	"errors"

	"github.com/google/uuid"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/models/system"
	"github.com/unicrm/server/internal/utils"
	"gorm.io/gorm"
)

type UserService struct{}

var UserServiceApp = new(UserService)

// Register 注册
func (userService *UserService) Register(u *system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(globals.UNICRM_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	// 附加UUID，密码加密
	u.UUID = uuid.Must(uuid.NewV6())
	u.Password, _ = utils.BcryptHash(u.Password)
	err = globals.UNICRM_DB.Create(&u).Error
	return *u, err
}

// Login 登录
func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if globals.UNICRM_DB == nil {
		return nil, errors.New("数据库连接失败")
	}
	var user system.SysUser
	err = globals.UNICRM_DB.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return userInter, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("密码错误")
	}
	return &user, err
}
