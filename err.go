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

	// ErrInvalidClientMgrState - invalid ClientMgrState
	ErrInvalidClientMgrState = errors.New("invalid ClientMgrState")
)
