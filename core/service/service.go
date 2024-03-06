package service

type serviceGroup struct {
	UserService
}

var Group = new(serviceGroup)
