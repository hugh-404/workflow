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
				"Variable": "test text",
				"StringEquals": "test text",
				"Next": "Success"
			  },
			  {
				"And": [
				  {
					"Variable": "test text",
					"StringEndWith": "text"
				  }
				],
				"Next": "Success"
			  }
			]
		  },
		  "Success": {
			"Type": "Success",
			"End": true
		  }
		}
	  }
	`
)