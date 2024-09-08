package functions

import (
	"code/database"
	"code/repository"
	"fmt"
)

func AuthLogin(token string) (err error, username string) {
	fmt.Println(token)

	res, err := DecodeJWT(token)
	if err != nil {
		return err, ""
	}

	if res["isLogin"] == false {
		fmt.Println("user is not login")
		return fmt.Errorf("authorization not good"), ""
	}

	fmt.Println(res["data"])

	data := res["data"].(map[string]interface{})

	fmt.Println(data)

	err = repository.KeepLogin(database.DbConnection, data["username"].(string))

	if err != nil {
		fmt.Println("Keep login error")
		return fmt.Errorf("authorization not good"), ""
	}

	return nil, data["username"].(string)
}
