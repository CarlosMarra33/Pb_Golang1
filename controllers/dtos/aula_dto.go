package dtos

type Aula_dto struct {
	Materia     string `json:"materia"`
	ProfessorId uint   `json:"student"`
	Alunos      []uint `json:"alunos"`
}
