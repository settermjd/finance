package banking

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLoanCalculation(t *testing.T) {

    loan := new(Loan)

	Convey("Given a loan term of 30 years", t, func() {
		loan.LoanTerm = 30

		Convey("And a monthly repayment period", func() {
			loan.RepaymentPeriod = "monthly"

			Convey("There should be 360 payments in total", func() {
			    So(loan.GetRepaymentPeriods(), ShouldEqual, 360)
			})
		})
	})

	Convey("Given a monthly repayment period", t, func() {
		loan.RepaymentPeriod = "monthly"

		Convey("There should be 12 payments per year", func() {
		    So(loan.GetTotalPaymentsPerYear(), ShouldEqual, 12)
		})
	})

	Convey("Given a weekly repayment period", t, func() {
		loan.RepaymentPeriod = "weekly"

		Convey("There should be 52 payments per year", func() {
		    So(loan.GetTotalPaymentsPerYear(), ShouldEqual, 52)
		})
	})

	Convey("Given a quarterly repayment period", t, func() {
		loan.RepaymentPeriod = "quarterly"

		Convey("There should be 4 payments per year", func() {
		    So(loan.GetTotalPaymentsPerYear(), ShouldEqual, 4)
		})
	})

	Convey("Given a yearly repayment period", t, func() {
		loan.RepaymentPeriod = "yearly"

		Convey("There should be 1 payment per year", func() {
		    So(loan.GetTotalPaymentsPerYear(), ShouldEqual, 1)
		})
	})
}
