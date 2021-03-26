package classroom

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/classroom"
)

const classroomURL = api.BASE_URL + "classroom/"

func CreateClassroom(P classroom.Classroom) (classroom.Classroom, error) {
	entity := classroom.Classroom{}
	resp, _ := api.Rest().SetBody(P).Post(classroomURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateClassroom(P classroom.Classroom) (classroom.Classroom, error) {
	entity := classroom.Classroom{}
	resp, _ := api.Rest().SetBody(P).Post(classroomURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadClassroom(id string) (classroom.Classroom, error) {
	entity := classroom.Classroom{}
	resp, _ := api.Rest().Get(classroomURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteClassroom(id string) (classroom.Classroom, error) {
	entity := classroom.Classroom{}
	resp, _ := api.Rest().Get(classroomURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadClassrooms() ([]classroom.Classroom, error) {
	entity := []classroom.Classroom{}
	resp, _ := api.Rest().Get(classroomURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
