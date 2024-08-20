module Tours

go 1.20

require (
	logic v1.0.0
)

replace (
	logic v1.0.0 => ..src/FreeTraining/Tours/logic
)