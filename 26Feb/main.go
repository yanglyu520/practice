package main

import (
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
)

func main() {
	newRan, _ := uuid.NewRandom()
	fmt.Println(newRan)
	bl := []Bfmc{
		{
			ID:                pgtype.UUID{Bytes: newRan, Valid: true},
			CreatedAt:         1609459200, // Example Unix timestamp
			Name:              "Bfmc One",
			Abbreviation:      "BFMC1",
			ShapefileLocation: "/path/to/shapefile1",
		},
		{
			ID:                pgtype.UUID{Bytes: newRan, Valid: true},
			CreatedAt:         1609545600,
			Name:              "Bfmc Two",
			Abbreviation:      "BFMC2",
			ShapefileLocation: "/path/to/shapefile2",
		},
	}
	Bl := transformBFMCList(bl)

	fmt.Println(Bl)

}

type Bfmc struct {
	ID                pgtype.UUID
	CreatedAt         int64
	Name              string
	Abbreviation      string
	ShapefileLocation string
}

type BFMC struct {
	ID           string `json:"id"`
	CreatedAt    int64  `json:"createdAt"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

// encodeUUID converts a uuid byte array to UUID standard string form.
func encodeUUID(src [16]byte) string {
	var buf [36]byte

	hex.Encode(buf[0:8], src[:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], src[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], src[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], src[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], src[10:])

	return string(buf[:])
}

func transformBFMCList(bl []Bfmc) []BFMC {
	result := lo.Map(bl, func(b Bfmc, index int) BFMC {
		if !b.ID.Valid {
			return BFMC{}
		}
		return BFMC{
			ID:           string(b.ID.Bytes[:]),
			CreatedAt:    b.CreatedAt,
			Name:         b.Name,
			Abbreviation: b.Abbreviation,
		}
	})

	return result
}
