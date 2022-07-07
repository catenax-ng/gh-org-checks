package github.repository

protection = data {
	resp := github.request("GET /repos/{owner}/{repo}/branches/{branch}/protection", {
		"owner": input.owner.login,
		"repo": input.name,
		"branch": input.default_branch,
	})

	resp.status == 200

	data := resp.body
}

violation_default_branch_not_protected {
	not protection
}
#
#violation_default_branch_pull_not_required {
#	not protection.required_pull_request_reviews
#}
#
#violation_default_branch_approvals_not_required {
#	not protection.required_pull_request_reviews.required_approving_review_count
#}
#
#violation_default_branch_approvals_not_required {
#	protection.required_pull_request_reviews.required_approving_review_count < 1
#}
#
#violation_default_branch_code_owners_reviews_not_required {
#	not protection.required_pull_request_reviews.require_code_owner_reviews
#}
#
#violation_default_branch_status_checks_not_required {
#	not protection.required_status_checks
#}
#
#violation_default_branch_up_to_date_not_required {
#	not protection.required_status_checks.strict
#}
