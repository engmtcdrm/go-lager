module example.com/example

go 1.25

replace github.com/engmtcdrm/go-lager => ../

require github.com/engmtcdrm/go-lager v0.0.0-20251107030436-9d0c558b42b2

require (
	github.com/engmtcdrm/go-ansi v1.0.1
	github.com/engmtcdrm/go-pardon v0.0.0-20251015210019-f3dbc0f5b83a
	github.com/engmtcdrm/go-prettyprint v1.2.0
)

require (
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/term v0.34.0 // indirect
)
