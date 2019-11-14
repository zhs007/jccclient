package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/proto"
)

// taobaoProduct - taobao product
func (client *Client) taobaoProduct(ctx context.Context, hostname string, itemid string,
	timeout int) (*jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_TAOBAO,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Taobao{
			Taobao: &jarviscrawlercore.RequestTaobao{
				Mode:   jarviscrawlercore.TaobaoMode_TBM_PRODUCT,
				ItemID: itemid,
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

// TaobaoProduct - taobao product
func (client *Client) TaobaoProduct(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.TaobaoProduct, error) {

	hostname := "taobao.com"

	reply, err := client.taobaoProduct(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_TAOBAO {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	taobao, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Taobao)
	if !isok {
		return nil, ErrNoReplyTmall
	}

	if taobao.Taobao.Mode != jarviscrawlercore.TaobaoMode_TBM_PRODUCT {
		return nil, ErrInvalidTaobaoMode
	}

	replyproduct, isok := (taobao.Taobao.Reply).(*jarviscrawlercore.ReplyTaobao_Product)
	if !isok {
		return nil, ErrNoReplyTaobaoProduct
	}

	return replyproduct.Product, nil
}
