package main

import (
	"github.com/google/uuid"
	"math/rand"
	"sort"
)

type triangle struct {
	generation  int
	coordinates [6]float32
	power       float32
	color       int
	uuid        uuid.UUID
}

func genRandomTriangle() triangle {
	var t triangle
	t.uuid = uuid.New()
	for i := 0; i < 6; i++ {
		t.coordinates[i] = rand.Float32() * 400
	}
	temp1 := (t.coordinates[2] - t.coordinates[0]) * (t.coordinates[5] - t.coordinates[1])
	temp2 := (t.coordinates[4] - t.coordinates[0]) * (t.coordinates[3] - t.coordinates[1])

	t.power = 0.5 * abs(temp1-temp2)

	return t
}

func genRandomTriangles(cnt int) []triangle {
	r := make([]triangle, 0)
	for i := 0; i < cnt; i++ {
		t := genRandomTriangle()
		r = append(r, t)
	}
	return r
}

func sortAliveTs() {
	sort.Slice(aliveTs, func(i, j int) bool {
		return aliveTs[i].power > aliveTs[j].power
	})
}

func killLastTriangle() {
	if len(aliveTs) == 0 {
		return
	}

	deathTs = append(deathTs, aliveTs[len(aliveTs)-1])

	if len(aliveTs) > 0 {
		aliveTs = aliveTs[:len(aliveTs)-1]
	}
}

func addNewRandomTriangle() {
	aliveTs = append(aliveTs, genRandomTriangle())
}
