package cmd

import (
	"os"

	e "github.com/CHW0218/goloxy/errors"
	"github.com/CHW0218/goloxy/vm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func App() (app *cobra.Command) {
	app = &cobra.Command{
		Use:   "golox [FILE]",
		Args:  cobra.MaximumNArgs(1),
		Short: "golox: A Lox interpreter in Go.",
	}
	app.Flags().SortFlags = true

	defaultVerbosityStr := "INFO"
	verbosity := app.Flags().StringP("verbosity", "v", defaultVerbosityStr, "logging verbosity")

	app.Run = func(_ *cobra.Command, args []string) {
		verbosityLvl, err := logrus.ParseLevel(*verbosity)
		if err != nil {
			verbosityLvl, _ = logrus.ParseLevel(defaultVerbosityStr)
		}
		logrus.SetLevel(verbosityLvl)
		logrus.SetFormatter(&easy.Formatter{LogFormat: "%lvl% %msg%\n"})

		if err := appMain(args); err != nil {
			logrus.Fatalln(err)
			os.Exit(1)
		}
	}
	return
}

func appMain(args []string) error {
	vm_ := vm.NewVM()

	switch len(args) {
	case 0:
		return vm_.REPL()
	case 1:
		src, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}
		_, err = vm_.Interpret(string(src), false)
		if err != nil {
			return err
		}
	default:
		panic(e.Unreachable)
	}
	return nil
}
