package model

type Projeto struct {
	Nome        string
	Id          ProjetoId
	Descricao   string
	DataInicio  string
	DataFim     string
	Status      string
	Responsavel UserId
}
