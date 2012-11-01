package otf

type DefaultTable struct {
	tag  TAG
	body []byte
}

func (t *DefaultTable) Tag() TAG {
	return t.tag
}

func (t *DefaultTable) Bytes() []byte {
	return t.body
}
