package usecase

import (
	"req_parallel/heavyapiload"
)

type heavyApiLoad struct {
	heavyApiLoadRepository heavyapiload.HeavyApiLoadRepository
}

func NewHeavyApiLoadUseCase(heavyApiLoadRepository heavyapiload.HeavyApiLoadRepository) *heavyApiLoad {
	return &heavyApiLoad{heavyApiLoadRepository: heavyApiLoadRepository}
}

func (h *heavyApiLoad) Load() (string, error) {
	req, err := h.heavyApiLoadRepository.Test()

	return req, err
}
