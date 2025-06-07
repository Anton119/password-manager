package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"main.go/account"
	"main.go/encrypter"
	"main.go/files"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": CreateAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("err env")
	}
	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())

Menu:
	for {

		variant := promtData(
			"1) Создать аккаунт",
			"2) Найти аккаунт по URL",
			"3) Найти аккаунт по логину",
			"4) Удалить аккаунт",
			"5) Выйти",
		)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}

		menuFunc(vault)
		/*switch variant {
		case "1":
			CreateAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}*/
	}

	CreateAccount(vault)
}

func CreateAccount(vault *account.VaultWithDb) {
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

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promtData("Введите URL")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outpurResults(&accounts)

}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promtData("Введите логин")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})

	outpurResults(&accounts)

}

func outpurResults(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		fmt.Println("Аккаунтов не найдено")
	}
	for _, account := range *accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promtData([]string{"Введите URL"})
	isDeleted := vault.DeleteAccount(url)
	if isDeleted {
		fmt.Println("Удалено")
	} else {
		fmt.Println("Не найдено")
	}
}

func promtData(prompt ...any) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
