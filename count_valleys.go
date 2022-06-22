package countvalleys

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	altitude := 0
	valleyCount := 0
	for _, step := range path {
		if string(step) == "D" {
			altitude--
		} else {
			altitude++
			if altitude == 0 {
				valleyCount++
			}
		}
	}
	return int32(valleyCount)
}
