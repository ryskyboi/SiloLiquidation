package querySilos

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"example.com/m/src/check"
)

func visit(path string, di fs.DirEntry, err error) error {
	fmt.Printf("Visited: %s\n", path)
	if err != nil {
		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		return err
	}

	if di.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Printf("failed to read dir: %v\n", err)
			return err
		}

		jsonFile, err := CreateJson(fmt.Sprintf("%v/state.json", path))
		check.Check(err)
		defer jsonFile.Close()
		make_borrow := false
		make_repay := false
		make_deposit := false
		var borrow, deposit, repay []Transaction
		for _, file := range files {
			if !file.IsDir() {
				if file.Name() == "borrow" {
					borrow = fileHandler(path, file)
					make_borrow = true
				} else if file.Name() == "repay" {
					repay = fileHandler(path, file)
					make_repay = true
				} else if file.Name() == "deposit" {
					deposit = fileHandler(path, file)
					make_deposit = true
				}
			}
		}
		if make_borrow && make_repay && make_deposit {
			state := siloState(borrow, repay, deposit)
			writeToJsonFile(state, jsonFile)
		}
	}

	return nil
}

func fileHandler(path string, file fs.DirEntry) []Transaction {
	name := file.Name()
	fileData, err := os.ReadFile(fmt.Sprintf("%v/%v", path, name))
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err)
		return []Transaction{}
	}
	var data []Transaction
	json.Unmarshal(fileData, &data)
	return data
}

func siloState(borrow []Transaction, repay []Transaction, deposit []Transaction) Account {
	accounts := make(Account, 0)
	for len(borrow) > 0 || len(repay) > 0 || len(deposit) > 0 {
		next_repay := _handleTx(repay)
		next_borrow := _handleTx(borrow)
		next_deposit := _handleTx(deposit)
		if min(next_repay.Block, next_borrow.Block, next_deposit.Block) == next_deposit.Block {
			accounts = addDepositTx(next_deposit, accounts)
			deposit = decrementTx(deposit)
		}
		if min(next_repay.Block, next_borrow.Block, next_deposit.Block) == next_borrow.Block {
			accounts = addBorrowTx(next_borrow, accounts)
			borrow = decrementTx(borrow)
		}
		if min(next_repay.Block, next_borrow.Block, next_deposit.Block) == next_repay.Block {
			accounts = addRepayTx(next_repay, accounts)
			repay = decrementTx(repay)
		}
	}
	return accounts
}

func decrementTx(tx []Transaction) []Transaction {
	if len(tx) > 0 {
		return tx[1:]
	}
	return tx
}

func _handleTx(tx []Transaction) Transaction {
	if len(tx) > 0 {
		return tx[0]
	}
	return Transaction{Block: uint64(2 ^ 63)}
}

func addBorrowTx(borrow Transaction, accounts Account) Account {
	if _, ok := accounts[borrow.User]; ok {
		accounts[borrow.User].Owe[borrow.Asset] += borrow.Amount
	} else {
		accounts[borrow.User] = Contents{
			Address: borrow.User,
			Owe: Balance{
				borrow.Asset: borrow.Amount,
			},
			Own: Balance{},
		}
	}
	return accounts
}

func addDepositTx(deposit Transaction, accounts Account) Account {
	if _, ok := accounts[deposit.User]; ok {
		accounts[deposit.User].Own[deposit.Asset] += deposit.Amount
	} else {
		accounts[deposit.User] = Contents{
			Address: deposit.User,
			Own: Balance{
				deposit.Asset: deposit.Amount,
			},
			Owe: Balance{},
		}
	}
	return accounts
}

func addRepayTx(repay Transaction, accounts Account) Account {
	if _, ok := accounts[repay.User]; ok {
		accounts[repay.User].Owe[repay.Asset] -= repay.Amount
	} else {
		panic("Repay transaction without borrow")
	}
	return accounts
}

func ProduceState() {
	root := "siloState"
	err := filepath.WalkDir(root, visit)
	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", root, err)
	}
}
