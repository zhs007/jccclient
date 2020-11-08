package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// hostInvesting - 6vdy.org
const hostInvesting = "investing.com"

// investingAssets - investing assets
func (client *Client) investingAssets(ctx context.Context, hostname string, url string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_INVESTING,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Investing{
			Investing: &jarviscrawlercore.RequestInvesting{
				Mode: jarviscrawlercore.InvestingMode_INVESTINGMODE_ASSETS,
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

// InvestingAssets - investing assets
func (client *Client) InvestingAssets(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.ReplyInvesting_Assets, error) {

	hostname := hostInvesting

	_, reply, err := client.investingAssets(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_INVESTING {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	investing, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Investing)
	if !isok {
		return nil, ErrNoReplyP6vdy
	}

	if investing.Investing.Mode != jarviscrawlercore.InvestingMode_INVESTINGMODE_ASSETS {
		return nil, ErrInvalidInvestingMode
	}

	assets, isok := (investing.Investing.Reply).(*jarviscrawlercore.ReplyInvesting_Assets)
	if !isok {
		return nil, ErrNoReplyInvestingAssets
	}

	return assets, nil
}

// investingHD - investing HD
func (client *Client) investingHD(ctx context.Context, hostname string, url string, sd string, ed string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_INVESTING,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Investing{
			Investing: &jarviscrawlercore.RequestInvesting{
				Mode:      jarviscrawlercore.InvestingMode_INVESTINGMODE_HD,
				Url:       url,
				StartData: sd,
				EndData:   ed,
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

// InvestingHD - investing HD
func (client *Client) InvestingHD(ctx context.Context, url string, sd string, ed string, timeout int) (
	*jarviscrawlercore.ReplyInvesting_Asset, error) {

	hostname := hostInvesting

	_, reply, err := client.investingHD(ctx, hostname, url, sd, ed, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_INVESTING {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	investing, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Investing)
	if !isok {
		return nil, ErrNoReplyInvesting
	}

	if investing.Investing.Mode != jarviscrawlercore.InvestingMode_INVESTINGMODE_HD {
		return nil, ErrInvalidInvestingMode
	}

	res, isok := (investing.Investing.Reply).(*jarviscrawlercore.ReplyInvesting_Asset)
	if !isok {
		return nil, ErrNoReplyInvestingAsset
	}

	return res, nil
}
