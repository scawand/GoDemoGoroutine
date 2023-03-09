package module

import (
	"math/rand"
	"time"
)

type Action interface {
	LancerDe()
}

type Joueur struct {
	SesDe [1]De
	Nom   string
}

type De struct {
	ID     int
	NbFace int
}

type ApiDe struct {
	ID    int
	Value int
}

func (jr Joueur) LancerDe(nbFace int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(nbFace)
	return randomNumber
}

