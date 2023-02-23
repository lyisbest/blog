package constant

import "blog/utils"

var CookieResolveError = utils.BlogError{ErrorCode: -1002, ErrorMessage: "cookie resolve failed"}

var CookieError = utils.BlogError{ErrorCode: -1003, ErrorMessage: "cookie is false"}

var UserNotExistError = utils.BlogError{ErrorCode: -1004, ErrorMessage: "User doesn't exist"}

var LogFailError = utils.BlogError{ErrorCode: -1005, ErrorMessage: "log failed"}
