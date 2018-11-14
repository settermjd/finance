package banking

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLoanCalculation(t *testing.T) {
    loanTerm := 30

	Convey("Given a loan term of 30 years and a monthly repayment period", t, func() {
        loan := NewLoan(loanTerm, "monthly")

        Convey("There should be 360 payments in total", func() {
            So(loan.GetRepaymentPeriods(), ShouldEqual, 360)
        })
	})

	Convey("Given a loan term of 30 years and a weekly repayment period", t, func() {
        loan := NewLoan(loanTerm, "weekly")

		Convey("There should be 52 payments per year", func() {
		    So(loan.GetTotalPaymentsPerYear(), ShouldEqual, 52)
		})
	})

	Convey("Given a loan term of 30 years and a fortnightly repayment period", t, func() {
        loan := NewLoan(loanTerm, "fortnightly")

		Convey("There should be 26 payments per year", func() {
		    So(loan.GetTotalPaymentsPerYear(), ShouldEqual, 26)
		})
	})

	Convey("Given a loan term of 30 years and a quarterly repayment period", t, func() {
        loan := NewLoan(loanTerm, "quarterly")

		Convey("There should be 4 payments per year", func() {
		    So(loan.GetTotalPaymentsPerYear(), ShouldEqual, 4)
		})
	})

	Convey("Given a loan term of 30 years and a yearly repayment period", t, func() {
        loan := NewLoan(loanTerm, "yearly")

		Convey("There should be 1 payment per year", func() {
		    So(loan.GetTotalPaymentsPerYear(), ShouldEqual, 1)
		})
	})
}
