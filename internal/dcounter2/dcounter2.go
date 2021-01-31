package dcounter2

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

// Attr ...
type Attr struct {
	ID    int
	Value int
}

// ReadDataset ...
func ReadDataset(filename string) (map[int]string, []Attr, []Attr, []Attr, []Attr) {

	f, _ := os.Open(filename)
	r := csv.NewReader(f)
	fmt.Println("start")

	fmt.Println("reading...")
	recs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("constructing index...")

	ni := map[int]string{}

	n := len(recs) - 1

	pcSlice := make([]Attr, n)
	cnSlice := make([]Attr, n)
	hiSlice := make([]Attr, n)
	piSlice := make([]Attr, n)

	for i, row := range recs {
		if i == 0 {
			continue
		}

		id, _ := strconv.Atoi(row[0])
		pc, _ := strconv.Atoi(row[2])
		cn, _ := strconv.Atoi(row[3])
		hi, _ := strconv.Atoi(row[4])

		pi, _ := strconv.ParseFloat(row[5], 64)
		pi = math.Trunc(pi)

		pcSlice[i-1] = Attr{
			ID:    id,
			Value: pc,
		}

		cnSlice[i-1] = Attr{
			ID:    id,
			Value: cn,
		}

		hiSlice[i-1] = Attr{
			ID:    id,
			Value: hi,
		}

		piSlice[i-1] = Attr{
			ID:    id,
			Value: int(pi),
		}

		ni[id] = row[1]

	}

	sort.Slice(pcSlice, func(i, j int) bool {
		return pcSlice[i].Value > pcSlice[j].Value
	})

	sort.Slice(cnSlice, func(i, j int) bool {
		return cnSlice[i].Value > cnSlice[j].Value
	})

	sort.Slice(hiSlice, func(i, j int) bool {
		return hiSlice[i].Value > hiSlice[j].Value
	})

	sort.Slice(piSlice, func(i, j int) bool {
		return piSlice[i].Value > piSlice[j].Value
	})

	return ni, pcSlice, cnSlice, hiSlice, piSlice
}

// IndexRow ...
type IndexRow struct {
	ID      int
	PCindex int
	CNindex int
	HIindex int
	PIindex int
}

// DCounter2 ...
type DCounter2 struct {
	NameIndex map[int]string
	PC        []Attr
	CN        []Attr
	HI        []Attr
	PI        []Attr
	RowIndex  map[int]IndexRow
}

// New ...
func New(filename string) *DCounter2 {

	nameIndex, a, b, c, d := ReadDataset(filename)

	rowIndex := map[int]IndexRow{}

	for i := 0; i < len(a); i++ {

		id := a[i].ID
		if _, ok := rowIndex[id]; !ok {
			rowIndex[id] = IndexRow{
				ID:      id,
				PCindex: i,
			}
		}

		r := rowIndex[id]
		r.PCindex = i
		rowIndex[id] = r

		//
		id = b[i].ID
		if _, ok := rowIndex[id]; !ok {
			rowIndex[id] = IndexRow{
				ID:      id,
				CNindex: i,
			}
		}
		r = rowIndex[id]
		r.CNindex = i
		rowIndex[id] = r

		//
		id = c[i].ID
		if _, ok := rowIndex[id]; !ok {
			rowIndex[id] = IndexRow{
				ID:      id,
				HIindex: i,
			}
		}
		r = rowIndex[id]
		r.HIindex = i
		rowIndex[id] = r

		//
		id = d[i].ID
		if _, ok := rowIndex[id]; !ok {
			rowIndex[id] = IndexRow{
				ID:      id,
				PIindex: i,
			}
		}
		r = rowIndex[id]
		r.PIindex = i
		rowIndex[id] = r

	}

	return &DCounter2{
		NameIndex: nameIndex,
		PC:        a,
		CN:        b,
		HI:        c,
		PI:        d,
		RowIndex:  rowIndex,
	}
}
