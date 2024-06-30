package cmd_manager

import "fmt"

type CMDManager struct{

} 


func(cmd *CMDManager) ReadLines()([]string, error){
	var prices []string
	
	for{
		fmt.Print("Please enter price: ")
		var price string
		fmt.Scan(&price)
		if price == "exit"{
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func(cmd *CMDManager) WriteLines(data interface{}) (error){
	fmt.Println(data)
	return nil
}

func New()*CMDManager{
	return &CMDManager{}
}