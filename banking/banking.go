package banking

// Loan is a small struct to model a typical loan
type Loan struct {
    LoanTerm int
    RepaymentPeriod string
}

// GetRepaymentPeriods calculate and return the total payment periods for the
// term of the loan, based on the loan term and repayment period,
func (l *Loan) GetRepaymentPeriods() int {
    return l.LoanTerm * l.GetTotalPaymentsPerYear()
}

// GetTotalPaymentsPerYear calculates the number of payments per year, based on
// the repayment period.
func (l *Loan) GetTotalPaymentsPerYear() int {
    var repaymentsPerYear int

    switch l.RepaymentPeriod {
        case "fortnightly": repaymentsPerYear = 26
        case "monthly": repaymentsPerYear = 12
        case "quarterly": repaymentsPerYear = 4
        case "weekly": repaymentsPerYear = 52
        case "yearly": repaymentsPerYear = 1
    }

    return repaymentsPerYear
}

