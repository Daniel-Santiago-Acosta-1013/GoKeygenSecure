package evaluator

func EvaluatePassword(password string) string {
    switch {
    case len(password) < 6:
        return "Muy débil"
    case len(password) < 8:
        return "Débil"
    case len(password) < 10:
        return "Moderada"
    default:
        return "Fuerte"
    }
}
