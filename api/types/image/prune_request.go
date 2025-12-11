package image

// PruneRequest contains the request body for POST /images/prune
//
// This struct is used for API version 1.53 and later.
// Earlier API versions use query parameters instead.
type PruneRequest struct {
	// Filters is a JSON-encoded set of filter arguments.
	//
	// Available filters:
	//   - dangling=<boolean>: When set to true (or 1), prune only unused and
	//     untagged images. When set to false (or 0), all unused images are pruned.
	//   - until=<timestamp>: Prune images created before this timestamp.
	//     The timestamp can be Unix timestamps, date formatted timestamps,
	//     or Go duration strings (e.g. 10m, 1h30m).
	//   - label: Prune images with (or without, in case label!=... is used)
	//     the specified labels.
	Filters map[string]map[string]bool `json:"Filters,omitempty"`
}
