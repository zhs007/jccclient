package jccclient

import (
	"context"

	jarviscrawlercore "github.com/zhs007/jccclient/pb"
)

// manhuadbAuthor - manhuadb author
func (client *Client) manhuadbAuthor(ctx context.Context, hostname string,
	authorid string, timeout int) (
	string, *jarviscrawlercore.ReplyCrawler, error) {

	if client.cfg != nil {
		client.Hosts.OnTaskStart(ctx, hostname, client.cfg)
	}

	req := &jarviscrawlercore.RequestCrawler{
		Token:       client.token,
		CrawlerType: jarviscrawlercore.CrawlerType_CT_MANHUADB,
		Timeout:     int32(timeout),
		CrawlerParam: &jarviscrawlercore.RequestCrawler_Manhuadb{
			Manhuadb: &jarviscrawlercore.RequestManhuaDB{
				Mode:     jarviscrawlercore.ManhuaDBMode_MHDB_AUTHOR,
				Authorid: authorid,
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

// ManhuaDBAuthor - ManhuaDB author
func (client *Client) ManhuaDBAuthor(ctx context.Context, authorid string, timeout int) (
	*jarviscrawlercore.ManhuaDBAuthor, error) {

	hostname := "manhuadb.com"

	_, reply, err := client.manhuadbAuthor(ctx, hostname, authorid, timeout)
	if err != nil {
		return nil, err
	}

	if reply.CrawlerType != jarviscrawlercore.CrawlerType_CT_MANHUADB {
		return nil, ErrInvalidCrawlerType
	}

	if reply.CrawlerResult == nil {
		return nil, ErrNoCrawlerResult
	}

	manhuadb, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Manhuadb)
	if !isok {
		return nil, ErrNoReplyManhuaDB
	}

	if manhuadb.Manhuadb.Mode != jarviscrawlercore.ManhuaDBMode_MHDB_AUTHOR {
		return nil, ErrInvalidManhuaDBMode
	}

	replymanhuadb, isok := (manhuadb.Manhuadb.Reply).(*jarviscrawlercore.ReplyManhuaDB_Author)
	if !isok {
		return nil, ErrNoReplyManhuaDBAuthor
	}

	return replymanhuadb.Author, nil
}
