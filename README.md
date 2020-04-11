# jsonformatting
simple helper for formatted JSON saving to file

### func ConvertToFormatJSONWithError
    ConvertToFormatJSONWithError(data []byte) []byte
If the incoming data is JSON, the function returns a formatted slice byte. Otherwise it returns error.

### func ConvertToFormatJSON
    ConvertToFormatJSON(data []byte) []byte
If the incoming data is JSON, the function returns a formatted slice byte. Otherwise it returns the original slice.

*Use only if you are sure that the incoming data is in the correct JSON format!*

