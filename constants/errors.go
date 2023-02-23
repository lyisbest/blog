package constants

import "blog/utils"

var QueryError = utils.BlogError{ErrorCode: -1, ErrorMessage: "query failed!"}

var ResolveError = utils.BlogError{ErrorCode: -1001, ErrorMessage: "parameter resolve failed"}
