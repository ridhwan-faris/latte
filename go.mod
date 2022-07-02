module github.com/ridhwan-faris/latte

go 1.18

require github.com/go-playground/validator/v10 v10.11.0

require (
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/sys v0.0.0-20220701225701-179beb0bd1a1 // indirect
	golang.org/x/text v0.3.7 // indirect
)

retract (
	v1.0.4 // Contains retractions only.
	v1.0.0 // Published accidentally.
)
