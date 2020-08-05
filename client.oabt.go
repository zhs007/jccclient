package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// oabtPage - oabt page
func (client *Client) oabtPage(ctx context.Context, hostname string,
	pageindex int32, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_OABT,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Oabt{
			Oabt: &jarviscrawlercore.RequestOABT{
				Mode:      jarviscrawlercore.OABTMode_OABTM_PAGE,
				PageIndex: pageindex,
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

// OABTPage - OABT Page
func (client *Client) OABTPage(ctx context.Context, pageindex int32, timeout int) (
	*jarviscrawlercore.ReplyOABT_Page, error) {

	hostname := "oabt008.com"

	_, reply, err := client.oabtPage(ctx, hostname, pageindex, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_OABT {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	oabt, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Oabt)
	if !isok {
		return nil, ErrNoReplyOABT
	}

	if oabt.Oabt.Mode != jarviscrawlercore.OABTMode_OABTM_PAGE {
		return nil, ErrInvalidOABTMode
	}

	replyoabtpage, isok := (oabt.Oabt.Reply).(*jarviscrawlercore.ReplyOABT_Page)
	if !isok {
		return nil, ErrNoReplyOABTPage
	}

	return replyoabtpage, nil
}
