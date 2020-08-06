package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// HOST - 6vdy.org
const HOST = "6vdy.org"

// p6vdyMovies - 6vdy movies
func (client *Client) p6vdyMovies(ctx context.Context, hostname string, url string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_6VDY,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_P6Vdy{
			P6Vdy: &jarviscrawlercore.RequestP6Vdy{
				Mode: jarviscrawlercore.P6VdyMode_P6VDY_MOVIES,
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

// P6vdyMovies - 6vdy movie
func (client *Client) P6vdyMovies(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.ReplyP6Vdy_Movies, error) {

	hostname := HOST

	_, reply, err := client.p6vdyMovies(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_6VDY {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	p6vdy, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_P6Vdy)
	if !isok {
		return nil, ErrNoReplyP6vdy
	}

	if p6vdy.P6Vdy.Mode != jarviscrawlercore.P6VdyMode_P6VDY_MOVIES {
		return nil, ErrInvalidHao6vMode
	}

	movies, isok := (p6vdy.P6Vdy.Reply).(*jarviscrawlercore.ReplyP6Vdy_Movies)
	if !isok {
		return nil, ErrNoReplyP6vdyMovies
	}

	return movies, nil
}

// p6vdyMovie - 6vdy movie
func (client *Client) p6vdyMovie(ctx context.Context, hostname string, url string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_6VDY,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_P6Vdy{
			P6Vdy: &jarviscrawlercore.RequestP6Vdy{
				Mode: jarviscrawlercore.P6VdyMode_P6VDY_MOVIE,
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

// P6vdyMovie - 6vdy movie
func (client *Client) P6vdyMovie(ctx context.Context, url string, timeout int) (
	*jarviscrawlercore.ReplyP6Vdy_Movie, error) {

	hostname := HOST

	_, reply, err := client.p6vdyMovie(ctx, hostname, url, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_6VDY {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	p6vdy, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_P6Vdy)
	if !isok {
		return nil, ErrNoReplyP6vdy
	}

	if p6vdy.P6Vdy.Mode != jarviscrawlercore.P6VdyMode_P6VDY_MOVIE {
		return nil, ErrInvalidP6vdyMode
	}

	res, isok := (p6vdy.P6Vdy.Reply).(*jarviscrawlercore.ReplyP6Vdy_Movie)
	if !isok {
		return nil, ErrNoReplyP6vdyMovie
	}

	return res, nil
}
