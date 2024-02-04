package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/wire"
)

func decodeTransactionHex(transactionHex string) (*wire.MsgTx, error) {
	// Decode the transaction hex
	serializedTx, err := hex.DecodeString(transactionHex)
	if err != nil {
		return nil, err
	}

	// Deserialize the transaction
	var decodedTransaction wire.MsgTx
	err = decodedTransaction.Deserialize(bytes.NewReader(serializedTx))
	if err != nil {
		return nil, err
	}

	return &decodedTransaction, nil
}

func main() {
	transactionHex := "020000000001010ccc140e766b5dbc884ea2d780c5e91e4eb77597ae64288a42575228b79e234900000000000000000002bd37060000000000225120245091249f4f29d30820e5f36e1e5d477dc3386144220bd6f35839e94de4b9cae81c00000000000016001416d31d7632aa17b3b316b813c0a3177f5b6150200140838a1f0f1ee607b54abf0a3f55792f6f8d09c3eb7a9fa46cd4976f2137ca2e3f4a901e314e1b827c3332d7e1865ffe1d7ff5f5d7576a9000f354487a09de44cd00000000"

	// Decode transaction
	decodedTransaction, err := decodeTransactionHex(transactionHex)
	if err != nil {
		log.Fatal("Error decoding transaction hex:", err)
	}

	printTransactionDetails(decodedTransaction)
}

func printTransactionDetails(decodedTransaction *wire.MsgTx) {
	fmt.Println("=========================================================================== Transaction Details ========================================================================")
	fmt.Printf("Version: %d\n", decodedTransaction.Version)
	fmt.Printf("Locktime: %d\n", decodedTransaction.LockTime)

	printInputs(decodedTransaction.TxIn)
	printOutputs(decodedTransaction.TxOut)

}

func printInputs(inputs []*wire.TxIn) {
	fmt.Println("==================================================================================== Inputs ================================================================================")
	for _, txin := range inputs {
		fmt.Printf("\tTransaction ID: %s\n", txin.PreviousOutPoint.Hash)
		fmt.Printf("\tOutput Index: %d\n", txin.PreviousOutPoint.Index)
		fmt.Printf("\tScript Sig: %s\n", hex.EncodeToString(txin.SignatureScript))
		fmt.Printf("\tSequence: %d\n", txin.Sequence)
		fmt.Println()
	}
}

func printOutputs(outputs []*wire.TxOut) {
	fmt.Println("================================================================================= Outputs =================================================================================")
	for i, txout := range outputs {
		// Convert satoshis to bitcoins
		value := float64(txout.Value) / 1e8
		fmt.Printf("\tOutput %d:\n", i)
		fmt.Printf("\t\tValue: %.8f BTC\n", value)
		fmt.Printf("\t\tScript Pub Key: %s\n", hex.EncodeToString(txout.PkScript))
		fmt.Println()
	}
}
