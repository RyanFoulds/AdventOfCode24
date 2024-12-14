package filesystem

type file struct {
	startingIndex int
	fileId        int
	length        int
}
type gap struct {
	startingIndex int
	length        int
}

type SimpleFs []int

func (f SimpleFs) Checksum() int {
	length := len(f)
	// Set J to be the last 'file' blocks
	j := length - 1
	if j%2 != 0 {
		j -= 1
	}

	sum := 0
	explodedIndex := 0
	for i := 0; i <= j; i++ {
		if i%2 == 0 {
			// Even case, just add 'em
			fileId := i / 2
			val := f[i]
			indexSum := val*explodedIndex + (val*(val-1))/2

			sum += fileId * indexSum
			explodedIndex += val
		} else {
			// Odd case, fill the gaps from the end of the list....
			blocksNeeded := f[i]
			for k := 0; k < blocksNeeded; k++ {
				for f[j] <= 0 {
					j -= 2
				}

				f[j] -= 1
				f[i] -= 1
				sum += explodedIndex * (j / 2)
				explodedIndex += 1
			}
		}
	}
	return sum
}

type FSystem struct {
	files []file
	gaps  []gap
}

func (f *FSystem) MoveFiles() {
	for j := len(f.files) - 1; j > 0; j-- {
		for i, gap := range f.gaps {
			if gap.startingIndex > f.files[j].startingIndex {
				break
			}
			fileSize := f.files[j].length
			if fileSize <= gap.length {
				f.files[j].startingIndex = gap.startingIndex // Move the file
				f.gaps[i].startingIndex += fileSize          // Adjust the gap
				f.gaps[i].length -= fileSize

				// Remove the gap if needed.
				if f.gaps[i].length == 0 {
					f.gaps = append(f.gaps[:i], f.gaps[i+1:]...)
				}
				break
			}
		}
	}

}

func (f FSystem) Checksum() int {
	sum := 0
	for _, file := range f.files {
		sum += file.fileId * (file.startingIndex*file.length + (file.length*(file.length-1))/2)
	}
	return sum
}

func SimpleFromString(s string) SimpleFs {
	length := len(s)
	ints := make([]int, length)
	for i, char := range s {
		ints[i] = int(char) - int('0')
	}
	return ints
}

func FromString(s string) FSystem {
	return parseFs(SimpleFromString(s))
}

func parseFs(fs []int) FSystem {
	files := make([]file, 0)
	gaps := make([]gap, 0)
	length := len(fs)
	expandedIndex := 0

	for i := 0; i < length; i++ {
		if i%2 == 0 {
			files = append(files, file{expandedIndex, i / 2, fs[i]})
			expandedIndex += fs[i]
		} else {
			gaps = append(gaps, gap{expandedIndex, fs[i]})
			expandedIndex += fs[i]
		}
	}
	return FSystem{files, gaps}
}
