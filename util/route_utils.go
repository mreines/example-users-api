package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const errMsgRequiredMissing = "required parameter is missing"

// EncodeJSONResponse uses the json encoder to write an interface to the http response with an optional status code
func EncodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if status != nil {
		w.WriteHeader(*status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return json.NewEncoder(w).Encode(i)
}

// ReadFormFileToTempFile reads file data from a request form and writes it to a temporary file
func ReadFormFileToTempFile(r *http.Request, key string) (*os.File, error) {
	_, fileHeader, err := r.FormFile(key)
	if err != nil {
		return nil, err
	}

	return readFileHeaderToTempFile(fileHeader)
}

// ReadFormFilesToTempFiles reads files array data from a request form and writes it to a temporary files
func ReadFormFilesToTempFiles(r *http.Request, key string) ([]*os.File, error) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return nil, err
	}

	files := make([]*os.File, 0, len(r.MultipartForm.File[key]))

	for _, fileHeader := range r.MultipartForm.File[key] {
		file, err := readFileHeaderToTempFile(fileHeader)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

// readFileHeaderToTempFile reads multipart.FileHeader and writes it to a temporary file
func readFileHeaderToTempFile(fileHeader *multipart.FileHeader) (*os.File, error) {
	formFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	defer formFile.Close()

	fileBytes, err := ioutil.ReadAll(formFile)
	if err != nil {
		return nil, err
	}

	file, err := ioutil.TempFile("", fileHeader.Filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	file.Write(fileBytes)

	return file, nil
}

// parseInt64Parameter parses a string parameter to an int64.
func parseInt64Parameter(param string, required bool) (int64, error) {
	if param == "" {
		if required {
			return 0, errors.New(errMsgRequiredMissing)
		}

		return 0, nil
	}

	return strconv.ParseInt(param, 10, 64)
}

// parseInt32Parameter parses a string parameter to an int32.
func ParseInt32Parameter(param string, required bool) (int32, error) {
	if param == "" {
		if required {
			return 0, errors.New(errMsgRequiredMissing)
		}

		return 0, nil
	}

	val, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		return -1, err
	}

	return int32(val), nil
}

// parseBoolParameter parses a string parameter to a bool
func ParseBoolParameter(param string) (bool, error) {
	val, err := strconv.ParseBool(param)
	if err != nil {
		return false, err
	}

	return bool(val), nil
}

// parseInt64ArrayParameter parses a string parameter containing array of values to []int64.
func ParseInt64ArrayParameter(param, delim string, required bool) ([]int64, error) {
	if param == "" {
		if required {
			return nil, errors.New(errMsgRequiredMissing)
		}

		return nil, nil
	}

	str := strings.Split(param, delim)
	ints := make([]int64, len(str))

	for i, s := range str {
		if v, err := strconv.ParseInt(s, 10, 64); err != nil {
			return nil, err
		} else {
			ints[i] = v
		}
	}

	return ints, nil
}

// parseInt32ArrayParameter parses a string parameter containing array of values to []int32.
func ParseInt32ArrayParameter(param, delim string, required bool) ([]int32, error) {
	if param == "" {
		if required {
			return nil, errors.New(errMsgRequiredMissing)
		}

		return nil, nil
	}

	str := strings.Split(param, delim)
	ints := make([]int32, len(str))

	for i, s := range str {
		if v, err := strconv.ParseInt(s, 10, 32); err != nil {
			return nil, err
		} else {
			ints[i] = int32(v)
		}
	}

	return ints, nil
}
