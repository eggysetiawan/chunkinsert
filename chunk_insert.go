package chunkinsert

func ChunkInsert(dest []interface{}, cls func(chunkData []interface{}), chunkSize int) {
	// Calculate the number of chunks
	numChunks := (len(dest) + chunkSize - 1) / chunkSize
	// Start a transaction

	// Iterate through the chunks
	for i := 0; i < numChunks; i++ {
		// Calculate the start and end indices for the current chunk
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize
		if endIndex > len(dest) {
			endIndex = len(dest)
		}

		// Extract the current chunk dest
		cd := dest[startIndex:endIndex]

		// Process the chunk
		cls(cd)
	}
}
