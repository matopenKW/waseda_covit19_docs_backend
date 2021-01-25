package impl

import (
	"net/url"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type RequestImpl interface {
	SetRequest(url.Values)
	Validate() error
	Execute(repository.Connection) (ResponceImpl, error)
}
