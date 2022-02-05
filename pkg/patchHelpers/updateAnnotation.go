package patchHelpers

import "github.com/muhammedkaya/tcpdump-webhook/pkg"

func UpdateAnnotation(
	target, added map[string]string,
) (patch []pkg.PatchOperation) {
	for key, value := range added {
		if target == nil || target[key] == "" {
			target = map[string]string{}

			patch = append(patch, pkg.PatchOperation{
				Op:   pkg.PatchOperationAdd,
				Path: "/metadata/annotations",
				Value: map[string]string{
					key: value,
				},
			})
		} else {
			patch = append(patch, pkg.PatchOperation{
				Op:    pkg.PatchOperationReplace,
				Path:  "/metadata/annotations/" + key,
				Value: value,
			})
		}
	}

	return patch
}
