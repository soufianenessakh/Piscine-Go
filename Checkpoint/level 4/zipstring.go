package checkpoint

func intToString(n int) string {
    if n == 0 {
        return "0"
    }

    digits := []byte{}
    for n > 0 {
        digits = append([]byte{byte('0' + n%10)}, digits...)
        n /= 10
    }

    return string(digits)
}

func ZipString(s string) string {
    result := ""
    count := 1

    for i := 0; i < len(s); i++ {
        if i+1 < len(s) && s[i] == s[i+1] {
            count++
        } else {
            result += intToString(count)
            result += string(s[i])
            count = 1
        }
    }

    return result
}