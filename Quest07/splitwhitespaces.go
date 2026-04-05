package Quest07

func SplitWhiteSpaces(s string) []string {
    output := []string{}
    word := ""

    for _, char := range s {
        if char != ' ' && char != '\t' && char != '\n' {
            word += string(char)
        } else {
            if word != "" {
                output = append(output, word)
                word = ""
            }
        }
    }

    if word != "" {
        output = append(output, word)
    }

    return output
}