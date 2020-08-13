package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// getJDActive - get jd active
func (client *Client) getJDActive(ctx context.Context, hostname string, url string,
	timeout int) (string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_JD,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Jd{
			Jd: &jarviscrawlercore.RequestJD{
				Mode: jarviscrawlercore.JDMode_JDM_ACTIVE,
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

// GetJDActive - get jd active
func (client *Client) GetJDActive(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.JDActive, error) {

	hostname := "jd.com"

	_, reply, err := client.getJDActive(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_JD {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	jd, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Jd)
	if !isok {
		return nil, ErrNoReplyJD
	}

	if jd.Jd.Mode != jarviscrawlercore.JDMode_JDM_ACTIVE {
		return nil, ErrInvalidJDMode
	}

	replyactive, isok := (jd.Jd.Reply).(*jarviscrawlercore.ReplyJD_Active)
	if !isok {
		return nil, ErrNoReplyJDActive
	}

	return replyactive.Active, nil
}

// getJDProduct - get jd product
func (client *Client) getJDProduct(ctx context.Context, hostname string, url string,
	timeout int) (string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_JD,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Jd{
			Jd: &jarviscrawlercore.RequestJD{
				Mode: jarviscrawlercore.JDMode_JDM_PRODUCT,
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

// GetJDProduct - get jd product
func (client *Client) GetJDProduct(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.JDProduct, error) {

	hostname := "jd.com"

	_, reply, err := client.getJDProduct(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_JD {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	jd, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Jd)
	if !isok {
		return nil, ErrNoReplyJD
	}

	if jd.Jd.Mode != jarviscrawlercore.JDMode_JDM_PRODUCT {
		return nil, ErrInvalidJDMode
	}

	replyproduct, isok := (jd.Jd.Reply).(*jarviscrawlercore.ReplyJD_Product)
	if !isok {
		return nil, ErrNoReplyJDActive
	}

	return replyproduct.Product, nil
}

// getJDActivePage - get jd activepage
func (client *Client) getJDActivePage(ctx context.Context, hostname string, url string,
	timeout int) (string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_JD,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Jd{
			Jd: &jarviscrawlercore.RequestJD{
				Mode: jarviscrawlercore.JDMode_JDM_ACTIVEPAGE,
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

// GetJDActivePage - get jd activepage
func (client *Client) GetJDActivePage(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.JDActive, error) {

	hostname := "jd.com"

	_, reply, err := client.getJDActivePage(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_JD {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	jd, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Jd)
	if !isok {
		return nil, ErrNoReplyJD
	}

	if jd.Jd.Mode != jarviscrawlercore.JDMode_JDM_ACTIVEPAGE {
		return nil, ErrInvalidJDMode
	}

	replyactive, isok := (jd.Jd.Reply).(*jarviscrawlercore.ReplyJD_Active)
	if !isok {
		return nil, ErrNoReplyJDActive
	}

	return replyactive.Active, nil
}
