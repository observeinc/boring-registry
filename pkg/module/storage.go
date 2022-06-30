package module

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"path"
)

const (
	DefaultArchiveFormat = "tar.gz"
)

var (
	archiveFormat = DefaultArchiveFormat
)

func SetArchiveFormat(newFormat string) {
	archiveFormat = newFormat
}

// Storage represents the repository of Terraform modules.
type Storage interface {
	GetModule(ctx context.Context, namespace, name, provider, version string) (Module, error)
	ListModuleVersions(ctx context.Context, namespace, name, provider string) ([]Module, error)
	UploadModule(ctx context.Context, namespace, name, provider, version string, body io.Reader) (Module, error)
}

func storagePrefix(prefix, namespace, name, provider string, urlEncode bool) string {
	delimiter := "="
	if urlEncode {
		urlEncodeAll(&delimiter, &namespace, &name, &provider)
	}
	return path.Join(
		prefix,
		fmt.Sprintf("namespace%s%s", delimiter, namespace),
		fmt.Sprintf("name%s%s", delimiter, name),
		fmt.Sprintf("provider%s%s", delimiter, provider),
	)
}

func storagePath(prefix, namespace, name, provider, version string, urlEncode bool) string {
	delimiter := "="
	if urlEncode {
		urlEncodeAll(&delimiter, &namespace, &name, &provider, &version)
	}
	return path.Join(
		prefix,
		fmt.Sprintf("namespace%s%s", delimiter, namespace),
		fmt.Sprintf("name%s%s", delimiter, name),
		fmt.Sprintf("provider%s%s", delimiter, provider),
		fmt.Sprintf("version%s%s", delimiter, version),
		fmt.Sprintf("%s-%s-%s-%s.%s", namespace, name, provider, version, archiveFormat),
	)
}

func urlEncodeAll(strings ...*string) {
	for _, s := range strings {
		*s = url.QueryEscape(*s)
	}
}
