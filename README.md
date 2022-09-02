
# Pars Time

Parse and validate time from string 


## Test

To test this project run

```bash
  go test -v 
```


## Sample output
### Output with valid string time
```bash
Input: "22:5:25"
Output: {22 5 25}
```
### Output with invalid string time
```bash
Input: "25:10:25"
Output: 25:10:25: Hour out of range : 0 <= hour <= 23: -1, error should be nil
```

## Usage/Examples

```go
data, err := ParseTime("13:25:00")
```





