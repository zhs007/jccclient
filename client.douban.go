package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/proto"
)

// doubanSearch - douban search
func (client *Client) doubanSearch(ctx context.Context, hostname string,
	doubanType jarviscrawlercore.DoubanType, text string, timeout int) (
	*jarviscrawlercore.ReplyCrawler, error) {

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

	reply, err := client.RequestCrawler(ctx, req)
	if err != nil {
		if client.cfg != nil {
			client.Hosts.OnTaskEnd(ctx, hostname, true, client.cfg)
		}

		return nil, err
	}

	if reply == nil {
		if client.cfg != nil {
			client.Hosts.OnTaskEnd(ctx, hostname, true, client.cfg)
		}

		return nil, ErrNoReplyCrawler
	}

	if client.cfg != nil {
		client.Hosts.OnTaskEnd(ctx, hostname, false, client.cfg)
	}

	return reply, nil
}

// DoubanSearch - douban search
func (client *Client) DoubanSearch(ctx context.Context,
	doubanType jarviscrawlercore.DoubanType, text string, timeout int) (
	*jarviscrawlercore.DoubanSearch, error) {

	hostname := "douban.com"

	reply, err := client.doubanSearch(ctx, hostname, doubanType, text, timeout)
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
