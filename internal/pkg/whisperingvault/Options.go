package whisperingvault

// making GO provide a *float64 is strange .. .and hard.  this constant is a shim to acomplish that.
var integerOptionValueTwo = 2.0
var integerOptionValueOne = 1.0
var integerOptionNegative99 = -99.0

const (
	OptionPool    string = "pool"
	OptionSkill   string = "skill"
	OptionCap     string = "cap"
	OptionComment string = "comment"

	ResultSelectedSide string = "selectedside"
	ResultRolls        string = "rolls"
	ResultTotal        string = "total"
)
