package cache

import "time"

//set token
func SetToken(tokenstr string,id int32) error  {
	if err:=rdb.Set(ctx,tokenstr,id,72*time.Hour).Err();err!=nil{
		return err
	}
	return nil
}

//get token(通过token获取id)
func GetToken(tokenstr string) (id string,err error) {
	return rdb.Get(ctx,tokenstr).Result()
}

// 删除用户登录token实现登出
func DelUserToken(tokenstr string) error {
	return rdb.Del(ctx,tokenstr).Err()
}