module github.com/patha2702/moduleExample

go 1.21.1

require (
	github.com/patha2702/cryptit v0.0.0-00010101000000-000000000000
	github.com/pborman/uuid v1.2.1
)

require github.com/google/uuid v1.0.0 // indirect

replace github.com/patha2702/cryptit => ../cryptit
