package dtos

type PresencaAluno struct {
	IdPresenca uint   `json:"idPresenca"`
    IdAluno    int   `json:"alunoId"`
    IdAula     int   `json:"aulaId"`
    Tipo       string `json:"tipo"`
}
