package main

import (
	"encoding/base64"
	"github.com/go-test/deep"
	"testing"
	"time"
)

func TestBlock(t *testing.T) {
	testBlock := NewBlockchain().
		NewBlock([]byte("proof"), "prevhash", Timestamp(time.Now().UnixNano()))
	jsonData := testBlock.
		ToJSON()

	var block Block
	block.ParseJSON(jsonData)

	if diff := deep.Equal(*testBlock, block); diff != nil {
		t.Error(diff)
	}
}

func TestValidProofHash_False(t *testing.T) {
	phsEncoded := [8]string{"Zunhoge", "zUnhoge", "zunhoge", "zunhoGe", "dOKohoge", "Kihoge", "yOhoge", "shyhoge"}
	var phs [8]ProofHash
	for i, value := range phsEncoded {
		phs[i], _ = base64.StdEncoding.DecodeString(value)
	}
	f, n := ValidProofHash(phs)
	if f {
		t.Errorf("False positive of ValidProof: num %d", n)
	}
}

func TestValidProofHash_True0(t *testing.T) {
	phsEncoded := [8]string{"Zunhoge", "zUnhoge", "zunhoge", "zunhoGe", "dOKohoge", "Kihoge", "yOhoge", "shIhoge"}
	var phs [8]ProofHash
	for i, value := range phsEncoded {
		phs[i], _ = base64.StdEncoding.DecodeString(value)
	}
	f, n := ValidProofHash(phs)
	if !f {
		t.Error("False negative of ValidProof")
	}
	t.Logf("num %d", n)
}

func TestValidProofHash_True1(t *testing.T) {
	phsEncoded := [8]string{"zUnhoge", "zunhoGe", "dOKohoge", "Kihoge", "yOhoge", "shIhoge", "zunhoge", "Zunhoge"}
	var phs [8]ProofHash
	for i, value := range phsEncoded {
		phs[i], _ = base64.StdEncoding.DecodeString(value)
	}
	f, n := ValidProofHash(phs)
	if !f {
		t.Error("False negative of ValidProof")
	}
	t.Logf("num %d", n)
}

func TestCreateZndkValidFunc_True(t *testing.T) {
	f := CreateZndkValidFunc("zun")
	if !f("Zunhogehoge") {
		t.Error("False negative of ValidProof")
	}
}

func TestCreateZndkValidFunc_False(t *testing.T) {
	f := CreateZndkValidFunc("zun")
	if f("Zinhogehoge") {
		t.Error("False negative of ValidProof")
	}
}

func TestCreateZndkValidFunc_Null(t *testing.T) {
	f := CreateZndkValidFunc("zun")
	if f("") {
		t.Error("False negative of ValidProof")
	}
}

func TestHash(t *testing.T) {
	if diff := deep.Equal(Hash(Proof([]byte("proof1")),
		Proof([]byte("lastProof"))),
		Hash(Proof([]byte("proof2")),
			Proof([]byte("lastProof")))); diff == nil {
		t.Error("TestHash is invalid: Hashes which must not be matched are matched")
	}
}
