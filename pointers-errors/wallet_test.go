package pointerserrors

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBitcoin(t *testing.T) {
	Convey("Should print Bitcoin in format '<number> BTC'", t, func() {
		got := Bitcoin(10).String()

		So(got, ShouldEqual, "10 BTC")
	})
}

func TestWallet(t *testing.T) {

	Convey("Should deposit amount to Wallet", t, func() {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	Convey("Should withdraw amount from Wallet", t, func() {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(10))
		So(err, ShouldEqual, nil)
		assertBalance(t, wallet, Bitcoin(10))
	})

	Convey("Should check for sufficient funds before withdraw", t, func() {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(Bitcoin(20))
		So(err, ShouldBeError, ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, expectedBalance Bitcoin) {
	So(wallet.Balance(), ShouldEqual, expectedBalance)
}
