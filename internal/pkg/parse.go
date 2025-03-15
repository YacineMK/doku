package pkg

import (
	"fmt"
	"strings"
	"github.com/YacineMK/doku/internal/types"
)

func ParseRequest(requestData string) (*types.HttpRequest, error) {
	var req types.HttpRequest

	parts := strings.SplitN(requestData, "\r\n\r\n", 2)
	headerPart := parts[0]
	bodyPart := ""
	if len(parts) > 1 {
		bodyPart = parts[1]
	}

	lines := strings.Split(headerPart, "\r\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("invalid request format")
	}

	requestLineParts := strings.Fields(lines[0])
	if len(requestLineParts) < 3 {
		return nil, fmt.Errorf("invalid request line")
	}
	req.Method = requestLineParts[0]
	req.Path = requestLineParts[1]
	req.Headers = make(map[string]string)

	for _, line := range lines[1:] {
		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) == 2 {
			req.Headers[strings.TrimSpace(headerParts[0])] = strings.TrimSpace(headerParts[1])
		}
	}

	req.Body = bodyPart

	return &req, nil
}
