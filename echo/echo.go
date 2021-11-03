package echo

import "io"

func Do(rwc io.ReadWriter) error {
	_, err := io.Copy(rwc, rwc)
	return err
}
