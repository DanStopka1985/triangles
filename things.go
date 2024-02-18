package main

import (
	"github.com/google/uuid"
	"image/color"
	"math"
	"math/rand"
	"sort"
	"time"
)

var (
	aliveTs             []triangle
	deathTs             []triangle
	mutationShareChance         = 100
	maxPopulation               = 1000
	evolutionSpeed              = time.Millisecond * 3
	side                float32 = 400
	startTsCnt                  = 10
	startTsCntDefault           = 10
)

type triangle struct {
	generation  int
	genes       [6]float32
	power       float64
	color       color.NRGBA
	uuid        uuid.UUID
	haveMutagen bool
}

func createTriangleChild(p triangle) triangle {
begin:
	r := p
	r.haveMutagen = false
	var mutagenIx = rand.Intn(6)
	if mutagenIx == 0 {

	}
	for i := 0; i < 6; i++ {
		if i != mutagenIx {
			r.genes[i] = p.genes[i]
		} else if rand.Intn(mutationShareChance) == 0 {
			r.haveMutagen = true
			if rand.Intn(2) == 1 { //random mutation delta
				r.genes[i] = p.genes[i] + rand.Float32()*10
			} else {
				r.genes[i] = p.genes[i] - rand.Float32()*10
			}
			if r.genes[i] < 0 || r.genes[i] > side {
				goto begin //if exit from window range - try new mutation
			}
		}
	}
	r.generation = p.generation + 1
	r.power = getPower(r)

	return r
}

func getPower(t triangle) float64 {
	return 0.5 * float64(math.Abs((float64(t.genes[2])-float64(t.genes[0]))*(float64(t.genes[5])-float64(t.genes[1]))-
		(float64(t.genes[4])-float64(t.genes[0]))*(float64(t.genes[3])-float64(t.genes[1]))))
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
		t.genes[i] = rand.Float32() * 400
	}

	t.power = getPower(t)
	t.generation = 0
	t.color = _color[rand.Intn(6)]

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

//func killLastTriangle() {
//	if len(aliveTs) == 0 {
//		return
//	}
//
//	deathTs = append(deathTs, aliveTs[len(aliveTs)-1])
//
//	if len(aliveTs) > 0 {
//		aliveTs = aliveTs[:len(aliveTs)-1]
//	}
//}

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
