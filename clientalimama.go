package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/proto"
)

// alimamaKeepalive - alimama keepalive
func (client *Client) alimamaKeepalive(ctx context.Context, hostname string,
	timeout int) (*jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_ALIMAMA,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Alimama{
			Alimama: &jarviscrawlercore.RequestAlimama{
				Mode: jarviscrawlercore.AlimamaMode_ALIMMM_KEEPALIVE,
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

// alimamaGetTop - alimama gettop
func (client *Client) alimamaGetTop(ctx context.Context, hostname string,
	timeout int) (*jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_ALIMAMA,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Alimama{
			Alimama: &jarviscrawlercore.RequestAlimama{
				Mode: jarviscrawlercore.AlimamaMode_ALIMMM_GETTOP,
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

// alimamaSearch - alimama search
func (client *Client) alimamaSearch(ctx context.Context, hostname string, text string,
	timeout int) (*jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_ALIMAMA,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Alimama{
			Alimama: &jarviscrawlercore.RequestAlimama{
				Mode: jarviscrawlercore.AlimamaMode_ALIMMM_SEARCH,
				Text: text,
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

// AlimamaKeepalive - alimama keepalive
func (client *Client) AlimamaKeepalive(ctx context.Context, timeout int) error {

	hostname := "alimama.com"

	reply, err := client.alimamaKeepalive(ctx, hostname, timeout)
	if err != nil {
		return err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_ALIMAMA {
		return ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return ErrNoCrawlerResult
	}

	alimama, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Alimama)
	if !isok {
		return ErrNoReplyAlimama
	}

	if alimama.Alimama.Mode != jarviscrawlercore.AlimamaMode_ALIMMM_KEEPALIVE {
		return ErrInvalidAlimamaMode
	}

	return nil
}

// AlimamaGetTop - alimama gettop
func (client *Client) AlimamaGetTop(ctx context.Context, timeout int) (
	*jarviscrawlercore.AlimamaTopInfo, error) {

	hostname := "alimama.com"

	reply, err := client.alimamaGetTop(ctx, hostname, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_ALIMAMA {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	alimama, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Alimama)
	if !isok {
		return nil, ErrNoReplyAlimama
	}

	if alimama.Alimama.Mode != jarviscrawlercore.AlimamaMode_ALIMMM_GETTOP {
		return nil, ErrInvalidAlimamaMode
	}

	replytopinfo, isok := (alimama.Alimama.Reply).(*jarviscrawlercore.ReplyAlimama_TopInfo)
	if !isok {
		return nil, ErrNoReplyAlimamaGetTop
	}

	return replytopinfo.TopInfo, nil
}

// AlimamaSearch - alimama search
func (client *Client) AlimamaSearch(ctx context.Context, text string, timeout int) (
	*jarviscrawlercore.AlimamaProducts, error) {

	hostname := "alimama.com"

	reply, err := client.alimamaSearch(ctx, hostname, text, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_ALIMAMA {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	alimama, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Alimama)
	if !isok {
		return nil, ErrNoReplyAlimama
	}

	if alimama.Alimama.Mode != jarviscrawlercore.AlimamaMode_ALIMMM_SEARCH {
		return nil, ErrInvalidAlimamaMode
	}

	replyproducts, isok := (alimama.Alimama.Reply).(*jarviscrawlercore.ReplyAlimama_Products)
	if !isok {
		return nil, ErrNoReplyAlimamaProducts
	}

	return replyproducts.Products, nil
}
