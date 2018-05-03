package backup

// Archiver represents type capable of archiving and
// restoring files.
type Archiver interface {
	Archive(src, dest string) error
}
