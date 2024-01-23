package res

import (
	"encoding/json"
	"errors"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, data interface{}, status int, headers http.Header  ) error {
  resBytes, err := json.Marshal(data)
  if err != nil {
    return errors.New("Error Encoding Data. Please check the data passed.")
  }
  resBytes = append(resBytes, '\n')
  for key, value := range headers {
    w.Header()[key] = value
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  w.Write(resBytes)
  return nil
}
