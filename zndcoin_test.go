package main

import (
	"encoding/json"
	//"github.com/go-test/deep"
	"testing"
	"time"
)

func TestBlock(t *testing.T) {
	testBlock := NewBlockchain().
		NewBlock([]byte("proof"), "prevhash", Timestamp(time.Now().UnixNano()))
	jsonData := testBlock.
		ToJSON()

	var block Block
	json.Unmarshal(jsonData, &block)

	//if diff := deep.Equal(*testBlock, block); diff != nil {
	//	t.Error(diff)
	//}
}

func TestValidProofHash_False(t *testing.T) {
	f, n := ValidProofHash([8]ProofHash{"Zunhoge", "zUnhoge", "zunhoge", "zunhoGe", "dOKohoge", "Kihoge", "yOhoge", "shyhoge"})
	if f {
		t.Errorf("False positive of ValidProof: num %d", n)
	}
}

func TestValidProofHash_True0(t *testing.T) {
	f, n := ValidProofHash([8]ProofHash{"Zunhoge", "zUnhoge", "zunhoge", "zunhoGe", "dOKohoge", "Kihoge", "yOhoge", "shIhoge"})
	if !f {
		t.Error("False negative of ValidProof")
	}
	t.Logf("num %d", n)
}

func TestValidProofHash_True1(t *testing.T) {
	f, n := ValidProofHash([8]ProofHash{"zUnhoge", "zunhoGe", "dOKohoge", "Kihoge", "yOhoge", "shIhoge", "zunhoge", "Zunhoge"})
	if !f {
		t.Error("False negative of ValidProof")
	}
	t.Logf("num %d", n)
}

func TestCreateZndkValidFunc_True(t *testing.T) {
	f := CreateZndkValidFunc("zun")
	if !f([]byte("Zunhogehoge")) {
		t.Error("False negative of ValidProof")
	}
}

func TestCreateZndkValidFunc_False(t *testing.T) {
	f := CreateZndkValidFunc("zun")
	if f([]byte("Zinhogehoge")) {
		t.Error("False negative of ValidProof")
	}
}
