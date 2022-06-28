package id

/**
 * @Author  Flagship
 * @Date  2022/4/10 14:54
 * @Description
 */

// UserID defines account id object.
type UserID string

func (a UserID) String() string {
	return string(a)
}
