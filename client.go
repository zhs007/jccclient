package jccclient

import (
	"bytes"
	"context"
	"errors"
	"io"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
	jarviscrawlercore "github.com/zhs007/jccclient/pb"
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
	Version  string
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

	_, reply, err := client.analyzePage(ctx, hostname, url, viewport, options)
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

// RequestCrawler -
func (client *Client) RequestCrawler(ctx context.Context, req *jarviscrawlercore.RequestCrawler) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	version := ""

	err := client.startRequest()
	if err != nil {
		client.reset()

		return version, nil, err
	}

	stream, err := client.client.RequestCrawler(ctx, req)
	if err != nil {
		client.reset()

		return version, nil, err
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

			version = in.Version

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

		return version, nil, recverr
	}

	if len(lstrect) == 1 {
		return version, lstrect[0].ReplyCrawler, nil
	}

	var lstbytes [][]byte
	for _, v := range lstrect {
		lstbytes = append(lstbytes, v.Data)
	}

	buf := bytes.Join(lstbytes, nil)

	rc := &jarviscrawlercore.ReplyCrawler{}
	err = proto.Unmarshal(buf, rc)
	if err != nil {
		return version, nil, err
	}

	return version, rc, nil

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
	options *AnalyzePageOptions) (string, *jarviscrawlercore.ReplyCrawler, error) {

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
