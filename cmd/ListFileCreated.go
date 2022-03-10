package cmd

//List the files created during the lifetime of the file. It prints out the list of files to os.Stdout, and return an error.
func ListFilesCreated() (Files, error) {
	print(dataFiles)
	return dataFiles, nil
}
