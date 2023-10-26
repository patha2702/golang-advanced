package decrypt


func Decrypt(data string) string {
        decryptedData := ""
        for _, char := range data {
                decryptedData += string(char - 3)
        }
        return decryptedData
}
