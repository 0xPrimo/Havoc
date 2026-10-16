package events

import (
	"encoding/base64"
	"strings"
	"time"

	"Havoc/pkg/packager"
)

var Gate gate

func (g gate) SendStageless(Format string, payload []byte) packager.Package {
	return packager.Package{
		Head: packager.Head{
			Event: packager.Type.Gate.Type,
			Time:  time.Now().Format("02/01/2006 15:04:05"),
		},

		Body: packager.Body{
			SubEvent: packager.Type.Gate.Stageless,
			Info: map[string]any{
				"PayloadArray": base64.StdEncoding.EncodeToString(payload),
				"Format":       Format,
				"FileName":     Format,
			},
		},
	}
}

func (g gate) SendUDRL(Format string, Config map[string]any, Payload []byte) packager.Package {
	var ServiceName string = ""
	if strings.Contains(Format, ".svc.exe") {
		ServiceName = Config["Service Name"].(string)
	}

	return packager.Package{
		Head: packager.Head{
			Event: packager.Type.Gate.Type,
			Time:  time.Now().Format("02/01/2006 15:04:05"),
		},

		Body: packager.Body{
			SubEvent: packager.Type.Gate.UserDefinedReflectiveLoader,
			Info: map[string]any{
				"PayloadArray": base64.StdEncoding.EncodeToString(Payload),
				"Format":       Format,
				"FileName":     Format,
				"ServiceName":  ServiceName,
			},
		},
	}
}

func (g gate) SendConsoleMessage(MsgType, text string) packager.Package {
	return packager.Package{
		Head: packager.Head{
			Event: packager.Type.Gate.Type,
			Time:  time.Now().Format("02/01/2006 15:04:05"),
		},

		Body: packager.Body{
			SubEvent: packager.Type.Gate.Stageless,
			Info: map[string]any{
				"MessageType": MsgType,
				"Message":     text,
			},
		},
	}
}
