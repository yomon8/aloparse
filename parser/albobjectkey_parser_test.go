package parser

import (
	"testing"
)

func TestAlbObjectKeyParser(t *testing.T) {
	cases := []struct {
		object    string
		accountID string
		region    string
		albName   string
		albIP     string
		filename  string
	}{
		{
			"s3://my-bucket/prefix/AWSLogs/123456789012/elasticloadbalancing/us-east-2/2016/05/01/123456789012_elasticloadbalancing_us-east-2_my-loadbalancer_20140215T2340Z_172.160.1.192_20sg8hgm.log.gz",
			"123456789012",
			"us-east-2",
			"my-loadbalancer",
			"172.160.1.192",
			"123456789012_elasticloadbalancing_us-east-2_my-loadbalancer_20140215T2340Z_172.160.1.192_20sg8hgm.log.gz",
		},
	}
	p := NewALBObjectKeyParser()
	for _, c := range cases {
		o, err := p.Parse(c.object)
		if err != nil {
			t.Fatalf("parse error: %#v\n", err)
		}
		if o.AccountID != c.accountID {
			t.Fatalf("parse error AccountId:%s:%s\n", o.AccountID, c.accountID)
		}
		if o.Region != c.region {
			t.Fatalf("parse error Regin:%s:%s\n", o.Region, c.region)
		}
		if o.ALBName != c.albName {
			t.Fatalf("parse error ALBName:%s:%s\n", o.ALBName, c.albName)
		}
		if o.ALBIP != c.albIP {
			t.Fatalf("parse error ALBIP:%s:%s\n", o.ALBIP, c.albIP)
		}
		if o.FileName != c.filename {
			t.Fatalf("parse error FileName:%s:%s\n", o.FileName, c.filename)
		}
	}

}
