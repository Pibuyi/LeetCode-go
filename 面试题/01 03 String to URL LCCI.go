func replaceSpaces(S string, length int) string {
    S = S[:length]
    S = strings.ReplaceAll(S, " ", "%20")
    return S
}
