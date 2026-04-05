package piscine

func Capitalize(s string) string {
    output := ""
    sing := false
    for _, char := range s {
        if char == '+' {
            sing = true
            output += string(char)
            continue
        }
        if sing && char >= 'a' && char <= 'z' {
            char = char - 32
            sing = false
		}
		if sing && char>='0' && char<='9'{
			sing=false
		}
        output += string(char)
    }
    return output
}
