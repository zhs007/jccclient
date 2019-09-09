package jccclient

import (
	"bytes"
	"context"
	"errors"
	"io"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
	jarviscrawlercore "github.com/zhs007/jccclient/proto"
)

// Client - client
type Client struct {
	Hosts    *HostInfoCollection
	Running  bool
	cfg      *Config
	servAddr string
	token    string
	conn     *grpc.ClientConn
	client   jarviscrawlercore.JarvisCrawlerServiceClient
	tags     []string
}

// NewClient - new Client
func NewClient(db *DB, servAddr string, token string) *Client {
	return &Client{
		Hosts:    NewHostInfoCollection(db, servAddr),
		servAddr: servAddr,
		token:    token,
	}
}

// reset - reset
func (client *Client) reset() {
	if client.conn != nil {
		client.conn.Close()
	}

	client.conn = nil
	client.client = nil
}

// AnalyzePage - analyze a web page
func (client *Client) AnalyzePage(ctx context.Context, url string, viewport *Viewport,
	options *AnalyzePageOptions) (*jarviscrawlercore.ReplyAnalyzePage, error) {

	hostname := ""

	if client.cfg != nil {
		var err error
		hostname, err = GetHostName(url)
		if err != nil {
			return nil, err
		}
	}

	reply, err := client.analyzePage(ctx, hostname, url, viewport, options)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_ANALYZEPAGE {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	ap, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_AnalyzePage)
	if !isok {
		return nil, ErrNoReplyAnalyzePage
	}

	rap := ap.AnalyzePage

	return rap, nil
}

// GetGeoIP - get ip geolocation
func (client *Client) GetGeoIP(ctx context.Context, ip string, platform string) (*jarviscrawlercore.ReplyGeoIP, error) {

	hostname := "www.ipvoid.com"

	reply, err := client.getGeoIP(ctx, hostname, ip, platform)
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

// GetTechInAsiaJob - get techinasia job
func (client *Client) GetTechInAsiaJob(ctx context.Context, jobcode string, timeout int) (
	*jarviscrawlercore.TechInAsiaJob, error) {

	hostname := "www.techinasia.com"

	reply, err := client.getTechInAsiaJob(ctx, hostname, jobcode, timeout)
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

	reply, err := client.getTechInAsiaJobList(ctx, hostname, maintag, subtag, jobnums, timeout)
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

	reply, err := client.getTechInAsiaJobTagList(ctx, hostname, maintag, timeout)
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

	reply, err := client.getTechInAsiaCompany(ctx, hostname, companycode, timeout)
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

// GetSteepAndCheapProducts - get steep&cheap products
func (client *Client) GetSteepAndCheapProducts(ctx context.Context, url string, page int, timeout int) (
	*jarviscrawlercore.SteepAndCheapProducts, error) {

	hostname := "www.steepandcheap.com"

	reply, err := client.getSteepAndCheapProducts(ctx, hostname, url, page, timeout)
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

	reply, err := client.getSteepAndCheapProduct(ctx, hostname, url, timeout)
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

// GetJRJFunds - get jrj funds
func (client *Client) GetJRJFunds(ctx context.Context, timeout int) (
	*jarviscrawlercore.JRJFunds, error) {

	hostname := "www.jrj.com"

	reply, err := client.getJRJFunds(ctx, hostname, timeout)
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

	reply, err := client.getJRJFund(ctx, hostname, code, timeout)
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

	reply, err := client.getJRJFundManager(ctx, hostname, code, timeout)
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

	reply, err := client.getJRJFundValue(ctx, hostname, code, year, timeout)
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

// RequestCrawler -
func (client *Client) RequestCrawler(ctx context.Context, req *jarviscrawlercore.RequestCrawler) (
	*jarviscrawlercore.ReplyCrawler, error) {

	err := client.startRequest()
	if err != nil {
		client.reset()

		return nil, err
	}

	stream, err := client.client.RequestCrawler(ctx, req)
	if err != nil {
		client.reset()

		return nil, err
	}

	var recverr error
	var lstrect []*jarviscrawlercore.ReplyCrawlerStream
	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)

				return
			}

			if err != nil {
				recverr = err

				close(waitc)

				return
			}

			if in.Error != "" {
				recverr = errors.New(in.Error)

				close(waitc)

				return
			}

			lstrect = append(lstrect, in)
		}
	}()

	<-waitc

	if recverr != nil {
		client.reset()

		return nil, recverr
	}

	if len(lstrect) == 1 {
		return lstrect[0].ReplyCrawler, nil
	}

	var lstbytes [][]byte
	for _, v := range lstrect {
		lstbytes = append(lstbytes, v.Data)
	}

	buf := bytes.Join(lstbytes, nil)

	rc := &jarviscrawlercore.ReplyCrawler{}
	err = proto.Unmarshal(buf, rc)
	if err != nil {
		return nil, err
	}

	return rc, nil

}

// startRequest -
func (client *Client) startRequest() error {
	if client.servAddr == "" {
		return ErrNoServAddr
	}

	if client.token == "" {
		return ErrNoToken
	}

	if client.conn == nil || client.client == nil {
		conn, err := grpc.Dial(client.servAddr, grpc.WithInsecure())
		if err != nil {
			return err
		}

		client.conn = conn

		client.client = jarviscrawlercore.NewJarvisCrawlerServiceClient(conn)
	}

	return nil
}

// analyzePage - analyze a web page
func (client *Client) analyzePage(ctx context.Context, hostname string, url string, viewport *Viewport,
	options *AnalyzePageOptions) (*jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_ANALYZEPAGE,
		CrawlerParam: &jarviscrawlercore.RequestCrawler_AnalyzePage{
			AnalyzePage: &jarviscrawlercore.AnalyzePage{
				Url:               url,
				Delay:             int32(options.ScreenshotsDelay),
				ViewportWidth:     int32(viewport.Width),
				ViewportHeight:    int32(viewport.Height),
				DeviceScaleFactor: viewport.DeviceScaleFactor,
				IsMobile:          viewport.IsMobile,
				IsLandscape:       viewport.IsLandscape,
				NeedScreenshots:   options.NeedScreenshots,
				NeedLogs:          options.NeedLogs,
				Timeout:           int32(options.Timeout),
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

// getGeoIP - get ip geolocation
func (client *Client) getGeoIP(ctx context.Context, hostname string, ip string,
	platform string) (*jarviscrawlercore.ReplyCrawler, error) {

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

// getTechInAsiaJob - get techinasia job
func (client *Client) getTechInAsiaJob(ctx context.Context, hostname string, jobcode string, timeout int) (
	*jarviscrawlercore.ReplyCrawler, error) {

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

// getTechInAsiaCompany - get techinasia company
func (client *Client) getTechInAsiaCompany(ctx context.Context, hostname string, companycode string, timeout int) (
	*jarviscrawlercore.ReplyCrawler, error) {

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

// getTechInAsiaJobList - get techinasia jobs
func (client *Client) getTechInAsiaJobList(ctx context.Context, hostname string, maintag string, subtag string,
	jobnums int, timeout int) (*jarviscrawlercore.ReplyCrawler, error) {

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

// getTechInAsiaJobTagList - get techinasia job tag list
func (client *Client) getTechInAsiaJobTagList(ctx context.Context, hostname string, maintag string, timeout int) (
	*jarviscrawlercore.ReplyCrawler, error) {

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

// getSteepAndCheapProducts - get steepandcheap products
func (client *Client) getSteepAndCheapProducts(ctx context.Context, hostname string, url string, page int, timeout int) (
	*jarviscrawlercore.ReplyCrawler, error) {

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

// getSteepAndCheapProduct - get steepandcheap product
func (client *Client) getSteepAndCheapProduct(ctx context.Context, hostname string, url string, timeout int) (
	*jarviscrawlercore.ReplyCrawler, error) {

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

// getJRJFunds - get jrj funds
func (client *Client) getJRJFunds(ctx context.Context, hostname string, timeout int) (
	*jarviscrawlercore.ReplyCrawler, error) {

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

// getJRJFund - get jrj fund
func (client *Client) getJRJFund(ctx context.Context, hostname string, code string, timeout int) (
	*jarviscrawlercore.ReplyCrawler, error) {

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

// getJRJFundManager - get jrj fund manager
func (client *Client) getJRJFundManager(ctx context.Context, hostname string, code string, timeout int) (
	*jarviscrawlercore.ReplyCrawler, error) {

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

// getJRJFundValue - get jrj fund value
func (client *Client) getJRJFundValue(ctx context.Context, hostname string, code string, year string,
	timeout int) (*jarviscrawlercore.ReplyCrawler, error) {

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
