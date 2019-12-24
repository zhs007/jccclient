package jccclient

// Tags - tags
type Tags struct {
	Tag     string
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
