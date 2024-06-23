package converter

import "backend/model"

func UserToToken(Token string) *model.UserResponse {
	return &model.UserResponse{
		Token: Token,
	}
}
