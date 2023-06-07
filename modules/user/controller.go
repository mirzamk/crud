package user

import (
	"crud/dto"
)

type ControllerUser interface {
	CreateUser(req UserParam) (any, error)
	GetUserById(id uint) (FindUser, error)
	UpdateUser(req UserParam, id uint) (any, error)
	DeleteUser(email string) (any, error)
}

type controllerUser struct {
	useCaseUser UseCaseUser
}

func (uc controllerUser) CreateUser(req UserParam) (any, error) {
	user, err := uc.useCaseUser.CreateUser(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create User",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: UserParam{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		},
	}
	return res, err
}

func (uc controllerUser) GetUserById(id uint) (FindUser, error) {
	var res FindUser
	user, err := uc.useCaseUser.GetUserById(id)
	if err != nil {
		return FindUser{}, err
	}
	res.Data = user
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success Getting User",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, err
}

func (uc controllerUser) UpdateUser(req UserParam, id uint) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.useCaseUser.UpdateUser(req, id)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success update"
	res.MessageTitle = "update"

	return res, nil
}

func (uc controllerUser) DeleteUser(email string) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.useCaseUser.DeleteUser(email)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success Delete"
	res.MessageTitle = "Delete"

	return res, nil
}
