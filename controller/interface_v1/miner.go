package interface_v1

import "api/controller/task"

type apiMiner struct {
}

func GetMiner() *apiMiner {
	return new(apiMiner)
}

func (a *apiMiner) Start(threads int) bool {
	return task.GetMiner().Start(threads)
}

func (a *apiMiner) Stop() bool {
	return task.GetMiner().Stop()
}

func (a *apiMiner) SetBase(address string) bool {
	return task.GetMiner().Set(address)
}
