package console

import (
	"fmt"
	"io"

	"github.com/logrusorgru/aurora"
	"github.com/oslokommune/okctl/pkg/client"
	"github.com/oslokommune/okctl/pkg/client/store"
	"github.com/theckman/yacspin"
)

type argoCDReport struct {
	console *Console
}

func (r *argoCDReport) CreateArgoCD(cd *client.ArgoCD, reports []*store.Report) error {
	var actions []store.Action // nolint: prealloc

	for _, report := range reports {
		actions = append(actions, report.Actions...)
	}

	description := fmt.Sprintf("%s (url: %s, org: %s, team: %s)",
		aurora.Blue("argocd"),
		cd.ArgoURL,
		cd.GithubOauthApp.Organisation,
		cd.GithubOauthApp.Team.Name,
	)

	return r.console.Report(actions, "argocd", description)
}

// NewArgoCDReport returns an initialised reporter
func NewArgoCDReport(out io.Writer, exit chan struct{}, spinner *yacspin.Spinner) client.ArgoCDReport {
	return &argoCDReport{
		console: New(out, exit, spinner),
	}
}