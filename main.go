package main 

import (	
	"crypto/sha256"	
	"encoding/hex"
	"encoding/json"
	"fmt"
)


type Block struct {
	Index 			int
	Data 			string
	PreviousHash	string
	Hash			string
}

func calculateHash(block Block)(string, error) {
	payload := struct {
		Index 			int
		Data 			string
		PreviousHash	string
	}{
		Index: 			block.Index,
		Data:			block.Data,
		PreviousHash:	block.PreviousHash,
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(encoded)
	return hex.EncodeToString(sum[:]), nil
}


func main () {
	block := Block {
		Index: 		0,
		Data: 		"Rindu rimbun rindiani",
		PreviousHash: "",
	}

	hash, err := calculateHash(block)
	if err != nil {
		fmt.Println("Gagal menghitung hash:", err)
		return
	}

	block.Hash = hash 

	fmt.Printf("Index : %d\n", block.Index)
	fmt.Printf("Data  : %s\n", block.Data)
	fmt.Printf("Hash  : %s\n", block.Hash)
}
