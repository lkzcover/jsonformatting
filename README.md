# jsonformatting
simple helper for formatted JSON saving to file

### func ConvertToFormatJSON
    ConvertToFormatJSON(data []byte) []byte
If the incoming data is JSON, the function returns a formatted slice byte. Otherwise it returns the original slice.