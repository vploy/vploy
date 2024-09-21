package playstore

import "github.com/Cyberpull/gokit"

const endpoint string = "https://androidpublisher.googleapis.com"

type initializable interface {
	initialize()
}

func d[T initializable](x T) T {
	x.initialize()
	return x
}

func url(paths ...any) string {
	paths = append([]any{endpoint}, paths...)
	return gokit.Join("/", paths...)
}
