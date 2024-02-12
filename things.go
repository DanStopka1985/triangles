package main

import (
	"github.com/google/uuid"
	"math/rand"
	"sort"
)

var (
	aliveTs             []triangle
	deathTs             []triangle
	mutationShareChance = 1
	maxPopulation       = 10
)

type triangle struct {
	generation  int
	coordinates [6]float32
	power       float32
	color       int
	uuid        uuid.UUID
	haveMutagen bool
}

func createTriangleChild(p triangle) triangle {
begin:
	r := p
	r.haveMutagen = false
	var genMutantIndex = rand.Intn(6)
	if genMutantIndex == 0 {

	}
	for i := 0; i < 6; i++ {
		if i != genMutantIndex {
			r.coordinates[i] = p.coordinates[i]
		} else if rand.Intn(mutationShareChance) == 0 {
			r.haveMutagen = true
			if rand.Intn(1) == 0 { //random mutation delta
				r.coordinates[i] = p.coordinates[i] + rand.Float32()*5
			} else {
				r.coordinates[i] = p.coordinates[i] - rand.Float32()*5
			}
			if r.coordinates[i] < 0 || r.coordinates[i] > 400 {
				goto begin //if exit from window range - try new mutation
			}
		}
	}
	r.generation = p.generation + 1
	r.power = getPower(r)

	return r
}

func getPower(t triangle) float32 {
	return 0.5 * abs((t.coordinates[2]-t.coordinates[0])*(t.coordinates[5]-t.coordinates[1])-
		(t.coordinates[4]-t.coordinates[0])*(t.coordinates[3]-t.coordinates[1]))
}

func createNewGeneration() {
	sortAliveTs()
	cnt := len(aliveTs)
	for i := 0; i < cnt; i++ {
		aliveTs = append(aliveTs, createTriangleChild(aliveTs[i]))

	}
}

func genRandomTriangle() triangle {
	var t triangle
	t.uuid = uuid.New()
	for i := 0; i < 6; i++ {
		t.coordinates[i] = rand.Float32() * 400
	}

	t.power = 0.5 * abs((t.coordinates[2]-t.coordinates[0])*(t.coordinates[5]-t.coordinates[1])-
		(t.coordinates[4]-t.coordinates[0])*(t.coordinates[3]-t.coordinates[1]))
	t.generation = 0

	return t
}

/*func genRandomTriangles(cnt int) []triangle {
	r := make([]triangle, 0)
	for i := 0; i < cnt; i++ {
		t := genRandomTriangle()
		r = append(r, t)
	}
	return r
}*/

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

func naturalSelection() {
	if len(aliveTs) <= maxPopulation {
		return
	}

	sortAliveTs()
	deathTs = append(deathTs, aliveTs[maxPopulation+1:]...)
	aliveTs = aliveTs[:maxPopulation]

}

func addNewRandomTriangle() {
	aliveTs = append(aliveTs, genRandomTriangle())
}
