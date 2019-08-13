package jccclient

import (
	"context"
	"io"

	"google.golang.org/grpc"

	jarviscrawlercore "github.com/zhs007/jccclient/proto"
)

// Client - client
type Client struct {
	servAddr string
	token    string
	conn     *grpc.ClientConn
	client   jarviscrawlercore.JarvisCrawlerServiceClient
}

// NewClient - new Client
func NewClient(servAddr string, token string) *Client {
	return &Client{
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

// AnalyzePage -
func (client *Client) AnalyzePage(ctx context.Context, url string, delay int, w int, h int,
	deviceScaleFactor float32, isMobile bool, isLandscape bool) (

	*jarviscrawlercore.ReplyAnalyzePage, error) {

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_ANALYZEPAGE,
		CrawlerParam: &jarviscrawlercore.RequestCrawler_AnalyzePage{
			AnalyzePage: &jarviscrawlercore.AnalyzePage{
				Url:               url,
				Delay:             int32(delay),
				ViewportWidth:     int32(w),
				ViewportHeight:    int32(h),
				DeviceScaleFactor: deviceScaleFactor,
				IsMobile:          isMobile,
				IsLandscape:       isLandscape,
			},
		},
	}

	reply, err := client.RequestCrawler(ctx, req)
	if err != nil {
		return nil, err
	}

	if reply == nil {
		return nil, ErrNoReplyCrawler
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

	return nil, nil
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
