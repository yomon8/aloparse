package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type LogEntry struct {
	Type                   string
	Timestamp              time.Time
	ALB                    string
	Client                 string
	ClientPort             string
	Target                 string
	TargetPort             string
	RequestProcessingTime  string
	TargetProcessingTime   string
	ResponseProcessingTime string
	ElbStatusCode          string
	TargetStatusCode       string
	ReceivedBytes          string
	SentBytes              string
	Method                 string
	URL                    string
	Protocol               string
	UserAgent              string
	SslCipher              string
	SslProtocol            string
	TargetGroupArn         string
	TraceID                string
	DomainName             string
	ChosenCertArn          string
}

// ALBLogParser can parse AWS ALB Access Log
type ALBLogParser struct {
	timeFormat string
	regexp     *regexp.Regexp
}

// NewALBLogParser create AlbLogParser instance
// refs below url
// http://docs.aws.amazon.com/en_us/elasticloadbalancing/latest/application/load-balancer-access-logs.html
func NewALBLogParser() *ALBLogParser {
	a := new(ALBLogParser)
	a.regexp = regexp.MustCompile(`([^\"].+?)\"(.+?)\" \"(.+?)\" ([^\"].+?)\"(.+?)\"(.*)`)
	a.timeFormat = "2006-01-02T15:04:05.000000Z"
	return a
}

// Parse log entry text
func (a *ALBLogParser) Parse(line string) (*LogEntry, error) {
	var (
		domainName    string
		chosenCertArn string
		target        string
		targetPort    string
	)

	reg := a.regexp.FindStringSubmatch(line)
	if len(reg) < 6 {
		return nil, ErrRegexp
	}
	s := strings.Split(reg[1], " ")
	dt, err := time.Parse(a.timeFormat, s[1])
	if err != nil {
		return nil, errors.New("time parse error")
	}
	c := strings.Split(s[3], ":")
	if len(c) != 2 {
		return nil, errors.New(fmt.Sprintf("invalid client:%#v", s[3]))
	}
	t := strings.Split(s[4], ":")
	if len(t) == 2 {
		target = t[0]
		targetPort = t[1]
	} else {
		target = "-"
		targetPort = "-"
	}
	r := strings.Split(reg[2], " ")
	ssl := strings.Split(reg[4], " ")
	if len(reg) == 7 {
		d := strings.Split(reg[6], " ")
		if len(d) == 3 {
			domainName = strings.Trim(d[1], "\"")
			chosenCertArn = strings.Trim(d[2], "\"")
		}
	}
	return &LogEntry{
		Type:                   s[0],
		Timestamp:              dt,
		ALB:                    s[2],
		Client:                 c[0],
		ClientPort:             c[1],
		Target:                 target,
		TargetPort:             targetPort,
		RequestProcessingTime:  s[5],
		TargetProcessingTime:   s[6],
		ResponseProcessingTime: s[7],
		ElbStatusCode:          s[8],
		TargetStatusCode:       s[9],
		ReceivedBytes:          s[10],
		SentBytes:              s[11],
		Method:                 r[0],
		URL:                    r[1],
		Protocol:               r[2],
		UserAgent:              reg[3],
		SslCipher:              ssl[0],
		SslProtocol:            ssl[1],
		TargetGroupArn:         ssl[2],
		TraceID:                reg[5],
		DomainName:             domainName,
		ChosenCertArn:          chosenCertArn,
	}, nil
}
