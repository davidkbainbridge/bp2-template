package hooks

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"
)

func GetMyIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // down
		}
		if iface.Flags&(net.FlagLoopback) != 0 {
			continue // broadcast or loopback
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		var ip net.IP
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() || ip.IsMulticast() {
				continue
			}
			if ip.To4() == nil {
				continue // not v4
			}
			return ip.String(), nil
		}
	}
	return "", nil
}

func HandleHooks(update chan interface{}) {
	hooker := http.NewServeMux()

	hooker.HandleFunc("/api/v1/hook/southbound-update", func(w http.ResponseWriter, r *http.Request) {
		log.Println("SOUNDBOUND-UPDATE")
	})
	hooker.HandleFunc("/api/v1/hook/peer-status", func(w http.ResponseWriter, r *http.Request) {
		var data interface{}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Printf("ERROR: unable to decode peering data: %s\n", err)
			return
		}
		b, _ := json.MarshalIndent(data, "", "    ")
		log.Printf("PEER STATUS: %v\n", string(b))

		ip, err := GetMyIP()
		if err != nil {
			log.Printf("ERROR: unable to determine IP address: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"command": "peer-join",
			"data": map[string]interface{}{
				"ip": ip,
			},
		}
		if b, err := json.Marshal(response); err != nil {
			log.Printf("ERROR: Unable to marshal response: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			log.Printf("RETURN: %s\n", string(b))
			w.Write(b)
		}
	})
	hooker.HandleFunc("/api/v1/hook/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		log.Println("HEARTBEAT")
	})
	hooker.HandleFunc("/api/v1/hook/peer-update", func(w http.ResponseWriter, r *http.Request) {
		var data interface{}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Printf("ERROR: unable to decode peering data: %s\n", err)
			return
		}
		b, _ := json.MarshalIndent(data, "", "    ")
		log.Printf("PEER UPDATE: %v\n", string(b))

		ip, err := GetMyIP()
		if err != nil {
			log.Printf("ERROR: unable to determine IP address: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"command": "peer-join",
			"data": map[string]interface{}{
				"ip": ip,
			},
		}
		if b, err := json.Marshal(response); err != nil {
			log.Printf("ERROR: Unable to marshal response: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			log.Printf("RETURN: %s\n", string(b))
			w.Write(b)
		}
	})

	s := &http.Server{
		Addr:           ":6789",
		Handler:        hooker,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
