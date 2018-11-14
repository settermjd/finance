package banking

// Loan is a small struct to model a typical loan
type Loan struct {
    LoanTerm int
    RepaymentPeriod string
    paymentPeriodToAmount map[string]int
}

// NewLoan returns a new, initialised Loan struct
func NewLoan(loanTerm int, repaymentPeriod string) *Loan {
    var paymentPeriodToAmount = map[string]int {
        "weekly": 52,
        "fortnightly": 26,
        "monthly": 12,
        "quarterly": 4,
        "yearly": 1,
    }
    return &Loan {
        LoanTerm: loanTerm,
        RepaymentPeriod: repaymentPeriod,
        paymentPeriodToAmount: paymentPeriodToAmount,
    }
}

// GetRepaymentPeriods calculate and return the total payment periods for the
// term of the loan, based on the loan term and repayment period,
func (l *Loan) GetRepaymentPeriods() int {
    return l.LoanTerm * l.GetTotalPaymentsPerYear()
}

// GetTotalPaymentsPerYear returns the number of payments per year, based on
// the specified repayment period.
func (l *Loan) GetTotalPaymentsPerYear() int {
    return l.paymentPeriodToAmount[l.RepaymentPeriod]
}

