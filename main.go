package main

import prng "prng/src"

const rngFileName = "prng_file.bin"

func main() {
	prng.CreatePRNGFile(rngFileName, 1024)
}
