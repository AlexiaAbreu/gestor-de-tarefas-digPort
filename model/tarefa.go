package model

type Tarefa struct {
	Nome        string
	Id          string
	Descricao   string
	DataInicio  string
	DataFim     string
	Status      string
	Prioridade  string
	ProjetoID   ProjetoId
	Responsavel UserId
}
