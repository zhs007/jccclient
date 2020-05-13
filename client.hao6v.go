package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// hao6vNewest - hao6v newest
func (client *Client) hao6vNewest(ctx context.Context, hostname string, timeout int) (
	*jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_HAO6V,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Hao6V{
			Hao6V: &jarviscrawlercore.RequestHao6V{
				Mode: jarviscrawlercore.Hao6VMode_H6VM_NEWPAGE,
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

// Hao6vNewest - haop6v newest
func (client *Client) Hao6vNewest(ctx context.Context, timeout int) (
	*jarviscrawlercore.ReplyHao6V_Newpage, error) {

	hostname := "hao6v.com"

	reply, err := client.hao6vNewest(ctx, hostname, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_HAO6V {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	hao6v, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Hao6V)
	if !isok {
		return nil, ErrNoReplyHao6v
	}

	if hao6v.Hao6V.Mode != jarviscrawlercore.Hao6VMode_H6VM_NEWPAGE {
		return nil, ErrInvalidHao6vMode
	}

	newpage, isok := (hao6v.Hao6V.Reply).(*jarviscrawlercore.ReplyHao6V_Newpage)
	if !isok {
		return nil, ErrNoReplyHao6vNewPage
	}

	return newpage, nil
}

// hao6vRes - hao6v res
func (client *Client) hao6vRes(ctx context.Context, hostname string, url string, timeout int) (
	*jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_HAO6V,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Hao6V{
			Hao6V: &jarviscrawlercore.RequestHao6V{
				Mode: jarviscrawlercore.Hao6VMode_H6VM_RESPAGE,
				Url:  url,
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

// Hao6vRes - haop6v res
func (client *Client) Hao6vRes(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.ReplyHao6V_Res, error) {

	hostname := "hao6v.com"

	reply, err := client.hao6vRes(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_HAO6V {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	hao6v, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Hao6V)
	if !isok {
		return nil, ErrNoReplyHao6v
	}

	if hao6v.Hao6V.Mode != jarviscrawlercore.Hao6VMode_H6VM_RESPAGE {
		return nil, ErrInvalidHao6vMode
	}

	res, isok := (hao6v.Hao6V.Reply).(*jarviscrawlercore.ReplyHao6V_Res)
	if !isok {
		return nil, ErrNoReplyHao6vResPage
	}

	return res, nil
}
