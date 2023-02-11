package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	var data string
	var mapData = make(map[string]int)

	for i:=0 ; i<len(transactions); i++ {
		if transactions[i].Type == "expense" {
			mapData[transactions[i].Date] -= transactions[i].Amount
		} else {
			mapData[transactions[i].Date] += transactions[i].Amount
		}
	}

	keys := make([]string, 0, len(mapData))
	
	for k := range mapData {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if mapData[k] < 0 {
			data += fmt.Sprintf("%s;expense;%d\n", k, (mapData[k]) * -1)
		} else {
			data += fmt.Sprintf("%s;income;%d\n", k, mapData[k])
		}
	}

	if len(data) > 2 {
		data = data[:len(data)-1]
	}

	err := ioutil.WriteFile(path, []byte(data), 0644)
	if err != nil {
		return err
	}

	return nil
	// file, err := os.Create(path)
    // if err != nil {
    //     return err
    // }
    // defer file.Close()

    // w := bufio.NewWriter(file)
    // for _, line := range transactions {
    //     fmt.Fprintln(w, line)
    // }

    // return w.Flush()
	
	// return errors.New("not implemented") // TODO: replace this
}


func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"01/01/2021", "expense", 30000},
		{"01/01/2021", "income", 20000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
