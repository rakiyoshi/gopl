package bank

var deposits = make(chan int) // send the amount of deposits
var balances = make(chan int) // receive the balances

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func teller() {
	var balance int // balance is closed inside teller go routine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start monitor go routine
}
