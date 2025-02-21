package http

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Response struct {
	Version string

	StatusCode []byte
	Reason     []byte
	StatusLine []byte
	Headers    []*HeaderKV
	Body       []byte

	buf bytes.Buffer
}

func (r *Response) GetStatusCode() int {
	code, err := strconv.Atoi(string(r.StatusCode))
	if err != nil {
		return -1
	}
	return code
}

// Len implements the Response buffer length method.
func (r *Response) Len() int {
	return r.buf.Len()
}

// Parse implements the Response Parse method.
func (r *Response) Parse(b []byte) (int, error) {
	n, err := r.buf.Read(b)
	if err != nil {
		return n, err
	}
	return r.parse()
}

// Read implements the Response Read method.
func (r *Response) Write(b []byte) (int, error) {
	nwrite, err := r.buf.Write(b)
	if err != nil {
		return nwrite, err
	}

	return r.parse()
}

// WriteTo implements the Response WriteTo method.
func (r *Response) WriteTo(w io.Writer) (n int64, err error) {
	return r.buf.WriteTo(w)
}

// Read implements the Response Read method.
func (r *Response) Read(b []byte) (int, error) {
	return r.buf.Read(b)
}

// ReadFrom implements the Response ReadFrom method.
func (r *Response) ReadFrom(in io.Reader) (int64, error) {
	nread, err := r.buf.ReadFrom(in)
	if err != nil {
		return nread, err
	}

	n, err := r.parse()

	return int64(n), err
}

// Close implements the Response Read method.
func (r *Response) Close() error {
	return nil
}

// String implements the Response String method.
func (r *Response) String() string {
	headerCnt := len(r.Headers)
	bodyLength := len(r.Body)
	return fmt.Sprintf("Response status code: %s http version: %s headers number: %d body length: %d\nResponse Line: %s\n", string(r.StatusCode), r.Version, headerCnt, bodyLength, string(r.StatusLine))
}

// ReadConn reads a complete HTTP response from the connection, handles gzip decoding and chunked transfer encoding if necessary.
func (r *Response) ReadConn(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// 读取并忽略状态行
	_, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read status line: %w", err)
	}

	// Read the headers
	headers, err := readHeaders(reader)
	if err != nil {
		return nil, err
	}
	r.Headers = headers

	// Check if the response is chunked
	isChunked := strings.ToLower(r.Find("Transfer-Encoding")) == "chunked"

	// Read the body
	var body []byte
	if isChunked {
		body, err = readChunkedBody(reader)
	} else {
		body, err = io.ReadAll(reader)
	}
	if err != nil {
		return nil, err
	}

	r.Body = body

	// Check if we need to decompress the body
	if strings.ToLower(r.Find("Content-Encoding")) == "gzip" {
		gzReader, err := gzip.NewReader(bytes.NewReader(r.Body))
		if err != nil {
			return nil, err
		}
		defer gzReader.Close()
		decompressed, err := io.ReadAll(gzReader)
		if err != nil {
			return nil, err
		}
		r.Body = decompressed
	}

	return r.Body, nil
}

// readHeaders reads the HTTP headers from the reader.
func readHeaders(reader *bufio.Reader) ([]*HeaderKV, error) {
	var headers []*HeaderKV
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid header line: %s", line)
		}
		headers = append(headers, &HeaderKV{Key: []byte(parts[0]), Value: []byte(parts[1])})
	}
	return headers, nil
}

// readChunkedBody reads the body of a chunked transfer encoded response.
func readChunkedBody(reader *bufio.Reader) ([]byte, error) {
	var body bytes.Buffer
	for {
		chunkSizeLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		chunkSize, err := strconv.ParseUint(strings.TrimRight(chunkSizeLine, "\r\n"), 16, 64)
		if err != nil {
			return nil, err
		}
		if chunkSize == 0 {
			break
		}
		chunk := make([]byte, chunkSize)
		_, err = io.ReadFull(reader, chunk)
		if err != nil {
			return nil, err
		}
		body.Write(chunk)
		// Read the trailing newline after the chunk
		_, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
	}
	return body.Bytes(), nil
}

// Helper method to find a header by key
func (h *Response) Find(key string) string {
	for _, header := range h.Headers {
		if strings.ToLower(string(header.Key)) == strings.ToLower(key) {
			return string(header.Value)
		}
	}
	return ""
}
