module github.com/lemolatoon/omikuji

go 1.21.0

require (
	github.com/lemolatoon/core v0.0.0-00010101000000-000000000000
	github.com/lemolatoon/discord v0.0.0-00010101000000-000000000000
)

replace (
	github.com/lemolatoon/core => ./core
	github.com/lemolatoon/discord => ./discord
)
