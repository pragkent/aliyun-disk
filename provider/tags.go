package provider

type TagResourceType string

const (
	TagResourceImage    = TagResourceType("image")
	TagResourceInstance = TagResourceType("instance")
	TagResourceSnapshot = TagResourceType("snapshot")
	TagResourceDisk     = TagResourceType("disk")
)

type AddTagsArgs struct {
	ResourceId   string
	ResourceType TagResourceType
	Tag          map[string]string
}
