package services

func chunkStrings(items []string, size int) [][]string {

	if size <= 0 {
		return nil
	}

	var chunks [][]string

	for i := 0; i < len(items); i += size {

		end := i + size

		if end > len(items) {
			end = len(items)
		}

		chunks = append(
			chunks,
			items[i:end],
		)
	}

	return chunks
}