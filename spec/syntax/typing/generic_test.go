package typing

import (
	"fmt"
	"testing"
)

type UserEntity struct {
	UserName string
	ID       string
}

// LIMIX_INFO 定义泛型类型
type CrudRepository[T any, ID string | int] struct {
	Element T
}

// LIMIX_INFO 定义泛型方法
func (repo *CrudRepository[T, ID]) Save(entity T) error {
	return nil
}

// LIMIX_INFO 定义泛型别名
type UserRepository = CrudRepository[UserEntity, string]

// 测试泛型类型
func TestTypingGeneric(t *testing.T) {
	TestTypingSpec(t)

	userRepo := UserRepository{
		Element: UserEntity{
			UserName: "limix",
			ID:       "limix",
		},
	}
	fmt.Printf("userRepo: %v\n", userRepo)
	userRepo.Save(userRepo.Element)
}
