package main

import pb "github.com/nabin3/userInfo/proto"

type UserID struct {
	UserId string `json:"user_id"`
}

func generatedUserIdToResponseUserId(user_id *pb.UserID) UserID {
	return UserID{
		UserId: user_id.Id,
	}
}

type User struct {
	Id        string  `json:"id"`
	Fname     string  `json:"fname"`
	City      string  `json:"city"`
	Phone     string  `json:"phone"`
	Height    float32 `json:"height"`
	IsMarried bool    `json:"is_married"`
}

func retrievedUserToResponseUser(user *pb.User) User {
	return User{
		Id:        user.Id,
		Fname:     user.Fname,
		City:      user.City,
		Phone:     user.Phone,
		Height:    user.Height,
		IsMarried: user.Ismarried,
	}
}

func retrievedUserListToResponseUserList(userList *pb.UserList) []User {
	user_list := make([]User, 0)

	for _, each_user := range userList.Users {
		user_list = append(user_list, retrievedUserToResponseUser(each_user))
	}

	return user_list
}
