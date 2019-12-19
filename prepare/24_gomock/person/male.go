package person

//go:generate mockgen -destination=../mock/male_mock.go -package=mock github.com/qiuhoude/go-web/prepare/24_gomock/person Male

// cd ..
// go generate ./...
type Male interface {
	Get(id int64) error
}
