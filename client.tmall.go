package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// tmallProduct - tmall product
func (client *Client) tmallProduct(ctx context.Context, hostname string, url string,
	timeout int) (string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_TMALL,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Tmall{
			Tmall: &jarviscrawlercore.RequestTmall{
				Mode: jarviscrawlercore.TmallMode_TMM_PRODUCT,
				Url:  url,
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

// TmallProduct - tmall product
func (client *Client) TmallProduct(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.TmallProduct, error) {

	hostname := "tmall.com"

	_, reply, err := client.tmallProduct(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_TMALL {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	tmall, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Tmall)
	if !isok {
		return nil, ErrNoReplyTmall
	}

	if tmall.Tmall.Mode != jarviscrawlercore.TmallMode_TMM_PRODUCT {
		return nil, ErrInvalidTmallMode
	}

	replyproduct, isok := (tmall.Tmall.Reply).(*jarviscrawlercore.ReplyTmall_Product)
	if !isok {
		return nil, ErrNoReplyTmallProduct
	}

	return replyproduct.Product, nil
}

// tmallMobileProduct - tmall mobile product
func (client *Client) tmallMobileProduct(ctx context.Context, hostname string,
	url string, device string, timeout int) (string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_TMALL,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Tmall{
			Tmall: &jarviscrawlercore.RequestTmall{
				Mode:   jarviscrawlercore.TmallMode_TMM_MOBILEPRODUCT,
				Url:    url,
				Device: device,
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

// TmallMobileProduct - tmall mobile product
func (client *Client) TmallMobileProduct(ctx context.Context, url string,
	device string, timeout int) (
	*jarviscrawlercore.TmallProduct, error) {

	hostname := "tmall.com"

	_, reply, err := client.tmallMobileProduct(ctx, hostname, url, device, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_TMALL {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	tmall, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Tmall)
	if !isok {
		return nil, ErrNoReplyTmall
	}

	if tmall.Tmall.Mode != jarviscrawlercore.TmallMode_TMM_PRODUCT {
		return nil, ErrInvalidTmallMode
	}

	replyproduct, isok := (tmall.Tmall.Reply).(*jarviscrawlercore.ReplyTmall_Product)
	if !isok {
		return nil, ErrNoReplyTmallProduct
	}

	return replyproduct.Product, nil
}
