package ctlog

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/g0rbe/gmod/net/dns"
	"github.com/g0rbe/gmod/slicer"
	gct "github.com/google/certificate-transparency-go"
	"github.com/google/certificate-transparency-go/client"
	"github.com/google/certificate-transparency-go/jsonclient"
)

type Log struct {
	Name   string // Log name
	URI    string // Log URI
	PubKey string // Base64 encoded string of the log's public key
}

// A list of logs used by Chrome.
// From: https://source.chromium.org/chromium/chromium/src/+/main:components/certificate_transparency/data/log_list.json
var Logs = []Log{
	{Name: "Argon2022", URI: "https://ct.googleapis.com/logs/argon2022/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEeIPc6fGmuBg6AJkv/z7NFckmHvf/OqmjchZJ6wm2qN200keRDg352dWpi7CHnSV51BpQYAj1CQY5JuRAwrrDwg=="},
	{Name: "Argon2023", URI: "https://ct.googleapis.com/logs/argon2023/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE0JCPZFJOQqyEti5M8j13ALN3CAVHqkVM4yyOcKWCu2yye5yYeqDpEXYoALIgtM3TmHtNlifmt+4iatGwLpF3eA=="},
	{Name: "Argon2024", URI: "https://ct.googleapis.com/logs/us1/argon2024/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHblsqctplMVc5ramA7vSuNxUQxcomQwGAVAdnWTAWUYr3MgDHQW0LagJ95lB7QT75Ve6JgT2EVLOFGU7L3YrwA=="},
	{Name: "Xenon2022", URI: "https://ct.googleapis.com/logs/xenon2022/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE+WS9FSxAYlCVEzg8xyGwOrmPonoV14nWjjETAIdZvLvukPzIWBMKv6tDNlQjpIHNrUcUt1igRPpqoKDXw2MeKw=="},
	{Name: "Xenon2023", URI: "https://ct.googleapis.com/logs/xenon2023/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEchY+C+/vzj5g3ZXLY3q5qY1Kb2zcYYCmRV4vg6yU84WI0KV00HuO/8XuQqLwLZPjwtCymeLhQunSxgAnaXSuzg=="},
	{Name: "Xenon2024", URI: "https://ct.googleapis.com/logs/eu1/xenon2024/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEuWDgNB415GUAk0+QCb1a7ETdjA/O7RE+KllGmjG2x5n33O89zY+GwjWlPtwpurvyVOKoDIMIUQbeIW02UI44TQ=="},
	{Name: "Nimbus2022", URI: "https://ct.cloudflare.com/logs/nimbus2022/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAESLJHTlAycmJKDQxIv60pZG8g33lSYxYpCi5gteI6HLevWbFVCdtZx+m9b+0LrwWWl/87mkNN6xE0M4rnrIPA/w=="},
	{Name: "Nimbus2023", URI: "https://ct.cloudflare.com/logs/nimbus2023/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEi/8tkhjLRp0SXrlZdTzNkTd6HqmcmXiDJz3fAdWLgOhjmv4mohvRhwXul9bgW0ODgRwC9UGAgH/vpGHPvIS1qA=="},
	{Name: "Nimbus2024", URI: "https://ct.cloudflare.com/logs/nimbus2024/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEd7Gbe4/mizX+OpIpLayKjVGKJfyTttegiyk3cR0zyswz6ii5H+Ksw6ld3Ze+9p6UJd02gdHrXSnDK0TxW8oVSA=="},
	{Name: "Yeti2023", URI: "https://yeti2023.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfQ0DsdWYitzwFTvG3F4Nbj8Nv5XIVYzQpkyWsU4nuSYlmcwrAp6m092fsdXEw6w1BAeHlzaqrSgNfyvZaJ9y0Q=="},
	{Name: "Yeti2024", URI: "https://yeti2024.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV7jBbzCkfy7k8NDZYGITleN6405Tw7O4c4XBGA0jDliE0njvm7MeLBrewY+BGxlEWLcAd2AgGnLYgt6unrHGSw=="},
	{Name: "Yeti2025", URI: "https://yeti2025.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE35UAXhDBAfc34xB00f+yypDtMplfDDn+odETEazRs3OTIMITPEy1elKGhj3jlSR82JGYSDvw8N8h8bCBWlklQw=="},
	{Name: "Nessie2023", URI: "https://nessie2023.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEXu8iQwSCRSf2CbITGpUpBtFVt8+I0IU0d1C36Lfe1+fbwdaI0Z5FktfM2fBoI1bXBd18k2ggKGYGgdZBgLKTg=="},
	{Name: "Nessie2024", URI: "https://nessie2024.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAELfyieza/VpHp/j/oPfzDp+BhUuos6QWjnycXgQVwa4FhRIr4OxCAQu0DLwBQIfxBVISjVNUusnoWSyofK2YEKw=="},
	{Name: "Nessie2025", URI: "https://nessie2025.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8vDwp4uBLgk5O59C2jhEX7TM7Ta72EN/FklXhwR/pQE09+hoP7d4H2BmLWeadYC3U6eF1byrRwZV27XfiKFvOA=="},
	{Name: "Sabre", URI: "https://sabre.ct.comodo.com/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8m/SiQ8/xfiHHqtls9m7FyOMBg4JVZY9CgiixXGz0akvKD6DEL8S0ERmFe9U4ZiA0M4kbT5nmuk3I85Sk4bagA=="},
	//{Name: "Mammoth", URI: "https://mammoth.ct.comodo.com/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7+R9dC4VFbbpuyOL+yy14ceAmEf7QGlo/EmtYU6DRzwat43f/3swtLr/L8ugFOOt1YU/RFmMjGCL17ixv66MZw=="},
	{Name: "Oak2022", URI: "https://oak.ct.letsencrypt.org/2022/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhjyxDVIjWt5u9sB/o2S8rcGJ2pdZTGA8+IpXhI/tvKBjElGE5r3de4yAfeOPhqTqqc+o7vPgXnDgu/a9/B+RLg=="},
	{Name: "Oak2023", URI: "https://oak.ct.letsencrypt.org/2023/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEsz0OeL7jrVxEXJu+o4QWQYLKyokXHiPOOKVUL3/TNFFquVzDSer7kZ3gijxzBp98ZTgRgMSaWgCmZ8OD74mFUQ=="},
	{Name: "Oak2024H1", URI: "https://oak.ct.letsencrypt.org/2024h1/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEVkPXfnvUcre6qVG9NpO36bWSD+pet0Wjkv3JpTyArBog7yUvuOEg96g6LgeN5uuk4n0kY59Gv5RzUo2Wrqkm/Q=="},
	{Name: "Oak2024H2", URI: "https://oak.ct.letsencrypt.org/2024h2/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE13PWU0fp88nVfBbC1o9wZfryUTapE4Av7fmU01qL6E8zz8PTidRfWmaJuiAfccvKu5+f81wtHqOBWa+Ss20waA=="},
	{Name: "TrustAsia2022", URI: "https://ct.trustasia.com/log2022/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEu1LyFs+SC8555lRtwjdTpPX5OqmzBewdvRbsMKwu+HliNRWOGtgWLuRIa/bGE/GWLlwQ/hkeqBi4Dy3DpIZRlw=="},
	{Name: "TrustAsia2023", URI: "https://ct.trustasia.com/log2023/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEpBFS2xdBTpDUVlESMFL4mwPPTJ/4Lji18Vq6+ji50o8agdqVzDPsIShmxlY+YDYhINnUrF36XBmhBX3+ICP89Q=="},
	{Name: "TrustAsia2024-2", URI: "https://ct2024.trustasia.com/log2024/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEp2TieYE/YdfsxvhlKB2gtGYzwyXVCpV4nI/+pCrYj35y4P6of/ixLYXAjhJ0DS+Mq9d/eh7ZhDM56P2JX5ZICA=="},
}

// LogByName returns the Log from Logs based on the name.
// The name is case insensitive!
//
// Returns nil if the Log not found.
func LogByName(name string) *Log {

	name = strings.ToLower(name)

	for i := range Logs {
		if strings.ToLower(Logs[i].Name) == name {
			return &Logs[i]
		}
	}

	return nil
}

// Size returns the log size of log url.
func Size(url string) (int64, error) {

	c, err := client.New(url, http.DefaultClient, jsonclient.Options{})
	if err != nil {
		return 0, fmt.Errorf("failed to create client: %w", err)
	}

	sth, err := c.GetSTH(context.TODO())
	if err != nil {
		return 0, fmt.Errorf("failed to get SignedTreeHead: %w", err)
	}

	return int64(sth.TreeSize), nil
}

// NumLeft returns the number of log entries left.
func NumLeft(url string, index int64) (int64, error) {

	size, err := Size(url)
	if err != nil {
		return 0, fmt.Errorf("failed to get size: %w", err)
	}

	return size - index, nil
}

// MaxBatchSize returns the max batch size that can retrieved in a single query.
func MaxBatchSize(url string) (int64, error) {

	c, err := client.New(url, http.DefaultClient, jsonclient.Options{})
	if err != nil {
		return 0, fmt.Errorf("failed to create client: %w", err)
	}

	resp, err := c.GetRawEntries(context.TODO(), 0, 10000)
	if err != nil {
		return 0, fmt.Errorf("failed to get raw entries: %w", err)
	}

	return int64(len(resp.Entries)), nil
}

// GetDomains returns the domains parsed from the log's certificates.
// start is the start index and fetch as many log entries as possible with one query.
// The returned int64 counts the number of parsed log entries.
//
// This function use IsValid() and append only the unique entries.
func GetDomains(url string, start int64) ([]string, int64, error) {

	c, err := client.New(url, http.DefaultClient, jsonclient.Options{})
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create client: %w", err)
	}

	resp, err := c.GetRawEntries(context.TODO(), start, start+10000)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get raw entries: %w", err)
	}

	var n int64
	r := make([]string, 0, len(resp.Entries))

	for i := range resp.Entries {

		// index for LogEntryFromLeaf() doenst matter, use 0.
		e, err := gct.LogEntryFromLeaf(0, &resp.Entries[i])
		if err != nil {
			return r, n, fmt.Errorf("failed to convert leaf entry to log: %w", err)
		}

		n++

		if e.X509Cert != nil {

			if dns.IsDomain(e.X509Cert.Subject.CommonName) {
				r = slicer.AppendUnique(r, e.X509Cert.Subject.CommonName)
			}

			for ii := range e.X509Cert.DNSNames {
				if dns.IsDomain(e.X509Cert.DNSNames[ii]) {
					r = slicer.AppendUnique(r, e.X509Cert.DNSNames[ii])
				}
			}

			for ii := range e.X509Cert.PermittedDNSDomains {
				if dns.IsDomain(e.X509Cert.PermittedDNSDomains[ii]) {
					r = slicer.AppendUnique(r, e.X509Cert.PermittedDNSDomains[ii])
				}
			}

			for ii := range e.X509Cert.ExcludedDNSDomains {
				if dns.IsDomain(e.X509Cert.ExcludedDNSDomains[ii]) {
					r = slicer.AppendUnique(r, e.X509Cert.ExcludedDNSDomains[ii])
				}
			}
		}

		if e.Precert != nil && e.Precert.TBSCertificate != nil {

			if dns.IsDomain(e.Precert.TBSCertificate.Subject.CommonName) {
				r = slicer.AppendUnique(r, e.Precert.TBSCertificate.Subject.CommonName)
			}

			for ii := range e.Precert.TBSCertificate.DNSNames {
				if dns.IsDomain(e.Precert.TBSCertificate.DNSNames[ii]) {
					r = slicer.AppendUnique(r, e.Precert.TBSCertificate.DNSNames[ii])
				}
			}

			for ii := range e.Precert.TBSCertificate.PermittedDNSDomains {
				if dns.IsDomain(e.Precert.TBSCertificate.PermittedDNSDomains[ii]) {
					r = slicer.AppendUnique(r, e.Precert.TBSCertificate.PermittedDNSDomains[ii])
				}
			}

			for ii := range e.Precert.TBSCertificate.ExcludedDNSDomains {
				if dns.IsDomain(e.Precert.TBSCertificate.ExcludedDNSDomains[ii]) {
					r = slicer.AppendUnique(r, e.Precert.TBSCertificate.ExcludedDNSDomains[ii])
				}
			}
		}
	}

	return r, n, nil
}
