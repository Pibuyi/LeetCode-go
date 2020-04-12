func entityParser(text string) string {
    text = strings.Replace(text, "&quot;", "\"",  -1)
    text = strings.Replace(text, "&apos;", "'",  -1)
    text = strings.Replace(text, "&amp;", "&",  -1)
    text = strings.Replace(text, "&gt;", ">",  -1)
    text = strings.Replace(text, "&lt;", "<", -1)
    text = strings.Replace(text, "&frasl;", "/", -1)
    return text
}
