package security

type PasswordHasher interface {
	Hash(password string) (string, error)

	Check(
		password string,
		hash string,
	) bool
}
