package ex01

type WithdrawResult struct {
	amount int
	ok     bool
}

var deposits = make(chan int) // send the amount of deposits
var balances = make(chan int) // receive the balances
var withdrawResult = make(chan WithdrawResult)

func Deposit(amount int) { deposits <- amount }

func Withdraw(amount int) (withdrawAmount int, ok bool) {
	deposits <- amount
	result := <-withdrawResult
	return result.amount, result.ok
}

func Balance() int { return <-balances }

func teller() {
	var balance int // balance is closed inside teller go routine
	for {
		select {
		case amount := <-deposits:
			if amount < 0 {
				if balance+amount < 0 {
					withdrawResult <- WithdrawResult{0, false}
				} else {
					withdrawResult <- WithdrawResult{balance + amount, true}
				}
			}
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start monitor go routine
}
