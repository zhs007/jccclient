package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// GetTechInAsiaJob - get techinasia job
func (client *Client) GetTechInAsiaJob(ctx context.Context, jobcode string, timeout int) (
	*jarviscrawlercore.TechInAsiaJob, error) {

	hostname := "www.techinasia.com"

	_, reply, err := client.getTechInAsiaJob(ctx, hostname, jobcode, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_TECHINASIA {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	techinasia, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Techinasia)
	if !isok {
		return nil, ErrNoReplyTechInAsia
	}

	if techinasia.Techinasia.Mode != jarviscrawlercore.TechInAsiaMode_TIAM_JOB {
		return nil, ErrInvalidTechInAsiaMode
	}

	replyjob, isok := (techinasia.Techinasia.Reply).(*jarviscrawlercore.ReplyTechInAsia_Job)
	if !isok {
		return nil, ErrNoReplyTechInAsiaJob
	}

	return replyjob.Job, nil
}

// GetTechInAsiaJobList - get techinasia jobs
func (client *Client) GetTechInAsiaJobList(ctx context.Context, maintag string, subtag string, jobnums int,
	timeout int) (*jarviscrawlercore.TechInAsiaJobList, error) {

	hostname := "www.techinasia.com"

	_, reply, err := client.getTechInAsiaJobList(ctx, hostname, maintag, subtag, jobnums, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_TECHINASIA {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	techinasia, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Techinasia)
	if !isok {
		return nil, ErrNoReplyTechInAsia
	}

	if techinasia.Techinasia.Mode != jarviscrawlercore.TechInAsiaMode_TIAM_JOBLIST {
		return nil, ErrInvalidTechInAsiaMode
	}

	replyjobs, isok := (techinasia.Techinasia.Reply).(*jarviscrawlercore.ReplyTechInAsia_Jobs)
	if !isok {
		return nil, ErrNoReplyTechInAsiaJobs
	}

	return replyjobs.Jobs, nil
}

// GetTechInAsiaJobTagList - get techinasia job tag list
func (client *Client) GetTechInAsiaJobTagList(ctx context.Context, maintag string, timeout int) (*jarviscrawlercore.TechInAsiaJobTagList, error) {

	hostname := "www.techinasia.com"

	_, reply, err := client.getTechInAsiaJobTagList(ctx, hostname, maintag, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_TECHINASIA {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	techinasia, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Techinasia)
	if !isok {
		return nil, ErrNoReplyTechInAsia
	}

	if techinasia.Techinasia.Mode != jarviscrawlercore.TechInAsiaMode_TIAM_JOBTAG {
		return nil, ErrInvalidTechInAsiaMode
	}

	replytags, isok := (techinasia.Techinasia.Reply).(*jarviscrawlercore.ReplyTechInAsia_Tags)
	if !isok {
		return nil, ErrNoReplyTechInAsiaJobTags
	}

	return replytags.Tags, nil
}

// GetTechInAsiaCompany - get techinasia company
func (client *Client) GetTechInAsiaCompany(ctx context.Context, companycode string, timeout int) (
	*jarviscrawlercore.TechInAsiaCompany, error) {

	hostname := "www.techinasia.com"

	_, reply, err := client.getTechInAsiaCompany(ctx, hostname, companycode, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_TECHINASIA {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	techinasia, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Techinasia)
	if !isok {
		return nil, ErrNoReplyTechInAsia
	}

	if techinasia.Techinasia.Mode != jarviscrawlercore.TechInAsiaMode_TIAM_COMPANY {
		return nil, ErrInvalidTechInAsiaMode
	}

	replycompany, isok := (techinasia.Techinasia.Reply).(*jarviscrawlercore.ReplyTechInAsia_Company)
	if !isok {
		return nil, ErrNoReplyTechInAsiaCompany
	}

	return replycompany.Company, nil
}

// getTechInAsiaJob - get techinasia job
func (client *Client) getTechInAsiaJob(ctx context.Context, hostname string, jobcode string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_TECHINASIA,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Techinasia{
			Techinasia: &jarviscrawlercore.RequestTechInAsia{
				Mode:    jarviscrawlercore.TechInAsiaMode_TIAM_JOB,
				JobCode: jobcode,
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

// getTechInAsiaCompany - get techinasia company
func (client *Client) getTechInAsiaCompany(ctx context.Context, hostname string, companycode string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_TECHINASIA,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Techinasia{
			Techinasia: &jarviscrawlercore.RequestTechInAsia{
				Mode:        jarviscrawlercore.TechInAsiaMode_TIAM_COMPANY,
				CompanyCode: companycode,
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

// getTechInAsiaJobList - get techinasia jobs
func (client *Client) getTechInAsiaJobList(ctx context.Context, hostname string, maintag string, subtag string,
	jobnums int, timeout int) (string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_TECHINASIA,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Techinasia{
			Techinasia: &jarviscrawlercore.RequestTechInAsia{
				Mode:      jarviscrawlercore.TechInAsiaMode_TIAM_JOBLIST,
				JobNums:   int32(jobnums),
				JobTag:    maintag,
				JobSubTag: subtag,
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

// getTechInAsiaJobTagList - get techinasia job tag list
func (client *Client) getTechInAsiaJobTagList(ctx context.Context, hostname string, maintag string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_TECHINASIA,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Techinasia{
			Techinasia: &jarviscrawlercore.RequestTechInAsia{
				Mode:   jarviscrawlercore.TechInAsiaMode_TIAM_JOBTAG,
				JobTag: maintag,
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
