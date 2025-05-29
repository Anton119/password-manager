package main

import (
	"fmt"

	"main.go/account"
)

func main() {
	vault, _ := account.NewVault()
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			CreateAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}

	CreateAccount(vault)
}

func getMenu() int {
	var variant int

	fmt.Println("1) Создать аккаунт")
	fmt.Println("2) Найти аккаунт")
	fmt.Println("3) Удалить аккаунт")
	fmt.Println("4) Выйти")
	fmt.Scan(&variant)
	return variant

}

func CreateAccount(vault *account.Vault) {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите url")
	url = "http://" + url

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или login")
		return
	}

	vault.AddAccount(*myAccount)

}

func findAccount(vault *account.Vault) {
	url := promtData("Введите URL")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		fmt.Println("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}

}

func deleteAccount(vault *account.Vault) {
	url := promtData("Введите URL")
	isDeleted := vault.DeleteAccount(url)
	if isDeleted {
		fmt.Println("Удалено")
	} else {
		fmt.Println("Не найдено")
	}
}

func promtData(promt string) string {
	fmt.Print(promt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
