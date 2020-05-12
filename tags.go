package jccclient

// Tags - tags
//		Tag: 如果Tag不为空，目标一定要包含这个Tag
//		Or: 如果Or不为空，目标至少要包含一个Or
//		And: 如果And不为空，目标必须包含全部And
//		Exclude: 如果Exclude不为空，目标不嫩包含任何And
type Tags struct {
	Tag     string
	Or      []string
	And     []string
	Exclude []string
}

func findtags(tag string, tags []string) bool {
	for _, v := range tags {
		if v == tag {
			return true
		}
	}

	return false
}

// IsMatch - is match tags
func (tags *Tags) IsMatch(client *Client) bool {

	if tags.Tag != "" && !findtags(tags.Tag, client.tags) {
		return false
	}

	if len(tags.Or) > 0 {
		hasor := false
		for _, v := range tags.Or {
			if findtags(v, client.tags) {
				hasor = true

				break
			}
		}

		if !hasor {
			return false
		}
	}

	for _, v := range tags.And {
		if !findtags(v, client.tags) {
			return false
		}
	}

	for _, v := range tags.Exclude {
		if findtags(v, client.tags) {
			return false
		}
	}

	return true
}
