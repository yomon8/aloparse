package parser

import (
	"testing"
)

func TestAlbLogParser(t *testing.T) {
	cases := []struct {
		line                   string
		typestring             string
		timestamp              string
		alb                    string
		client                 string
		clientPort             string
		target                 string
		targetPort             string
		requestProcessingTime  string
		targetProcessingTime   string
		responseProcessingTime string
		elbStatusCode          string
		targetStatusCode       string
		receivedBytes          string
		sentBytes              string
		method                 string
		url                    string
		protocol               string
		userAgent              string
		sslCipher              string
		sslProtocol            string
		targetGroupArn         string
		traceID                string
		domainName             string
		chosenCertArn          string
	}{
		{
			"http 2017-09-09T00:55:00.228196Z app/albname/12a34bc6d78e9f0 111.222.10.240:57965 192.168.131.118:80 0.020 0.019 0.001 200 400 623 111603 \"GET http://www.host.com:80/path/?a=1&b=2&c=3 HTTP/1.1\" \"Mozilla/5.0 (iPhone; CPU iPhone OS 10_2_1 like Mac OS X) AppleWebKit/602.4.6 (KHTML, like Gecko) Version/10.0 Mobile/14D27 Safari/602.1\" ECDHE-RSA-AES128-GCM-SHA256 TLSv1.2 arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:targetgroup/albname/abcdefgh12345678 \"Root=1-11111be1-1110300f111111111aa1a1aa\" \"www.domain.co.jp\" \"arn:aws:iam::12345567890:server-certificate/domain.co.jp\"",
			"http",
			"2017-09-09T00:55:00.228196Z",
			"app/albname/12a34bc6d78e9f0",
			"111.222.10.240",
			"57965",
			"192.168.131.118",
			"80",
			"0.020",
			"0.019",
			"0.001",
			"200",
			"400",
			"623",
			"111603",
			"GET",
			"http://www.host.com:80/path/?a=1&b=2&c=3",
			"HTTP/1.1",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 10_2_1 like Mac OS X) AppleWebKit/602.4.6 (KHTML, like Gecko) Version/10.0 Mobile/14D27 Safari/602.1",
			"ECDHE-RSA-AES128-GCM-SHA256",
			"TLSv1.2",
			"arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:targetgroup/albname/abcdefgh12345678",
			"Root=1-11111be1-1110300f111111111aa1a1aa",
			"www.domain.co.jp",
			"arn:aws:iam::12345567890:server-certificate/domain.co.jp",
		}, {
			"http 2016-08-10T22:08:42.945958Z app/my-loadbalancer/50dc6c495c0c9188 192.168.131.39:2817 10.0.0.1:80 0.000 0.001 0.000 200 200 34 366 \"GET http://www.example.com:80/ HTTP/1.1\" \"curl/7.46.0\" - - arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067 \"Root=1-58337262-36d228ad5d99923122bbe354\" - -",
			"http",
			"2016-08-10T22:08:42.945958Z",
			"app/my-loadbalancer/50dc6c495c0c9188",
			"192.168.131.39",
			"2817",
			"10.0.0.1",
			"80",
			"0.000",
			"0.001",
			"0.000",
			"200",
			"200",
			"34",
			"366",
			"GET",
			"http://www.example.com:80/",
			"HTTP/1.1",
			"curl/7.46.0",
			"-",
			"-",
			"arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067",
			"Root=1-58337262-36d228ad5d99923122bbe354",
			"-",
			"-",
		}, {
			"https 2016-08-10T23:39:43.065466Z app/my-loadbalancer/50dc6c495c0c9188 192.168.131.39:2817 10.0.0.1:80 0.086 0.048 0.037 200 200 0 57 \"GET https://www.example.com:443/ HTTP/1.1\" \"curl/7.46.0\" ECDHE-RSA-AES128-GCM-SHA256 TLSv1.2 arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067 \"Root=1-58337281-1d84f3d73c47ec4e58577259\" www.example.com arn:aws:acm:us-east-2:123456789012:certificate/12345678-1234-1234-1234-123456789012",
			"https",
			"2016-08-10T23:39:43.065466Z",
			"app/my-loadbalancer/50dc6c495c0c9188",
			"192.168.131.39",
			"2817",
			"10.0.0.1",
			"80",
			"0.086",
			"0.048",
			"0.037",
			"200",
			"200",
			"0",
			"57",
			"GET",
			"https://www.example.com:443/",
			"HTTP/1.1",
			"curl/7.46.0",
			"ECDHE-RSA-AES128-GCM-SHA256",
			"TLSv1.2",
			"arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067",
			"Root=1-58337281-1d84f3d73c47ec4e58577259",
			"www.example.com",
			"arn:aws:acm:us-east-2:123456789012:certificate/12345678-1234-1234-1234-123456789012",
		}, {
			"h2 2016-08-10T00:10:33.145057Z app/my-loadbalancer/50dc6c495c0c9188 10.0.1.252:48160 10.0.0.66:9000 0.000 0.002 0.000 200 200 5 257 \"GET https://10.0.2.105:773/ HTTP/2.0\" \"curl/7.46.0\" ECDHE-RSA-AES128-GCM-SHA256 TLSv1.2 arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067 \"Root=1-58337327-72bd00b0343d75b906739c42\" - -",
			"h2",
			"2016-08-10T00:10:33.145057Z",
			"app/my-loadbalancer/50dc6c495c0c9188",
			"10.0.1.252",
			"48160",
			"10.0.0.66",
			"9000",
			"0.000",
			"0.002",
			"0.000",
			"200",
			"200",
			"5",
			"257",
			"GET",
			"https://10.0.2.105:773/",
			"HTTP/2.0",
			"curl/7.46.0",
			"ECDHE-RSA-AES128-GCM-SHA256",
			"TLSv1.2",
			"arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067",
			"Root=1-58337327-72bd00b0343d75b906739c42",
			"-",
			"-",
		}, {
			"ws 2016-08-10T00:32:08.923954Z app/my-loadbalancer/50dc6c495c0c9188 10.0.0.140:40914 10.0.1.192:8010 0.001 0.003 0.000 101 101 218 587 \"GET http://10.0.0.30:80/ HTTP/1.1\" \"-\" - - arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067 \"Root=1-58337364-23a8c76965a2ef7629b185e3\" - -",
			"ws",
			"2016-08-10T00:32:08.923954Z",
			"app/my-loadbalancer/50dc6c495c0c9188",
			"10.0.0.140",
			"40914",
			"10.0.1.192",
			"8010",
			"0.001",
			"0.003",
			"0.000",
			"101",
			"101",
			"218",
			"587",
			"GET",
			"http://10.0.0.30:80/",
			"HTTP/1.1",
			"-",
			"-",
			"-",
			"arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067",
			"Root=1-58337364-23a8c76965a2ef7629b185e3",
			"-",
			"-",
		}, {
			"https 2016-08-10T23:39:43.065466Z app/my-loadbalancer/50dc6c495c0c9188 192.168.131.39:2817 10.0.0.1:80 0.086 0.048 0.037 200 200 0 57 \"GET https://www.example.com:443/ HTTP/1.1\" \"curl/7.46.0\" ECDHE-RSA-AES128-GCM-SHA256 TLSv1.2 arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067 \"Root=1-58337281-1d84f3d73c47ec4e58577259\"",
			"https",
			"2016-08-10T23:39:43.065466Z",
			"app/my-loadbalancer/50dc6c495c0c9188",
			"192.168.131.39",
			"2817",
			"10.0.0.1",
			"80",
			"0.086",
			"0.048",
			"0.037",
			"200",
			"200",
			"0",
			"57",
			"GET",
			"https://www.example.com:443/",
			"HTTP/1.1",
			"curl/7.46.0",
			"ECDHE-RSA-AES128-GCM-SHA256",
			"TLSv1.2",
			"arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067",
			"Root=1-58337281-1d84f3d73c47ec4e58577259",
			"",
			"",
		},
	}

	p := NewALBLogParser()
	for _, c := range cases {
		entry, err := p.Parse(c.line)
		if err != nil {
			t.Fatalf("parse error:%#v\n", err)
		}
		if entry.Type != c.typestring {
			t.Fatalf("parse error Type:%s\n", entry.Type)
		}
		if entry.Timestamp.Format(p.timeFormat) != c.timestamp {
			t.Fatalf("parse error Time:%v\n", entry.Timestamp)
		}
		if entry.ALB != c.alb {
			t.Fatalf("parse error Type:%s\n", entry.ALB)
		}
		if entry.Client != c.client {
			t.Fatalf("parse error Client:%s\n", entry.Client)
		}
		if entry.ClientPort != c.clientPort {
			t.Fatalf("parse error Type:%s\n", entry.ClientPort)
		}
		if entry.Target != c.target {
			t.Fatalf("parse error Target:%s\n", entry.Target)
		}
		if entry.TargetPort != c.targetPort {
			t.Fatalf("parse error TargetPort:%s\n", entry.TargetPort)
		}
		if entry.RequestProcessingTime != c.requestProcessingTime {
			t.Fatalf("parse error RequestProcessingTime:%s\n", entry.RequestProcessingTime)
		}
		if entry.TargetProcessingTime != c.targetProcessingTime {
			t.Fatalf("parse error TargetProcessingTime:%s\n", entry.TargetProcessingTime)
		}
		if entry.ResponseProcessingTime != c.responseProcessingTime {
			t.Fatalf("parse error ResponseProcessingTime:%s\n", entry.ResponseProcessingTime)
		}
		if entry.ElbStatusCode != c.elbStatusCode {
			t.Fatalf("parse error ElbStatusCode:%s\n", entry.ElbStatusCode)
		}
		if entry.TargetStatusCode != c.targetStatusCode {
			t.Fatalf("parse error TargetStatusCode:%s\n", entry.TargetStatusCode)
		}
		if entry.ReceivedBytes != c.receivedBytes {
			t.Fatalf("parse error ReceivedBytes:%s\n", entry.ReceivedBytes)
		}
		if entry.SentBytes != c.sentBytes {
			t.Fatalf("parse error SentBytes:%s\n", entry.SentBytes)
		}
		if entry.Method != c.method {
			t.Fatalf("parse error Method:%s\n", entry.Method)
		}
		if entry.URL != c.url {
			t.Fatalf("parse error Url:%s\n", entry.URL)
		}
		if entry.Protocol != c.protocol {
			t.Fatalf("parse error Protocol:%s:%s\n", c.protocol, entry.Protocol)
		}
		if entry.UserAgent != c.userAgent {
			t.Fatalf("parse error UserAgent:%s:%s\n", c.userAgent, entry.UserAgent)
		}
		if entry.SslCipher != c.sslCipher {
			t.Fatalf("parse error SslCipher:%s\n", entry.SslCipher)
		}
		if entry.SslProtocol != c.sslProtocol {
			t.Fatalf("parse error SslProtocol:%s\n", entry.SslProtocol)
		}
		if entry.TargetGroupArn != c.targetGroupArn {
			t.Fatalf("parse error TargetGroupArn:%s\n", entry.TargetGroupArn)
		}
		if entry.TraceID != c.traceID {
			t.Fatalf("parse error TraceId:%s\n", entry.TraceID)
		}
		if entry.DomainName != c.domainName {
			t.Fatalf("parse error DomainName:%s:%s\n", c.domainName, entry.DomainName)
		}
		if entry.ChosenCertArn != c.chosenCertArn {
			t.Fatalf("parse error ChosenCertArn:%s\n", entry.ChosenCertArn)
		}
	}
}
