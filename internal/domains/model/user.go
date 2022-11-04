package model

import "github.com/matsunaga-project-2022/course-chat/internal/domains/valueobject"

type User struct {
	ID          string
	Name        string
	DisplayName string
	Icon        valueobject.Image
	Role        string
	Status      string
	IsOnline    bool
}

/*
ID (非公開)

名前

表示名

Email

アバター画像

役職

ステータス

オンラインかどうか
*/
