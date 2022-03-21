package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type ExRand struct {
	seed int64
	*rand.Rand
}

func NewExRand(seed int) *ExRand {
	s := int64(seed)
	return &ExRand{
		seed: s,
		Rand: rand.New(rand.NewSource(s)),
	}
}

func Import(path string, v interface{}) {
	f, err := os.Open(path)
	if err != nil {
		panic("can't import " + path)
	}

	json.NewDecoder(f).Decode(v)
}

func Export(name string, v interface{}) {
	filename := fmt.Sprintf("%[1]ss/%[2]v-%[1]s.json", name, time.Now().UnixMilli())
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")
	if err = enc.Encode(v); err != nil {
		panic(err)
	}
}

func (e *ExRand) Seed(_ int64) {
	panic("you cannot re-seed an exrand")
}

func (e *ExRand) UnmarshalJSON(bytes []byte) error {
	var seed int64
	if err := json.Unmarshal(bytes, &seed); err != nil {
		return err
	}
	e.Rand.Seed(seed)
	return nil
}

func (e *ExRand) MarshalJSON() ([]byte, error) {
	if b, err := json.Marshal(e.seed); err != nil {
		return nil, err
	} else {
		return b, nil
	}
}

func (e *ExRand) Uniform(a, b float64) float64 {
	return a + (b-a)*rand.Float64()
}

func (e *ExRand) IntRange(a, b int) int {
	return a + rand.Intn(b-a)
}

func (e *ExRand) Roll(chance float64) bool {
	if chance <= 0 {
		return false
	}
	if chance >= 1 {
		return true
	}

	return e.Float64() < chance
}
