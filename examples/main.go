package main

import (
	"github.com/engmtcdrm/go-eggy"
	pp "github.com/engmtcdrm/go-prettyprint"

	"github.com/engmtcdrm/go-lager/examples/internal"
)

func main() {
	eggy.NewExamplePrompt(internal.AllExamples).
		Title(pp.Yellow("Examples of Lager Logging")).
		Show()
}
