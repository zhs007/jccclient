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

	// ErrNoReplyCrawler - no ReplyCrawler
	ErrNoReplyCrawler = errors.New("no ReplyCrawler")
	// ErrInvalidCrawlerType - invalid crealertype
	ErrInvalidCrawlerType = errors.New("invalid crealertype")
	// ErrNoCrawlerResult - no CrawlerResult
	ErrNoCrawlerResult = errors.New("no CrawlerResult")

	// ErrNoReplyAnalyzePage - invalid ReplyAnalyzePage
	ErrNoReplyAnalyzePage = errors.New("invalid ReplyAnalyzePage")
)
