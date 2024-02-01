package cmd

import "fmt"

func printfBuilder(verbose bool) func(format string, a ...interface{}) {
	if verbose {
		return func(format string, a ...interface{}) {
			fmt.Printf(format, a...)
		}
	}
	return func(format string, a ...interface{}) {}
}

func printlnBuilder(verbose bool) func(a ...interface{}) {
	if verbose {
		return func(a ...interface{}) {
			fmt.Println(a...)
		}
	}
	return func(a ...interface{}) {}
}
