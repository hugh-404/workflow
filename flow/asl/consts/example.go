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
			"Next": "Choice"
		  },
		  "Choice": {
			"Type": "Choice",
			"Branches": [
			  {
				"Variable": "$$.Var2",
				"StringEquals": "test atext",
				"Next": "Fail"
			  },
			  {
				"Variable": "$$.Var1",
				"NumericEquals": 9,
				"Next": "Fail"
			  },
			  {
				"And": [
				  {
					"Variable": "$$.Var3.InnerS",
					"StringContains": "text"
				  }
				],
				"Next": "Success"
			  }
			]
		  },
		  "Success": {
			"Type": "Success",
			"End": true
		  },
		  "Fail": {
			"Type": "Success",
			"End": true
		  }
		}
	  }
	`
)