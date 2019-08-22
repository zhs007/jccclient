package jccclient

import "errors"

var (
	// ErrNoServAddr - There is no ServAddr
	ErrNoServAddr = errors.New("There is no ServAddr")
	// ErrNoToken - There is no token
	ErrNoToken = errors.New("There is no token")
	// ErrDuplicateServAddr - Duplicate servaddr
	ErrDuplicateServAddr = errors.New("Duplicate servaddr")

	// ErrNoReplyCrawler - no ReplyCrawler
	ErrNoReplyCrawler = errors.New("no ReplyCrawler")
	// ErrInvalidCrawlerType - invalid crealertype
	ErrInvalidCrawlerType = errors.New("invalid crealertype")
	// ErrNoCrawlerResult - no CrawlerResult
	ErrNoCrawlerResult = errors.New("no CrawlerResult")

	// ErrNoReplyAnalyzePage - invalid ReplyAnalyzePage
	ErrNoReplyAnalyzePage = errors.New("invalid ReplyAnalyzePage")
)
