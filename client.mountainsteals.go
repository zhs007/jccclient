package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/proto"
)

// mountainstealsSale - mountainsteals sale
func (client *Client) mountainstealsSale(ctx context.Context, hostname string,
	url string, timeout int) (*jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_MOUNTAINSTEALS,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Mountainsteals{
			Mountainsteals: &jarviscrawlercore.RequestMountainSteals{
				Mode: jarviscrawlercore.MountainStealsMode_MSM_SALE,
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

// mountainstealsProduct - mountainsteals product
func (client *Client) mountainstealsProduct(ctx context.Context, hostname string,
	url string, timeout int) (*jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_MOUNTAINSTEALS,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Mountainsteals{
			Mountainsteals: &jarviscrawlercore.RequestMountainSteals{
				Mode: jarviscrawlercore.MountainStealsMode_MSM_PRODUCT,
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

// MountainStealsSale - mountainsteals sale
func (client *Client) MountainStealsSale(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.MountainStealsSale, error) {

	hostname := "mountainsteals.com"

	reply, err := client.mountainstealsSale(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_MOUNTAINSTEALS {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	ms, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Mountainsteals)
	if !isok {
		return nil, ErrNoReplyTmall
	}

	if ms.Mountainsteals.Mode != jarviscrawlercore.MountainStealsMode_MSM_SALE {
		return nil, ErrInvalidMountainstealsMode
	}

	replysale, isok := (ms.Mountainsteals.Reply).(*jarviscrawlercore.ReplyMountainSteals_Sale)
	if !isok {
		return nil, ErrNoReplyMountainstealsSale
	}

	return replysale.Sale, nil
}

// MountainStealsProduct - mountainsteals product
func (client *Client) MountainStealsProduct(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.MountainStealsProduct, error) {

	hostname := "mountainsteals.com"

	reply, err := client.mountainstealsProduct(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_MOUNTAINSTEALS {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	ms, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Mountainsteals)
	if !isok {
		return nil, ErrNoReplyTmall
	}

	if ms.Mountainsteals.Mode != jarviscrawlercore.MountainStealsMode_MSM_PRODUCT {
		return nil, ErrInvalidMountainstealsMode
	}

	replyproduct, isok := (ms.Mountainsteals.Reply).(*jarviscrawlercore.ReplyMountainSteals_Product)
	if !isok {
		return nil, ErrNoReplyMountainstealsProduct
	}

	return replyproduct.Product, nil
}
