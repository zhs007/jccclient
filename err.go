package jccclient

import "errors"

var (
	// ErrNoServAddr - There is no ServAddr
	ErrNoServAddr = errors.New("There is no ServAddr")
	// ErrNoToken - There is no token
	ErrNoToken = errors.New("There is no token")
	// ErrDuplicateServAddr - Duplicate servaddr
	ErrDuplicateServAddr = errors.New("Duplicate servaddr")
	// ErrNoHostname - There is no hostname
	ErrNoHostname = errors.New("There is no hostname")
	// ErrNoHostInfo - There is no hostinfo
	ErrNoHostInfo = errors.New("There is no hostinfo")
	// ErrNoDBPath - There is no dbpath
	ErrNoDBPath = errors.New("There is no dbpath")
	// ErrNoDBEngine - There is no dbengine
	ErrNoDBEngine = errors.New("There is no dbengine")
	// ErrNoTag - There is no tag
	ErrNoTag = errors.New("There is no tag")

	// ErrInvalidTask - invalid task
	ErrInvalidTask = errors.New("invalid task")

	// ErrNoReplyCrawler - no ReplyCrawler
	ErrNoReplyCrawler = errors.New("no ReplyCrawler")
	// ErrInvalidCrawlerType - invalid crealertype
	ErrInvalidCrawlerType = errors.New("invalid crealertype")
	// ErrNoCrawlerResult - no CrawlerResult
	ErrNoCrawlerResult = errors.New("no CrawlerResult")

	// ErrNoReplyAnalyzePage - invalid ReplyAnalyzePage
	ErrNoReplyAnalyzePage = errors.New("invalid ReplyAnalyzePage")

	// ErrNoReplyGeoip - invalid ReplyGeoip
	ErrNoReplyGeoip = errors.New("invalid ReplyGeoip")

	// ErrNoReplyTechInAsia - invalid ReplyTechInAsia
	ErrNoReplyTechInAsia = errors.New("invalid ReplyTechInAsia")
	// ErrInvalidTechInAsiaMode - invalid techinasia mode
	ErrInvalidTechInAsiaMode = errors.New("invalid techinasia mode")
	// ErrNoReplyTechInAsiaJob - invalid ReplyTechInAsiaJob
	ErrNoReplyTechInAsiaJob = errors.New("invalid ReplyTechInAsiaJob")
	// ErrNoReplyTechInAsiaJobs - invalid ReplyTechInAsiaJobs
	ErrNoReplyTechInAsiaJobs = errors.New("invalid ReplyTechInAsiaJobs")
	// ErrNoReplyTechInAsiaJobTags - invalid ReplyTechInAsiaJobTags
	ErrNoReplyTechInAsiaJobTags = errors.New("invalid ReplyTechInAsiaJobTags")
	// ErrNoReplyTechInAsiaCompany - invalid ReplyTechInAsiaCompany
	ErrNoReplyTechInAsiaCompany = errors.New("invalid ReplyTechInAsiaCompany")

	// ErrNoReplySteepAndCheap - invalid ReplySteepAndCheap
	ErrNoReplySteepAndCheap = errors.New("invalid ReplySteepAndCheap")
	// ErrInvalidSteepAndCheapMode - invalid steepandcheap mode
	ErrInvalidSteepAndCheapMode = errors.New("invalid steepandcheap mode")
	// ErrNoReplySteepAndCheapProducts - invalid ReplySteepAndCheapProducts
	ErrNoReplySteepAndCheapProducts = errors.New("invalid ReplySteepAndCheapProducts")
	// ErrNoReplySteepAndCheapProduct - invalid ReplySteepAndCheapProduct
	ErrNoReplySteepAndCheapProduct = errors.New("invalid ReplySteepAndCheapProduct")

	// ErrNoReplyJRJ - invalid ReplyJRJ
	ErrNoReplyJRJ = errors.New("invalid ReplyJRJ")
	// ErrInvalidJRJMode - invalid jrj mode
	ErrInvalidJRJMode = errors.New("invalid jrj mode")
	// ErrNoReplyJRJFunds - invalid ReplyJRJFunds
	ErrNoReplyJRJFunds = errors.New("invalid ReplyJRJFunds")
	// ErrNoReplyJRJFund - invalid ErrNoReplyJRJFund
	ErrNoReplyJRJFund = errors.New("invalid NoReplyJRJFund")
	// ErrNoReplyJRJFundValue - invalid ErrNoReplyJRJFundValue
	ErrNoReplyJRJFundValue = errors.New("invalid ErrNoReplyJRJFundValue")

	// ErrNoReplyJD - invalid ReplyJD
	ErrNoReplyJD = errors.New("invalid ReplyJD")
	// ErrInvalidJDMode - invalid jd mode
	ErrInvalidJDMode = errors.New("invalid jd mode")
	// ErrNoReplyJDActive - no ReplyJDActive
	ErrNoReplyJDActive = errors.New("no ReplyJDActive")
	// ErrNoReplyJDProduct - no ReplyJDProduct
	ErrNoReplyJDProduct = errors.New("no ReplyJDProduct")

	// ErrNoReplyAlimama - invalid ReplyAlimama
	ErrNoReplyAlimama = errors.New("invalid ReplyAlimama")
	// ErrInvalidAlimamaMode - invalid alimama mode
	ErrInvalidAlimamaMode = errors.New("invalid alimama mode")
	// ErrNoReplyAlimamaGetTop - no ReplyAlimamaGetTop
	ErrNoReplyAlimamaGetTop = errors.New("no ReplyAlimamaGetTop")
	// ErrNoReplyAlimamaProducts - no ReplyAlimamaProducts
	ErrNoReplyAlimamaProducts = errors.New("no ReplyAlimamaProducts")

	// ErrNoReplyTmall - invalid ReplyTmall
	ErrNoReplyTmall = errors.New("invalid ReplyTmall")
	// ErrInvalidTmallMode - invalid tmall mode
	ErrInvalidTmallMode = errors.New("invalid tmall mode")
	// ErrNoReplyTmallProduct - no ReplyTmallProduct
	ErrNoReplyTmallProduct = errors.New("no ReplyTmallProduct")

	// ErrNoReplyTaobao - invalid ReplyTaobao
	ErrNoReplyTaobao = errors.New("invalid ReplyTaobao")
	// ErrInvalidTaobaoMode - invalid taobao mode
	ErrInvalidTaobaoMode = errors.New("invalid taobao mode")
	// ErrNoReplyTaobaoProduct - no ReplyTaobaoProduct
	ErrNoReplyTaobaoProduct = errors.New("no ReplyTaobaoProduct")

	// ErrNoReplyMountainsteals - invalid ReplyMountainsteals
	ErrNoReplyMountainsteals = errors.New("invalid ReplyMountainsteals")
	// ErrInvalidMountainstealsMode - invalid mountainsteals mode
	ErrInvalidMountainstealsMode = errors.New("invalid mountainsteals mode")
	// ErrNoReplyMountainstealsProduct - no ReplyMountainstealsProduct
	ErrNoReplyMountainstealsProduct = errors.New("no ReplyMountainstealsProduct")
	// ErrNoReplyMountainstealsSale - no ReplyMountainstealsSale
	ErrNoReplyMountainstealsSale = errors.New("no ReplyMountainstealsSale")

	// ErrInvalidClientMgrState - invalid ClientMgrState
	ErrInvalidClientMgrState = errors.New("invalid ClientMgrState")
)
