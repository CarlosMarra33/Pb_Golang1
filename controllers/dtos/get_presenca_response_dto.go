package dtos

type GetPresencaAula struct {
	AlunoId    uint   `json:"alunoId"`
	AulaId     uint   `json:"aulaId"`
	Tipo       string `json:"tipoPresenca"`
	DataCreate int    `json:"created"`
	DataUpdate int    `json:"updated"`
}
