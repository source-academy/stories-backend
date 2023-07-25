package userviews

import "github.com/source-academy/stories-backend/model"

type BatchCreateView struct {
	Count int64 `json:"createdUserCount"`
	IDs   []int `json:"createdUserIDs"`
}

// FIXME: users should ideally be a []model.User
func BatchCreateFrom(users []*model.User, count int64) BatchCreateView {
	ids := make([]int, len(users))
	for i, user := range users {
		ids[i] = int(user.ID)
	}
	return BatchCreateView{
		Count: count,
		IDs:   ids,
	}
}
