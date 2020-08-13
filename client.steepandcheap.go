package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// GetSteepAndCheapProducts - get steep&cheap products
func (client *Client) GetSteepAndCheapProducts(ctx context.Context, url string, page int, timeout int) (
	*jarviscrawlercore.SteepAndCheapProducts, error) {

	hostname := "www.steepandcheap.com"

	_, reply, err := client.getSteepAndCheapProducts(ctx, hostname, url, page, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_STEEPANDCHEAP {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	steepandcheap, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Steepandcheap)
	if !isok {
		return nil, ErrNoReplySteepAndCheap
	}

	if steepandcheap.Steepandcheap.Mode != jarviscrawlercore.SteepAndCheapMode_SACM_PRODUCTS {
		return nil, ErrInvalidSteepAndCheapMode
	}

	replyproducts, isok := (steepandcheap.Steepandcheap.Reply).(*jarviscrawlercore.ReplySteepAndCheap_Products)
	if !isok {
		return nil, ErrNoReplySteepAndCheapProducts
	}

	return replyproducts.Products, nil
}

// GetSteepAndCheapProduct - get steep&cheap product
func (client *Client) GetSteepAndCheapProduct(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.SteepAndCheapProduct, error) {

	hostname := "www.steepandcheap.com"

	_, reply, err := client.getSteepAndCheapProduct(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_STEEPANDCHEAP {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	steepandcheap, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Steepandcheap)
	if !isok {
		return nil, ErrNoReplySteepAndCheap
	}

	if steepandcheap.Steepandcheap.Mode != jarviscrawlercore.SteepAndCheapMode_SACM_PRODUCT {
		return nil, ErrInvalidSteepAndCheapMode
	}

	replyproduct, isok := (steepandcheap.Steepandcheap.Reply).(*jarviscrawlercore.ReplySteepAndCheap_Product)
	if !isok {
		return nil, ErrNoReplySteepAndCheapProduct
	}

	return replyproduct.Product, nil
}

// getSteepAndCheapProducts - get steepandcheap products
func (client *Client) getSteepAndCheapProducts(ctx context.Context, hostname string, url string, page int, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_STEEPANDCHEAP,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Steepandcheap{
			Steepandcheap: &jarviscrawlercore.RequestSteepAndCheap{
				Mode: jarviscrawlercore.SteepAndCheapMode_SACM_PRODUCTS,
				Url:  url,
				Page: int32(page),
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

// getSteepAndCheapProduct - get steepandcheap product
func (client *Client) getSteepAndCheapProduct(ctx context.Context, hostname string, url string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_STEEPANDCHEAP,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Steepandcheap{
			Steepandcheap: &jarviscrawlercore.RequestSteepAndCheap{
				Mode: jarviscrawlercore.SteepAndCheapMode_SACM_PRODUCT,
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
