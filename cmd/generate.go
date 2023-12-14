package cmd

import (
    "fmt"
    "GoKeygenSecure/pkg/generator"
    "github.com/spf13/cobra"
)

// Flags para la configuración de la contraseña
var (
    length      int
    includeNums bool
    includeSyms bool
    easyToSay   bool
    easyToRead  bool
)

var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Genera una contraseña segura",
    Run: func(cmd *cobra.Command, args []string) {
        password := generator.GeneratePassword(length, includeNums, includeSyms, easyToSay, easyToRead)
        fmt.Println("Contraseña generada:", password)
    },
}

func init() {
    rootCmd.AddCommand(generateCmd)
    generateCmd.Flags().IntVarP(&length, "length", "l", 12, "longitud de la contraseña")
    generateCmd.Flags().BoolVarP(&includeNums, "nums", "n", true, "incluye números en la contraseña")
    generateCmd.Flags().BoolVarP(&includeSyms, "syms", "s", true, "incluye símbolos en la contraseña")
    generateCmd.Flags().BoolVarP(&easyToSay, "easy-to-say", "e", false, "fácil de decir (evita caracteres confusos)")
    generateCmd.Flags().BoolVarP(&easyToRead, "easy-to-read", "r", false, "fácil de leer (evita caracteres confusos)")
}
