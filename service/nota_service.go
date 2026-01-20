package service

import (
	"ProjetoNota/model"
	"ProjetoNota/repository"
)

type NotaService struct {
	repository repository.NotaRepository
}

func NewNotaService(repository repository.NotaRepository) NotaService {
	return NotaService{
		repository: repository,
	}
}

func (ns *NotaService) GetNotas() ([]model.Nota, error) {
	return ns.repository.GetNotas()
}

func (ns *NotaService) AddNota(nota model.Nota) (model.Nota, error) {

	notaId, err := ns.repository.AddNota(nota)
	if err != nil {
		return model.Nota{}, err
	}
	nota.Id = notaId

	return nota, nil
}

func (ns *NotaService) GetNota(id int) (*model.Nota, error) {
	return ns.repository.GetNota(id)
}
