package example

import "time"

//go:generate go run ../../../ error go_definition -t . -o . --codes Unknown,NotFound,BadRequest

type WorkKind string

type WorkingTime struct {
	WorkKind WorkKind
	StartAt  time.Time
	EndAt    time.Time
}

type WorkingTimeSpec struct {
	allowedWorkKinds []WorkKind
}

func (s WorkingTimeSpec) Validate(workingTime WorkingTime) Error {
	if s.IsAllowedWorkKind(workingTime.WorkKind) {
		//errcode InvalidWorkKind,workKind WorkKind
		return nil // TODO
	}

	//errcode InvalidWorkingTime,startAt int64,endAt int64
	return nil
}

func (s WorkingTimeSpec) IsAllowedWorkKind(workKind WorkKind) bool {
	for _, w := range s.allowedWorkKinds {
		if w == workKind {
			return true
		}
	}
	return false
}
