package dtos

type Aula_dto struct {
	Materia     string `json:"materia"`
	ProfessorId uint   `json:"professor_id"`
	Alunos      []uint `json:"alunos"`
}
