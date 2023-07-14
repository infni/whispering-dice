package exalted

// making GO provide a *float64 is strange .. .and hard.  this constant is a shim to acomplish that.
var IntegerOptionValueOne = 1.0
var IntegerOptionValueZero = 0.0

const (
	OptionPool               string = "pool"
	OptionComment            string = "comment"
	OptionAutomaticSuccesses string = "as"
	OptionDoubleSuccesses    string = "db"
	OptionRerollOnes         string = "re1"
	OptionRerollTwos         string = "re2"
	OptionRerollThrees       string = "re3"
	OptionRerollFours        string = "re4"
	OptionRerollFives        string = "re5"
	OptionRerollSixes        string = "re6"

	ResultSelectedSide string = "selectedside"
	ResultRolls        string = "rolls"
	ResultTotal        string = "total"
)
