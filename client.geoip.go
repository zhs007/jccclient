package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// GetGeoIP - get ip geolocation
func (client *Client) GetGeoIP(ctx context.Context, ip string, platform string) (*jarviscrawlercore.ReplyGeoIP, error) {

	hostname := "www.ipvoid.com"

	_, reply, err := client.getGeoIP(ctx, hostname, ip, platform)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_GEOIP {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	ap, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Geoip)
	if !isok {
		return nil, ErrNoReplyGeoip
	}

	rap := ap.Geoip

	return rap, nil
}

// getGeoIP - get ip geolocation
func (client *Client) getGeoIP(ctx context.Context, hostname string, ip string,
	platform string) (string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_GEOIP,
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Geoip{
			Geoip: &jarviscrawlercore.RequestGeoIP{
				Ip:       ip,
				Platform: platform,
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
