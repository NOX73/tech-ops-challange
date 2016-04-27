package main

import (
	"errors"
	"fmt"

	"github.com/NOX73/tech-ops-challenge/db"
	"github.com/abiosoft/ishell"
)

var (
	ArgumentsError = errors.New("Arguments Error")
)

func main() {
	shell := ishell.New()
	shell.ShowPrompt(false)
	base := db.New()

	shell.Register("SET", func(args ...string) (string, error) {
		if len(args) < 2 {
			return "", ArgumentsError
		}
		base.Set(args[0], args[1])
		return "", nil
	})

	shell.Register("GET", func(args ...string) (string, error) {
		if len(args) < 1 {
			return "", ArgumentsError
		}
		val, _ := base.Get(args[0])
		return val, nil
	})

	shell.Register("UNSET", func(args ...string) (string, error) {
		if len(args) < 1 {
			return "", ArgumentsError
		}
		base.Unset(args[0])
		return "", nil
	})

	shell.Register("NUMEQUALTO", func(args ...string) (string, error) {
		if len(args) < 1 {
			return "", ArgumentsError
		}
		count := base.NumEqualTo(args[0])
		return fmt.Sprintf("%d", count), nil
	})

	shell.Register("BEGIN", func(args ...string) (string, error) {
		base.Begin()
		return "", nil
	})

	shell.Register("ROLLBACK", func(args ...string) (string, error) {
		return "", base.Rollback()
	})

	shell.Register("COMMIT", func(args ...string) (string, error) {
		base.Commit()
		return "", nil
	})

	shell.Register("END", func(args ...string) (string, error) {
		shell.Stop()
		return "", nil
	})

	shell.Start()
}
