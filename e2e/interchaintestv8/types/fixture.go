package types

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"strings"
)

type Fixture struct {
	Vkey         string `json:"vkey"`
	Proof        string `json:"proof"`
	PublicValues string `json:"publicValues"`
}

func (f *Fixture) DecodedProof() []byte {
	return decodeHex(f.Proof)
}

func (f *Fixture) DecodedPublicValues() []byte {
	return decodeHex(f.PublicValues)
}

func decodeHex(s string) []byte {
	stripped := strings.TrimPrefix(s, "0x")

	bz, err := hex.DecodeString(stripped)
	if err != nil {
		panic(err)
	}

	return bz
}

func GetPlonkFixture() Fixture {
	return getFixture("../../fixtures/plonk-fixture.json")
}

func GetGroth16Fixture() Fixture {
	return getFixture("../../fixtures/groth16-fixture.json")
}

func getFixture(path string) Fixture {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var fixture Fixture
	err = json.NewDecoder(file).Decode(&fixture)
	if err != nil {
		panic(err)
	}

	return fixture
}
