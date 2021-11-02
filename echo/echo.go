package echo

import "io"

func Do(rwc io.ReadWriteCloser) error {
	_, err := io.Copy(rwc, rwc)
	if err != nil {
		return err
	}
	return rwc.Close()
}
