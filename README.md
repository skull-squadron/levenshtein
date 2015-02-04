# Levenshtein edit distance

Because everybody that visits wikipedia feels obliged to push out trivial code.

Both case-sensitive and case-insensitive functions.

```go
import "github.com/steakknife/levenshtein"

levenshtein.Distance("xyz", "x") // 2
levenshtein.DistanceCaseInsensitive("xyz", "X") // 2

```

## License

MIT

## Copyright

Copyright (C) 2015 Barry Allard
