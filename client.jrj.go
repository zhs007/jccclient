package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// GetJRJFunds - get jrj funds
func (client *Client) GetJRJFunds(ctx context.Context, timeout int) (
	*jarviscrawlercore.JRJFunds, error) {

	hostname := "www.jrj.com"

	_, reply, err := client.getJRJFunds(ctx, hostname, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_JRJ {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	jrj, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Jrj)
	if !isok {
		return nil, ErrNoReplyJRJ
	}

	if jrj.Jrj.Mode != jarviscrawlercore.JRJMode_JRJM_FUNDS {
		return nil, ErrInvalidJRJMode
	}

	replyfunds, isok := (jrj.Jrj.Reply).(*jarviscrawlercore.ReplyJRJ_Funds)
	if !isok {
		return nil, ErrNoReplyJRJFunds
	}

	return replyfunds.Funds, nil
}

// GetJRJFund - get jrj fund
func (client *Client) GetJRJFund(ctx context.Context, code string, timeout int) (
	*jarviscrawlercore.JRJFund, error) {

	hostname := "www.jrj.com"

	_, reply, err := client.getJRJFund(ctx, hostname, code, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_JRJ {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	jrj, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Jrj)
	if !isok {
		return nil, ErrNoReplyJRJ
	}

	if jrj.Jrj.Mode != jarviscrawlercore.JRJMode_JRJM_FUND {
		return nil, ErrInvalidJRJMode
	}

	replyfund, isok := (jrj.Jrj.Reply).(*jarviscrawlercore.ReplyJRJ_Fund)
	if !isok {
		return nil, ErrNoReplyJRJFund
	}

	return replyfund.Fund, nil
}

// GetJRJFundManager - get jrj fund manager
func (client *Client) GetJRJFundManager(ctx context.Context, code string, timeout int) (
	*jarviscrawlercore.JRJFund, error) {

	hostname := "www.jrj.com"

	_, reply, err := client.getJRJFundManager(ctx, hostname, code, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_JRJ {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	jrj, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Jrj)
	if !isok {
		return nil, ErrNoReplyJRJ
	}

	if jrj.Jrj.Mode != jarviscrawlercore.JRJMode_JRJM_FUNDMANAGER {
		return nil, ErrInvalidJRJMode
	}

	replyfund, isok := (jrj.Jrj.Reply).(*jarviscrawlercore.ReplyJRJ_Fund)
	if !isok {
		return nil, ErrNoReplyJRJFund
	}

	return replyfund.Fund, nil
}

// GetJRJFundValue - get jrj fund value
func (client *Client) GetJRJFundValue(ctx context.Context, code string, year string, timeout int) (
	*jarviscrawlercore.JRJFundValue, error) {

	hostname := "www.jrj.com"

	_, reply, err := client.getJRJFundValue(ctx, hostname, code, year, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_JRJ {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	jrj, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Jrj)
	if !isok {
		return nil, ErrNoReplyJRJ
	}

	if jrj.Jrj.Mode != jarviscrawlercore.JRJMode_JRJM_FUNDVALUE {
		return nil, ErrInvalidJRJMode
	}

	replyfundvalue, isok := (jrj.Jrj.Reply).(*jarviscrawlercore.ReplyJRJ_FundValue)
	if !isok {
		return nil, ErrNoReplyJRJFundValue
	}

	return replyfundvalue.FundValue, nil
}

// getJRJFunds - get jrj funds
func (client *Client) getJRJFunds(ctx context.Context, hostname string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_JRJ,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Jrj{
			Jrj: &jarviscrawlercore.RequestJRJ{
				Mode: jarviscrawlercore.JRJMode_JRJM_FUNDS,
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

// getJRJFund - get jrj fund
func (client *Client) getJRJFund(ctx context.Context, hostname string, code string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_JRJ,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Jrj{
			Jrj: &jarviscrawlercore.RequestJRJ{
				Mode:     jarviscrawlercore.JRJMode_JRJM_FUND,
				Fundcode: code,
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

// getJRJFundManager - get jrj fund manager
func (client *Client) getJRJFundManager(ctx context.Context, hostname string, code string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_JRJ,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Jrj{
			Jrj: &jarviscrawlercore.RequestJRJ{
				Mode:     jarviscrawlercore.JRJMode_JRJM_FUNDMANAGER,
				Fundcode: code,
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

// getJRJFundValue - get jrj fund value
func (client *Client) getJRJFundValue(ctx context.Context, hostname string, code string, year string,
	timeout int) (string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_JRJ,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Jrj{
			Jrj: &jarviscrawlercore.RequestJRJ{
				Mode:     jarviscrawlercore.JRJMode_JRJM_FUNDVALUE,
				Fundcode: code,
				Year:     year,
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
