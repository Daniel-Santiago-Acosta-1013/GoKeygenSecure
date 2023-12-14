package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "GoKeygenSecure",
    Short: "GoKeygenSecure es una herramienta de generación y evaluación de contraseñas",
    Long: `Una herramienta completa de línea de comandos para generar contraseñas seguras
y evaluar la seguridad de las contraseñas existentes.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func init() {
}
