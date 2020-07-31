package cache

type redis struct {
}

func (r *redis) Get(string) (string, error) {

	return "", nil
}

func (r *redis) Set(string, string) (bool, error) {
	return false, nil
}

func (r *redis) SetEx(string, string, int) (bool, error) {
	return false, nil
}

func (r *redis) Delete(string) (bool, error) {
	return false, nil
}

func (r *redis) SetTimeout(string, int) (bool, error) {
	return false, nil
}
