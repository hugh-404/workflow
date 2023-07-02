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
			  "Result": {"result":{"code":"200","description":"desc"}}
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