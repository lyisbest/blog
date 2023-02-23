package constant

import "blog/utils"

var BlogCreateFailError = utils.BlogError{ErrorCode: -2001, ErrorMessage: "blog create failed"}
var BlogDeleteFailError = utils.BlogError{ErrorCode: -2002, ErrorMessage: "blog delete failed"}
var BlogUpdateFailError = utils.BlogError{ErrorCode: -2003, ErrorMessage: "blog update failed"}
var BlogGetFailError = utils.BlogError{ErrorCode: -2004, ErrorMessage: "blog get failed"}
var ListBlogGetFailError = utils.BlogError{ErrorCode: -2005, ErrorMessage: "list blog get failed"}
var BlogLogGetFailError = utils.BlogError{ErrorCode: -2006, ErrorMessage: "blog log get failed"}
