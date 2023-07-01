package consts

const (
	SimpleResultAsl = `
	{
		"Comment": "A simple minimal example of the States language",
		"StartAt": "Result",
		"States": {
		  "Result": {
			"Type": "Result",
			"Param": {
			  "Result": {
				"A": "ResultA",
				"B": "ResultB"
			  }
			},
			"Next": "Success"
		  },
		  "Success": {
			"Type": "Success",
			"End": true
		  }
		}
	  }
	`
)