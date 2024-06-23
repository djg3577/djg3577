func isPalindrome(s string) bool {
    var leftByte, rightByte byte
    l := 0
    r := len(s)-1

    for l < r {
        leftByte = getCorrectByte(s[l])
        rightByte = getCorrectByte(s[r])

        if leftByte == 0 {
            l++
        }

        if rightByte == 0 {
            r--
        }

        if leftByte == 0 || rightByte == 0 {
            continue
        }

        if leftByte != rightByte{
            return false
        }

        l++
        r--
    }
    return true
}

func getCorrectByte(b byte) byte {
    if b > 64 && b < 91{
        b += 32
    }

    if (b > 47 && b < 58) || (b > 96 && b < 123) {
        return b
    }

    return 0
}
