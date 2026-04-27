package repository

func nullIfEmpty(value string) any {
	if value == "" {
		return nil
	}

	return value
}
