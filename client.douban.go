package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// doubanSearch - douban search
func (client *Client) doubanSearch(ctx context.Context, hostname string,
	doubanType jarviscrawlercore.DoubanType, text string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_DOUBAN,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Douban{
			Douban: &jarviscrawlercore.RequestDouban{
				Mode:       jarviscrawlercore.DoubanMode_DBM_SEARCH,
				Text:       text,
				DoubanType: doubanType,
			},
		},
	}

	version, reply, err := client.RequestCrawler(ctx, req)
	if err != nil {
		if client.cfg != nil {
			client.Hosts.OnTaskEnd(ctx, hostname, true, client.cfg)
		}

		return version, nil, err
	}

	if reply == nil {
		if client.cfg != nil {
			client.Hosts.OnTaskEnd(ctx, hostname, true, client.cfg)
		}

		return version, nil, ErrNoReplyCrawler
	}

	if client.cfg != nil {
		client.Hosts.OnTaskEnd(ctx, hostname, false, client.cfg)
	}

	return version, reply, nil
}

// DoubanSearch - douban search
func (client *Client) DoubanSearch(ctx context.Context,
	doubanType jarviscrawlercore.DoubanType, text string, timeout int) (
	*jarviscrawlercore.DoubanSearch, error) {

	hostname := "douban.com"

	_, reply, err := client.doubanSearch(ctx, hostname, doubanType, text, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_DOUBAN {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	douban, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Douban)
	if !isok {
		return nil, ErrNoReplyDouban
	}

	if douban.Douban.Mode != jarviscrawlercore.DoubanMode_DBM_SEARCH {
		return nil, ErrInvalidDoubanMode
	}

	replysearch, isok := (douban.Douban.Reply).(*jarviscrawlercore.ReplyDouban_Search)
	if !isok {
		return nil, ErrNoReplyDoubanSearch
	}

	return replysearch.Search, nil
}

// doubanBook - douban book
func (client *Client) doubanBook(ctx context.Context, hostname string,
	id string, timeout int) (string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_DOUBAN,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Douban{
			Douban: &jarviscrawlercore.RequestDouban{
				Mode: jarviscrawlercore.DoubanMode_DBM_BOOK,
				Id:   id,
			},
		},
	}

	version, reply, err := client.RequestCrawler(ctx, req)
	if err != nil {
		if client.cfg != nil {
			client.Hosts.OnTaskEnd(ctx, hostname, true, client.cfg)
		}

		return version, nil, err
	}

	if reply == nil {
		if client.cfg != nil {
			client.Hosts.OnTaskEnd(ctx, hostname, true, client.cfg)
		}

		return version, nil, ErrNoReplyCrawler
	}

	if client.cfg != nil {
		client.Hosts.OnTaskEnd(ctx, hostname, false, client.cfg)
	}

	return version, reply, nil
}

// DoubanBook - douban book
func (client *Client) DoubanBook(ctx context.Context,
	id string, timeout int) (*jarviscrawlercore.DoubanBook, error) {

	hostname := "douban.com"

	_, reply, err := client.doubanBook(ctx, hostname, id, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_DOUBAN {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	douban, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Douban)
	if !isok {
		return nil, ErrNoReplyDouban
	}

	if douban.Douban.Mode != jarviscrawlercore.DoubanMode_DBM_BOOK {
		return nil, ErrInvalidDoubanMode
	}

	replybook, isok := (douban.Douban.Reply).(*jarviscrawlercore.ReplyDouban_Book)
	if !isok {
		return nil, ErrNoReplyDoubanBook
	}

	return replybook.Book, nil
}
