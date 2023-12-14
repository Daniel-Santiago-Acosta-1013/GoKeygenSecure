package cmd

import (
    "bufio"
    "fmt"
    "os"
    "GoKeygenSecure/pkg/evaluator"
    "github.com/spf13/cobra"
)

var evaluateCmd = &cobra.Command{
    Use:   "evaluate",
    Short: "Evalúa la seguridad de una contraseña",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Print("Ingresa la contraseña para evaluar: ")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        password := scanner.Text()

        securityLevel := evaluator.EvaluatePassword(password)
        fmt.Println("Nivel de seguridad de la contraseña:", securityLevel)
    },
}

func init() {
    rootCmd.AddCommand(evaluateCmd)
}
